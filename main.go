package main

import (
	"errors"
	"fmt"
	"os/exec"
)

func cmd(name string, args ...string) (string, error) {
	c := exec.Command(name, args...)
	output, err := c.Output()
	if err != nil {
		return "", errors.New(err.Error())
	}
	return string(output), nil
}

func main() {
	//containerId := exec.Command("docker", "ps", "-aqf")

	//containerId := exec.Command("bash", "-c", "docker info|grep Runtime")
	s, err := cmd("bash", "-c", "docker info|grep Runtime")
	if err != nil {
		return
	}
	fmt.Println(s)

	//cid, err := containerId.Output()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(string(cid))
	//c := string(cid)
	//cmd := fmt.Sprintf("docker inspect %v", c)
	//inspectCmd := exec.Command(cmd)
	//if err != nil {
	//	return
	//}
	//o, err := inspectCmd.Output()
	//fmt.Println(string(o))
	//fmt.Println(containerId.Args)

}
