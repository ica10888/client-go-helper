package kubectl

import "k8s.io/apimachinery/pkg/apis/meta/v1"

func (i *CronJob) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.BatchV1beta1().CronJobs(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *AuditSink) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AuditregistrationV1alpha1().AuditSinks().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *MutatingWebhookConfiguration) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ValidatingWebhookConfiguration) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Job) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.BatchV1().Jobs(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *CertificateSigningRequest) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CertificatesV1beta1().CertificateSigningRequests().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *RoleBinding) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.RbacV1().RoleBindings(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Role) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.RbacV1().Roles(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ClusterRoleBinding) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.RbacV1().ClusterRoleBindings().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ClusterRole) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.RbacV1().ClusterRoles().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Lease) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoordinationV1beta1().Leases(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *NetworkPolicy) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.NetworkingV1().NetworkPolicies(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *StatefulSet) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AppsV1().StatefulSets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ControllerRevision) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AppsV1().ControllerRevisions(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *DaemonSet) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AppsV1().DaemonSets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Deployment) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AppsV1().Deployments(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ReplicaSet) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AppsV1().ReplicaSets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *HorizontalPodAutoscaler) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PodSecurityPolicy) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.PolicyV1beta1().PodSecurityPolicies().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}



func (i *PodDisruptionBudget) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.PolicyV1beta1().PodDisruptionBudgets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}




func (i *Service) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().Services(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PersistentVolume) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().PersistentVolumes().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PodTemplate) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().PodTemplates(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ReplicationController) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().ReplicationControllers(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ResourceQuota) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().ResourceQuotas(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ComponentStatus) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().ComponentStatuses().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Namespace) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().Namespaces().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PersistentVolumeClaim) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().PersistentVolumeClaims(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ConfigMap) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().ConfigMaps(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Node) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().Nodes().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Pod) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().Pods(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Secret) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().Secrets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ServiceAccount) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().ServiceAccounts(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}


func (i *Event) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().Events(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *LimitRange) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.CoreV1().LimitRanges(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PodPreset) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.SettingsV1alpha1().PodPresets(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *InitializerConfiguration) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}


func (i *Ingress) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.ExtensionsV1beta1().Ingresses(i.Namespace).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}


func (i *PriorityClass) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.SchedulingV1beta1().PriorityClasses().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *StorageClass) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.StorageV1().StorageClasses().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *VolumeAttachment) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.StorageV1().VolumeAttachments().Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}
