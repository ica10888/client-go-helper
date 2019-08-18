package kubectl

import (
	"client-go-helper/pkg/kubectl/client"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *cronJob) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.BatchV1beta1().CronJobs(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *auditSink) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AuditregistrationV1alpha1().AuditSinks().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *mutatingWebhookConfiguration) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *validatingWebhookConfiguration) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *job) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.BatchV1().Jobs(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *certificateSigningRequest) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CertificatesV1beta1().CertificateSigningRequests().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *role) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.RbacV1().Roles(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *clusterRoleBinding) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.RbacV1().ClusterRoleBindings().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *clusterRole) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.RbacV1().ClusterRoles().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *roleBinding) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.RbacV1().RoleBindings(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *lease) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoordinationV1beta1().Leases(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *networkPolicy) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.NetworkingV1().NetworkPolicies(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *deployment) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().Deployments(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *replicaSet) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().ReplicaSets(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *statefulSet) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().StatefulSets(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *controllerRevision) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().ControllerRevisions(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *daemonSet) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().DaemonSets(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *horizontalPodAutoscaler) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *podDisruptionBudget) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.PolicyV1beta1().PodDisruptionBudgets(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *podSecurityPolicy) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.PolicyV1beta1().PodSecurityPolicies().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *configMap) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ConfigMaps(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *persistentVolume) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().PersistentVolumes().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *pod) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Pods(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *resourceQuota) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ResourceQuotas(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *componentStatus) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ComponentStatuses().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *replicationController) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ReplicationControllers(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *limitRange) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().LimitRanges(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *namespace) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Namespaces().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *service) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Services(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *event) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Events(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *node) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Nodes().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *persistentVolumeClaim) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().PersistentVolumeClaims(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *podTemplate) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().PodTemplates(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *secret) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Secrets(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *serviceAccount) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ServiceAccounts(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *podPreset) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.SettingsV1alpha1().PodPresets(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *initializerConfiguration) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ingress) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.ExtensionsV1beta1().Ingresses(i.Namespace).Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *priorityClass) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.SchedulingV1beta1().PriorityClasses().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *storageClass) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.StorageV1().StorageClasses().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *volumeAttachment) Delete(opts *v1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}
	err = clientset.StorageV1().VolumeAttachments().Delete(i.Name, opts)
	if err != nil {
		return err
	}
	return nil
}
