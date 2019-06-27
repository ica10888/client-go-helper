package kubectl

import (
	corev1 "k8s.io/api/core/v1"
)

//auditregistration.v1alpha1
type AuditSink struct {
	Name     string

}


//admissionregistration.v1beta1
type ValidatingWebhookConfiguration struct {
	Name     string

}

type MutatingWebhookConfiguration struct {
	Name     string

}


//batch.v1
type Job struct {
	Name     string
	Namespace     string
}


//certificates.v1beta1
type CertificateSigningRequest struct {
	Name     string

}


//rbac.v1
type RoleBinding struct {
	Name     string
	Namespace     string
}

type Role struct {
	Name     string
	Namespace     string
}

type ClusterRoleBinding struct {
	Name     string

}

type ClusterRole struct {
	Name     string

}


//coordination.v1beta1
type Lease struct {
	Name     string
	Namespace     string
}


//networking.v1
type NetworkPolicy struct {
	Name     string
	Namespace     string
}


//apps.v1
type ControllerRevision struct {
	Name     string
	Namespace     string
}

type DaemonSet struct {
	Name     string
	Namespace     string
}

type Deployment struct {
	Name     string
	Namespace     string
}

type ReplicaSet struct {
	Name     string
	Namespace     string
}

type StatefulSet struct {
	Name     string
	Namespace     string
}


//autoscaling.v2beta2
type HorizontalPodAutoscaler struct {
	Name     string
	Namespace     string
}


//policy.v1beta1
type Eviction struct {
	Name     string
	Namespace     string
}

type PodDisruptionBudget struct {
	Name     string
	Namespace     string
}

type PodSecurityPolicy struct {
	Name     string

}


//authorization.v1
type SelfSubjectAccessReview struct {
	Name     string

}

type SelfSubjectRulesReview struct {
	Name     string

}

type SubjectAccessReview struct {
	Name     string

}

type LocalSubjectAccessReview struct {
	Name     string
	Namespace     string
}


// v1
type ResourceQuota struct {
	Name     string
	Namespace     string
}

type ComponentStatus struct {
	Name     string

}

type ConfigMap struct {
	Name     string
	Namespace     string
}

type LimitRange struct {
	Name     string
	Namespace     string
}

type Namespace struct {
	Name     string

}

type Pod struct {
	Name     string
	Namespace     string
	ContainerName string
}

type ReplicationController struct {
	Name     string
	Namespace     string
}

type Endpoints struct {
	Name     string
	Namespace     string
}

type Event struct {
	Name     string
	Namespace     string
}

type PersistentVolumeClaim struct {
	Name     string
	Namespace     string
}

type PersistentVolume struct {
	Name     string

}

type ServiceAccount struct {
	Name     string
	Namespace     string
}

type Node struct {
	Name     string

}

type PodTemplate struct {
	Name     string
	Namespace     string
}

type Secret struct {
	Name     string
	Namespace     string
}

type Service struct {
	Name     string
	Namespace     string
}


//settings.v1alpha1
type PodPreset struct {
	Name     string
	Namespace     string
}


//admissionregistration.v1alpha1
type InitializerConfiguration struct {
	Name     string

}


//authentication.v1
type TokenReview struct {
	Name     string

}



type Ingress struct {
	Name     string
	Namespace     string
}



//scheduling.v1beta1
type PriorityClass struct {
	Name     string

}


//storage.v1
type StorageClass struct {
	Name     string

}

type VolumeAttachment struct {
	Name     string

}


//batch.v1beta1
type BatchV1beta1CronJob struct {
	Name     string
	Namespace     string
}


//certificates.v1beta1
type CertificatesV1beta1CertificateSigningRequest struct {
	Name     string

}


//coordination.v1beta1
type CoordinationV1beta1Lease struct {
	Name     string
	Namespace     string
}


//networking.v1
type NetworkingV1NetworkPolicy struct {
	Name     string
	Namespace     string
}


//authorization.v1beta1
type AuthorizationV1beta1SubjectAccessReview struct {
	Name     string

}

type AuthorizationV1beta1LocalSubjectAccessReview struct {
	Name     string
	Namespace     string
}

type AuthorizationV1beta1SelfSubjectAccessReview struct {
	Name     string

}

type AuthorizationV1beta1SelfSubjectRulesReview struct {
	Name     string

}


//extensions.v1beta1
type ExtensionsV1beta1ReplicaSet struct {
	Name     string
	Namespace     string
}

type ExtensionsV1beta1DaemonSet struct {
	Name     string
	Namespace     string
}

type ExtensionsV1beta1Deployment struct {
	Name     string
	Namespace     string
}

type ExtensionsV1beta1Ingress struct {
	Name     string
	Namespace     string
}

type ExtensionsV1beta1PodSecurityPolicy struct {
	Name     string

}


//rbac.v1
type RbacV1ClusterRoleBinding struct {
	Name     string

}

type RbacV1ClusterRole struct {
	Name     string

}

type RbacV1RoleBinding struct {
	Name     string
	Namespace     string
}

type RbacV1Role struct {
	Name     string
	Namespace     string
}


//apps.v1beta2
type AppsV1beta2ControllerRevision struct {
	Name     string
	Namespace     string
}

type AppsV1beta2DaemonSet struct {
	Name     string
	Namespace     string
}

type AppsV1beta2Deployment struct {
	Name     string
	Namespace     string
}

type AppsV1beta2ReplicaSet struct {
	Name     string
	Namespace     string
}

type AppsV1beta2StatefulSet struct {
	Name     string
	Namespace     string
}


//authentication.v1
type AuthenticationV1TokenReview struct {
	Name     string

}


