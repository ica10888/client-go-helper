package test

import (
	"client-go-helper/pkg/kubectl"
	"fmt"
	"testing"
)

func TestWriteDeleteCode(t *testing.T)  {

	clientset, _ := kubectl.InitClient()

	//输出转输入
	writeDeleteCode(GetFunctionName2(clientset.BatchV1beta1()),"BatchV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.AuditregistrationV1alpha1()),"AuditregistrationV1alpha1")

	writeDeleteCode(GetFunctionName2(clientset.AdmissionregistrationV1beta1()),"AdmissionregistrationV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.BatchV1()),"BatchV1")

	writeDeleteCode(GetFunctionName2(clientset.CertificatesV1beta1()),"CertificatesV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.RbacV1()),"RbacV1")

	writeDeleteCode(GetFunctionName2(clientset.CoordinationV1beta1()),"CoordinationV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.NetworkingV1()),"NetworkingV1")

	writeDeleteCode(GetFunctionName2(clientset.AppsV1()),"AppsV1")

	writeDeleteCode(GetFunctionName2(clientset.AutoscalingV2beta2()),"AutoscalingV2beta2")

	writeDeleteCode(GetFunctionName2(clientset.PolicyV1beta1()),"PolicyV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.AuthorizationV1()),"AuthorizationV1")

	writeDeleteCode(GetFunctionName2(clientset.CoreV1()),"CoreV1")

	writeDeleteCode(GetFunctionName2(clientset.SettingsV1alpha1()),"SettingsV1alpha1")

	writeDeleteCode(GetFunctionName2(clientset.AdmissionregistrationV1alpha1()),"AdmissionregistrationV1alpha1")

	writeDeleteCode(GetFunctionName2(clientset.AuthenticationV1()),"AuthenticationV1")

	writeDeleteCode(GetFunctionName2(clientset.ExtensionsV1beta1()),"ExtensionsV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.SchedulingV1beta1()),"SchedulingV1beta1")

	writeDeleteCode(GetFunctionName2(clientset.StorageV1()),"StorageV1")
}



func TestWriteGetCode(t *testing.T)  {

	clientset, _ := kubectl.InitClient()

	//输出转输入
	writeGetCode(GetFunctionName2(clientset.BatchV1beta1()),"BatchV1beta1")

	writeGetCode(GetFunctionName2(clientset.AuditregistrationV1alpha1()),"AuditregistrationV1alpha1")

	writeGetCode(GetFunctionName2(clientset.AdmissionregistrationV1beta1()),"AdmissionregistrationV1beta1")

	writeGetCode(GetFunctionName2(clientset.BatchV1()),"BatchV1")

	writeGetCode(GetFunctionName2(clientset.CertificatesV1beta1()),"CertificatesV1beta1")

	writeGetCode(GetFunctionName2(clientset.RbacV1()),"RbacV1")

	writeGetCode(GetFunctionName2(clientset.CoordinationV1beta1()),"CoordinationV1beta1")

	writeGetCode(GetFunctionName2(clientset.NetworkingV1()),"NetworkingV1")

	writeGetCode(GetFunctionName2(clientset.AppsV1()),"AppsV1")

	writeGetCode(GetFunctionName2(clientset.AutoscalingV2beta2()),"AutoscalingV2beta2")

	writeGetCode(GetFunctionName2(clientset.PolicyV1beta1()),"PolicyV1beta1")

	writeGetCode(GetFunctionName2(clientset.AuthorizationV1()),"AuthorizationV1")

	writeGetCode(GetFunctionName2(clientset.CoreV1()),"CoreV1")

	writeGetCode(GetFunctionName2(clientset.SettingsV1alpha1()),"SettingsV1alpha1")

	writeGetCode(GetFunctionName2(clientset.AdmissionregistrationV1alpha1()),"AdmissionregistrationV1alpha1")

	writeGetCode(GetFunctionName2(clientset.AuthenticationV1()),"AuthenticationV1")

	writeGetCode(GetFunctionName2(clientset.ExtensionsV1beta1()),"ExtensionsV1beta1")

	writeGetCode(GetFunctionName2(clientset.SchedulingV1beta1()),"SchedulingV1beta1")

	writeGetCode(GetFunctionName2(clientset.StorageV1()),"StorageV1")
}


func writeDeleteCode (args map[string]bool,api string){
	for  v,vbool := range args {
		if v!= "" {
			var vs = v + "s"
			if  v[len(v)-1:]== "s"  {
				vs = v + "es"
			}
			if v[len(v)-1:]== "y" {
				vs = v[:len(v)-1] + "ies"
			}

			var ns = "i.Namespace"
			if vbool == false {
				ns = ""
			}
			fmt.Printf(`
func (i *%s) Delete (opts *v1.DeleteOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	err = clientset.%s().%s(%s).Delete(i.Name,opts)
	if err != nil {
		return err
	}
	return nil
}
`, v, api, vs, ns)
		}
	}


}




func writeGetCode (args map[string]bool,api string){
	for  v,vbool := range args {
		if v!= "" {
			var vs = v + "s"
			if  v[len(v)-1:]== "s"  {
				vs = v + "es"
			}
			if v[len(v)-1:]== "y" {
				vs = v[:len(v)-1] + "ies"
			}

			var ns = "i.Namespace"
			if vbool == false {
				ns = ""
			}
			fmt.Printf(`
func (i *%s) Get (opts *v1.GetOptions) (error) {
	var clientset, err  = InitClient()
	if err != nil {
		return err
	}
	_ , err = clientset.%s().%s(%s).Get(i.Name,*opts)
	if err != nil {
		return err
	}
	return nil
}
`, v, api, vs, ns)
		}
	}


}





func TestWriteListCode(t *testing.T)  {

	clientset, _ := kubectl.InitClient()

	//输出转输入
	writeListCode(GetFunctionName2(clientset.BatchV1beta1()),"BatchV1beta1")

	writeListCode(GetFunctionName2(clientset.AuditregistrationV1alpha1()),"AuditregistrationV1alpha1")

	writeListCode(GetFunctionName2(clientset.AdmissionregistrationV1beta1()),"AdmissionregistrationV1beta1")

	writeListCode(GetFunctionName2(clientset.BatchV1()),"BatchV1")

	writeListCode(GetFunctionName2(clientset.CertificatesV1beta1()),"CertificatesV1beta1")

	writeListCode(GetFunctionName2(clientset.RbacV1()),"RbacV1")

	writeListCode(GetFunctionName2(clientset.CoordinationV1beta1()),"CoordinationV1beta1")

	writeListCode(GetFunctionName2(clientset.NetworkingV1()),"NetworkingV1")

	writeListCode(GetFunctionName2(clientset.AppsV1()),"AppsV1")

	writeListCode(GetFunctionName2(clientset.AutoscalingV2beta2()),"AutoscalingV2beta2")

	writeListCode(GetFunctionName2(clientset.PolicyV1beta1()),"PolicyV1beta1")

	writeListCode(GetFunctionName2(clientset.AuthorizationV1()),"AuthorizationV1")

	writeListCode(GetFunctionName2(clientset.CoreV1()),"CoreV1")

	writeListCode(GetFunctionName2(clientset.SettingsV1alpha1()),"SettingsV1alpha1")

	writeListCode(GetFunctionName2(clientset.AdmissionregistrationV1alpha1()),"AdmissionregistrationV1alpha1")

	writeListCode(GetFunctionName2(clientset.AuthenticationV1()),"AuthenticationV1")

	writeListCode(GetFunctionName2(clientset.ExtensionsV1beta1()),"ExtensionsV1beta1")

	writeListCode(GetFunctionName2(clientset.SchedulingV1beta1()),"SchedulingV1beta1")

	writeListCode(GetFunctionName2(clientset.StorageV1()),"StorageV1")
}


func writeListCode (args map[string]bool,api string){
	for  v,vbool := range args {
		if v!= "" {
			var vs = v + "s"
			if  v[len(v)-1:]== "s"  {
				vs = v + "es"
			}
			if v[len(v)-1:]== "y" {
				vs = v[:len(v)-1] + "ies"
			}

			var ns = "i.Namespace"
			if vbool == false {
				ns = ""
			}
			fmt.Printf(`
func (i *%s) GetAll(opts *v1.ListOptions) ([]%s.%s, error) {
	var clientset, err  = InitClient()
	if err != nil {
		return nil,err
	}
	%sList, err := clientset.%s().%s(%s).List(*opts)
	if err != nil {
		return nil,err
	}
	if i.Name == "" {
		return %sList.Items, nil
	} else {
		var items []%s.%s
		for _, v := range %sList.Items {
			match, err := regexp.Match(i.Name, []byte(v.ObjectMeta.Name))
			if err != nil {
				return nil,err
			}
			if match {
				items = append(items, v)
			}
		}
		return items, nil
	}
	}

`, v, lcfirst(api), v,lcfirst(v) , api,vs, ns,lcfirst(v),lcfirst(api), v,lcfirst(v))
		}
	}


}