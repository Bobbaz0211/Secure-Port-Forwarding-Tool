package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const defCommand string = "ssh"

var baseArgs = []string{
	"-N",
	"-o",
	"ExitOnForwardFaliure=yes",
	"-o",
	"StreamLocalBindUnlink=yes",
	"-o",
	"ServerAliveInterval=5",
	"-o",
	"ServerAliveCountMax=1",
	"-L",
}

func setupExec(target string, metadata []string) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	var cmdErr error
	var cmdArgs []string = nil

	switch metadata[1] {
	case "tcp":
		cmdArgs = append(baseArgs, fmt.Sprintf("%s:localhost:%s",
			metadata[0], metadata[0]))

	default:
		fmt.Println(metadata[1])
		return nil, cmdErr

	}

	cmdArgs = append(cmdArgs, metadata[2])
	cmd = exec.Command(defCommand, cmdArgs...)
	cmdErr = cmd.Start()

	if cmdErr != nil {
		return nil, cmdErr
	}

	return cmd, cmdErr
}

func processViewer(target string, metadata []string) {

	var cmd *exec.Cmd
	var cmdErr error

	for true {
		//exec and start process
		// wait for exit
		//sleep 5s

		cmd, cmdErr = setupExec(target, metadata)

		if cmd == nil || cmdErr != nil {
			fmt.Println("Cannot spawn process\n")

			time.Sleep(time.Duration(time.Duration(5) * time.Second))
			continue
		}

		fmt.Println("Spawned process for '%s' as %d\n")

		cmd.Wait()

		fmt.Println("Process for '%s' died, Waiting for %ds before restarting...\n")

		time.Sleep(time.Duration(5) * time.Second)

	}
}

func main() {

	var cfgContainer map[string]interface{}
	var cfg = map[string][]string{}
	var start string
	var first bool = false

	data, err := os.ReadFile("./pfToolConfig.json")
	if err != nil {
		fmt.Printf("Error Reading Config file!\n")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &cfgContainer)
	if err != nil {
		fmt.Printf("Error Unmarshlling Json\n")
		os.Exit(1)
	}

	for i := range cfgContainer {

		switch interfaceType := cfgContainer[i].(type) {

		case []interface{}:

			for j := 0; j < len(interfaceType); j++ {

				switch valueType := interfaceType[j].(type) {

				case string:
					cfg[i] = append(cfg[i], valueType)
				default:
					continue
				}
			}
		default:
			continue
		}
	}

	for i := range cfg {
		if !first {
			first = true
			start = i
			continue
		}
		go processViewer(i, cfg[i])
	}

	processViewer(start, cfg[start])

}
