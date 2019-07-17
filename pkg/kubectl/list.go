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
	cronJobList, err := clientset.BatchV1beta1().CronJobs(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return cronJobList.Items, nil
	} else {
		var items = cronJobList.Items
		for _, v := range cronJobList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	auditSinkList, err := clientset.AuditregistrationV1alpha1().AuditSinks().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return auditSinkList.Items, nil
	} else {
		var items = auditSinkList.Items
		for _, v := range auditSinkList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	mutatingWebhookConfigurationList, err := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return mutatingWebhookConfigurationList.Items, nil
	} else {
		var items = mutatingWebhookConfigurationList.Items
		for _, v := range mutatingWebhookConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	validatingWebhookConfigurationList, err := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return validatingWebhookConfigurationList.Items, nil
	} else {
		var items = validatingWebhookConfigurationList.Items
		for _, v := range validatingWebhookConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	jobList, err := clientset.BatchV1().Jobs(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return jobList.Items, nil
	} else {
		var items = jobList.Items
		for _, v := range jobList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	certificateSigningRequestList, err := clientset.CertificatesV1beta1().CertificateSigningRequests().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return certificateSigningRequestList.Items, nil
	} else {
		var items = certificateSigningRequestList.Items
		for _, v := range certificateSigningRequestList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	roleList, err := clientset.RbacV1().Roles(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return roleList.Items, nil
	} else {
		var items = roleList.Items
		for _, v := range roleList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	clusterRoleBindingList, err := clientset.RbacV1().ClusterRoleBindings().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return clusterRoleBindingList.Items, nil
	} else {
		var items = clusterRoleBindingList.Items
		for _, v := range clusterRoleBindingList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	clusterRoleList, err := clientset.RbacV1().ClusterRoles().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return clusterRoleList.Items, nil
	} else {
		var items = clusterRoleList.Items
		for _, v := range clusterRoleList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	roleBindingList, err := clientset.RbacV1().RoleBindings(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return roleBindingList.Items, nil
	} else {
		var items = roleBindingList.Items
		for _, v := range roleBindingList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	leaseList, err := clientset.CoordinationV1beta1().Leases(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return leaseList.Items, nil
	} else {
		var items = leaseList.Items
		for _, v := range leaseList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	networkPolicyList, err := clientset.NetworkingV1().NetworkPolicies(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return networkPolicyList.Items, nil
	} else {
		var items = networkPolicyList.Items
		for _, v := range networkPolicyList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	replicaSetList, err := clientset.AppsV1().ReplicaSets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return replicaSetList.Items, nil
	} else {
		var items = replicaSetList.Items
		for _, v := range replicaSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	statefulSetList, err := clientset.AppsV1().StatefulSets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return statefulSetList.Items, nil
	} else {
		var items = statefulSetList.Items
		for _, v := range statefulSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	controllerRevisionList, err := clientset.AppsV1().ControllerRevisions(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return controllerRevisionList.Items, nil
	} else {
		var items = controllerRevisionList.Items
		for _, v := range controllerRevisionList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	daemonSetList, err := clientset.AppsV1().DaemonSets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return daemonSetList.Items, nil
	} else {
		var items = daemonSetList.Items
		for _, v := range daemonSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	deploymentList, err := clientset.AppsV1().Deployments(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return deploymentList.Items, nil
	} else {
		var items = deploymentList.Items
		for _, v := range deploymentList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	horizontalPodAutoscalerList, err := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return horizontalPodAutoscalerList.Items, nil
	} else {
		var items = horizontalPodAutoscalerList.Items
		for _, v := range horizontalPodAutoscalerList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	podDisruptionBudgetList, err := clientset.PolicyV1beta1().PodDisruptionBudgets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podDisruptionBudgetList.Items, nil
	} else {
		var items = podDisruptionBudgetList.Items
		for _, v := range podDisruptionBudgetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	podSecurityPolicyList, err := clientset.PolicyV1beta1().PodSecurityPolicies().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podSecurityPolicyList.Items, nil
	} else {
		var items = podSecurityPolicyList.Items
		for _, v := range podSecurityPolicyList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	componentStatusList, err := clientset.CoreV1().ComponentStatuses().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return componentStatusList.Items, nil
	} else {
		var items = componentStatusList.Items
		for _, v := range componentStatusList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	configMapList, err := clientset.CoreV1().ConfigMaps(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return configMapList.Items, nil
	} else {
		var items = configMapList.Items
		for _, v := range configMapList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	eventList, err := clientset.CoreV1().Events(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return eventList.Items, nil
	} else {
		var items = eventList.Items
		for _, v := range eventList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	namespaceList, err := clientset.CoreV1().Namespaces().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return namespaceList.Items, nil
	} else {
		var items = namespaceList.Items
		for _, v := range namespaceList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	persistentVolumeList, err := clientset.CoreV1().PersistentVolumes().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return persistentVolumeList.Items, nil
	} else {
		var items = persistentVolumeList.Items
		for _, v := range persistentVolumeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	podTemplateList, err := clientset.CoreV1().PodTemplates(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podTemplateList.Items, nil
	} else {
		var items = podTemplateList.Items
		for _, v := range podTemplateList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	nodeList, err := clientset.CoreV1().Nodes().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return nodeList.Items, nil
	} else {
		var items = nodeList.Items
		for _, v := range nodeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	persistentVolumeClaimList, err := clientset.CoreV1().PersistentVolumeClaims(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return persistentVolumeClaimList.Items, nil
	} else {
		var items = persistentVolumeClaimList.Items
		for _, v := range persistentVolumeClaimList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	resourceQuotaList, err := clientset.CoreV1().ResourceQuotas(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return resourceQuotaList.Items, nil
	} else {
		var items = resourceQuotaList.Items
		for _, v := range resourceQuotaList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	secretList, err := clientset.CoreV1().Secrets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return secretList.Items, nil
	} else {
		var items = secretList.Items
		for _, v := range secretList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	limitRangeList, err := clientset.CoreV1().LimitRanges(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return limitRangeList.Items, nil
	} else {
		var items = limitRangeList.Items
		for _, v := range limitRangeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	podList, err := clientset.CoreV1().Pods(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podList.Items, nil
	} else {
		var items = podList.Items
		for _, v := range podList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	serviceAccountList, err := clientset.CoreV1().ServiceAccounts(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return serviceAccountList.Items, nil
	} else {
		var items = serviceAccountList.Items
		for _, v := range serviceAccountList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	replicationControllerList, err := clientset.CoreV1().ReplicationControllers(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return replicationControllerList.Items, nil
	} else {
		var items = replicationControllerList.Items
		for _, v := range replicationControllerList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	serviceList, err := clientset.CoreV1().Services(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return serviceList.Items, nil
	} else {
		var items = serviceList.Items
		for _, v := range serviceList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	podPresetList, err := clientset.SettingsV1alpha1().PodPresets(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return podPresetList.Items, nil
	} else {
		var items = podPresetList.Items
		for _, v := range podPresetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	initializerConfigurationList, err := clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return initializerConfigurationList.Items, nil
	} else {
		var items = initializerConfigurationList.Items
		for _, v := range initializerConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	ingressList, err := clientset.ExtensionsV1beta1().Ingresses(i.Namespace).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return ingressList.Items, nil
	} else {
		var items = ingressList.Items
		for _, v := range ingressList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	priorityClassList, err := clientset.SchedulingV1beta1().PriorityClasses().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return priorityClassList.Items, nil
	} else {
		var items = priorityClassList.Items
		for _, v := range priorityClassList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	volumeAttachmentList, err := clientset.StorageV1().VolumeAttachments().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return volumeAttachmentList.Items, nil
	} else {
		var items = volumeAttachmentList.Items
		for _, v := range volumeAttachmentList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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
	storageClassList, err := clientset.StorageV1().StorageClasses().List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return storageClassList.Items, nil
	} else {
		var items = storageClassList.Items
		for _, v := range storageClassList.Items {
			match, err := regexp.Match(i.Name, []byte(v.Name))
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