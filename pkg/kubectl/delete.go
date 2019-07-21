package kubectl

import "k8s.io/apimachinery/pkg/apis/meta/v1"

func (i *CronJob) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.BatchV1beta1().CronJobs(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *AuditSink) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AuditregistrationV1alpha1().AuditSinks().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *MutatingWebhookConfiguration) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ValidatingWebhookConfiguration) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AdmissionregistrationV1beta1().ValidatingWebhookConfigurations().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Job) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.BatchV1().Jobs(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *CertificateSigningRequest) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CertificatesV1beta1().CertificateSigningRequests().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Role) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.RbacV1().Roles(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ClusterRoleBinding) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.RbacV1().ClusterRoleBindings().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ClusterRole) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.RbacV1().ClusterRoles().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *RoleBinding) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.RbacV1().RoleBindings(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Lease) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoordinationV1beta1().Leases(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *NetworkPolicy) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.NetworkingV1().NetworkPolicies(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Deployment) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().Deployments(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ReplicaSet) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().ReplicaSets(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *StatefulSet) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().StatefulSets(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ControllerRevision) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().ControllerRevisions(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *DaemonSet) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AppsV1().DaemonSets(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *HorizontalPodAutoscaler) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AutoscalingV2beta2().HorizontalPodAutoscalers(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}


func (i *PodDisruptionBudget) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.PolicyV1beta1().PodDisruptionBudgets(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PodSecurityPolicy) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.PolicyV1beta1().PodSecurityPolicies().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}


func (i *ConfigMap) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ConfigMaps(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PersistentVolume) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().PersistentVolumes().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Pod) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Pods(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ResourceQuota) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ResourceQuotas(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ComponentStatus) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ComponentStatuses().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ReplicationController) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ReplicationControllers(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *LimitRange) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().LimitRanges(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Namespace) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Namespaces().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Service) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Services(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Event) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Events(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Node) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Nodes().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PersistentVolumeClaim) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().PersistentVolumeClaims(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *PodTemplate) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().PodTemplates(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *Secret) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().Secrets(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *ServiceAccount) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.CoreV1().ServiceAccounts(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}


func (i *PodPreset) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.SettingsV1alpha1().PodPresets(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *InitializerConfiguration) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.AdmissionregistrationV1alpha1().InitializerConfigurations().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}


func (i *Ingress) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.ExtensionsV1beta1().Ingresses(i.Namespace).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}


func (i *PriorityClass) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.SchedulingV1beta1().PriorityClasses().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *StorageClass) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.StorageV1().StorageClasses().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}

func (i *VolumeAttachment) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.StorageV1().VolumeAttachments().Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}