package test

import (
	"client-go-helper/pkg/kubectl"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unicode"
)

// 注意！！！corev1 -> v1
func TestCreateCode(t *testing.T)  {


	clientset, _ := kubectl.InitClient()
	maps := GetFunctionName1(clientset)
	for k, _ := range maps {
		fmt.Printf(`
	writeCode(GetFunctionName2(clientset.%s()),"%s")
`,k,k)
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	for v, _ := range maps {
		comma:= strings.Index(lcfirst(v), "V")
		v2 := lcfirst(v)[:comma]
		v3 :=  lcfirst(lcfirst(v)[comma:])
		fmt.Printf("	%s           \"k8s.io/api/%s/%s\"",lcfirst(v),v2,v3)
		fmt.Printf("\n")
	}


	//输出转输入
	writeCode(GetFunctionName2(clientset.BatchV1beta1()),"BatchV1beta1")

	writeCode(GetFunctionName2(clientset.CertificatesV1beta1()),"CertificatesV1beta1")

	writeCode(GetFunctionName2(clientset.CoordinationV1beta1()),"CoordinationV1beta1")

	writeCode(GetFunctionName2(clientset.NetworkingV1()),"NetworkingV1")

	writeCode(GetFunctionName2(clientset.AuthorizationV1beta1()),"AuthorizationV1beta1")

	writeCode(GetFunctionName2(clientset.ExtensionsV1beta1()),"ExtensionsV1beta1")

	writeCode(GetFunctionName2(clientset.RbacV1()),"RbacV1")

	writeCode(GetFunctionName2(clientset.AppsV1beta2()),"AppsV1beta2")

	writeCode(GetFunctionName2(clientset.AuthenticationV1()),"AuthenticationV1")

	writeCode(GetFunctionName2(clientset.RbacV1alpha1()),"RbacV1alpha1")

	writeCode(GetFunctionName2(clientset.SettingsV1alpha1()),"SettingsV1alpha1")

	writeCode(GetFunctionName2(clientset.AdmissionregistrationV1beta1()),"AdmissionregistrationV1beta1")

	writeCode(GetFunctionName2(clientset.EventsV1beta1()),"EventsV1beta1")

	writeCode(GetFunctionName2(clientset.PolicyV1beta1()),"PolicyV1beta1")

	writeCode(GetFunctionName2(clientset.SchedulingV1alpha1()),"SchedulingV1alpha1")

	writeCode(GetFunctionName2(clientset.StorageV1beta1()),"StorageV1beta1")

	writeCode(GetFunctionName2(clientset.BatchV1()),"BatchV1")

	writeCode(GetFunctionName2(clientset.RbacV1beta1()),"RbacV1beta1")

	writeCode(GetFunctionName2(clientset.StorageV1()),"StorageV1")

	writeCode(GetFunctionName2(clientset.AuditregistrationV1alpha1()),"AuditregistrationV1alpha1")

	writeCode(GetFunctionName2(clientset.AuthenticationV1beta1()),"AuthenticationV1beta1")

	writeCode(GetFunctionName2(clientset.AutoscalingV2beta1()),"AutoscalingV2beta1")

	writeCode(GetFunctionName2(clientset.SchedulingV1beta1()),"SchedulingV1beta1")

	writeCode(GetFunctionName2(clientset.AppsV1()),"AppsV1")

	writeCode(GetFunctionName2(clientset.AppsV1beta1()),"AppsV1beta1")

	writeCode(GetFunctionName2(clientset.AuthorizationV1()),"AuthorizationV1")

	writeCode(GetFunctionName2(clientset.AutoscalingV1()),"AutoscalingV1")

	writeCode(GetFunctionName2(clientset.AutoscalingV2beta2()),"AutoscalingV2beta2")

	writeCode(GetFunctionName2(clientset.BatchV2alpha1()),"BatchV2alpha1")

	writeCode(GetFunctionName2(clientset.CoreV1()),"CoreV1")

	writeCode(GetFunctionName2(clientset.StorageV1alpha1()),"StorageV1alpha1")

	writeCode(GetFunctionName2(clientset.AdmissionregistrationV1alpha1()),"AdmissionregistrationV1alpha1")



}

func lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}



func GetFunctionName1(i interface{}) map[string]bool {
	maps := map[string]bool{}
	for num := 0;num < reflect.ValueOf(i).NumMethod();num++ {

		str := reflect.ValueOf(i).Method(num).String()
		if str[8:9] == "v"{
			begin := strings.Index(str, ".")
			end := strings.Index(str, "Interface")
			maps[str[begin+1:end]] = true
		}
	}
	return maps
}
func GetFunctionName2(i interface{}) map[string]bool {
	maps := map[string]bool{}
	for num := 0;num < reflect.ValueOf(i).NumMethod();num++ {

		str := reflect.ValueOf(i).Method(num).String()
		begin := strings.Index(str, ".")
		end := strings.Index(str, "Interface")
		if str[1:7] == "func()" {
			maps[str[begin+1:end]] = false
		} else {
			maps[str[begin+1:end]] = true
		}
	}
	return maps
}


func writeCode (args map[string]bool,api string){
	if api != "CoreV1" {
		comma := strings.Index(lcfirst(api), "V")
		v2 := lcfirst(api)[:comma]
		v3 := lcfirst(lcfirst(api)[comma:])
		fmt.Printf(`
	case "%s/%s":
	switch kapi.Kind {
`, v2, v3)
	} else {
		fmt.Printf(`
	case "v1":
	switch kapi.Kind {
`)
	}


	for  v,vbool := range args {
		if v!= "" {
			var vs = v + "s"
			if  v[len(v)-1:]== "s"  {
				vs = v + "es"
			}
			if v[len(v)-1:]== "y" {
				vs = v[:len(v)-1] + "ies"
			}

			var ns = "namespace"
			if vbool == false {
				ns = ""
			}

			var _ = func() {
				fmt.Printf(`
		case "%s":
			%s := obj.(*%s.%s)
			_, e = clientset.%s().%s(%s).Create(%s)
			if e != nil {
				if apierrs.IsAlreadyExists(e) {
					_ ,e := clientset.%s().%s(%s).Update(%s)
					if e != nil{
						fmt.Printf("updated,but something wrong happend ,%%+v",e)
					}
				} else {
					fmt.Printf("something wrong happend ,%%+v",e)
				}
			}
`, v, lcfirst(v), lcfirst(api), v, api, vs, ns,lcfirst(v), api, vs, ns,lcfirst(v))
			}
			//applyPrintf()


			var _ = func() {
				fmt.Printf(`
		case "%s":
			%s := obj.(*%s.%s)
			_, e = clientset.%s().%s(%s).Create(%s)
			if e != nil {
				return e
			}
`, v, lcfirst(v), lcfirst(api), v, api, vs, ns,lcfirst(v))
			}
			//createPrintf()
			var listPrintf = func() {
				fmt.Printf(`
		case "%s":
			%sList, e := clientset.%s().%s(%s).List(*opts)
			if e != nil {
				return nil,e
			}
			var items []string
			for _,i := range %sList.Items {
				json, err := json.Marshal(i)
				if err != nil {
					return nil,fmt.Errorf("json Unmarshal err")
				}
				data ,_ := yaml.JSONToYAML(json)
				items = append(items, string(data))
			}
			return items,nil
`, v, lcfirst(v),api, vs, ns,lcfirst(v))
			}
			listPrintf()


		}
	}
	fmt.Printf(`
	default:
		return fmt.Errorf("not support a kind : %%s in  apiVersion: %%s",kapi.Kind,kapi.ApiVersion)
	}
`)


}


