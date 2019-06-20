package kubectl

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)



func Get(opts *v1.ListOptions,kapi *Kubeapi ,namespace string) (error) {
	// 校验和放入结构体
	clientset, e := InitClient()
	if e != nil {
		return fmt.Errorf("something wrong happend ,%s", e)
	}

	switch kapi.ApiVersion {
	case "batch/v1beta1":
		switch kapi.Kind {

		case "CronJob":
			_, e = clientset.BatchV1beta1().CronJobs(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "certificates/v1beta1":
		switch kapi.Kind {

		case "CertificateSigningRequest":
			_, e = clientset.CertificatesV1beta1().CertificateSigningRequests().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "coordination/v1beta1":
		switch kapi.Kind {

		case "Lease":
			_, e = clientset.CoordinationV1beta1().Leases(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "networking/v1":
		switch kapi.Kind {

		case "NetworkPolicy":
			_, e = clientset.NetworkingV1().NetworkPolicies(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "authorization/v1beta1":
		switch kapi.Kind {

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "extensions/v1beta1":
		switch kapi.Kind {

		case "DaemonSet":
			_, e = clientset.ExtensionsV1beta1().DaemonSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Deployment":
			dl, e := clientset.ExtensionsV1beta1().Deployments(namespace).List(*opts)
			if e != nil {
				return e
			}
			fmt.Print(dl)

		case "Ingress":
			_, e = clientset.ExtensionsV1beta1().Ingresses(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "PodSecurityPolicy":
			_, e = clientset.ExtensionsV1beta1().PodSecurityPolicies().List(*opts)
			if e != nil {
				return e
			}

		case "ReplicaSet":
			_, e = clientset.ExtensionsV1beta1().ReplicaSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "rbac/v1":
		switch kapi.Kind {

		case "RoleBinding":
			_, e = clientset.RbacV1().RoleBindings(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Role":
			_, e = clientset.RbacV1().Roles(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ClusterRoleBinding":
			_, e = clientset.RbacV1().ClusterRoleBindings().List(*opts)
			if e != nil {
				return e
			}

		case "ClusterRole":
			_, e = clientset.RbacV1().ClusterRoles().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1beta2":
		switch kapi.Kind {

		case "ReplicaSet":
			_, e = clientset.AppsV1beta2().ReplicaSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "StatefulSet":
			_, e = clientset.AppsV1beta2().StatefulSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ControllerRevision":
			_, e = clientset.AppsV1beta2().ControllerRevisions(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "DaemonSet":
			_, e = clientset.AppsV1beta2().DaemonSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Deployment":
			_, e = clientset.AppsV1beta2().Deployments(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "authentication/v1":
		switch kapi.Kind {

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "rbac/v1alpha1":
		switch kapi.Kind {

		case "ClusterRole":
			_, e = clientset.RbacV1alpha1().ClusterRoles().List(*opts)
			if e != nil {
				return e
			}

		case "RoleBinding":
			_, e = clientset.RbacV1alpha1().RoleBindings(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Role":
			_, e = clientset.RbacV1alpha1().Roles(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ClusterRoleBinding":
			_, e = clientset.RbacV1alpha1().ClusterRoleBindings().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "settings/v1alpha1":
		switch kapi.Kind {

		case "PodPreset":
			_, e = clientset.SettingsV1alpha1().PodPresets(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "admissionregistration/v1beta1":
		switch kapi.Kind {

		case "ValidatingWebhookConfiguration":
			_, e = clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(*opts)
			if e != nil {
				return e
			}

		case "MutatingWebhookConfiguration":
			_, e = clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "events/v1beta1":
		switch kapi.Kind {

		case "Event":
			_, e = clientset.EventsV1beta1().Events(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "policy/v1beta1":
		switch kapi.Kind {

		case "PodDisruptionBudget":
			_, e = clientset.PolicyV1beta1().PodDisruptionBudgets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "PodSecurityPolicy":
			_, e = clientset.PolicyV1beta1().PodSecurityPolicies().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "scheduling/v1alpha1":
		switch kapi.Kind {

		case "PriorityClass":
			_, e = clientset.SchedulingV1alpha1().PriorityClasses().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1beta1":
		switch kapi.Kind {

		case "StorageClass":
			_, e = clientset.StorageV1beta1().StorageClasses().List(*opts)
			if e != nil {
				return e
			}

		case "VolumeAttachment":
			_, e = clientset.StorageV1beta1().VolumeAttachments().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "batch/v1":
		switch kapi.Kind {

		case "Job":
			_, e = clientset.BatchV1().Jobs(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "rbac/v1beta1":
		switch kapi.Kind {

		case "ClusterRoleBinding":
			_, e = clientset.RbacV1beta1().ClusterRoleBindings().List(*opts)
			if e != nil {
				return e
			}

		case "ClusterRole":
			_, e = clientset.RbacV1beta1().ClusterRoles().List(*opts)
			if e != nil {
				return e
			}

		case "RoleBinding":
			_, e = clientset.RbacV1beta1().RoleBindings(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Role":
			_, e = clientset.RbacV1beta1().Roles(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1":
		switch kapi.Kind {

		case "StorageClass":
			_, e = clientset.StorageV1().StorageClasses().List(*opts)
			if e != nil {
				return e
			}

		case "VolumeAttachment":
			_, e = clientset.StorageV1().VolumeAttachments().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "auditregistration/v1alpha1":
		switch kapi.Kind {

		case "AuditSink":
			_, e = clientset.AuditregistrationV1alpha1().AuditSinks().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "authentication/v1beta1":
		switch kapi.Kind {

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "autoscaling/v2beta1":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			_, e = clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "scheduling/v1beta1":
		switch kapi.Kind {

		case "PriorityClass":
			_, e = clientset.SchedulingV1beta1().PriorityClasses().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1":
		switch kapi.Kind {

		case "DaemonSet":
			_, e = clientset.AppsV1().DaemonSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Deployment":
			_, e = clientset.AppsV1().Deployments(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ReplicaSet":
			_, e = clientset.AppsV1().ReplicaSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "StatefulSet":
			_, e = clientset.AppsV1().StatefulSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ControllerRevision":
			_, e = clientset.AppsV1().ControllerRevisions(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1beta1":
		switch kapi.Kind {

		case "Deployment":
			_, e = clientset.AppsV1beta1().Deployments(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "StatefulSet":
			_, e = clientset.AppsV1beta1().StatefulSets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ControllerRevision":
			_, e = clientset.AppsV1beta1().ControllerRevisions(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "authorization/v1":
		switch kapi.Kind {

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "autoscaling/v1":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			_, e = clientset.AutoscalingV1().HorizontalPodAutoscalers(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "autoscaling/v2beta2":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			_, e = clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "batch/v2alpha1":
		switch kapi.Kind {

		case "CronJob":
			_, e = clientset.BatchV2alpha1().CronJobs(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "v1":
		switch kapi.Kind {


		case "Node":
			_, e = clientset.CoreV1().Nodes().List(*opts)
			if e != nil {
				return e
			}

		case "Secret":
			_, e = clientset.CoreV1().Secrets(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ResourceQuota":
			_, e = clientset.CoreV1().ResourceQuotas(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ServiceAccount":
			_, e = clientset.CoreV1().ServiceAccounts(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ComponentStatus":
			_, e = clientset.CoreV1().ComponentStatuses().List(*opts)
			if e != nil {
				return e
			}

		case "ConfigMap":
			_, e = clientset.CoreV1().ConfigMaps(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Namespace":
			_, e = clientset.CoreV1().Namespaces().List(*opts)
			if e != nil {
				return e
			}

		case "PersistentVolume":
			_, e = clientset.CoreV1().PersistentVolumes().List(*opts)
			if e != nil {
				return e
			}

		case "PodTemplate":
			_, e = clientset.CoreV1().PodTemplates(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "PersistentVolumeClaim":
			_, e = clientset.CoreV1().PersistentVolumeClaims(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Service":
			_, e = clientset.CoreV1().Services(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Event":
			_, e = clientset.CoreV1().Events(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "LimitRange":
			_, e = clientset.CoreV1().LimitRanges(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "Pod":
			_, e = clientset.CoreV1().Pods(namespace).List(*opts)
			if e != nil {
				return e
			}

		case "ReplicationController":
			_, e = clientset.CoreV1().ReplicationControllers(namespace).List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1alpha1":
		switch kapi.Kind {

		case "VolumeAttachment":
			_, e = clientset.StorageV1alpha1().VolumeAttachments().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "admissionregistration/v1alpha1":
		switch kapi.Kind {

		case "InitializerConfiguration":
			_, e = clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().List(*opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	}

	return nil

}

