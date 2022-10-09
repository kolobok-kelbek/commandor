package main

import (
	"fmt"
	"log"
	"os/exec"
)

func execute(command string) {
	bashExecPath, err := exec.LookPath("bash")
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(bashExecPath, "-c", command)
	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))

	if err != nil {
		log.Fatal(err)
	}
}
