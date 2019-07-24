package kubectl

import (
	admissionregistrationV1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	admissionregistrationV1beta1 "k8s.io/api/admissionregistration/v1beta1"
	appsV1 "k8s.io/api/apps/v1"
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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)



func (i *CronJob) Get (opts *v1.GetOptions) (batchV1beta1.CronJob,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return batchV1beta1.CronJob{},err
	}
	cronJob , err := clientset.BatchV1beta1().CronJobs(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return batchV1beta1.CronJob{},err
	}
	return *cronJob,nil
}

func (i *AuditSink) Get (opts *v1.GetOptions) (auditregistrationV1alpha1.AuditSink,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return auditregistrationV1alpha1.AuditSink{},err
	}
	auditSink , err := clientset.AuditregistrationV1alpha1().AuditSinks().Get(i.Name,*opts)
	if err != nil {
		return auditregistrationV1alpha1.AuditSink{},err
	}
	return *auditSink,nil
}

func (i *ValidatingWebhookConfiguration) Get (opts *v1.GetOptions) (admissionregistrationV1beta1.ValidatingWebhookConfiguration,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return admissionregistrationV1beta1.ValidatingWebhookConfiguration{},err
	}
	validatingWebhookConfiguration , err := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().Get(i.Name,*opts)
	if err != nil {
		return admissionregistrationV1beta1.ValidatingWebhookConfiguration{},err
	}
	return *validatingWebhookConfiguration,nil
}

func (i *MutatingWebhookConfiguration) Get (opts *v1.GetOptions) (admissionregistrationV1beta1.MutatingWebhookConfiguration,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return admissionregistrationV1beta1.MutatingWebhookConfiguration{},err
	}
	mutatingWebhookConfiguration , err := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Get(i.Name,*opts)
	if err != nil {
		return admissionregistrationV1beta1.MutatingWebhookConfiguration{},err
	}
	return *mutatingWebhookConfiguration,nil
}

func (i *Job) Get (opts *v1.GetOptions) (batchV1.Job,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return batchV1.Job{},err
	}
	job , err := clientset.BatchV1().Jobs(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return batchV1.Job{},err
	}
	return *job,nil
}

func (i *CertificateSigningRequest) Get (opts *v1.GetOptions) (certificatesV1beta1.CertificateSigningRequest,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return certificatesV1beta1.CertificateSigningRequest{},err
	}
	certificateSigningRequest , err := clientset.CertificatesV1beta1().CertificateSigningRequests().Get(i.Name,*opts)
	if err != nil {
		return certificatesV1beta1.CertificateSigningRequest{},err
	}
	return *certificateSigningRequest,nil
}

func (i *ClusterRoleBinding) Get (opts *v1.GetOptions) (rbacV1.ClusterRoleBinding,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return rbacV1.ClusterRoleBinding{},err
	}
	clusterRoleBinding , err := clientset.RbacV1().ClusterRoleBindings().Get(i.Name,*opts)
	if err != nil {
		return rbacV1.ClusterRoleBinding{},err
	}
	return *clusterRoleBinding,nil
}

func (i *ClusterRole) Get (opts *v1.GetOptions) (rbacV1.ClusterRole,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return rbacV1.ClusterRole{},err
	}
	clusterRole , err := clientset.RbacV1().ClusterRoles().Get(i.Name,*opts)
	if err != nil {
		return rbacV1.ClusterRole{},err
	}
	return *clusterRole,nil
}

func (i *RoleBinding) Get (opts *v1.GetOptions) (rbacV1.RoleBinding,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return rbacV1.RoleBinding{},err
	}
	roleBinding , err := clientset.RbacV1().RoleBindings(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return rbacV1.RoleBinding{},err
	}
	return *roleBinding,nil
}

func (i *Role) Get (opts *v1.GetOptions) (rbacV1.Role,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return rbacV1.Role{},err
	}
	role , err := clientset.RbacV1().Roles(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return rbacV1.Role{},err
	}
	return *role,nil
}

func (i *Lease) Get (opts *v1.GetOptions) (coordinationV1beta1.Lease,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coordinationV1beta1.Lease{},err
	}
	lease , err := clientset.CoordinationV1beta1().Leases(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coordinationV1beta1.Lease{},err
	}
	return *lease,nil
}

func (i *NetworkPolicy) Get (opts *v1.GetOptions) (networkingV1.NetworkPolicy,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return networkingV1.NetworkPolicy{},err
	}
	networkPolicy , err := clientset.NetworkingV1().NetworkPolicies(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return networkingV1.NetworkPolicy{},err
	}
	return *networkPolicy,nil
}

