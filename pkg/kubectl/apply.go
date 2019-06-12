package kubectl


import (
"fmt"
yaml2 "gopkg.in/yaml.v2"

admissionregistrationV1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
admissionregistrationV1beta1 "k8s.io/api/admissionregistration/v1beta1"
appsV1 "k8s.io/api/apps/v1"
appsV1beta1 "k8s.io/api/apps/v1beta1"
appsV1beta2 "k8s.io/api/apps/v1beta2"
auditregistrationV1alpha1 "k8s.io/api/auditregistration/v1alpha1"
authenticationV1 "k8s.io/api/authentication/v1"
authenticationV1beta1 "k8s.io/api/authentication/v1beta1"
authorizationV1 "k8s.io/api/authorization/v1"
authorizationV1beta1 "k8s.io/api/authorization/v1beta1"
autoscalingV1 "k8s.io/api/autoscaling/v1"
autoscalingV2beta1 "k8s.io/api/autoscaling/v2beta1"
autoscalingV2beta2 "k8s.io/api/autoscaling/v2beta2"
batchV1 "k8s.io/api/batch/v1"
batchV1beta1 "k8s.io/api/batch/v1beta1"
batchV2alpha1 "k8s.io/api/batch/v2alpha1"
certificatesV1beta1 "k8s.io/api/certificates/v1beta1"
coordinationV1beta1 "k8s.io/api/coordination/v1beta1"
coreV1 "k8s.io/api/core/v1"
eventsV1beta1 "k8s.io/api/events/v1beta1"
extensionsV1beta1 "k8s.io/api/extensions/v1beta1"
networkingV1 "k8s.io/api/networking/v1"
policyV1beta1 "k8s.io/api/policy/v1beta1"
rbacV1 "k8s.io/api/rbac/v1"
rbacV1alpha1 "k8s.io/api/rbac/v1alpha1"
rbacV1beta1 "k8s.io/api/rbac/v1beta1"
schedulingV1alpha1 "k8s.io/api/scheduling/v1alpha1"
schedulingV1beta1 "k8s.io/api/scheduling/v1beta1"
settingsV1alpha1 "k8s.io/api/settings/v1alpha1"
storageV1 "k8s.io/api/storage/v1"
storageV1alpha1 "k8s.io/api/storage/v1alpha1"
storageV1beta1 "k8s.io/api/storage/v1beta1"

apierrs "k8s.io/apimachinery/pkg/api/errors"
"k8s.io/client-go/kubernetes/scheme"
)

type kubeapi struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Yaml       string
}

