package crossplane

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/intelops/go-common/credentials"
	"github.com/intelops/go-common/logging"
	captenstore "github.com/kube-tarian/kad/capten/agent/internal/capten-store"
	"github.com/kube-tarian/kad/capten/agent/internal/temporalclient"
	"github.com/kube-tarian/kad/capten/agent/internal/workers"

	"github.com/kube-tarian/kad/capten/agent/internal/pb/captenpluginspb"

	"github.com/kube-tarian/kad/capten/common-pkg/credential"
	"github.com/kube-tarian/kad/capten/common-pkg/k8s"
	"github.com/kube-tarian/kad/capten/model"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

var (
	readyStatusType = "ready"

	clusterNotReadyStatus    = "NotReady"
	clusterReadyStatus       = "Ready"
	clusterDeletingStatus    = "Deleting"
	readyStatusValue         = "True"
	NorReadyStatusValue      = "False"
	clusterSecretName        = "%s-cluster"
	kubeConfig               = "kubeconfig"
	k8sEndpoint              = "endpoint"
	k8sClusterCA             = "clusterCA"
	managedClusterEntityName = "managedcluster"
)

var (
	cgvk = schema.GroupVersionResource{Group: "prodready.cluster", Version: "v1alpha1", Resource: "clusterclaims"}
)

type ClusterClaimSyncHandler struct {
	log     logging.Logger
	tc      *temporalclient.Client
	dbStore *captenstore.Store
}

func NewClusterClaimSyncHandler(log logging.Logger, dbStore *captenstore.Store) (*ClusterClaimSyncHandler, error) {
	tc, err := temporalclient.NewClient(log)
	if err != nil {
		return nil, err
	}

	return &ClusterClaimSyncHandler{log: log, dbStore: dbStore, tc: tc}, nil
}

func registerK8SClusterClaimWatcher(log logging.Logger, dbStore *captenstore.Store, dynamicClient dynamic.Interface) error {
	obj, err := NewClusterClaimSyncHandler(log, dbStore)
	if err != nil {
		return err
	}
	return k8s.RegisterDynamicInformers(obj, dynamicClient, cgvk)
}

func getClusterClaimObj(obj any) (*model.ClusterClaim, error) {
	clusterClaimByte, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var clObj model.ClusterClaim
	err = json.Unmarshal(clusterClaimByte, &clObj)
	if err != nil {
		return nil, err
	}

	return &clObj, nil
}

func (h *ClusterClaimSyncHandler) OnAdd(obj interface{}) {
	h.log.Info("Crossplane ClusterCliam Add Callback")

	newCcObj, err := getClusterClaimObj(obj)
	if newCcObj == nil {
		h.log.Errorf("failed to read ClusterCliam object, %v", err)
		return
	}

	if err = h.updateManagedClusters([]model.ClusterClaim{*newCcObj}); err != nil {
		h.log.Errorf("failed to update ClusterCliam object, %v", err)
		return
	}

	h.log.Info("cluster-claims resources synched")
}

func (h *ClusterClaimSyncHandler) OnUpdate(oldObj, newObj interface{}) {
	h.log.Info("Crossplane ClusterCliam Update Callback")

	prevObj, err := getClusterClaimObj(oldObj)
	if prevObj == nil {
		h.log.Errorf("failed to read ClusterCliam old object %v", err)
		return
	}

	newCcObj, err := getClusterClaimObj(oldObj)
	if newCcObj == nil {
		h.log.Errorf("failed to read ClusterCliam new object %v", err)
		return
	}

	if err = h.updateManagedClusters([]model.ClusterClaim{*newCcObj}); err != nil {
		h.log.Errorf("failed to update ClusterCliam object, %v", err)
		return
	}

	h.log.Info("cluster-claims resources synched")
}

func (h *ClusterClaimSyncHandler) OnDelete(obj interface{}) {
	h.log.Info("Crossplane ClusterCliam Delete Callback")

	newCcObj, err := getClusterClaimObj(obj)
	if newCcObj == nil {
		h.log.Errorf("failed to read ClusterCliam object, %v", err)
		return
	}

	if err = h.deleteManagedClusters(*newCcObj); err != nil {
		h.log.Errorf("failed to delete ClusterCliam object, %v", err)
		return
	}

}

func (h *ClusterClaimSyncHandler) Sync() error {
	h.log.Debug("started to sync ClusterCliam resources")

	k8sclient, err := k8s.NewK8SClient(h.log)
	if err != nil {
		return fmt.Errorf("failed to initalize k8s client: %v", err)
	}

	objList, err := k8sclient.DynamicClient.ListAllNamespaceResource(context.TODO(), cgvk)
	if err != nil {
		return fmt.Errorf("failed to list cluster claim resources, %v", err)
	}

	clusterClaimByte, err := json.Marshal(objList)
	if err != nil {
		return fmt.Errorf("failed to marshal cluster claim resources, %v", err)
	}

	var clObj model.ClusterClaimList
	err = json.Unmarshal(clusterClaimByte, &clObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal cluster claim resources, %v", err)
	}

	if err = h.updateManagedClusters(clObj.Items); err != nil {
		return fmt.Errorf("failed to update clusters in DB, %v", err)
	}
	h.log.Info("cluster-claims resources synched")

	if err := h.syncClusterClaimsWithDB(clObj.Items); err != nil {
		return fmt.Errorf("failed to sync clusters in DB, %v", err)
	}

	return nil
}

