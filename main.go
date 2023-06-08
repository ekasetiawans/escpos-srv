package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	printers []*Printer
)

func init() {
	b, err := ioutil.ReadFile("config.json")
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

	printers = make([]*Printer, len(config.Printers))
	for i, printer := range config.Printers {
		printers[i] = NewPrinter(printer.Name, printer.Address)
	}
}

func main() {
	for _, printer := range printers {
		printer.Attach(pool)
	}

	panic(router.Run(":8989"))
}