func (i *DaemonSet) Get (opts *v1.GetOptions) (appsV1.DaemonSet,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return appsV1.DaemonSet{},err
	}
	daemonSet , err := clientset.AppsV1().DaemonSets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return appsV1.DaemonSet{},err
	}
	return *daemonSet,nil
}

func (i *Deployment) Get (opts *v1.GetOptions) (appsV1.Deployment,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return appsV1.Deployment{},err
	}
	deployment , err := clientset.AppsV1().Deployments(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return appsV1.Deployment{},err
	}
	return *deployment,nil
}

func (i *ReplicaSet) Get (opts *v1.GetOptions) (appsV1.ReplicaSet,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return appsV1.ReplicaSet{},err
	}
	replicaSet , err := clientset.AppsV1().ReplicaSets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return appsV1.ReplicaSet{},err
	}
	return *replicaSet,nil
}

func (i *StatefulSet) Get (opts *v1.GetOptions) (appsV1.StatefulSet,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return appsV1.StatefulSet{},err
	}
	statefulSet , err := clientset.AppsV1().StatefulSets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return appsV1.StatefulSet{},err
	}
	return *statefulSet,nil
}

func (i *ControllerRevision) Get (opts *v1.GetOptions) (appsV1.ControllerRevision,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return appsV1.ControllerRevision{},err
	}
	controllerRevision , err := clientset.AppsV1().ControllerRevisions(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return appsV1.ControllerRevision{},err
	}
	return *controllerRevision,nil
}

func (i *HorizontalPodAutoscaler) Get (opts *v1.GetOptions) (autoscalingV2beta2.HorizontalPodAutoscaler,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return autoscalingV2beta2.HorizontalPodAutoscaler{},err
	}
	horizontalPodAutoscaler , err := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return autoscalingV2beta2.HorizontalPodAutoscaler{},err
	}
	return *horizontalPodAutoscaler,nil
}



func (i *PodDisruptionBudget) Get (opts *v1.GetOptions) (policyV1beta1.PodDisruptionBudget,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return policyV1beta1.PodDisruptionBudget{},err
	}
	podDisruptionBudget , err := clientset.PolicyV1beta1().PodDisruptionBudgets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return policyV1beta1.PodDisruptionBudget{},err
	}
	return *podDisruptionBudget,nil
}

func (i *PodSecurityPolicy) Get (opts *v1.GetOptions) (policyV1beta1.PodSecurityPolicy,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return policyV1beta1.PodSecurityPolicy{},err
	}
	podSecurityPolicy , err := clientset.PolicyV1beta1().PodSecurityPolicies().Get(i.Name,*opts)
	if err != nil {
		return policyV1beta1.PodSecurityPolicy{},err
	}
	return *podSecurityPolicy,nil
}


func (i *ResourceQuota) Get (opts *v1.GetOptions) (coreV1.ResourceQuota,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.ResourceQuota{},err
	}
	resourceQuota , err := clientset.CoreV1().ResourceQuotas(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.ResourceQuota{},err
	}
	return *resourceQuota,nil
}

func (i *Secret) Get (opts *v1.GetOptions) (coreV1.Secret,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.Secret{},err
	}
	secret , err := clientset.CoreV1().Secrets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.Secret{},err
	}
	return *secret,nil
}

func (i *ComponentStatus) Get (opts *v1.GetOptions) (coreV1.ComponentStatus,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.ComponentStatus{},err
	}
	componentStatus , err := clientset.CoreV1().ComponentStatuses().Get(i.Name,*opts)
	if err != nil {
		return coreV1.ComponentStatus{},err
	}
	return *componentStatus,nil
}

func (i *Event) Get (opts *v1.GetOptions) (coreV1.Event,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.Event{},err
	}
	event , err := clientset.CoreV1().Events(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.Event{},err
	}
	return *event,nil
}

func (i *Pod) Get (opts *v1.GetOptions) (coreV1.Pod,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.Pod{},err
	}
	pod , err := clientset.CoreV1().Pods(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.Pod{},err
	}
	return *pod,nil
}

func (i *PersistentVolume) Get (opts *v1.GetOptions) (coreV1.PersistentVolume,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.PersistentVolume{},err
	}
	persistentVolume , err := clientset.CoreV1().PersistentVolumes().Get(i.Name,*opts)
	if err != nil {
		return coreV1.PersistentVolume{},err
	}
	return *persistentVolume,nil
}

func (i *Service) Get (opts *v1.GetOptions) (coreV1.Service,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.Service{},err
	}
	service , err := clientset.CoreV1().Services(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.Service{},err
	}
	return *service,nil
}