func (h *ClusterClaimSyncHandler) updateManagedClusters(clusterCliams []model.ClusterClaim) error {
	k8sclient, err := k8s.NewK8SClient(h.log)
	if err != nil {
		return fmt.Errorf("failed to get k8s client, %v", err)
	}

	clusters, err := h.getManagedClusters()
	if err != nil {
		return fmt.Errorf("failed to get managed clusters from DB, %v", err)
	}

	for _, clusterCliam := range clusterCliams {
		h.log.Infof("processing cluster claim %s", clusterCliam.Metadata.Name)
		readyStatus := h.getClusterClaimStatus(clusterCliam.Status.Conditions)
		h.log.Infof("cluster claim %s status: %s-%s-%s", clusterCliam.Metadata.Name,
			clusterCliam.Status.NodePoolStatus, clusterCliam.Status.ControlPlaneStatus, readyStatus)

		if !(strings.EqualFold(clusterCliam.Status.NodePoolStatus, "active") &&
			strings.EqualFold(clusterCliam.Status.ControlPlaneStatus, "active")) {
			h.log.Infof("cluster %s is not created", clusterCliam.Metadata.Name)
			return nil
		}

		managedCluster := &captenpluginspb.ManagedCluster{}
		managedCluster.ClusterName = clusterCliam.Metadata.Name

		clusterObj, ok := clusters[managedCluster.ClusterName]
		if !ok {
			managedCluster.Id = uuid.New().String()
		} else {
			h.log.Infof("found existing managed clusterId %s, updating", clusterObj.Id)
			managedCluster.Id = clusterObj.Id
			managedCluster.ClusterDeployStatus = clusterObj.ClusterDeployStatus
		}

		if strings.EqualFold(readyStatus, readyStatusValue) {
			managedCluster.ClusterDeployStatus = clusterReadyStatus
		} else {
			managedCluster.ClusterDeployStatus = clusterNotReadyStatus
		}

		secretName := fmt.Sprintf(clusterSecretName, clusterCliam.Spec.Id)
		resp, err := k8sclient.GetSecretData(clusterCliam.Metadata.Namespace, secretName)
		if err != nil {
			h.log.Errorf("failed to get secret %s/%s, %v", clusterCliam.Metadata.Namespace, secretName, err)
			continue
		}

		clusterEndpoint := resp.Data[k8sEndpoint]
		managedCluster.ClusterEndpoint = clusterEndpoint
		cred := map[string]string{}
		cred[kubeConfig] = resp.Data[kubeConfig]
		cred[k8sClusterCA] = resp.Data[k8sClusterCA]
		cred[k8sEndpoint] = clusterEndpoint

		err = credential.PutGenericCredential(context.TODO(), managedClusterEntityName, managedCluster.Id, cred)
		if err != nil {
			h.log.Errorf("failed to store credential for %s, %v", managedCluster.Id, err)
			continue
		}

		err = h.dbStore.UpsertManagedCluster(managedCluster)
		if err != nil {
			h.log.Info("failed to update information to db, %v", err)
			continue
		}
		h.log.Infof("updated the cluster claim %s with status %s", managedCluster.ClusterName, managedCluster.ClusterDeployStatus)

		if managedCluster.ClusterDeployStatus == clusterReadyStatus {
			err = h.triggerClusterUpdates(clusterCliam.Spec.Id, managedCluster.Id)
			if err != nil {
				h.log.Info("failed to trigger cluster update workflow, %v", err)
				continue
			}
			h.log.Infof("triggered cluster update workflow for cluster %s", managedCluster.ClusterName)
		}
	}
	return nil
}

func (h *ClusterClaimSyncHandler) getManagedClusters() (map[string]*captenpluginspb.ManagedCluster, error) {
	clusters, err := h.dbStore.GetManagedClusters()
	if err != nil {
		return nil, err
	}

	clusterEndpointMap := map[string]*captenpluginspb.ManagedCluster{}
	for _, cluster := range clusters {
		clusterEndpointMap[cluster.ClusterName] = cluster
	}
	return clusterEndpointMap, nil
}

func (h *ClusterClaimSyncHandler) triggerClusterUpdates(clusterName, managedClusterID string) error {
	proj, err := h.dbStore.GetCrossplaneProject()
	if err != nil {
		return err
	}

	ci := model.CrossplaneClusterUpdate{RepoURL: proj.GitProjectUrl, GitProjectId: proj.GitProjectId,
		ManagedClusterName: clusterName, ManagedClusterId: managedClusterID}
	wd := workers.NewConfig(h.tc, h.log)
	_, err = wd.SendEvent(context.TODO(), &model.ConfigureParameters{Resource: model.CrossPlaneResource, Action: model.CrossPlaneClusterUpdate}, ci)
	return err
}

