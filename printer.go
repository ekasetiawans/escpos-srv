package main

import "os/exec"

func Print(filename, printer string) error {
	cmd := exec.Command("lpr", "-r", "-P", printer, filename)
	return cmd.Run()
}
