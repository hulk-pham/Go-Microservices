package background

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	cron *gocron.Scheduler
}

func (m *Scheduler) Init() {
	m.cron = gocron.NewScheduler(time.UTC)
}

func (m *Scheduler) RegisterCron(timePattern string, task any) {
	m.cron.Cron(timePattern).Do(task)
}

func (m *Scheduler) Start() {
	m.cron.StartAsync()
}
