package kubectl

//batch.v1beta1
type cronJob struct {
	Name      string
	Namespace string
}

//auditregistration.v1alpha1
type auditSink struct {
	Name string
}

//admissionregistration.v1beta1
type validatingWebhookConfiguration struct {
	Name string
}

type mutatingWebhookConfiguration struct {
	Name string
}

//batch.v1
type job struct {
	Name      string
	Namespace string
}

//certificates.v1beta1
type certificateSigningRequest struct {
	Name string
}

//rbac.v1
type roleBinding struct {
	Name      string
	Namespace string
}

type role struct {
	Name      string
	Namespace string
}

type clusterRoleBinding struct {
	Name string
}

type clusterRole struct {
	Name string
}

//coordination.v1beta1
type lease struct {
	Name      string
	Namespace string
}

//networking.v1
type networkPolicy struct {
	Name      string
	Namespace string
}

//apps.v1
type controllerRevision struct {
	Name      string
	Namespace string
}

type daemonSet struct {
	Name      string
	Namespace string
}

type deployment struct {
	Name      string
	Namespace string
}

type replicaSet struct {
	Name      string
	Namespace string
}

type statefulSet struct {
	Name      string
	Namespace string
}

//autoscaling.v2beta2
type horizontalPodAutoscaler struct {
	Name      string
	Namespace string
}

//policy.v1beta1
type eviction struct {
	Name      string
	Namespace string
}

type podDisruptionBudget struct {
	Name      string
	Namespace string
}

type podSecurityPolicy struct {
	Name string
}

//authorization.v1
type selfSubjectAccessReview struct {
	Name string
}

type selfSubjectRulesReview struct {
	Name string
}

type subjectAccessReview struct {
	Name string
}

type localSubjectAccessReview struct {
	Name      string
	Namespace string
}

// v1
type resourceQuota struct {
	Name      string
	Namespace string
}

type componentStatus struct {
	Name string
}

type configMap struct {
	Name      string
	Namespace string
}

type limitRange struct {
	Name      string
	Namespace string
}

type namespace struct {
	Name string
}

type pod struct {
	Name          string
	Namespace     string
	ContainerName string
}

type replicationController struct {
	Name      string
	Namespace string
}

type endpoints struct {
	Name      string
	Namespace string
}

type event struct {
	Name      string
	Namespace string
}

type persistentVolumeClaim struct {
	Name      string
	Namespace string
}

type persistentVolume struct {
	Name string
}

type serviceAccount struct {
	Name      string
	Namespace string
}

type node struct {
	Name string
}

type podTemplate struct {
	Name      string
	Namespace string
}

type secret struct {
	Name      string
	Namespace string
}

type service struct {
	Name      string
	Namespace string
}

//settings.v1alpha1
type podPreset struct {
	Name      string
	Namespace string
}

//admissionregistration.v1alpha1
type initializerConfiguration struct {
	Name string
}

//authentication.v1
type tokenReview struct {
	Name string
}

type ingress struct {
	Name      string
	Namespace string
}

//scheduling.v1beta1
type priorityClass struct {
	Name string
}

//storage.v1
type storageClass struct {
	Name string
}

type volumeAttachment struct {
	Name string
}
