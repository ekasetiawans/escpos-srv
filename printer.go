package main

import (
	"log"
	"os/exec"
)

func Print(filename, printer string) error {
	output, err := exec.Command("lpr", "-r", "-P", printer, "-o", "portrait", "-o", "media=A4", filename).CombinedOutput()
	log.Println(output)

	return err
}
