package main

import (
	"encoding/json"
	"os"
)

var (
	printers map[string]*Printer
)

func init() {
	b, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	config := struct {
		Printers []*struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"printers"`
	}{}

	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}

	printers = make(map[string]*Printer, len(config.Printers))
	for _, printer := range config.Printers {
		printers[printer.Address] = NewPrinter(printer.Name, printer.Address)
	}
}

func main() {
	for _, printer := range printers {
		printer.Attach()
	}

	panic(router.Run(":8989"))
}
