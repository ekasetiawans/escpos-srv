package main

import (
	"os"
	"time"

	"github.com/knq/escpos"
)

func init() {

}

type PrintJob struct {
	Printer string
	Data    string
}

type Printer struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	isAttached bool
	Pool       *Pool
}

func NewPrinter(name, address string) *Printer {
	return &Printer{
		Name:       name,
		Address:    address,
		isAttached: false,
		Pool:       NewPool(),
	}
}

func (p *Printer) Detach() {
	p.isAttached = false
}

func (p *Printer) Attach() {
	go p.doAttach()
}

func (p *Printer) doAttach() {
	p.isAttached = true
	// infinite loops
	for p.isAttached {
		job := p.Pool.GetJob()
		if job == nil {
			time.Sleep(2 * time.Second)
			continue
		}

		p.Print(job.Data)
		time.Sleep(2 * time.Second)
	}
}

func (p *Printer) Print(data string) error {
	f, err := os.OpenFile(p.Address, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()

	pos := escpos.New(f)
	_, err = pos.WriteRaw([]byte(data))
	if err != nil {
		return err
	}

	return nil
}