func (i *PodTemplate) Get (opts *v1.GetOptions) (coreV1.PodTemplate,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.PodTemplate{},err
	}
	podTemplate , err := clientset.CoreV1().PodTemplates(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.PodTemplate{},err
	}
	return *podTemplate,nil
}

func (i *ReplicationController) Get (opts *v1.GetOptions) (coreV1.ReplicationController,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.ReplicationController{},err
	}
	replicationController , err := clientset.CoreV1().ReplicationControllers(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.ReplicationController{},err
	}
	return *replicationController,nil
}

func (i *ServiceAccount) Get (opts *v1.GetOptions) (coreV1.ServiceAccount,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.ServiceAccount{},err
	}
	serviceAccount , err := clientset.CoreV1().ServiceAccounts(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.ServiceAccount{},err
	}
	return *serviceAccount,nil
}

func (i *LimitRange) Get (opts *v1.GetOptions) (coreV1.LimitRange,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.LimitRange{},err
	}
	limitRange , err := clientset.CoreV1().LimitRanges(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.LimitRange{},err
	}
	return *limitRange,nil
}

func (i *Namespace) Get (opts *v1.GetOptions) (coreV1.Namespace,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.Namespace{},err
	}
	namespace , err := clientset.CoreV1().Namespaces().Get(i.Name,*opts)
	if err != nil {
		return coreV1.Namespace{},err
	}
	return *namespace,nil
}

func (i *Node) Get (opts *v1.GetOptions) (coreV1.Node,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.Node{},err
	}
	node , err := clientset.CoreV1().Nodes().Get(i.Name,*opts)
	if err != nil {
		return coreV1.Node{},err
	}
	return *node,nil
}

func (i *ConfigMap) Get (opts *v1.GetOptions) (coreV1.ConfigMap,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.ConfigMap{},err
	}
	configMap , err := clientset.CoreV1().ConfigMaps(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.ConfigMap{},err
	}
	return *configMap,nil
}



func (i *PersistentVolumeClaim) Get (opts *v1.GetOptions) (coreV1.PersistentVolumeClaim,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return coreV1.PersistentVolumeClaim{},err
	}
	persistentVolumeClaim , err := clientset.CoreV1().PersistentVolumeClaims(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return coreV1.PersistentVolumeClaim{},err
	}
	return *persistentVolumeClaim,nil
}

func (i *PodPreset) Get (opts *v1.GetOptions) (settingsV1alpha1.PodPreset,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return settingsV1alpha1.PodPreset{},err
	}
	podPreset , err := clientset.SettingsV1alpha1().PodPresets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return settingsV1alpha1.PodPreset{},err
	}
	return *podPreset,nil
}

func (i *InitializerConfiguration) Get (opts *v1.GetOptions) (admissionregistrationV1alpha1.InitializerConfiguration,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return admissionregistrationV1alpha1.InitializerConfiguration{},err
	}
	initializerConfiguration , err := clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().Get(i.Name,*opts)
	if err != nil {
		return admissionregistrationV1alpha1.InitializerConfiguration{},err
	}
	return *initializerConfiguration,nil
}



func (i *Ingress) Get (opts *v1.GetOptions) (extensionsV1beta1.Ingress,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return extensionsV1beta1.Ingress{},err
	}
	ingress , err := clientset.ExtensionsV1beta1().Ingresses(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return extensionsV1beta1.Ingress{},err
	}
	return *ingress,nil
}



func (i *PriorityClass) Get (opts *v1.GetOptions) (schedulingV1beta1.PriorityClass,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return schedulingV1beta1.PriorityClass{},err
	}
	priorityClass , err := clientset.SchedulingV1beta1().PriorityClasses().Get(i.Name,*opts)
	if err != nil {
		return schedulingV1beta1.PriorityClass{},err
	}
	return *priorityClass,nil
}

func (i *StorageClass) Get (opts *v1.GetOptions) (storageV1.StorageClass,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return storageV1.StorageClass{},err
	}
	storageClass , err := clientset.StorageV1().StorageClasses().Get(i.Name,*opts)
	if err != nil {
		return storageV1.StorageClass{},err
	}
	return *storageClass,nil
}

func (i *VolumeAttachment) Get (opts *v1.GetOptions) (storageV1.VolumeAttachment,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return storageV1.VolumeAttachment{},err
	}
	volumeAttachment , err := clientset.StorageV1().VolumeAttachments().Get(i.Name,*opts)
	if err != nil {
		return storageV1.VolumeAttachment{},err
	}
	return *volumeAttachment,nil
}