package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type Printer struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func getPrinters() []*Printer {
	// Execute lpstat -p command
	cmd := exec.Command("lpstat", "-p")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing lpstat -p: %s\n", err)
		return nil
	}

	// Parse the output
	printers := parseLPStatOutput(string(output))
	return printers
}

func parseLPStatOutput(output string) []*Printer {
	lines := strings.Split(output, "\n")
	printers := make([]*Printer, 0)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Example line: printer printer1 is idle.  enabled since Mon 01 Jan 2023 00:00:00 AM UTC -
		parts := strings.SplitN(line, " ", 4)
		if len(parts) < 4 {
			continue
		}

		printerName := parts[1]
		printerStatus := parts[3]

		printer := &Printer{
			Name:   printerName,
			Status: printerStatus,
		}

		printers = append(printers, printer)
	}

	return printers
}
