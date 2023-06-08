package main

import (
	"sync"
)

var (
	pool = &Pool{}
)

type Pool struct {
	Jobs  []*PrintJob
	mutex sync.Mutex
}

func (p *Pool) AddJob(job *PrintJob) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Jobs = append(p.Jobs, job)
}

func (p *Pool) GetJob() *PrintJob {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if len(p.Jobs) > 0 {
		job := p.Jobs[0]
		p.Jobs = p.Jobs[1:]
		return job
	}

	return nil
}
