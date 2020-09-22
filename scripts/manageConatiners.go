package scripts

import (
	"fmt"
	"os/exec"
)

// CreateNewContainer runs a python script to create a container
func CreateNewContainer(image, name string) {
	exec.Command("powershell.exe", "python.exe", "C:\\workspace\\go\\src\\server-spin\\scripts\\create_container.py "+image+" "+name).Run()

	fmt.Println("Image pulled and server has been spun up")
}

// ListContainers lists all containers in the Docker enviroment
func ListContainers() {
	exec.Command("powershell.exe", "python.exe", "C:\\workspace\\go\\src\\server-spin\\scripts\\list_containers.py").Run()
}
