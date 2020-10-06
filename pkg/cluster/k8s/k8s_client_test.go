package k8s

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"log"
	_ "perch/pkg/log"
	"testing"
)

func TestK8SClusterManager_InitK8SClusterClient(t *testing.T) {

	var (
		err error
	)
	clusterManager := K8SClusterManager{
		KubeClusterName: "dev",
		KubeConfig:      "E:\\WorksSpaces\\GoWorkSpaces\\perch\\configs\\dev\\cluster_config\\k8s_dev.config",
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
	selector := labels.NewSelector()
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

}
