package kubectl

//batch.v1beta1
func CronJob(name string, namespace string) cronJob {
	return cronJob{name, namespace}
}

//certificates.v1beta1
func CertificateSigningRequest(name string) certificateSigningRequest {
	return certificateSigningRequest{name}
}

//coordination.v1beta1
func Lease(name string, namespace string) lease {
	return lease{name, namespace}
}

//networking.v1
func NetworkPolicy(name string, namespace string) networkPolicy {
	return networkPolicy{name, namespace}
}

func Ingress(name string, namespace string) ingress {
	return ingress{name, namespace}
}

//rbac.v1
func ClusterRoleBinding(name string) clusterRoleBinding {
	return clusterRoleBinding{name}
}

func ClusterRole(name string) clusterRole {
	return clusterRole{name}
}

func RoleBinding(name string, namespace string) roleBinding {
	return roleBinding{name, namespace}
}

func Role(name string, namespace string) role {
	return role{name, namespace}
}

//settings.v1alpha1
func PodPreset(name string, namespace string) podPreset {
	return podPreset{name, namespace}
}

//admissionregistration.v1beta1
func MutatingWebhookConfiguration(name string) mutatingWebhookConfiguration {
	return mutatingWebhookConfiguration{name}
}

func ValidatingWebhookConfiguration(name string) validatingWebhookConfiguration {
	return validatingWebhookConfiguration{name}
}

//policy.v1beta1
func Eviction(name string, namespace string) eviction {
	return eviction{name, namespace}
}

func PodDisruptionBudget(name string, namespace string) podDisruptionBudget {
	return podDisruptionBudget{name, namespace}
}

func PodSecurityPolicy(name string) podSecurityPolicy {
	return podSecurityPolicy{name}
}

//batch.v1
func Job(name string, namespace string) job {
	return job{name, namespace}
}

//storage.v1
func StorageClass(name string) storageClass {
	return storageClass{name}
}

func VolumeAttachment(name string) volumeAttachment {
	return volumeAttachment{name}
}

//auditregistration.v1alpha1
func AuditSink(name string) auditSink {
	return auditSink{name}
}

//authentication.v1beta1
func TokenReview(name string) tokenReview {
	return tokenReview{name}
}

//scheduling.v1beta1
func PriorityClass(name string) priorityClass {
	return priorityClass{name}
}

//apps.v1
func ControllerRevision(name string, namespace string) controllerRevision {
	return controllerRevision{name, namespace}
}

func DaemonSet(name string, namespace string) daemonSet {
	return daemonSet{name, namespace}
}

func Deployment(name string, namespace string) deployment {
	return deployment{name, namespace}
}

func ReplicaSet(name string, namespace string) replicaSet {
	return replicaSet{name, namespace}
}

func StatefulSet(name string, namespace string) statefulSet {
	return statefulSet{name, namespace}
}

//authorization.v1
func LocalSubjectAccessReview(name string, namespace string) localSubjectAccessReview {
	return localSubjectAccessReview{name, namespace}
}

func SelfSubjectAccessReview(name string) selfSubjectAccessReview {
	return selfSubjectAccessReview{name}
}

func SelfSubjectRulesReview(name string) selfSubjectRulesReview {
	return selfSubjectRulesReview{name}
}

func SubjectAccessReview(name string) subjectAccessReview {
	return subjectAccessReview{name}
}

//autoscaling.v1
func HorizontalPodAutoscaler(name string, namespace string) horizontalPodAutoscaler {
	return horizontalPodAutoscaler{name, namespace}
}

// v1
func ComponentStatus(name string) componentStatus {
	return componentStatus{name}
}

func ResourceQuota(name string, namespace string) resourceQuota {
	return resourceQuota{name, namespace}
}

func ServiceAccount(name string, namespace string) serviceAccount {
	return serviceAccount{name, namespace}
}

func Service(name string, namespace string) service {
	return service{name, namespace}
}

func ConfigMap(name string, namespace string) configMap {
	return configMap{name, namespace}
}

func PersistentVolumeClaim(name string, namespace string) persistentVolumeClaim {
	return persistentVolumeClaim{name, namespace}
}

func ReplicationController(name string, namespace string) replicationController {
	return replicationController{name, namespace}
}

func Secret(name string, namespace string) secret {
	return secret{name, namespace}
}

func LimitRange(name string, namespace string) limitRange {
	return limitRange{name, namespace}
}

func Node(name string) node {
	return node{name}
}

func PersistentVolume(name string) persistentVolume {
	return persistentVolume{name}
}

func PodTemplate(name string, namespace string) podTemplate {
	return podTemplate{name, namespace}
}

func Endpoints(name string, namespace string) endpoints {
	return endpoints{name, namespace}
}

func Event(name string, namespace string) event {
	return event{name, namespace}
}

func Namespace(name string) namespace {
	return namespace{name}
}

func Pod(name string, namespace string, ContainerName ...string) pod {
	if len(ContainerName) == 0 {
		return pod{name, namespace, ""}
	} else if len(ContainerName) == 1 {
		return pod{name, namespace, ContainerName[0]}
	} else {
		panic("none or less one container")
	}
}

//admissionregistration.v1alpha1
func InitializerConfiguration(name string) initializerConfiguration {
	return initializerConfiguration{name}
}
