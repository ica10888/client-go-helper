package kubectl

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)




func (i *CronJob) Delete (opts *v1.DeleteOptions) (error) {
	e := clientset.BatchV1beta1().CronJobs(i.Namespace).Delete(i.Name,opts)
	if e != nil {
		return e
	}
	return nil
}



func Delete(opts *v1.DeleteOptions,kapi *Kubeapi ,name string ,namespace string) (error) {

	switch kapi.ApiVersion {
	case "batch/v1beta1":
		switch kapi.Kind {

		case "CronJob":
			e := clientset.BatchV1beta1().CronJobs(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "certificates/v1beta1":
		switch kapi.Kind {

		case "CertificateSigningRequest":
			e := clientset.CertificatesV1beta1().CertificateSigningRequests().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "coordination/v1beta1":
		switch kapi.Kind {

		case "Lease":
			e := clientset.CoordinationV1beta1().Leases(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "networking/v1":
		switch kapi.Kind {

		case "NetworkPolicy":
			e := clientset.NetworkingV1().NetworkPolicies(namespace).Delete(name,opts)
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

		case "Deployment":
			e := clientset.ExtensionsV1beta1().Deployments(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Ingress":
			e := clientset.ExtensionsV1beta1().Ingresses(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "PodSecurityPolicy":
			e := clientset.ExtensionsV1beta1().PodSecurityPolicies().Delete(name,opts)
			if e != nil {
				return e
			}

		case "ReplicaSet":
			e := clientset.ExtensionsV1beta1().ReplicaSets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "DaemonSet":
			e := clientset.ExtensionsV1beta1().DaemonSets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "rbac/v1":
		switch kapi.Kind {

		case "ClusterRoleBinding":
			e := clientset.RbacV1().ClusterRoleBindings().Delete(name,opts)
			if e != nil {
				return e
			}

		case "ClusterRole":
			e := clientset.RbacV1().ClusterRoles().Delete(name,opts)
			if e != nil {
				return e
			}

		case "RoleBinding":
			e := clientset.RbacV1().RoleBindings(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Role":
			e := clientset.RbacV1().Roles(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1beta2":
		switch kapi.Kind {

		case "Deployment":
			e := clientset.AppsV1beta2().Deployments(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ReplicaSet":
			e := clientset.AppsV1beta2().ReplicaSets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "StatefulSet":
			e := clientset.AppsV1beta2().StatefulSets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ControllerRevision":
			e := clientset.AppsV1beta2().ControllerRevisions(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "DaemonSet":
			e := clientset.AppsV1beta2().DaemonSets(namespace).Delete(name,opts)
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

		case "ClusterRoleBinding":
			e := clientset.RbacV1alpha1().ClusterRoleBindings().Delete(name,opts)
			if e != nil {
				return e
			}

		case "ClusterRole":
			e := clientset.RbacV1alpha1().ClusterRoles().Delete(name,opts)
			if e != nil {
				return e
			}

		case "RoleBinding":
			e := clientset.RbacV1alpha1().RoleBindings(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Role":
			e := clientset.RbacV1alpha1().Roles(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "settings/v1alpha1":
		switch kapi.Kind {

		case "PodPreset":
			e := clientset.SettingsV1alpha1().PodPresets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "admissionregistration/v1beta1":
		switch kapi.Kind {

		case "MutatingWebhookConfiguration":
			e := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Delete(name,opts)
			if e != nil {
				return e
			}

		case "ValidatingWebhookConfiguration":
			e := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "events/v1beta1":
		switch kapi.Kind {

		case "Event":
			e := clientset.EventsV1beta1().Events(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "policy/v1beta1":
		switch kapi.Kind {

		case "PodSecurityPolicy":
			e := clientset.PolicyV1beta1().PodSecurityPolicies().Delete(name,opts)
			if e != nil {
				return e
			}


		case "PodDisruptionBudget":
			e := clientset.PolicyV1beta1().PodDisruptionBudgets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "scheduling/v1alpha1":
		switch kapi.Kind {

		case "PriorityClass":
			e := clientset.SchedulingV1alpha1().PriorityClasses().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1beta1":
		switch kapi.Kind {

		case "StorageClass":
			e := clientset.StorageV1beta1().StorageClasses().Delete(name,opts)
			if e != nil {
				return e
			}

		case "VolumeAttachment":
			e := clientset.StorageV1beta1().VolumeAttachments().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "batch/v1":
		switch kapi.Kind {

		case "Job":
			e := clientset.BatchV1().Jobs(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "rbac/v1beta1":
		switch kapi.Kind {

		case "RoleBinding":
			e := clientset.RbacV1beta1().RoleBindings(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Role":
			e := clientset.RbacV1beta1().Roles(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ClusterRoleBinding":
			e := clientset.RbacV1beta1().ClusterRoleBindings().Delete(name,opts)
			if e != nil {
				return e
			}

		case "ClusterRole":
			e := clientset.RbacV1beta1().ClusterRoles().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1":
		switch kapi.Kind {

		case "StorageClass":
			e := clientset.StorageV1().StorageClasses().Delete(name,opts)
			if e != nil {
				return e
			}

		case "VolumeAttachment":
			e := clientset.StorageV1().VolumeAttachments().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "auditregistration/v1alpha1":
		switch kapi.Kind {

		case "AuditSink":
			e := clientset.AuditregistrationV1alpha1().AuditSinks().Delete(name,opts)
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
			e := clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "scheduling/v1beta1":
		switch kapi.Kind {

		case "PriorityClass":
			e := clientset.SchedulingV1beta1().PriorityClasses().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1":
		switch kapi.Kind {

		case "ReplicaSet":
			e := clientset.AppsV1().ReplicaSets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "StatefulSet":
			e := clientset.AppsV1().StatefulSets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ControllerRevision":
			e := clientset.AppsV1().ControllerRevisions(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "DaemonSet":
			e := clientset.AppsV1().DaemonSets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Deployment":
			e := clientset.AppsV1().Deployments(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1beta1":
		switch kapi.Kind {

		case "ControllerRevision":
			e := clientset.AppsV1beta1().ControllerRevisions(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Deployment":
			e := clientset.AppsV1beta1().Deployments(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "StatefulSet":
			e := clientset.AppsV1beta1().StatefulSets(namespace).Delete(name,opts)
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
			e := clientset.AutoscalingV1().HorizontalPodAutoscalers(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "autoscaling/v2beta2":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			e := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "batch/v2alpha1":
		switch kapi.Kind {

		case "CronJob":
			e := clientset.BatchV2alpha1().CronJobs(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "v1":
		switch kapi.Kind {

		case "Node":
			e := clientset.CoreV1().Nodes().Delete(name,opts)
			if e != nil {
				return e
			}

		case "Pod":
			e := clientset.CoreV1().Pods(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ResourceQuota":
			e := clientset.CoreV1().ResourceQuotas(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ComponentStatus":
			e := clientset.CoreV1().ComponentStatuses().Delete(name,opts)
			if e != nil {
				return e
			}

		case "LimitRange":
			e := clientset.CoreV1().LimitRanges(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Service":
			e := clientset.CoreV1().Services(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ConfigMap":
			e := clientset.CoreV1().ConfigMaps(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Namespace":
			e := clientset.CoreV1().Namespaces().Delete(name,opts)
			if e != nil {
				return e
			}

		case "PersistentVolumeClaim":
			e := clientset.CoreV1().PersistentVolumeClaims(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Secret":
			e := clientset.CoreV1().Secrets(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ServiceAccount":
			e := clientset.CoreV1().ServiceAccounts(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "PodTemplate":
			e := clientset.CoreV1().PodTemplates(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "ReplicationController":
			e := clientset.CoreV1().ReplicationControllers(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "Event":
			e := clientset.CoreV1().Events(namespace).Delete(name,opts)
			if e != nil {
				return e
			}

		case "PersistentVolume":
			e := clientset.CoreV1().PersistentVolumes().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1alpha1":
		switch kapi.Kind {

		case "VolumeAttachment":
			e := clientset.StorageV1alpha1().VolumeAttachments().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "admissionregistration/v1alpha1":
		switch kapi.Kind {

		case "InitializerConfiguration":
			e := clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().Delete(name,opts)
			if e != nil {
				return e
			}

		default:
			return fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}


	}


	return fmt.Errorf("Error")

}