func Apply(yaml string, namespace string) (error) {
	// 校验和放入结构体
	kapi := kubeapi{}
	yaml2.Unmarshal([]byte(yaml), &kapi)
	kapi.Yaml = yaml

	//序列化
	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode([]byte(yaml), nil, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	clientset, e := InitClient()
	if e != nil {
		fmt.Errorf("something wrong happend ,%s", e)
	}

	switch kapi.ApiVersion {
	case "batch/v1beta1":
		switch kapi.Kind {

		case "CronJob":
			cronJob := obj.(*batchV1beta1.CronJob)
			_, e = clientset.BatchV1beta1().CronJobs(namespace).Create(cronJob)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.BatchV1beta1().CronJobs(namespace).Update(cronJob)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "certificates/v1beta1":
		switch kapi.Kind {

		case "CertificateSigningRequest":
			certificateSigningRequest := obj.(*certificatesV1beta1.CertificateSigningRequest)
			_, e = clientset.CertificatesV1beta1().CertificateSigningRequests().Create(certificateSigningRequest)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CertificatesV1beta1().CertificateSigningRequests().Update(certificateSigningRequest)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "coordination/v1beta1":
		switch kapi.Kind {

		case "Lease":
			lease := obj.(*coordinationV1beta1.Lease)
			_, e = clientset.CoordinationV1beta1().Leases(namespace).Create(lease)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoordinationV1beta1().Leases(namespace).Update(lease)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "networking/v1":
		switch kapi.Kind {

		case "NetworkPolicy":
			networkPolicy := obj.(*networkingV1.NetworkPolicy)
			_, e = clientset.NetworkingV1().NetworkPolicies(namespace).Create(networkPolicy)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.NetworkingV1().NetworkPolicies(namespace).Update(networkPolicy)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "authorization/v1beta1":
		switch kapi.Kind {

		case "SelfSubjectAccessReview":
			selfSubjectAccessReview := obj.(*authorizationV1beta1.SelfSubjectAccessReview)
			_, e = clientset.AuthorizationV1beta1().SelfSubjectAccessReviews().Create(selfSubjectAccessReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthorizationV1beta1().SelfSubjectAccessReviews().Update(selfSubjectAccessReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "SelfSubjectRulesReview":
			selfSubjectRulesReview := obj.(*authorizationV1beta1.SelfSubjectRulesReview)
			_, e = clientset.AuthorizationV1beta1().SelfSubjectRulesReviews().Create(selfSubjectRulesReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthorizationV1beta1().SelfSubjectRulesReviews().Update(selfSubjectRulesReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "SubjectAccessReview":
			subjectAccessReview := obj.(*authorizationV1beta1.SubjectAccessReview)
			_, e = clientset.AuthorizationV1beta1().SubjectAccessReviews().Create(subjectAccessReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthorizationV1beta1().SubjectAccessReviews().Update(subjectAccessReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "LocalSubjectAccessReview":
			localSubjectAccessReview := obj.(*authorizationV1beta1.LocalSubjectAccessReview)
			_, e = clientset.AuthorizationV1beta1().LocalSubjectAccessReviews(namespace).Create(localSubjectAccessReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthorizationV1beta1().LocalSubjectAccessReviews(namespace).Update(localSubjectAccessReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "extensions/v1beta1":
		switch kapi.Kind {

		case "DaemonSet":
			daemonSet := obj.(*extensionsV1beta1.DaemonSet)
			_, e = clientset.ExtensionsV1beta1().DaemonSets(namespace).Create(daemonSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.ExtensionsV1beta1().DaemonSets(namespace).Update(daemonSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Deployment":
			deployment := obj.(*extensionsV1beta1.Deployment)
			_, e = clientset.ExtensionsV1beta1().Deployments(namespace).Create(deployment)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.ExtensionsV1beta1().Deployments(namespace).Update(deployment)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Ingress":
			ingress := obj.(*extensionsV1beta1.Ingress)
			_, e = clientset.ExtensionsV1beta1().Ingresses(namespace).Create(ingress)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.ExtensionsV1beta1().Ingresses(namespace).Update(ingress)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "PodSecurityPolicy":
			podSecurityPolicy := obj.(*extensionsV1beta1.PodSecurityPolicy)
			_, e = clientset.ExtensionsV1beta1().PodSecurityPolicies().Create(podSecurityPolicy)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.ExtensionsV1beta1().PodSecurityPolicies().Update(podSecurityPolicy)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ReplicaSet":
			replicaSet := obj.(*extensionsV1beta1.ReplicaSet)
			_, e = clientset.ExtensionsV1beta1().ReplicaSets(namespace).Create(replicaSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.ExtensionsV1beta1().ReplicaSets(namespace).Update(replicaSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "rbac/v1":
		switch kapi.Kind {

		case "ClusterRole":
			clusterRole := obj.(*rbacV1.ClusterRole)
			_, e = clientset.RbacV1().ClusterRoles().Create(clusterRole)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1().ClusterRoles().Update(clusterRole)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "RoleBinding":
			roleBinding := obj.(*rbacV1.RoleBinding)
			_, e = clientset.RbacV1().RoleBindings(namespace).Create(roleBinding)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1().RoleBindings(namespace).Update(roleBinding)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Role":
			role := obj.(*rbacV1.Role)
			_, e = clientset.RbacV1().Roles(namespace).Create(role)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1().Roles(namespace).Update(role)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ClusterRoleBinding":
			clusterRoleBinding := obj.(*rbacV1.ClusterRoleBinding)
			_, e = clientset.RbacV1().ClusterRoleBindings().Create(clusterRoleBinding)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1().ClusterRoleBindings().Update(clusterRoleBinding)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "apps/v1beta2":
		switch kapi.Kind {

		case "DaemonSet":
			daemonSet := obj.(*appsV1beta2.DaemonSet)
			_, e = clientset.AppsV1beta2().DaemonSets(namespace).Create(daemonSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1beta2().DaemonSets(namespace).Update(daemonSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Deployment":
			deployment := obj.(*appsV1beta2.Deployment)
			_, e = clientset.AppsV1beta2().Deployments(namespace).Create(deployment)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1beta2().Deployments(namespace).Update(deployment)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ReplicaSet":
			replicaSet := obj.(*appsV1beta2.ReplicaSet)
			_, e = clientset.AppsV1beta2().ReplicaSets(namespace).Create(replicaSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1beta2().ReplicaSets(namespace).Update(replicaSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "StatefulSet":
			statefulSet := obj.(*appsV1beta2.StatefulSet)
			_, e = clientset.AppsV1beta2().StatefulSets(namespace).Create(statefulSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1beta2().StatefulSets(namespace).Update(statefulSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ControllerRevision":
			controllerRevision := obj.(*appsV1beta2.ControllerRevision)
			_, e = clientset.AppsV1beta2().ControllerRevisions(namespace).Create(controllerRevision)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1beta2().ControllerRevisions(namespace).Update(controllerRevision)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "authentication/v1":
		switch kapi.Kind {

		case "TokenReview":
			tokenReview := obj.(*authenticationV1.TokenReview)
			_, e = clientset.AuthenticationV1().TokenReviews().Create(tokenReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthenticationV1().TokenReviews().Update(tokenReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "rbac/v1alpha1":
		switch kapi.Kind {

		case "Role":
			role := obj.(*rbacV1alpha1.Role)
			_, e = clientset.RbacV1alpha1().Roles(namespace).Create(role)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1alpha1().Roles(namespace).Update(role)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ClusterRoleBinding":
			clusterRoleBinding := obj.(*rbacV1alpha1.ClusterRoleBinding)
			_, e = clientset.RbacV1alpha1().ClusterRoleBindings().Create(clusterRoleBinding)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1alpha1().ClusterRoleBindings().Update(clusterRoleBinding)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ClusterRole":
			clusterRole := obj.(*rbacV1alpha1.ClusterRole)
			_, e = clientset.RbacV1alpha1().ClusterRoles().Create(clusterRole)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1alpha1().ClusterRoles().Update(clusterRole)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "RoleBinding":
			roleBinding := obj.(*rbacV1alpha1.RoleBinding)
			_, e = clientset.RbacV1alpha1().RoleBindings(namespace).Create(roleBinding)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1alpha1().RoleBindings(namespace).Update(roleBinding)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "settings/v1alpha1":
		switch kapi.Kind {

		case "PodPreset":
			podPreset := obj.(*settingsV1alpha1.PodPreset)
			_, e = clientset.SettingsV1alpha1().PodPresets(namespace).Create(podPreset)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.SettingsV1alpha1().PodPresets(namespace).Update(podPreset)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "admissionregistration/v1beta1":
		switch kapi.Kind {

		case "MutatingWebhookConfiguration":
			mutatingWebhookConfiguration := obj.(*admissionregistrationV1beta1.MutatingWebhookConfiguration)
			_, e = clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Create(mutatingWebhookConfiguration)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Update(mutatingWebhookConfiguration)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ValidatingWebhookConfiguration":
			validatingWebhookConfiguration := obj.(*admissionregistrationV1beta1.ValidatingWebhookConfiguration)
			_, e = clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().Create(validatingWebhookConfiguration)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().Update(validatingWebhookConfiguration)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "events/v1beta1":
		switch kapi.Kind {

		case "Event":
			event := obj.(*eventsV1beta1.Event)
			_, e = clientset.EventsV1beta1().Events(namespace).Create(event)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.EventsV1beta1().Events(namespace).Update(event)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "policy/v1beta1":
		switch kapi.Kind {

		case "PodDisruptionBudget":
			podDisruptionBudget := obj.(*policyV1beta1.PodDisruptionBudget)
			_, e = clientset.PolicyV1beta1().PodDisruptionBudgets(namespace).Create(podDisruptionBudget)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.PolicyV1beta1().PodDisruptionBudgets(namespace).Update(podDisruptionBudget)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "PodSecurityPolicy":
			podSecurityPolicy := obj.(*policyV1beta1.PodSecurityPolicy)
			_, e = clientset.PolicyV1beta1().PodSecurityPolicies().Create(podSecurityPolicy)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.PolicyV1beta1().PodSecurityPolicies().Update(podSecurityPolicy)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "scheduling/v1alpha1":
		switch kapi.Kind {

		case "PriorityClass":
			priorityClass := obj.(*schedulingV1alpha1.PriorityClass)
			_, e = clientset.SchedulingV1alpha1().PriorityClasses().Create(priorityClass)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.SchedulingV1alpha1().PriorityClasses().Update(priorityClass)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "storage/v1beta1":
		switch kapi.Kind {

		case "StorageClass":
			storageClass := obj.(*storageV1beta1.StorageClass)
			_, e = clientset.StorageV1beta1().StorageClasses().Create(storageClass)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.StorageV1beta1().StorageClasses().Update(storageClass)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "VolumeAttachment":
			volumeAttachment := obj.(*storageV1beta1.VolumeAttachment)
			_, e = clientset.StorageV1beta1().VolumeAttachments().Create(volumeAttachment)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.StorageV1beta1().VolumeAttachments().Update(volumeAttachment)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "batch/v1":
		switch kapi.Kind {

		case "Job":
			job := obj.(*batchV1.Job)
			_, e = clientset.BatchV1().Jobs(namespace).Create(job)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.BatchV1().Jobs(namespace).Update(job)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "rbac/v1beta1":
		switch kapi.Kind {

		case "ClusterRoleBinding":
			clusterRoleBinding := obj.(*rbacV1beta1.ClusterRoleBinding)
			_, e = clientset.RbacV1beta1().ClusterRoleBindings().Create(clusterRoleBinding)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1beta1().ClusterRoleBindings().Update(clusterRoleBinding)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ClusterRole":
			clusterRole := obj.(*rbacV1beta1.ClusterRole)
			_, e = clientset.RbacV1beta1().ClusterRoles().Create(clusterRole)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1beta1().ClusterRoles().Update(clusterRole)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "RoleBinding":
			roleBinding := obj.(*rbacV1beta1.RoleBinding)
			_, e = clientset.RbacV1beta1().RoleBindings(namespace).Create(roleBinding)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1beta1().RoleBindings(namespace).Update(roleBinding)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Role":
			role := obj.(*rbacV1beta1.Role)
			_, e = clientset.RbacV1beta1().Roles(namespace).Create(role)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.RbacV1beta1().Roles(namespace).Update(role)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "storage/v1":
		switch kapi.Kind {

		case "StorageClass":
			storageClass := obj.(*storageV1.StorageClass)
			_, e = clientset.StorageV1().StorageClasses().Create(storageClass)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.StorageV1().StorageClasses().Update(storageClass)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "VolumeAttachment":
			volumeAttachment := obj.(*storageV1.VolumeAttachment)
			_, e = clientset.StorageV1().VolumeAttachments().Create(volumeAttachment)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.StorageV1().VolumeAttachments().Update(volumeAttachment)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "auditregistration/v1alpha1":
		switch kapi.Kind {

		case "AuditSink":
			auditSink := obj.(*auditregistrationV1alpha1.AuditSink)
			_, e = clientset.AuditregistrationV1alpha1().AuditSinks().Create(auditSink)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AuditregistrationV1alpha1().AuditSinks().Update(auditSink)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "authentication/v1beta1":
		switch kapi.Kind {

		case "TokenReview":
			tokenReview := obj.(*authenticationV1beta1.TokenReview)
			_, e = clientset.AuthenticationV1beta1().TokenReviews().Create(tokenReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthenticationV1beta1().TokenReviews().Update(tokenReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "autoscaling/v2beta1":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			horizontalPodAutoscaler := obj.(*autoscalingV2beta1.HorizontalPodAutoscaler)
			_, e = clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).Create(horizontalPodAutoscaler)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AutoscalingV2beta1().HorizontalPodAutoscalers(namespace).Update(horizontalPodAutoscaler)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "scheduling/v1beta1":
		switch kapi.Kind {

		case "PriorityClass":
			priorityClass := obj.(*schedulingV1beta1.PriorityClass)
			_, e = clientset.SchedulingV1beta1().PriorityClasses().Create(priorityClass)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.SchedulingV1beta1().PriorityClasses().Update(priorityClass)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "apps/v1":
		switch kapi.Kind {

		case "DaemonSet":
			daemonSet := obj.(*appsV1.DaemonSet)
			_, e = clientset.AppsV1().DaemonSets(namespace).Create(daemonSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1().DaemonSets(namespace).Update(daemonSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Deployment":
			deployment := obj.(*appsV1.Deployment)
			_, e = clientset.AppsV1().Deployments(namespace).Create(deployment)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1().Deployments(namespace).Update(deployment)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ReplicaSet":
			replicaSet := obj.(*appsV1.ReplicaSet)
			_, e = clientset.AppsV1().ReplicaSets(namespace).Create(replicaSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1().ReplicaSets(namespace).Update(replicaSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "StatefulSet":
			statefulSet := obj.(*appsV1.StatefulSet)
			_, e = clientset.AppsV1().StatefulSets(namespace).Create(statefulSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1().StatefulSets(namespace).Update(statefulSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ControllerRevision":
			controllerRevision := obj.(*appsV1.ControllerRevision)
			_, e = clientset.AppsV1().ControllerRevisions(namespace).Create(controllerRevision)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1().ControllerRevisions(namespace).Update(controllerRevision)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "apps/v1beta1":
		switch kapi.Kind {

		case "ControllerRevision":
			controllerRevision := obj.(*appsV1beta1.ControllerRevision)
			_, e = clientset.AppsV1beta1().ControllerRevisions(namespace).Create(controllerRevision)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1beta1().ControllerRevisions(namespace).Update(controllerRevision)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Deployment":
			deployment := obj.(*appsV1beta1.Deployment)
			_, e = clientset.AppsV1beta1().Deployments(namespace).Create(deployment)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1beta1().Deployments(namespace).Update(deployment)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "StatefulSet":
			statefulSet := obj.(*appsV1beta1.StatefulSet)
			_, e = clientset.AppsV1beta1().StatefulSets(namespace).Create(statefulSet)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AppsV1beta1().StatefulSets(namespace).Update(statefulSet)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "authorization/v1":
		switch kapi.Kind {

		case "SelfSubjectRulesReview":
			selfSubjectRulesReview := obj.(*authorizationV1.SelfSubjectRulesReview)
			_, e = clientset.AuthorizationV1().SelfSubjectRulesReviews().Create(selfSubjectRulesReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthorizationV1().SelfSubjectRulesReviews().Update(selfSubjectRulesReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "SubjectAccessReview":
			subjectAccessReview := obj.(*authorizationV1.SubjectAccessReview)
			_, e = clientset.AuthorizationV1().SubjectAccessReviews().Create(subjectAccessReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthorizationV1().SubjectAccessReviews().Update(subjectAccessReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "LocalSubjectAccessReview":
			localSubjectAccessReview := obj.(*authorizationV1.LocalSubjectAccessReview)
			_, e = clientset.AuthorizationV1().LocalSubjectAccessReviews(namespace).Create(localSubjectAccessReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthorizationV1().LocalSubjectAccessReviews(namespace).Update(localSubjectAccessReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "SelfSubjectAccessReview":
			selfSubjectAccessReview := obj.(*authorizationV1.SelfSubjectAccessReview)
			_, e = clientset.AuthorizationV1().SelfSubjectAccessReviews().Create(selfSubjectAccessReview)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					//_ ,e := clientset.AuthorizationV1().SelfSubjectAccessReviews().Update(selfSubjectAccessReview)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "autoscaling/v1":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			horizontalPodAutoscaler := obj.(*autoscalingV1.HorizontalPodAutoscaler)
			_, e = clientset.AutoscalingV1().HorizontalPodAutoscalers(namespace).Create(horizontalPodAutoscaler)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AutoscalingV1().HorizontalPodAutoscalers(namespace).Update(horizontalPodAutoscaler)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "autoscaling/v2beta2":
		switch kapi.Kind {

		case "HorizontalPodAutoscaler":
			horizontalPodAutoscaler := obj.(*autoscalingV2beta2.HorizontalPodAutoscaler)
			_, e = clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace).Create(horizontalPodAutoscaler)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(namespace).Update(horizontalPodAutoscaler)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "batch/v2alpha1":
		switch kapi.Kind {

		case "CronJob":
			cronJob := obj.(*batchV2alpha1.CronJob)
			_, e = clientset.BatchV2alpha1().CronJobs(namespace).Create(cronJob)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.BatchV2alpha1().CronJobs(namespace).Update(cronJob)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "v1":
		switch kapi.Kind {

		case "ResourceQuota":
			resourceQuota := obj.(*coreV1.ResourceQuota)
			_, e = clientset.CoreV1().ResourceQuotas(namespace).Create(resourceQuota)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().ResourceQuotas(namespace).Update(resourceQuota)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ServiceAccount":
			serviceAccount := obj.(*coreV1.ServiceAccount)
			_, e = clientset.CoreV1().ServiceAccounts(namespace).Create(serviceAccount)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().ServiceAccounts(namespace).Update(serviceAccount)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ComponentStatus":
			componentStatus := obj.(*coreV1.ComponentStatus)
			_, e = clientset.CoreV1().ComponentStatuses().Create(componentStatus)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().ComponentStatuses().Update(componentStatus)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Node":
			node := obj.(*coreV1.Node)
			_, e = clientset.CoreV1().Nodes().Create(node)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().Nodes().Update(node)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "PersistentVolume":
			persistentVolume := obj.(*coreV1.PersistentVolume)
			_, e = clientset.CoreV1().PersistentVolumes().Create(persistentVolume)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().PersistentVolumes().Update(persistentVolume)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Pod":
			pod := obj.(*coreV1.Pod)
			_, e = clientset.CoreV1().Pods(namespace).Create(pod)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().Pods(namespace).Update(pod)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ReplicationController":
			replicationController := obj.(*coreV1.ReplicationController)
			_, e = clientset.CoreV1().ReplicationControllers(namespace).Create(replicationController)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().ReplicationControllers(namespace).Update(replicationController)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Service":
			service := obj.(*coreV1.Service)
			_, e = clientset.CoreV1().Services(namespace).Create(service)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().Services(namespace).Update(service)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Endpoints":
			endpoints := obj.(*coreV1.Endpoints)
			_, e = clientset.CoreV1().Endpoints(namespace).Create(endpoints)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().Endpoints(namespace).Update(endpoints)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Event":
			event := obj.(*coreV1.Event)
			_, e = clientset.CoreV1().Events(namespace).Create(event)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().Events(namespace).Update(event)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Namespace":
			namespace := obj.(*coreV1.Namespace)
			_, e = clientset.CoreV1().Namespaces().Create(namespace)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().Namespaces().Update(namespace)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "PersistentVolumeClaim":
			persistentVolumeClaim := obj.(*coreV1.PersistentVolumeClaim)
			_, e = clientset.CoreV1().PersistentVolumeClaims(namespace).Create(persistentVolumeClaim)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().PersistentVolumeClaims(namespace).Update(persistentVolumeClaim)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "PodTemplate":
			podTemplate := obj.(*coreV1.PodTemplate)
			_, e = clientset.CoreV1().PodTemplates(namespace).Create(podTemplate)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().PodTemplates(namespace).Update(podTemplate)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "ConfigMap":
			configMap := obj.(*coreV1.ConfigMap)
			_, e = clientset.CoreV1().ConfigMaps(namespace).Create(configMap)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().ConfigMaps(namespace).Update(configMap)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "LimitRange":
			limitRange := obj.(*coreV1.LimitRange)
			_, e = clientset.CoreV1().LimitRanges(namespace).Create(limitRange)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().LimitRanges(namespace).Update(limitRange)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		case "Secret":
			secret := obj.(*coreV1.Secret)
			_, e = clientset.CoreV1().Secrets(namespace).Create(secret)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.CoreV1().Secrets(namespace).Update(secret)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "storage/v1alpha1":
		switch kapi.Kind {

		case "VolumeAttachment":
			volumeAttachment := obj.(*storageV1alpha1.VolumeAttachment)
			_, e = clientset.StorageV1alpha1().VolumeAttachments().Create(volumeAttachment)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.StorageV1alpha1().VolumeAttachments().Update(volumeAttachment)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}

	case "admissionregistration/v1alpha1":
		switch kapi.Kind {

		case "InitializerConfiguration":
			initializerConfiguration := obj.(*admissionregistrationV1alpha1.InitializerConfiguration)
			_, e = clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().Create(initializerConfiguration)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_, e := clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().Update(initializerConfiguration)
					if e != nil {
						return err
					}
				} else {
					return e
				}
			}

		default:
			fmt.Printf("not support a kind : %s in  apiVersion: %s", kapi.Kind, kapi.ApiVersion)
		}
	}

	return nil

}
