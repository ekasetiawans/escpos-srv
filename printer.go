package main

import (
	"log"
	"os/exec"
)

func Print(filename, printer string) error {
	output, err := exec.Command("lpr", "-r", "-P", printer, filename).CombinedOutput()
	if err != nil {
		log.Println(output)
	}

	return err
}
