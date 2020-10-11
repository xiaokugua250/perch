package k8s

import (
	"fmt"
	"log"
	_ "perch/pkg/log"
	"testing"
)

func TestK8SClusterManager_InitK8SClusterClient(t *testing.T) {

	var (
		err error
	)
	clusterManager := ClusterManager{


		KubeConfig:   KUBEConfig{
			ClusterName:"k8s_dev",
			ConfigFile: "E:\\WorksSpaces\\GoWorkSpaces\\perch\\configs\\dev\\cluster_config\\kubernetes_clusters\\k8s_dev.config",
		},
	}
	if err = clusterManager.InitK8SClusterClient(); err != nil {
		log.Fatal(err)
	}
	/*nodes,err :=K8SClusterClientMap["dev"].CoreV1().Nodes().List(context.Background(),metav1.ListOptions{})
	if err!= nil{
		log.Println(err)
	}
	for _,node:=range nodes.Items{
		fmt.Println(node)
	}*/
	/*selector := labels.NewSelector()
	result, err := K8sResourceListWithInformer(ClusterClientMap["dev"], K8S_RESOURCE_NODE, selector)
	if err != nil {
		log.Println(err)
	}
	nodes, ok := result.([]*v1.Node)
	if ok {
		for _, node := range nodes {
			fmt.Println(node)
		}
	}
	fmt.Println("result is ", result)
	*/


}


func TestInitKubernetesCluster(t *testing.T) {
	if err :=InitKubernetesCluster();err!= nil{
		fmt.Println(err)
	}
}
