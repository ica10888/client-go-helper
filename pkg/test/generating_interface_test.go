package test

import (
	"client-go-helper/pkg/kubectl"
	"fmt"
	"strings"
	"testing"
)


// 注意！！！corev1 -> v1
func TestCreateInterfaceCode(t *testing.T)  {


	clientset, _ := kubectl.InitClient()
	maps := GetFunctionName1(clientset)
	for k, _ := range maps {
		fmt.Printf(`
	writeInterfaceCode(GetFunctionName2(clientset.%s()),"%s")
`,k,k)
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")

	for k, _ := range maps {
		fmt.Printf(`
	writeInterfaceCode2(GetFunctionName2(clientset.%s()),"%s")
`,k,k)
	}
}


func TestWriteInterfaceCode(t *testing.T)  {

	clientset, _ := kubectl.InitClient()

	//输出转输入
	writeInterfaceCode(GetFunctionName2(clientset.BatchV1beta1()),"BatchV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.CertificatesV1beta1()),"CertificatesV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.CoordinationV1beta1()),"CoordinationV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.NetworkingV1()),"NetworkingV1")

	writeInterfaceCode(GetFunctionName2(clientset.AuthorizationV1beta1()),"AuthorizationV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.ExtensionsV1beta1()),"ExtensionsV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.RbacV1()),"RbacV1")

	writeInterfaceCode(GetFunctionName2(clientset.AppsV1beta2()),"AppsV1beta2")

	writeInterfaceCode(GetFunctionName2(clientset.AuthenticationV1()),"AuthenticationV1")

	writeInterfaceCode(GetFunctionName2(clientset.RbacV1alpha1()),"RbacV1alpha1")

	writeInterfaceCode(GetFunctionName2(clientset.SettingsV1alpha1()),"SettingsV1alpha1")

	writeInterfaceCode(GetFunctionName2(clientset.AdmissionregistrationV1beta1()),"AdmissionregistrationV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.EventsV1beta1()),"EventsV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.PolicyV1beta1()),"PolicyV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.SchedulingV1alpha1()),"SchedulingV1alpha1")

	writeInterfaceCode(GetFunctionName2(clientset.StorageV1beta1()),"StorageV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.BatchV1()),"BatchV1")

	writeInterfaceCode(GetFunctionName2(clientset.RbacV1beta1()),"RbacV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.StorageV1()),"StorageV1")

	writeInterfaceCode(GetFunctionName2(clientset.AuditregistrationV1alpha1()),"AuditregistrationV1alpha1")

	writeInterfaceCode(GetFunctionName2(clientset.AuthenticationV1beta1()),"AuthenticationV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.AutoscalingV2beta1()),"AutoscalingV2beta1")

	writeInterfaceCode(GetFunctionName2(clientset.SchedulingV1beta1()),"SchedulingV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.AppsV1()),"AppsV1")

	writeInterfaceCode(GetFunctionName2(clientset.AppsV1beta1()),"AppsV1beta1")

	writeInterfaceCode(GetFunctionName2(clientset.AuthorizationV1()),"AuthorizationV1")

	writeInterfaceCode(GetFunctionName2(clientset.AutoscalingV1()),"AutoscalingV1")

	writeInterfaceCode(GetFunctionName2(clientset.AutoscalingV2beta2()),"AutoscalingV2beta2")

	writeInterfaceCode(GetFunctionName2(clientset.BatchV2alpha1()),"BatchV2alpha1")

	writeInterfaceCode(GetFunctionName2(clientset.CoreV1()),"CoreV1")

	writeInterfaceCode(GetFunctionName2(clientset.StorageV1alpha1()),"StorageV1alpha1")

	writeInterfaceCode(GetFunctionName2(clientset.AdmissionregistrationV1alpha1()),"AdmissionregistrationV1alpha1")

}



func writeInterfaceCode (args map[string]bool,api string){
	if api != "CoreV1" {
		comma := strings.Index(lcfirst(api), "V")
		v2 := lcfirst(api)[:comma]
		v3 := lcfirst(lcfirst(api)[comma:])
		fmt.Printf(`
//%s.%s`, v2, v3)
	} else {
		fmt.Printf(`
// v1`)
	}

	for  v,vbool := range args {
		if v!= "" {

			var ns = "Namespace     string"
			if vbool == false {
				ns = ""
			}

			fmt.Printf(`
type %s%s struct {
	Name     string
	%s
}
`, api, v,ns)


		}
	}
	fmt.Printf(`
`)
}