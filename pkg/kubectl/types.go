package kubectl

//batch.v1beta1
type CronJob struct {
	Name      string
	Namespace string
}

//auditregistration.v1alpha1
type AuditSink struct {
	Name string
}

//admissionregistration.v1beta1
type ValidatingWebhookConfiguration struct {
	Name string
}

type MutatingWebhookConfiguration struct {
	Name string
}

//batch.v1
type Job struct {
	Name      string
	Namespace string
}

//certificates.v1beta1
type CertificateSigningRequest struct {
	Name string
}

//rbac.v1
type RoleBinding struct {
	Name      string
	Namespace string
}

type Role struct {
	Name      string
	Namespace string
}

type ClusterRoleBinding struct {
	Name string
}

type ClusterRole struct {
	Name string
}

//coordination.v1beta1
type Lease struct {
	Name      string
	Namespace string
}

//networking.v1
type NetworkPolicy struct {
	Name      string
	Namespace string
}

//apps.v1
type ControllerRevision struct {
	Name      string
	Namespace string
}

type DaemonSet struct {
	Name      string
	Namespace string
}

type Deployment struct {
	Name      string
	Namespace string
}

type ReplicaSet struct {
	Name      string
	Namespace string
}

type StatefulSet struct {
	Name      string
	Namespace string
}

//autoscaling.v2beta2
type HorizontalPodAutoscaler struct {
	Name      string
	Namespace string
}

//policy.v1beta1
type Eviction struct {
	Name      string
	Namespace string
}

type PodDisruptionBudget struct {
	Name      string
	Namespace string
}

type PodSecurityPolicy struct {
	Name string
}

//authorization.v1
type SelfSubjectAccessReview struct {
	Name string
}

type SelfSubjectRulesReview struct {
	Name string
}

type SubjectAccessReview struct {
	Name string
}

type LocalSubjectAccessReview struct {
	Name      string
	Namespace string
}

// v1
type ResourceQuota struct {
	Name      string
	Namespace string
}

type ComponentStatus struct {
	Name string
}

type ConfigMap struct {
	Name      string
	Namespace string
}

type LimitRange struct {
	Name      string
	Namespace string
}

type Namespace struct {
	Name string
}

type Pod struct {
	Name          string
	Namespace     string
	ContainerName string
}

type ReplicationController struct {
	Name      string
	Namespace string
}

type Endpoints struct {
	Name      string
	Namespace string
}

type Event struct {
	Name      string
	Namespace string
}

type PersistentVolumeClaim struct {
	Name      string
	Namespace string
}

type PersistentVolume struct {
	Name string
}

type ServiceAccount struct {
	Name      string
	Namespace string
}

type Node struct {
	Name string
}

type PodTemplate struct {
	Name      string
	Namespace string
}

type Secret struct {
	Name      string
	Namespace string
}

type Service struct {
	Name      string
	Namespace string
}

//settings.v1alpha1
type PodPreset struct {
	Name      string
	Namespace string
}

//admissionregistration.v1alpha1
type InitializerConfiguration struct {
	Name string
}

//authentication.v1
type TokenReview struct {
	Name string
}

type Ingress struct {
	Name      string
	Namespace string
}

//scheduling.v1beta1
type PriorityClass struct {
	Name string
}

//storage.v1
type StorageClass struct {
	Name string
}

type VolumeAttachment struct {
	Name string
}
