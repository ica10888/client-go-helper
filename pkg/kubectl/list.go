package kubectl

import (
	admissionregistrationV1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	admissionregistrationV1beta1 "k8s.io/api/admissionregistration/v1beta1"
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"regexp"

	"client-go-helper/pkg/kubectl/client"
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

func (i *cronJob) GetAll(opts *v1.ListOptions) ([]batchV1beta1.CronJob, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	cronJobList, err := clientset.BatchV1beta1().CronJobs(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return cronJobList.Items, nil
	} else {
		var items []batchV1beta1.CronJob
		for _, v := range cronJobList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *auditSink) GetAll(opts *v1.ListOptions) ([]auditregistrationV1alpha1.AuditSink, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	auditSinkList, err := clientset.AuditregistrationV1alpha1().AuditSinks().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return auditSinkList.Items, nil
	} else {
		var items []auditregistrationV1alpha1.AuditSink
		for _, v := range auditSinkList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *mutatingWebhookConfiguration) GetAll(opts *v1.ListOptions) ([]admissionregistrationV1beta1.MutatingWebhookConfiguration, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	mutatingWebhookConfigurationList, err := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return mutatingWebhookConfigurationList.Items, nil
	} else {
		var items []admissionregistrationV1beta1.MutatingWebhookConfiguration
		for _, v := range mutatingWebhookConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *validatingWebhookConfiguration) GetAll(opts *v1.ListOptions) ([]admissionregistrationV1beta1.ValidatingWebhookConfiguration, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	validatingWebhookConfigurationList, err := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return validatingWebhookConfigurationList.Items, nil
	} else {
		var items []admissionregistrationV1beta1.ValidatingWebhookConfiguration
		for _, v := range validatingWebhookConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *job) GetAll(opts *v1.ListOptions) ([]batchV1.Job, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	jobList, err := clientset.BatchV1().Jobs(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return jobList.Items, nil
	} else {
		var items []batchV1.Job
		for _, v := range jobList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *certificateSigningRequest) GetAll(opts *v1.ListOptions) ([]certificatesV1beta1.CertificateSigningRequest, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	certificateSigningRequestList, err := clientset.CertificatesV1beta1().CertificateSigningRequests().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return certificateSigningRequestList.Items, nil
	} else {
		var items []certificatesV1beta1.CertificateSigningRequest
		for _, v := range certificateSigningRequestList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *clusterRoleBinding) GetAll(opts *v1.ListOptions) ([]rbacV1.ClusterRoleBinding, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	clusterRoleBindingList, err := clientset.RbacV1().ClusterRoleBindings().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return clusterRoleBindingList.Items, nil
	} else {
		var items []rbacV1.ClusterRoleBinding
		for _, v := range clusterRoleBindingList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *clusterRole) GetAll(opts *v1.ListOptions) ([]rbacV1.ClusterRole, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	clusterRoleList, err := clientset.RbacV1().ClusterRoles().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return clusterRoleList.Items, nil
	} else {
		var items []rbacV1.ClusterRole
		for _, v := range clusterRoleList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *roleBinding) GetAll(opts *v1.ListOptions) ([]rbacV1.RoleBinding, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	roleBindingList, err := clientset.RbacV1().RoleBindings(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return roleBindingList.Items, nil
	} else {
		var items []rbacV1.RoleBinding
		for _, v := range roleBindingList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *role) GetAll(opts *v1.ListOptions) ([]rbacV1.Role, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	roleList, err := clientset.RbacV1().Roles(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return roleList.Items, nil
	} else {
		var items []rbacV1.Role
		for _, v := range roleList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *lease) GetAll(opts *v1.ListOptions) ([]coordinationV1beta1.Lease, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	leaseList, err := clientset.CoordinationV1beta1().Leases(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return leaseList.Items, nil
	} else {
		var items []coordinationV1beta1.Lease
		for _, v := range leaseList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *networkPolicy) GetAll(opts *v1.ListOptions) ([]networkingV1.NetworkPolicy, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	networkPolicyList, err := clientset.NetworkingV1().NetworkPolicies(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return networkPolicyList.Items, nil
	} else {
		var items []networkingV1.NetworkPolicy
		for _, v := range networkPolicyList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *controllerRevision) GetAll(opts *v1.ListOptions) ([]appsV1.ControllerRevision, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	controllerRevisionList, err := clientset.AppsV1().ControllerRevisions(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return controllerRevisionList.Items, nil
	} else {
		var items []appsV1.ControllerRevision
		for _, v := range controllerRevisionList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *daemonSet) GetAll(opts *v1.ListOptions) ([]appsV1.DaemonSet, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	daemonSetList, err := clientset.AppsV1().DaemonSets(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return daemonSetList.Items, nil
	} else {
		var items []appsV1.DaemonSet
		for _, v := range daemonSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *deployment) GetAll(opts *v1.ListOptions) ([]appsV1.Deployment, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	deploymentList, err := clientset.AppsV1().Deployments(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return deploymentList.Items, nil
	} else {
		var items []appsV1.Deployment
		for _, v := range deploymentList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *replicaSet) GetAll(opts *v1.ListOptions) ([]appsV1.ReplicaSet, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	replicaSetList, err := clientset.AppsV1().ReplicaSets(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return replicaSetList.Items, nil
	} else {
		var items []appsV1.ReplicaSet
		for _, v := range replicaSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *statefulSet) GetAll(opts *v1.ListOptions) ([]appsV1.StatefulSet, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	statefulSetList, err := clientset.AppsV1().StatefulSets(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return statefulSetList.Items, nil
	} else {
		var items []appsV1.StatefulSet
		for _, v := range statefulSetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *horizontalPodAutoscaler) GetAll(opts *v1.ListOptions) ([]autoscalingV2beta2.HorizontalPodAutoscaler, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	horizontalPodAutoscalerList, err := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return horizontalPodAutoscalerList.Items, nil
	} else {
		var items []autoscalingV2beta2.HorizontalPodAutoscaler
		for _, v := range horizontalPodAutoscalerList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *podDisruptionBudget) GetAll(opts *v1.ListOptions) ([]policyV1beta1.PodDisruptionBudget, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	podDisruptionBudgetList, err := clientset.PolicyV1beta1().PodDisruptionBudgets(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return podDisruptionBudgetList.Items, nil
	} else {
		var items []policyV1beta1.PodDisruptionBudget
		for _, v := range podDisruptionBudgetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *podSecurityPolicy) GetAll(opts *v1.ListOptions) ([]policyV1beta1.PodSecurityPolicy, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	podSecurityPolicyList, err := clientset.PolicyV1beta1().PodSecurityPolicies().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return podSecurityPolicyList.Items, nil
	} else {
		var items []policyV1beta1.PodSecurityPolicy
		for _, v := range podSecurityPolicyList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *serviceAccount) GetAll(opts *v1.ListOptions) ([]coreV1.ServiceAccount, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	serviceAccountList, err := clientset.CoreV1().ServiceAccounts(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return serviceAccountList.Items, nil
	} else {
		var items []coreV1.ServiceAccount
		for _, v := range serviceAccountList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *componentStatus) GetAll(opts *v1.ListOptions) ([]coreV1.ComponentStatus, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	componentStatusList, err := clientset.CoreV1().ComponentStatuses().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return componentStatusList.Items, nil
	} else {
		var items []coreV1.ComponentStatus
		for _, v := range componentStatusList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *persistentVolumeClaim) GetAll(opts *v1.ListOptions) ([]coreV1.PersistentVolumeClaim, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	persistentVolumeClaimList, err := clientset.CoreV1().PersistentVolumeClaims(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return persistentVolumeClaimList.Items, nil
	} else {
		var items []coreV1.PersistentVolumeClaim
		for _, v := range persistentVolumeClaimList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *resourceQuota) GetAll(opts *v1.ListOptions) ([]coreV1.ResourceQuota, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	resourceQuotaList, err := clientset.CoreV1().ResourceQuotas(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return resourceQuotaList.Items, nil
	} else {
		var items []coreV1.ResourceQuota
		for _, v := range resourceQuotaList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *replicationController) GetAll(opts *v1.ListOptions) ([]coreV1.ReplicationController, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	replicationControllerList, err := clientset.CoreV1().ReplicationControllers(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return replicationControllerList.Items, nil
	} else {
		var items []coreV1.ReplicationController
		for _, v := range replicationControllerList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *secret) GetAll(opts *v1.ListOptions) ([]coreV1.Secret, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	secretList, err := clientset.CoreV1().Secrets(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return secretList.Items, nil
	} else {
		var items []coreV1.Secret
		for _, v := range secretList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *namespace) GetAll(opts *v1.ListOptions) ([]coreV1.Namespace, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	namespaceList, err := clientset.CoreV1().Namespaces().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return namespaceList.Items, nil
	} else {
		var items []coreV1.Namespace
		for _, v := range namespaceList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *persistentVolume) GetAll(opts *v1.ListOptions) ([]coreV1.PersistentVolume, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	persistentVolumeList, err := clientset.CoreV1().PersistentVolumes().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return persistentVolumeList.Items, nil
	} else {
		var items []coreV1.PersistentVolume
		for _, v := range persistentVolumeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *pod) GetAll(opts *v1.ListOptions) ([]coreV1.Pod, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	podList, err := clientset.CoreV1().Pods(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return podList.Items, nil
	} else {
		var items []coreV1.Pod
		for _, v := range podList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *node) GetAll(opts *v1.ListOptions) ([]coreV1.Node, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	nodeList, err := clientset.CoreV1().Nodes().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return nodeList.Items, nil
	} else {
		var items []coreV1.Node
		for _, v := range nodeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *podTemplate) GetAll(opts *v1.ListOptions) ([]coreV1.PodTemplate, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	podTemplateList, err := clientset.CoreV1().PodTemplates(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return podTemplateList.Items, nil
	} else {
		var items []coreV1.PodTemplate
		for _, v := range podTemplateList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *service) GetAll(opts *v1.ListOptions) ([]coreV1.Service, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	serviceList, err := clientset.CoreV1().Services(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return serviceList.Items, nil
	} else {
		var items []coreV1.Service
		for _, v := range serviceList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *configMap) GetAll(opts *v1.ListOptions) ([]coreV1.ConfigMap, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	configMapList, err := clientset.CoreV1().ConfigMaps(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return configMapList.Items, nil
	} else {
		var items []coreV1.ConfigMap
		for _, v := range configMapList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *event) GetAll(opts *v1.ListOptions) ([]coreV1.Event, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	eventList, err := clientset.CoreV1().Events(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return eventList.Items, nil
	} else {
		var items []coreV1.Event
		for _, v := range eventList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *limitRange) GetAll(opts *v1.ListOptions) ([]coreV1.LimitRange, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	limitRangeList, err := clientset.CoreV1().LimitRanges(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return limitRangeList.Items, nil
	} else {
		var items []coreV1.LimitRange
		for _, v := range limitRangeList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *podPreset) GetAll(opts *v1.ListOptions) ([]settingsV1alpha1.PodPreset, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	podPresetList, err := clientset.SettingsV1alpha1().PodPresets(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return podPresetList.Items, nil
	} else {
		var items []settingsV1alpha1.PodPreset
		for _, v := range podPresetList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *initializerConfiguration) GetAll(opts *v1.ListOptions) ([]admissionregistrationV1alpha1.InitializerConfiguration, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	initializerConfigurationList, err := clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return initializerConfigurationList.Items, nil
	} else {
		var items []admissionregistrationV1alpha1.InitializerConfiguration
		for _, v := range initializerConfigurationList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *ingress) GetAll(opts *v1.ListOptions) ([]extensionsV1beta1.Ingress, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	ingressList, err := clientset.ExtensionsV1beta1().Ingresses(i.Namespace).List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return ingressList.Items, nil
	} else {
		var items []extensionsV1beta1.Ingress
		for _, v := range ingressList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *priorityClass) GetAll(opts *v1.ListOptions) ([]schedulingV1beta1.PriorityClass, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	priorityClassList, err := clientset.SchedulingV1beta1().PriorityClasses().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return priorityClassList.Items, nil
	} else {
		var items []schedulingV1beta1.PriorityClass
		for _, v := range priorityClassList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *storageClass) GetAll(opts *v1.ListOptions) ([]storageV1.StorageClass, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	storageClassList, err := clientset.StorageV1().StorageClasses().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return storageClassList.Items, nil
	} else {
		var items []storageV1.StorageClass
		for _, v := range storageClassList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}

func (i *volumeAttachment) GetAll(opts *v1.ListOptions) ([]storageV1.VolumeAttachment, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return nil, err
	}
	volumeAttachmentList, err := clientset.StorageV1().VolumeAttachments().List(*opts)
	if err != nil {
		return nil, err
	}
	if i.Name == "" {
		return volumeAttachmentList.Items, nil
	} else {
		var items []storageV1.VolumeAttachment
		for _, v := range volumeAttachmentList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil, err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
}
