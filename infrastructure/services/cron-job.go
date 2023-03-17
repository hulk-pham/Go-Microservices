package services

import (
	"time"

	"github.com/go-co-op/gocron"
)

type SchedulerService struct {
	cron *gocron.Scheduler
}

var SchedulerInstance *SchedulerService

func InitSchedulerService() {
	SchedulerInstance = &SchedulerService{
		cron: gocron.NewScheduler(time.UTC),
	}
}

func (m *SchedulerService) RegisterCron(timePattern string, task any) {
	m.cron.Cron(timePattern).Do(task)
}

func (m *SchedulerService) Start() {
	m.cron.StartAsync()
}
