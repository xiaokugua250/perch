package k8scloud

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"log"
	model "perch/web/model/rbac"
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
								Requests: apiv1.ResourceList{
									"cpu":resource.MustParse("cpuRequest"),
									"memory": resource.MustParse("memoryRequest"),
								},
								Limits:   apiv1.ResourceList{
									"cpu":resource.MustParse("cpuLimit"),
									"memory": resource.MustParse("memoryLimit"),
								},
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
func  (k8sClient *clusterClientManager)SubmitServiceType(ns string, serviceName string, lables map[string]string)(*apiv1.Service,error) {
	var (
		k8sService *apiv1.Service
		err       error
	)

	k8sService = &apiv1.Service{

		ObjectMeta: metav1.ObjectMeta{
			Name:                       serviceName,
			Namespace:                  ns,
			Labels: lables,
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Name:       "terminal",
					Port:       8000,
					TargetPort: intstr.FromInt(8000),
				},
				{
					Name:       "ssh",
					Port:       22,
					TargetPort: intstr.FromInt(22),
				},

			},
			Selector: lables,
		},


}

	k8sService, err = k8sClient.clusterClient.CoreV1().Services(ns).Create(context.Background(),k8sService,metav1.CreateOptions{})
	return k8sService, err
}

/**
提交job类型资源
*/
func  (k8sClient *clusterClientManager)SubmitJobType(ns string, jobName string, lables map[string]string,
	imageName string, Image string,
	user model.AuthUser) {
	jobs := k8sClient.clusterClient.BatchV1().Jobs("default")
	var backOffLimit int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
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
								Requests: apiv1.ResourceList{
									"cpu":resource.MustParse("cpuRequest"),
									"memory": resource.MustParse("memoryRequest"),
								},
								Limits:   apiv1.ResourceList{
									"cpu":resource.MustParse("cpuLimit"),
									"memory": resource.MustParse("memoryLimit"),
								},
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
					RestartPolicy: apiv1.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("Failed to create K8s job.")
	}

	//print job details
	log.Println("Created K8s job successfully")
}

/**
提交cronjob类型资源
*/
func  (k8sClient *clusterClientManager)SubmitCronJobType(ns string, jobName string, lables map[string]string,
	imageName string, Image string,
	user model.AuthUser) {

	cronjob := &v1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: ns,
		},
		Spec: v1beta1.CronJobSpec{
			Schedule:          "* * * * *",
			ConcurrencyPolicy: v1beta1.ForbidConcurrent,
			JobTemplate: v1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: apiv1.PodTemplateSpec{
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
										Requests: apiv1.ResourceList{
											"cpu":resource.MustParse("cpuRequest"),
											"memory": resource.MustParse("memoryRequest"),
										},
										Limits:   apiv1.ResourceList{
											"cpu":resource.MustParse("cpuLimit"),
											"memory": resource.MustParse("memoryLimit"),
										},
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
			},
		},
	}
	_, err := k8sClient.clusterClient.BatchV1beta1().CronJobs("ns").Create(context.TODO(), cronjob, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("Failed to create K8s job.")
	}
}

func int32Ptr(i int32) *int32 { return &i }
