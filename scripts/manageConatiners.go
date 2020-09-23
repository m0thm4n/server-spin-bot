package scripts

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// CreateNewContainer runs a python script to create a container
func CreateNewContainer(image, name string) {
	// exec.Command("powershell.exe", "python.exe", "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\create_container.py "+image+" "+name).Run()
	pythonExecutable, err := exec.LookPath("python")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   pythonExecutable,
		Args:   []string{pythonExecutable, "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\create_container.py " + image + " " + name},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Image pulled and server has been spun up")
}

// ListContainers lists all containers in the Docker enviroment
func ListContainers() string {
	// cmd := exec.Command("powershell.exe", "python.exe", "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\list_containers.py").Run()

	pythonExecutable, err := exec.LookPath("python")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   pythonExecutable,
		Args:   []string{pythonExecutable, "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\list_all_containers.py"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	output := fmt.Sprint(cmd.Stdout)
	fmt.Println(output)

	return output
}
