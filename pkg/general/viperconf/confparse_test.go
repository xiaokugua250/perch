package viperconf

import (
	"fmt"
	"testing"
)

func TestInitYamlconfig(t *testing.T) {
	err := InitGeneralWebConfig("E:/WorksSpaces/GoWorkSpaces/perch/configs/web_config/admin.yaml")
	if err != nil {
		fmt.Print("error is ", err)
	}
	fmt.Printf("%v\n")
}
