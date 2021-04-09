package api

import (
	"context"
	"fmt"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeployments(){
	clientset := GetClientSet()
	deployment := clientset.AppsV1().Deployments(v12.NamespaceDefault)
	list, err := deployment.List(context.TODO(), v12.ListOptions{})
	if err != nil{
		fmt.Println(err)
	}
	for _, deploy := range list.Items {
		fmt.Println(deploy)
	}
}