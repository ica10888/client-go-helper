package kubectl

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"encoding/json"
	"github.com/ghodss/yaml"
)



func Get(opts *v1.ListOptions,kapi *Kubeapi ,namespace string) ( []string, error) {
	// 校验和放入结构体
	clientset, e := InitClient()
	if e != nil {
		return nil,fmt.Errorf("something wrong happend ,%s", e)
	}

	switch kapi.ApiVersion {
	case "batch/v1beta1":
		switch kapi.Kind {

		case "CronJob":
			cronJobList, e := clientset.BatchV1beta1().CronJobs(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range cronJobList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "certificates/v1beta1":
		switch kapi.Kind {

		case "CertificateSigningRequest":
			certificateSigningRequestList, e := clientset.CertificatesV1beta1().CertificateSigningRequests().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range certificateSigningRequestList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "coordination/v1beta1":
		switch kapi.Kind {

		case "Lease":
			leaseList, e := clientset.CoordinationV1beta1().Leases(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range leaseList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "networking/v1":
		switch kapi.Kind {

		case "NetworkPolicy":
			networkPolicyList, e := clientset.NetworkingV1().NetworkPolicies(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range networkPolicyList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "authorization/v1beta1":
		switch kapi.Kind {


		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "extensions/v1beta1":
		switch kapi.Kind {

		case "Deployment":
			deploymentList, e := clientset.ExtensionsV1beta1().Deployments(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range deploymentList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Ingress":
			ingressList, e := clientset.ExtensionsV1beta1().Ingresses(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range ingressList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "PodSecurityPolicy":
			podSecurityPolicyList, e := clientset.ExtensionsV1beta1().PodSecurityPolicies().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range podSecurityPolicyList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ReplicaSet":
			replicaSetList, e := clientset.ExtensionsV1beta1().ReplicaSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range replicaSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "DaemonSet":
			daemonSetList, e := clientset.ExtensionsV1beta1().DaemonSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range daemonSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "rbac/v1":
		switch kapi.Kind {

		case "ClusterRoleBinding":
			clusterRoleBindingList, e := clientset.RbacV1().ClusterRoleBindings().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range clusterRoleBindingList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ClusterRole":
			clusterRoleList, e := clientset.RbacV1().ClusterRoles().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range clusterRoleList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "RoleBinding":
			roleBindingList, e := clientset.RbacV1().RoleBindings(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range roleBindingList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Role":
			roleList, e := clientset.RbacV1().Roles(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range roleList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1beta2":
		switch kapi.Kind {

		case "Deployment":
			deploymentList, e := clientset.AppsV1beta2().Deployments(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range deploymentList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ReplicaSet":
			replicaSetList, e := clientset.AppsV1beta2().ReplicaSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range replicaSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "StatefulSet":
			statefulSetList, e := clientset.AppsV1beta2().StatefulSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range statefulSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ControllerRevision":
			controllerRevisionList, e := clientset.AppsV1beta2().ControllerRevisions(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range controllerRevisionList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "DaemonSet":
			daemonSetList, e := clientset.AppsV1beta2().DaemonSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range daemonSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "authentication/v1":
		switch kapi.Kind {

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "rbac/v1alpha1":
		switch kapi.Kind {

		case "RoleBinding":
			roleBindingList, e := clientset.RbacV1alpha1().RoleBindings(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range roleBindingList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Role":
			roleList, e := clientset.RbacV1alpha1().Roles(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range roleList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ClusterRoleBinding":
			clusterRoleBindingList, e := clientset.RbacV1alpha1().ClusterRoleBindings().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range clusterRoleBindingList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ClusterRole":
			clusterRoleList, e := clientset.RbacV1alpha1().ClusterRoles().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range clusterRoleList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "settings/v1alpha1":
		switch kapi.Kind {

		case "PodPreset":
			podPresetList, e := clientset.SettingsV1alpha1().PodPresets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range podPresetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "admissionregistration/v1beta1":
		switch kapi.Kind {

		case "MutatingWebhookConfiguration":
			mutatingWebhookConfigurationList, e := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range mutatingWebhookConfigurationList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ValidatingWebhookConfiguration":
			validatingWebhookConfigurationList, e := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range validatingWebhookConfigurationList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "events/v1beta1":
		switch kapi.Kind {

		case "Event":
			eventList, e := clientset.EventsV1beta1().Events(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range eventList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "policy/v1beta1":
		switch kapi.Kind {


		case "PodDisruptionBudget":
			podDisruptionBudgetList, e := clientset.PolicyV1beta1().PodDisruptionBudgets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range podDisruptionBudgetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "PodSecurityPolicy":
			podSecurityPolicyList, e := clientset.PolicyV1beta1().PodSecurityPolicies().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range podSecurityPolicyList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "scheduling/v1alpha1":
		switch kapi.Kind {

		case "PriorityClass":
			priorityClassList, e := clientset.SchedulingV1alpha1().PriorityClasses().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range priorityClassList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1beta1":
		switch kapi.Kind {

		case "StorageClass":
			storageClassList, e := clientset.StorageV1beta1().StorageClasses().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range storageClassList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "VolumeAttachment":
			volumeAttachmentList, e := clientset.StorageV1beta1().VolumeAttachments().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range volumeAttachmentList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "batch/v1":
		switch kapi.Kind {

		case "Job":
			jobList, e := clientset.BatchV1().Jobs(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range jobList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "rbac/v1beta1":
		switch kapi.Kind {

		case "ClusterRoleBinding":
			clusterRoleBindingList, e := clientset.RbacV1beta1().ClusterRoleBindings().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range clusterRoleBindingList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ClusterRole":
			clusterRoleList, e := clientset.RbacV1beta1().ClusterRoles().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range clusterRoleList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "RoleBinding":
			roleBindingList, e := clientset.RbacV1beta1().RoleBindings(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range roleBindingList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Role":
			roleList, e := clientset.RbacV1beta1().Roles(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range roleList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1":
		switch kapi.Kind {

		case "StorageClass":
			storageClassList, e := clientset.StorageV1().StorageClasses().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range storageClassList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "VolumeAttachment":
			volumeAttachmentList, e := clientset.StorageV1().VolumeAttachments().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range volumeAttachmentList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "auditregistration/v1alpha1":
		switch kapi.Kind {

		case "AuditSink":
			auditSinkList, e := clientset.AuditregistrationV1alpha1().AuditSinks().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range auditSinkList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "authentication/v1beta1":
		switch kapi.Kind {


		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "autoscaling/v2beta1":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			horizontalPodAutoscalerList, e := clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range horizontalPodAutoscalerList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "scheduling/v1beta1":
		switch kapi.Kind {

		case "PriorityClass":
			priorityClassList, e := clientset.SchedulingV1beta1().PriorityClasses().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range priorityClassList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1":
		switch kapi.Kind {

		case "ReplicaSet":
			replicaSetList, e := clientset.AppsV1().ReplicaSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range replicaSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "StatefulSet":
			statefulSetList, e := clientset.AppsV1().StatefulSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range statefulSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ControllerRevision":
			controllerRevisionList, e := clientset.AppsV1().ControllerRevisions(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range controllerRevisionList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "DaemonSet":
			daemonSetList, e := clientset.AppsV1().DaemonSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range daemonSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Deployment":
			deploymentList, e := clientset.AppsV1().Deployments(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range deploymentList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "apps/v1beta1":
		switch kapi.Kind {

		case "StatefulSet":
			statefulSetList, e := clientset.AppsV1beta1().StatefulSets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range statefulSetList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ControllerRevision":
			controllerRevisionList, e := clientset.AppsV1beta1().ControllerRevisions(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range controllerRevisionList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Deployment":
			deploymentList, e := clientset.AppsV1beta1().Deployments(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range deploymentList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "authorization/v1":
		switch kapi.Kind {

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "autoscaling/v1":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			horizontalPodAutoscalerList, e := clientset.AutoscalingV1().HorizontalPodAutoscalers(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range horizontalPodAutoscalerList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "autoscaling/v2beta2":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			horizontalPodAutoscalerList, e := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range horizontalPodAutoscalerList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "batch/v2alpha1":
		switch kapi.Kind {

		case "CronJob":
			cronJobList, e := clientset.BatchV2alpha1().CronJobs(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range cronJobList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "v1":
		switch kapi.Kind {

		case "ComponentStatus":
			componentStatusList, e := clientset.CoreV1().ComponentStatuses().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range componentStatusList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Namespace":
			namespaceList, e := clientset.CoreV1().Namespaces().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range namespaceList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ConfigMap":
			configMapList, e := clientset.CoreV1().ConfigMaps(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range configMapList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil


		case "Event":
			eventList, e := clientset.CoreV1().Events(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range eventList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Pod":
			podList, e := clientset.CoreV1().Pods(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range podList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "PersistentVolumeClaim":
			persistentVolumeClaimList, e := clientset.CoreV1().PersistentVolumeClaims(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range persistentVolumeClaimList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "PersistentVolume":
			persistentVolumeList, e := clientset.CoreV1().PersistentVolumes().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range persistentVolumeList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ReplicationController":
			replicationControllerList, e := clientset.CoreV1().ReplicationControllers(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range replicationControllerList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Service":
			serviceList, e := clientset.CoreV1().Services(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range serviceList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Secret":
			secretList, e := clientset.CoreV1().Secrets(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range secretList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ServiceAccount":
			serviceAccountList, e := clientset.CoreV1().ServiceAccounts(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range serviceAccountList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "LimitRange":
			limitRangeList, e := clientset.CoreV1().LimitRanges(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range limitRangeList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "Node":
			nodeList, e := clientset.CoreV1().Nodes().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range nodeList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "PodTemplate":
			podTemplateList, e := clientset.CoreV1().PodTemplates(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range podTemplateList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		case "ResourceQuota":
			resourceQuotaList, e := clientset.CoreV1().ResourceQuotas(namespace).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range resourceQuotaList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "storage/v1alpha1":
		switch kapi.Kind {

		case "VolumeAttachment":
			volumeAttachmentList, e := clientset.StorageV1alpha1().VolumeAttachments().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range volumeAttachmentList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}

	case "admissionregistration/v1alpha1":
		switch kapi.Kind {

		case "InitializerConfiguration":
			initializerConfigurationList, e := clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range initializerConfigurationList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil

		default:
			return nil,fmt.Errorf("not support a kind : %s in  apiVersion: %s",kapi.Kind,kapi.ApiVersion)
		}
	}

	return nil,fmt.Errorf("Error")

}

func List(opts *v1.ListOptions,kapi *Kubeapi ,namespace string) ( []string, error) {
	return Get(opts,kapi,namespace)

}