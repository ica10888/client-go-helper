package kubectl

import (
	admissionregistrationV1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	admissionregistrationV1beta1 "k8s.io/api/admissionregistration/v1beta1"
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	auditregistrationV1alpha1 "k8s.io/api/auditregistration/v1alpha1"
	autoscalingV2beta2 "k8s.io/api/autoscaling/v2beta2"
	batchV1 "k8s.io/api/batch/v1"
	batchV1beta1 "k8s.io/api/batch/v1beta1"
	certificatesV1beta1 "k8s.io/api/certificates/v1beta1"
	coordinationV1beta1 "k8s.io/api/coordination/v1beta1"
	coreV1 "k8s.io/api/core/v1"
	extensionsV1beta1 "k8s.io/api/extensions/v1beta1"
	networkingV1 "k8s.io/api/networking/v1"
	policyV1beta1 "k8s.io/api/policy/v1beta1"
	rbacV1 "k8s.io/api/rbac/v1"
	schedulingV1beta1 "k8s.io/api/scheduling/v1beta1"
	settingsV1alpha1 "k8s.io/api/settings/v1alpha1"
	storageV1 "k8s.io/api/storage/v1"
)

func (i *CronJob) List(opts *v1.ListOptions) ([]batchV1beta1.CronJob, error) {
	cronJobList, e := clientset.BatchV1beta1().CronJobs(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return cronJobList.Items,nil
}


func (i *AuditSink) List(opts *v1.ListOptions) ([]auditregistrationV1alpha1.AuditSink, error) {
	auditSinkList, e := clientset.AuditregistrationV1alpha1().AuditSinks().List(*opts)
	if e != nil {
		return nil,e
	}
	return auditSinkList.Items,nil
}


func (i *ValidatingWebhookConfiguration) List(opts *v1.ListOptions) ([]admissionregistrationV1beta1.ValidatingWebhookConfiguration, error) {
	validatingWebhookConfigurationList, e := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(*opts)
	if e != nil {
		return nil,e
	}
	return validatingWebhookConfigurationList.Items,nil
}


func (i *MutatingWebhookConfiguration) List(opts *v1.ListOptions) ([]admissionregistrationV1beta1.MutatingWebhookConfiguration, error) {
	mutatingWebhookConfigurationList, e := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().List(*opts)
	if e != nil {
		return nil,e
	}
	return mutatingWebhookConfigurationList.Items,nil
}


func (i *Job) List(opts *v1.ListOptions) ([]batchV1.Job, error) {
	jobList, e := clientset.BatchV1().Jobs(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return jobList.Items,nil
}


func (i *CertificateSigningRequest) List(opts *v1.ListOptions) ([]certificatesV1beta1.CertificateSigningRequest, error) {
	certificateSigningRequestList, e := clientset.CertificatesV1beta1().CertificateSigningRequests().List(*opts)
	if e != nil {
		return nil,e
	}
	return certificateSigningRequestList.Items,nil
}


func (i *RoleBinding) List(opts *v1.ListOptions) ([]rbacV1.RoleBinding, error) {
	roleBindingList, e := clientset.RbacV1().RoleBindings(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return roleBindingList.Items,nil
}


func (i *Role) List(opts *v1.ListOptions) ([]rbacV1.Role, error) {
	roleList, e := clientset.RbacV1().Roles(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return roleList.Items,nil
}


func (i *ClusterRoleBinding) List(opts *v1.ListOptions) ([]rbacV1.ClusterRoleBinding, error) {
	clusterRoleBindingList, e := clientset.RbacV1().ClusterRoleBindings().List(*opts)
	if e != nil {
		return nil,e
	}
	return clusterRoleBindingList.Items,nil
}


func (i *ClusterRole) List(opts *v1.ListOptions) ([]rbacV1.ClusterRole, error) {
	clusterRoleList, e := clientset.RbacV1().ClusterRoles().List(*opts)
	if e != nil {
		return nil,e
	}
	return clusterRoleList.Items,nil
}


func (i *Lease) List(opts *v1.ListOptions) ([]coordinationV1beta1.Lease, error) {
	leaseList, e := clientset.CoordinationV1beta1().Leases(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return leaseList.Items,nil
}


func (i *NetworkPolicy) List(opts *v1.ListOptions) ([]networkingV1.NetworkPolicy, error) {
	networkPolicyList, e := clientset.NetworkingV1().NetworkPolicies(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return networkPolicyList.Items,nil
}


func (i *StatefulSet) List(opts *v1.ListOptions) ([]appsV1.StatefulSet, error) {
	statefulSetList, e := clientset.AppsV1().StatefulSets(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return statefulSetList.Items,nil
}


func (i *ControllerRevision) List(opts *v1.ListOptions) ([]appsV1.ControllerRevision, error) {
	controllerRevisionList, e := clientset.AppsV1().ControllerRevisions(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return controllerRevisionList.Items,nil
}


func (i *DaemonSet) List(opts *v1.ListOptions) ([]appsV1.DaemonSet, error) {
	daemonSetList, e := clientset.AppsV1().DaemonSets(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return daemonSetList.Items,nil
}


func (i *Deployment) List(opts *v1.ListOptions) ([]appsV1.Deployment, error) {
	deploymentList, e := clientset.AppsV1().Deployments(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return deploymentList.Items,nil
}


func (i *ReplicaSet) List(opts *v1.ListOptions) ([]appsV1.ReplicaSet, error) {
	replicaSetList, e := clientset.AppsV1().ReplicaSets(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return replicaSetList.Items,nil
}


func (i *HorizontalPodAutoscaler) List(opts *v1.ListOptions) ([]autoscalingV2beta2.HorizontalPodAutoscaler, error) {
	horizontalPodAutoscalerList, e := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return horizontalPodAutoscalerList.Items,nil
}



func (i *PodDisruptionBudget) List(opts *v1.ListOptions) ([]policyV1beta1.PodDisruptionBudget, error) {
	podDisruptionBudgetList, e := clientset.PolicyV1beta1().PodDisruptionBudgets(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return podDisruptionBudgetList.Items,nil
}


func (i *PodSecurityPolicy) List(opts *v1.ListOptions) ([]policyV1beta1.PodSecurityPolicy, error) {
	podSecurityPolicyList, e := clientset.PolicyV1beta1().PodSecurityPolicies().List(*opts)
	if e != nil {
		return nil,e
	}
	return podSecurityPolicyList.Items,nil
}




func (i *Namespace) List(opts *v1.ListOptions) ([]coreV1.Namespace, error) {
	namespaceList, e := clientset.CoreV1().Namespaces().List(*opts)
	if e != nil {
		return nil,e
	}
	return namespaceList.Items,nil
}


func (i *Node) List(opts *v1.ListOptions) ([]coreV1.Node, error) {
	nodeList, e := clientset.CoreV1().Nodes().List(*opts)
	if e != nil {
		return nil,e
	}
	return nodeList.Items,nil
}


func (i *ReplicationController) List(opts *v1.ListOptions) ([]coreV1.ReplicationController, error) {
	replicationControllerList, e := clientset.CoreV1().ReplicationControllers(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return replicationControllerList.Items,nil
}


func (i *ConfigMap) List(opts *v1.ListOptions) ([]coreV1.ConfigMap, error) {
	configMapList, e := clientset.CoreV1().ConfigMaps(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return configMapList.Items,nil
}


func (i *Event) List(opts *v1.ListOptions) ([]coreV1.Event, error) {
	eventList, e := clientset.CoreV1().Events(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return eventList.Items,nil
}


func (i *Pod) List(opts *v1.ListOptions) ([]coreV1.Pod, error) {
	podList, e := clientset.CoreV1().Pods(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return podList.Items,nil
}


func (i *Secret) List(opts *v1.ListOptions) ([]coreV1.Secret, error) {
	secretList, e := clientset.CoreV1().Secrets(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return secretList.Items,nil
}


func (i *ComponentStatus) List(opts *v1.ListOptions) ([]coreV1.ComponentStatus, error) {
	componentStatusList, e := clientset.CoreV1().ComponentStatuses().List(*opts)
	if e != nil {
		return nil,e
	}
	return componentStatusList.Items,nil
}


func (i *PersistentVolumeClaim) List(opts *v1.ListOptions) ([]coreV1.PersistentVolumeClaim, error) {
	persistentVolumeClaimList, e := clientset.CoreV1().PersistentVolumeClaims(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return persistentVolumeClaimList.Items,nil
}


func (i *PodTemplate) List(opts *v1.ListOptions) ([]coreV1.PodTemplate, error) {
	podTemplateList, e := clientset.CoreV1().PodTemplates(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return podTemplateList.Items,nil
}


func (i *ResourceQuota) List(opts *v1.ListOptions) ([]coreV1.ResourceQuota, error) {
	resourceQuotaList, e := clientset.CoreV1().ResourceQuotas(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return resourceQuotaList.Items,nil
}


func (i *Service) List(opts *v1.ListOptions) ([]coreV1.Service, error) {
	serviceList, e := clientset.CoreV1().Services(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return serviceList.Items,nil
}


func (i *LimitRange) List(opts *v1.ListOptions) ([]coreV1.LimitRange, error) {
	limitRangeList, e := clientset.CoreV1().LimitRanges(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return limitRangeList.Items,nil
}


func (i *PersistentVolume) List(opts *v1.ListOptions) ([]coreV1.PersistentVolume, error) {
	persistentVolumeList, e := clientset.CoreV1().PersistentVolumes().List(*opts)
	if e != nil {
		return nil,e
	}
	return persistentVolumeList.Items,nil
}


func (i *ServiceAccount) List(opts *v1.ListOptions) ([]coreV1.ServiceAccount, error) {
	serviceAccountList, e := clientset.CoreV1().ServiceAccounts(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return serviceAccountList.Items,nil
}


func (i *PodPreset) List(opts *v1.ListOptions) ([]settingsV1alpha1.PodPreset, error) {
	podPresetList, e := clientset.SettingsV1alpha1().PodPresets(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return podPresetList.Items,nil
}


func (i *InitializerConfiguration) List(opts *v1.ListOptions) ([]admissionregistrationV1alpha1.InitializerConfiguration, error) {
	initializerConfigurationList, e := clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().List(*opts)
	if e != nil {
		return nil,e
	}
	return initializerConfigurationList.Items,nil
}





func (i *Ingress) List(opts *v1.ListOptions) ([]extensionsV1beta1.Ingress, error) {
	ingressList, e := clientset.ExtensionsV1beta1().Ingresses(i.Namespace).List(*opts)
	if e != nil {
		return nil,e
	}
	return ingressList.Items,nil
}



func (i *PriorityClass) List(opts *v1.ListOptions) ([]schedulingV1beta1.PriorityClass, error) {
	priorityClassList, e := clientset.SchedulingV1beta1().PriorityClasses().List(*opts)
	if e != nil {
		return nil,e
	}
	return priorityClassList.Items,nil
}


func (i *StorageClass) List(opts *v1.ListOptions) ([]storageV1.StorageClass, error) {
	storageClassList, e := clientset.StorageV1().StorageClasses().List(*opts)
	if e != nil {
		return nil,e
	}
	return storageClassList.Items,nil
}


func (i *VolumeAttachment) List(opts *v1.ListOptions) ([]storageV1.VolumeAttachment, error) {
	volumeAttachmentList, e := clientset.StorageV1().VolumeAttachments().List(*opts)
	if e != nil {
		return nil,e
	}
	return volumeAttachmentList.Items,nil
}