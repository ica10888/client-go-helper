package kubectl

import (
	admissionregistrationV1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	admissionregistrationV1beta1 "k8s.io/api/admissionregistration/v1beta1"
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"regexp"

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

func (i *CronJob) GetAll(opts *v1.ListOptions) ([]batchV1beta1.CronJob, error) {
	var clientset,_  = InitClient()
	cronJobList, err := clientset.BatchV1beta1().CronJobs(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return cronJobList.Items, nil
	} else {
		var items []batchV1beta1.CronJob
		for _, v := range cronJobList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *AuditSink) GetAll(opts *v1.ListOptions) ([]auditregistrationV1alpha1.AuditSink, error) {
	var clientset,_  = InitClient()
	auditSinkList, err := clientset.AuditregistrationV1alpha1().AuditSinks().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return auditSinkList.Items, nil
	} else {
		var items []auditregistrationV1alpha1.AuditSink
		for _, v := range auditSinkList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *MutatingWebhookConfiguration) GetAll(opts *v1.ListOptions) ([]admissionregistrationV1beta1.MutatingWebhookConfiguration, error) {
	var clientset,_  = InitClient()
	mutatingWebhookConfigurationList, err := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return mutatingWebhookConfigurationList.Items, nil
	} else {
		var items []admissionregistrationV1beta1.MutatingWebhookConfiguration
		for _, v := range mutatingWebhookConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *ValidatingWebhookConfiguration) GetAll(opts *v1.ListOptions) ([]admissionregistrationV1beta1.ValidatingWebhookConfiguration, error) {
	var clientset,_  = InitClient()
	validatingWebhookConfigurationList, err := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return validatingWebhookConfigurationList.Items, nil
	} else {
		var items []admissionregistrationV1beta1.ValidatingWebhookConfiguration
		for _, v := range validatingWebhookConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Job) GetAll(opts *v1.ListOptions) ([]batchV1.Job, error) {
	var clientset,_  = InitClient()
	jobList, err := clientset.BatchV1().Jobs(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return jobList.Items, nil
	} else {
		var items []batchV1.Job
		for _, v := range jobList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *CertificateSigningRequest) GetAll(opts *v1.ListOptions) ([]certificatesV1beta1.CertificateSigningRequest, error) {
	var clientset,_  = InitClient()
	certificateSigningRequestList, err := clientset.CertificatesV1beta1().CertificateSigningRequests().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return certificateSigningRequestList.Items, nil
	} else {
		var items []certificatesV1beta1.CertificateSigningRequest
		for _, v := range certificateSigningRequestList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *ClusterRoleBinding) GetAll(opts *v1.ListOptions) ([]rbacV1.ClusterRoleBinding, error) {
	var clientset,_  = InitClient()
	clusterRoleBindingList, err := clientset.RbacV1().ClusterRoleBindings().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return clusterRoleBindingList.Items, nil
	} else {
		var items []rbacV1.ClusterRoleBinding
		for _, v := range clusterRoleBindingList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *ClusterRole) GetAll(opts *v1.ListOptions) ([]rbacV1.ClusterRole, error) {
	var clientset,_  = InitClient()
	clusterRoleList, err := clientset.RbacV1().ClusterRoles().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return clusterRoleList.Items, nil
	} else {
		var items []rbacV1.ClusterRole
		for _, v := range clusterRoleList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *RoleBinding) GetAll(opts *v1.ListOptions) ([]rbacV1.RoleBinding, error) {
	var clientset,_  = InitClient()
	roleBindingList, err := clientset.RbacV1().RoleBindings(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return roleBindingList.Items, nil
	} else {
		var items []rbacV1.RoleBinding
		for _, v := range roleBindingList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Role) GetAll(opts *v1.ListOptions) ([]rbacV1.Role, error) {
	var clientset,_  = InitClient()
	roleList, err := clientset.RbacV1().Roles(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return roleList.Items, nil
	} else {
		var items []rbacV1.Role
		for _, v := range roleList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Lease) GetAll(opts *v1.ListOptions) ([]coordinationV1beta1.Lease, error) {
	var clientset,_  = InitClient()
	leaseList, err := clientset.CoordinationV1beta1().Leases(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return leaseList.Items, nil
	} else {
		var items []coordinationV1beta1.Lease
		for _, v := range leaseList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *NetworkPolicy) GetAll(opts *v1.ListOptions) ([]networkingV1.NetworkPolicy, error) {
	var clientset,_  = InitClient()
	networkPolicyList, err := clientset.NetworkingV1().NetworkPolicies(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return networkPolicyList.Items, nil
	} else {
		var items []networkingV1.NetworkPolicy
		for _, v := range networkPolicyList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *ControllerRevision) GetAll(opts *v1.ListOptions) ([]appsV1.ControllerRevision, error) {
	var clientset,_  = InitClient()
	controllerRevisionList, err := clientset.AppsV1().ControllerRevisions(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return controllerRevisionList.Items, nil
	} else {
		var items []appsV1.ControllerRevision
		for _, v := range controllerRevisionList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *DaemonSet) GetAll(opts *v1.ListOptions) ([]appsV1.DaemonSet, error) {
	var clientset,_  = InitClient()
	daemonSetList, err := clientset.AppsV1().DaemonSets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return daemonSetList.Items, nil
	} else {
		var items []appsV1.DaemonSet
		for _, v := range daemonSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Deployment) GetAll(opts *v1.ListOptions) ([]appsV1.Deployment, error) {
	var clientset,_  = InitClient()
	deploymentList, err := clientset.AppsV1().Deployments(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return deploymentList.Items, nil
	} else {
		var items []appsV1.Deployment
		for _, v := range deploymentList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *ReplicaSet) GetAll(opts *v1.ListOptions) ([]appsV1.ReplicaSet, error) {
	var clientset,_  = InitClient()
	replicaSetList, err := clientset.AppsV1().ReplicaSets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return replicaSetList.Items, nil
	} else {
		var items []appsV1.ReplicaSet
		for _, v := range replicaSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *StatefulSet) GetAll(opts *v1.ListOptions) ([]appsV1.StatefulSet, error) {
	var clientset,_  = InitClient()
	statefulSetList, err := clientset.AppsV1().StatefulSets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return statefulSetList.Items, nil
	} else {
		var items []appsV1.StatefulSet
		for _, v := range statefulSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *HorizontalPodAutoscaler) GetAll(opts *v1.ListOptions) ([]autoscalingV2beta2.HorizontalPodAutoscaler, error) {
	var clientset,_  = InitClient()
	horizontalPodAutoscalerList, err := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return horizontalPodAutoscalerList.Items, nil
	} else {
		var items []autoscalingV2beta2.HorizontalPodAutoscaler
		for _, v := range horizontalPodAutoscalerList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}





func (i *PodDisruptionBudget) GetAll(opts *v1.ListOptions) ([]policyV1beta1.PodDisruptionBudget, error) {
	var clientset,_  = InitClient()
	podDisruptionBudgetList, err := clientset.PolicyV1beta1().PodDisruptionBudgets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podDisruptionBudgetList.Items, nil
	} else {
		var items []policyV1beta1.PodDisruptionBudget
		for _, v := range podDisruptionBudgetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *PodSecurityPolicy) GetAll(opts *v1.ListOptions) ([]policyV1beta1.PodSecurityPolicy, error) {
	var clientset,_  = InitClient()
	podSecurityPolicyList, err := clientset.PolicyV1beta1().PodSecurityPolicies().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podSecurityPolicyList.Items, nil
	} else {
		var items []policyV1beta1.PodSecurityPolicy
		for _, v := range podSecurityPolicyList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}




func (i *ServiceAccount) GetAll(opts *v1.ListOptions) ([]coreV1.ServiceAccount, error) {
	var clientset,_  = InitClient()
	serviceAccountList, err := clientset.CoreV1().ServiceAccounts(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return serviceAccountList.Items, nil
	} else {
		var items []coreV1.ServiceAccount
		for _, v := range serviceAccountList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *ComponentStatus) GetAll(opts *v1.ListOptions) ([]coreV1.ComponentStatus, error) {
	var clientset,_  = InitClient()
	componentStatusList, err := clientset.CoreV1().ComponentStatuses().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return componentStatusList.Items, nil
	} else {
		var items []coreV1.ComponentStatus
		for _, v := range componentStatusList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *PersistentVolumeClaim) GetAll(opts *v1.ListOptions) ([]coreV1.PersistentVolumeClaim, error) {
	var clientset,_  = InitClient()
	persistentVolumeClaimList, err := clientset.CoreV1().PersistentVolumeClaims(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return persistentVolumeClaimList.Items, nil
	} else {
		var items []coreV1.PersistentVolumeClaim
		for _, v := range persistentVolumeClaimList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *ResourceQuota) GetAll(opts *v1.ListOptions) ([]coreV1.ResourceQuota, error) {
	var clientset,_  = InitClient()
	resourceQuotaList, err := clientset.CoreV1().ResourceQuotas(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return resourceQuotaList.Items, nil
	} else {
		var items []coreV1.ResourceQuota
		for _, v := range resourceQuotaList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}





func (i *ReplicationController) GetAll(opts *v1.ListOptions) ([]coreV1.ReplicationController, error) {
	var clientset,_  = InitClient()
	replicationControllerList, err := clientset.CoreV1().ReplicationControllers(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return replicationControllerList.Items, nil
	} else {
		var items []coreV1.ReplicationController
		for _, v := range replicationControllerList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Secret) GetAll(opts *v1.ListOptions) ([]coreV1.Secret, error) {
	var clientset,_  = InitClient()
	secretList, err := clientset.CoreV1().Secrets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return secretList.Items, nil
	} else {
		var items []coreV1.Secret
		for _, v := range secretList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Namespace) GetAll(opts *v1.ListOptions) ([]coreV1.Namespace, error) {
	var clientset,_  = InitClient()
	namespaceList, err := clientset.CoreV1().Namespaces().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return namespaceList.Items, nil
	} else {
		var items []coreV1.Namespace
		for _, v := range namespaceList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *PersistentVolume) GetAll(opts *v1.ListOptions) ([]coreV1.PersistentVolume, error) {
	var clientset,_  = InitClient()
	persistentVolumeList, err := clientset.CoreV1().PersistentVolumes().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return persistentVolumeList.Items, nil
	} else {
		var items []coreV1.PersistentVolume
		for _, v := range persistentVolumeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Pod) GetAll(opts *v1.ListOptions) ([]coreV1.Pod, error) {
	var clientset,_  = InitClient()
	podList, err := clientset.CoreV1().Pods(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podList.Items, nil
	} else {
		var items []coreV1.Pod
		for _, v := range podList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Node) GetAll(opts *v1.ListOptions) ([]coreV1.Node, error) {
	var clientset,_  = InitClient()
	nodeList, err := clientset.CoreV1().Nodes().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return nodeList.Items, nil
	} else {
		var items []coreV1.Node
		for _, v := range nodeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *PodTemplate) GetAll(opts *v1.ListOptions) ([]coreV1.PodTemplate, error) {
	var clientset,_  = InitClient()
	podTemplateList, err := clientset.CoreV1().PodTemplates(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podTemplateList.Items, nil
	} else {
		var items []coreV1.PodTemplate
		for _, v := range podTemplateList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Service) GetAll(opts *v1.ListOptions) ([]coreV1.Service, error) {
	var clientset,_  = InitClient()
	serviceList, err := clientset.CoreV1().Services(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return serviceList.Items, nil
	} else {
		var items []coreV1.Service
		for _, v := range serviceList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *ConfigMap) GetAll(opts *v1.ListOptions) ([]coreV1.ConfigMap, error) {
	var clientset,_  = InitClient()
	configMapList, err := clientset.CoreV1().ConfigMaps(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return configMapList.Items, nil
	} else {
		var items []coreV1.ConfigMap
		for _, v := range configMapList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *Event) GetAll(opts *v1.ListOptions) ([]coreV1.Event, error) {
	var clientset,_  = InitClient()
	eventList, err := clientset.CoreV1().Events(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return eventList.Items, nil
	} else {
		var items []coreV1.Event
		for _, v := range eventList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *LimitRange) GetAll(opts *v1.ListOptions) ([]coreV1.LimitRange, error) {
	var clientset,_  = InitClient()
	limitRangeList, err := clientset.CoreV1().LimitRanges(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return limitRangeList.Items, nil
	} else {
		var items []coreV1.LimitRange
		for _, v := range limitRangeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *PodPreset) GetAll(opts *v1.ListOptions) ([]settingsV1alpha1.PodPreset, error) {
	var clientset,_  = InitClient()
	podPresetList, err := clientset.SettingsV1alpha1().PodPresets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podPresetList.Items, nil
	} else {
		var items []settingsV1alpha1.PodPreset
		for _, v := range podPresetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *InitializerConfiguration) GetAll(opts *v1.ListOptions) ([]admissionregistrationV1alpha1.InitializerConfiguration, error) {
	var clientset,_  = InitClient()
	initializerConfigurationList, err := clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return initializerConfigurationList.Items, nil
	} else {
		var items []admissionregistrationV1alpha1.InitializerConfiguration
		for _, v := range initializerConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}





func (i *Ingress) GetAll(opts *v1.ListOptions) ([]extensionsV1beta1.Ingress, error) {
	var clientset,_  = InitClient()
	ingressList, err := clientset.ExtensionsV1beta1().Ingresses(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return ingressList.Items, nil
	} else {
		var items []extensionsV1beta1.Ingress
		for _, v := range ingressList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}




func (i *PriorityClass) GetAll(opts *v1.ListOptions) ([]schedulingV1beta1.PriorityClass, error) {
	var clientset,_  = InitClient()
	priorityClassList, err := clientset.SchedulingV1beta1().PriorityClasses().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return priorityClassList.Items, nil
	} else {
		var items []schedulingV1beta1.PriorityClass
		for _, v := range priorityClassList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *StorageClass) GetAll(opts *v1.ListOptions) ([]storageV1.StorageClass, error) {
	var clientset,_  = InitClient()
	storageClassList, err := clientset.StorageV1().StorageClasses().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return storageClassList.Items, nil
	} else {
		var items []storageV1.StorageClass
		for _, v := range storageClassList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}


func (i *VolumeAttachment) GetAll(opts *v1.ListOptions) ([]storageV1.VolumeAttachment, error) {
	var clientset,_  = InitClient()
	volumeAttachmentList, err := clientset.StorageV1().VolumeAttachments().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return volumeAttachmentList.Items, nil
	} else {
		var items []storageV1.VolumeAttachment
		for _, v := range volumeAttachmentList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}