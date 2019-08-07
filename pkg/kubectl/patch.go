package kubectl

import (
	batchV1beta1 "k8s.io/api/batch/v1beta1"
	types "k8s.io/apimachinery/pkg/types")

func (i *CronJob) Patch (data string,pt *types.PatchType) (batchV1beta1.CronJob,error) {
	var clientset, err  = InitClient()
	if err != nil {
		return batchV1beta1.CronJob{},err
	}
	cronJob , err := clientset.BatchV1beta1().CronJobs(i.Namespace).Patch(i.Name,*pt,[]byte(data))
	if err != nil {
		return batchV1beta1.CronJob{},err
	}
	return *cronJob,nil
}