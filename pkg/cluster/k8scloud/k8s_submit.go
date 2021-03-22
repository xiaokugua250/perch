package k8scloud

import (
	"context"
	"github.com/coreos/etcd/etcdserver/api"
	"github.com/coreos/etcd/etcdserver/auth"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
	model "perch/web/model/rbac"
	"strconv"
)

/**
提交deployment类型资源
*/
func (k8sClient *clusterClientManager) SubmitDeploymentType(ns string, dName string, lables map[string]string,
	imageName string, Image string,
	user model.AuthUser,
) (*appsv1.Deployment, error) {
	var (
		k8sDeploy *appsv1.Deployment
		err       error
	)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: dName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: lables,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: lables,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            imageName,
							Image:           Image,
							ImagePullPolicy: "Always",
							Command:         []string{},
							Env: []apiv1.EnvVar{
								{
									Name:  "USER",
									Value: user.UserName,
								},
								/*	{
									Name:  "HOME",
									Value: "/home/" + userName,
								},*/

							},
							Resources: apiv1.ResourceRequirements{
								Requests: requests,
								Limits:   requests,
							},
							VolumeMounts: []apiv1.VolumeMount{

								{
									Name:      "shm",
									MountPath: "/dev/shm",
								},
							},
							SecurityContext: &apiv1.SecurityContext{
								Capabilities: &apiv1.Capabilities{
									Add: []apiv1.Capability{
										"SYS_ADMIN",
									},
								},
							},
						},
					},
					Volumes: []apiv1.Volume{
						{
							Name: "home",
							VolumeSource: apiv1.VolumeSource{
								HostPath: &apiv1.HostPathVolumeSource{
									//	Path: "/home/" + userName,
								},
							},
						},

						{
							Name: "shm",
							VolumeSource: apiv1.VolumeSource{
								EmptyDir: &apiv1.EmptyDirVolumeSource{
									Medium: apiv1.StorageMediumMemory,
								},
							},
						},
					},
					SecurityContext: &apiv1.PodSecurityContext{
						RunAsUser:  &user.UserUID,
						RunAsGroup: &user.UserGID,
					},
				},
			},
		},
	}

	k8sDeploy, err = k8sClient.clusterClient.AppsV1().Deployments(ns).Create(context.TODO(), deployment, metav1.CreateOptions{})
	return k8sDeploy, err
}

/**
创建网络Service类型
*/
func SubmitServiceType() {

}

/**
提交job类型资源
*/
func SubmitJobType() {

}

/**
提交cronjob类型资源
*/
func SubmitCronJobType() {

}

func int32Ptr(i int32) *int32 { return &i }
