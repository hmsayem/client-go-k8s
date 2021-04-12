package api

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeployment(){
	clientset := GetClientSet()
	deployClient := clientset.AppsV1().Deployments(v12.NamespaceDefault)

	deployment := &v1.Deployment{
		ObjectMeta: v12.ObjectMeta{
			Name: "test-deploy",
		},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &v12.LabelSelector{
				MatchLabels: map[string]string{
					"app": "server",
				},
			},
			Template: v13.PodTemplateSpec{
				ObjectMeta: v12.ObjectMeta{
					Labels: map[string] string{
						"app": "server",
					},
				},
				Spec: v13.PodSpec{
					Containers: []v13.Container{
						{
							Name: "web-server",
							Image: "nginx",
							Ports: []v13.ContainerPort{

							},
						},
					},
				},
			},
		},
	}

	//create deployment
	res, err := deployClient.Create(context.TODO(), deployment, v12.CreateOptions{})
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("%q Created", res.GetObjectMeta().GetName())
	}
}

func GetDeployments(){
	clientset := GetClientSet()
	deployClient := clientset.AppsV1().Deployments(v12.NamespaceDefault)
	list, err := deployClient.List(context.TODO(), v12.ListOptions{})
	if err != nil{
		fmt.Println(err)
	}
	for _, deploy := range list.Items{
		fmt.Println(deploy.Name, "[ Ready: ", deploy.Status.ReadyReplicas, "|", "Available: ", deploy.Status.AvailableReplicas, "]")
	}
}

func UpdateDeployment(name string, replica int32, image string){
	clientset := GetClientSet()
	deploymentClient := clientset.AppsV1().Deployments(v12.NamespaceDefault)
	res, err := deploymentClient.Get(context.TODO(), name, v12.GetOptions{})
	if err != nil{
		fmt.Println(err)
	}
	res.Spec.Replicas = int32Ptr(replica)
	res.Spec.Template.Spec.Containers[0].Image = image
	_, err = deploymentClient.Update(context.TODO(), res, v12.UpdateOptions{})
	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Printf("%q Updated", res.GetObjectMeta().GetName())
	}
}

func DeleteDeployment(args []string){
	clientset := GetClientSet()
	deployClient := clientset.AppsV1().Deployments(v12.NamespaceDefault)
	deletePolicy := v12.DeletePropagationForeground
	for _, deployment := range args {
		err := deployClient.Delete(context.TODO(), deployment, v12.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}