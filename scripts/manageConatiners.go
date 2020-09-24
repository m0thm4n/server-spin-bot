package scripts

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// CreateNewContainer runs a python script to create a container
func CreateNewContainer(image, name, port string) {
	// exec.Command("powershell.exe", "python.exe", "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\create_container.py "+image+" "+name).Run()
	pythonExecutable, err := exec.LookPath("python")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   pythonExecutable,
		Args:   []string{pythonExecutable, "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\create_container.py", image, name, port},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Image pulled and server has been spun up")
}

// ListContainers lists all containers in the Docker enviroment
func ListContainers() []string {
	// cmd := exec.Command("powershell.exe", "python.exe", "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\list_containers.py").Run()
	var containers []string

	pythonExecutable, err := exec.LookPath("python")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path: pythonExecutable,
		Args: []string{pythonExecutable, "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\list_all_containers.py"},
	}

	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("Error obtaining stdin: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error obtaining stdout: %s", err.Error())
	}

	reader := bufio.NewReader(stdout)
	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			log.Printf("Reading the list of containers: %s", scanner.Text())
			containers = append(containers, scanner.Text())
			stdin.Write([]byte("some sample text\n"))
		}
	}(reader)

	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	cmd.Wait()

	return containers
}

// StartContainer starts a container that is already created
func StartContainer(id string) {
	pythonExecutable, err := exec.LookPath("python")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   pythonExecutable,
		Args:   []string{pythonExecutable, "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\start_container.py", id},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

// StopContainer stops a container that is currently running
func StopContainer(id string) {
	pythonExecutable, err := exec.LookPath("python")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   pythonExecutable,
		Args:   []string{pythonExecutable, "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\stop_container.py", id},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

// DeleteContainer deletes a stopped container only. It will not remove a running container
func DeleteContainer(id string) {
	pythonExecutable, err := exec.LookPath("python")
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   pythonExecutable,
		Args:   []string{pythonExecutable, "C:\\workspace\\go\\src\\server-spin-bot\\scripts\\delete_container.py", id},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