//rbac.v1alpha1
type RbacV1alpha1ClusterRoleBinding struct {
	Name     string

}

type RbacV1alpha1ClusterRole struct {
	Name     string

}

type RbacV1alpha1RoleBinding struct {
	Name     string
	Namespace     string
}

type RbacV1alpha1Role struct {
	Name     string
	Namespace     string
}


//settings.v1alpha1
type SettingsV1alpha1PodPreset struct {
	Name     string
	Namespace     string
}


//admissionregistration.v1beta1
type AdmissionregistrationV1beta1MutatingWebhookConfiguration struct {
	Name     string

}

type AdmissionregistrationV1beta1ValidatingWebhookConfiguration struct {
	Name     string

}


//events.v1beta1
type EventsV1beta1Event struct {
	Name     string
	Namespace     string
}


//policy.v1beta1
type PolicyV1beta1Eviction struct {
	Name     string
	Namespace     string
}

type PolicyV1beta1PodDisruptionBudget struct {
	Name     string
	Namespace     string
}

type PolicyV1beta1PodSecurityPolicy struct {
	Name     string

}


//scheduling.v1alpha1
type SchedulingV1alpha1PriorityClass struct {
	Name     string

}


//storage.v1beta1
type StorageV1beta1StorageClass struct {
	Name     string

}

type StorageV1beta1VolumeAttachment struct {
	Name     string

}


//batch.v1
type BatchV1Job struct {
	Name     string
	Namespace     string
}


//rbac.v1beta1
type RbacV1beta1ClusterRoleBinding struct {
	Name     string

}

type RbacV1beta1ClusterRole struct {
	Name     string

}

type RbacV1beta1RoleBinding struct {
	Name     string
	Namespace     string
}

type RbacV1beta1Role struct {
	Name     string
	Namespace     string
}


//storage.v1
type StorageV1StorageClass struct {
	Name     string

}

type StorageV1VolumeAttachment struct {
	Name     string

}


//auditregistration.v1alpha1
type AuditregistrationV1alpha1AuditSink struct {
	Name     string

}


//authentication.v1beta1
type AuthenticationV1beta1TokenReview struct {
	Name     string

}


//autoscaling.v2beta1
type AutoscalingV2beta1HorizontalPodAutoscaler struct {
	Name     string
	Namespace     string
}


//scheduling.v1beta1
type SchedulingV1beta1PriorityClass struct {
	Name     string

}


//apps.v1
type AppsV1Deployment struct {
	Name     string
	Namespace     string
}

type AppsV1ReplicaSet struct {
	Name     string
	Namespace     string
}

type AppsV1StatefulSet struct {
	Name     string
	Namespace     string
}

type AppsV1ControllerRevision struct {
	Name     string
	Namespace     string
}

type AppsV1DaemonSet struct {
	Name     string
	Namespace     string
}


//apps.v1beta1
type AppsV1beta1StatefulSet struct {
	Name     string
	Namespace     string
}

type AppsV1beta1ControllerRevision struct {
	Name     string
	Namespace     string
}

type AppsV1beta1Deployment struct {
	Name     string
	Namespace     string
}


//authorization.v1
type AuthorizationV1SelfSubjectAccessReview struct {
	Name     string

}

type AuthorizationV1SelfSubjectRulesReview struct {
	Name     string

}

type AuthorizationV1SubjectAccessReview struct {
	Name     string

}

type AuthorizationV1LocalSubjectAccessReview struct {
	Name     string
	Namespace     string
}


//autoscaling.v1
type AutoscalingV1HorizontalPodAutoscaler struct {
	Name     string
	Namespace     string
}


//autoscaling.v2beta2
type AutoscalingV2beta2HorizontalPodAutoscaler struct {
	Name     string
	Namespace     string
}


//batch.v2alpha1
type BatchV2alpha1CronJob struct {
	Name     string
	Namespace     string
}


// v1
type CoreV1ConfigMap struct {
	Name     string
	Namespace     string
}

type CoreV1Endpoints struct {
	Name     string
	Namespace     string
}

type CoreV1PersistentVolumeClaim struct {
	Name     string
	Namespace     string
}

type CoreV1ResourceQuota struct {
	Name     string
	Namespace     string
}

type CoreV1ServiceAccount struct {
	Name     string
	Namespace     string
}

type CoreV1LimitRange struct {
	Name     string
	Namespace     string
}

type CoreV1Node struct {
	Name     string

}

type CoreV1PodTemplate struct {
	Name     string
	Namespace     string
}

type CoreV1ReplicationController struct {
	Name     string
	Namespace     string
}

type CoreV1ComponentStatus struct {
	Name     string

}

type CoreV1Event struct {
	Name     string
	Namespace     string
}

type CoreV1Secret struct {
	Name     string
	Namespace     string
}

type CoreV1Service struct {
	Name     string
	Namespace     string
}

type CoreV1Namespace struct {
	Name     string

}

type CoreV1PersistentVolume struct {
	Name     string

}

type CoreV1Pod struct {
	Name     string
	Namespace     string
	ContainerName string
}


//storage.v1alpha1
type StorageV1alpha1VolumeAttachment struct {
	Name     string

}


//admissionregistration.v1alpha1
type AdmissionregistrationV1alpha1InitializerConfiguration struct {
	Name     string

}







type describe interface {
	Describe()(string ,error)
}


type get interface {
	Get()(error)
}

type exec interface {
	Exec(cmd []string)(error)
}



type cp interface {
	Cp(srcPath string, destPath string)(error)
}


type logs interface {
	Logs(podLogOpts *corev1.PodLogOptions)(string,error)
}




//used for apply/create
type Kubeapi struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Yaml       string
}
