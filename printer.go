package main

import (
	"log"
	"os/exec"
)

func Print(filename, printer string) error {
	output, err := exec.Command("lpr", "-r", "-P", printer, "-o", "portrait", "-o", "media=684x396", filename).CombinedOutput()
	log.Println(output)

	return err
}