func (h *ClusterClaimSyncHandler) getClusterClaimStatus(conditions []model.ClusterClaimCondition) (readyStatus string) {
	for _, condition := range conditions {
		switch strings.ToLower(condition.Type) {
		case readyStatusType:
			readyStatus = condition.Status
		}
	}
	return
}

func (h *ClusterClaimSyncHandler) deleteManagedClusters(clusterCliam model.ClusterClaim) error {

	clusters, err := h.getManagedClusters()
	if err != nil {
		return fmt.Errorf("failed to get managed clusters from DB, %v", err)
	}

	var isDeleteManagedCluster bool
	var managedCluster *captenpluginspb.ManagedCluster
	for _, v := range clusters {
		if v.ClusterName == clusterCliam.Metadata.Name {
			isDeleteManagedCluster = true
			managedCluster = v
			break
		}
	}

	if isDeleteManagedCluster {
		Cluster := &captenpluginspb.ManagedCluster{
			Id:                  managedCluster.Id,
			ClusterName:         managedCluster.ClusterName,
			ClusterEndpoint:     managedCluster.ClusterEndpoint,
			AppDeployStatus:     managedCluster.AppDeployStatus,
			ClusterDeployStatus: clusterDeletingStatus,
		}
		if err := h.dbStore.UpsertManagedCluster(Cluster); err != nil {
			return fmt.Errorf("failed to update managed cluster from DB, %v", err)
		}
	}

	err = h.triggerClusterDelete(clusterCliam.Spec.Id, managedCluster.Id, managedCluster.ClusterEndpoint)
	if err != nil {
		return fmt.Errorf("failed to trigger cluster delete workflow, %v", err)
	}

	if isDeleteManagedCluster {
		if err := h.dbStore.DeleteManagedClusterById(managedCluster.Id); err != nil {
			return fmt.Errorf("failed to delete managed cluster from DB, %v", err)
		}
	} else {
		h.log.Info("failed to delete managed cluster from DB, %s Cluster is not stored in ManagedClusters table", clusterCliam.Metadata.Name)
		return nil
	}

	h.log.Infof("triggered cluster delete workflow for cluster %s", managedCluster.ClusterName)

	return nil
}

func (h *ClusterClaimSyncHandler) deleteManagedClusterCredential(ctx context.Context, id string) error {
	credPath := fmt.Sprintf("%s/%s/%s", credentials.GenericCredentialType, managedClusterEntityName, id)
	credAdmin, err := credentials.NewCredentialAdmin(ctx)
	if err != nil {
		h.log.Audit("security", "storecred", "failed", "system", "failed to intialize credentials client for %s", credPath)
		h.log.Errorf("failed to delete credential for %s, %v", credPath, err)
		return err
	}

	err = credAdmin.DeleteCredential(ctx, credentials.GenericCredentialType, managedClusterEntityName, id)
	if err != nil {
		h.log.Audit("security", "storecred", "failed", "system", "failed to store crendential for %s", credPath)
		h.log.Errorf("failed to delete credential for %s, %v", credPath, err)
		return err
	}
	h.log.Audit("security", "storecred", "success", "system", "credential stored for %s", credPath)
	h.log.Infof("deleted credential for entity %s", credPath)
	return nil
}

func (h *ClusterClaimSyncHandler) triggerClusterDelete(clusterName, managedClusterID, clusterEndpoint string) error {
	proj, err := h.dbStore.GetCrossplaneProject()
	if err != nil {
		return err
	}

	ci := model.CrossplaneClusterUpdate{RepoURL: proj.GitProjectUrl, GitProjectId: proj.GitProjectId,
		ManagedClusterName: clusterName, ManagedClusterId: managedClusterID, ClusterEndpoint: clusterEndpoint}
	wd := workers.NewConfig(h.tc, h.log)
	_, err = wd.SendEvent(context.TODO(), &model.ConfigureParameters{Resource: model.CrossPlaneResource, Action: model.CrossPlaneProjectDelete}, ci)
	if err != nil {
		return err
	}

	if err = h.deleteManagedClusterCredential(context.TODO(), managedClusterID); err != nil {
		return fmt.Errorf("failed to delete credential for %s, %v", managedClusterID, err)
	}

	return nil
}

func (h *ClusterClaimSyncHandler) syncClusterClaimsWithDB(clusterClaims []model.ClusterClaim) error {

	clusters, err := h.getManagedClusters()
	if err != nil {
		return fmt.Errorf("failed to get managed clusters from DB, %v", err)
	}

	for _, cm := range clusterClaims {
		var isDeleteManagedCluster bool
		var clusterId string
		for _, c := range clusters {
			if c.ClusterName == cm.Metadata.Name {
				isDeleteManagedCluster = true
				clusterId = c.Id
				break
			}
		}

		if isDeleteManagedCluster {
			if err := h.dbStore.DeleteManagedClusterById(clusterId); err != nil {
				return fmt.Errorf("failed to delete managed clusters from DB, %v", err)
			}
		}
	}
	return nil
}
