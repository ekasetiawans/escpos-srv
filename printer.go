package main

import (
	"os"
	"time"

	"github.com/knq/escpos"
)

func init() {

}

type PrintJob struct {
	Data string
}

type Printer struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	isAttached bool
}

func NewPrinter(name, address string) *Printer {
	return &Printer{
		Name:       name,
		Address:    address,
		isAttached: false,
	}
}

func (p *Printer) Detach() {
	p.isAttached = false
}

func (p *Printer) Attach(pool *Pool) {
	go p.doAttach(pool)
}

func (p *Printer) doAttach(pool *Pool) {
	p.isAttached = true
	// infinite loops
	for p.isAttached {
		job := pool.GetJob()
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
