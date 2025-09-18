package server

import (
	"context"
	"time"

	"goal-advanced-layout/pkg/log"

	"github.com/go-co-op/gocron"
)

type Task struct {
	log       *log.Logger
	scheduler *gocron.Scheduler
}

func NewTask(log *log.Logger) *Task {
	return &Task{
		log: log,
	}
}

func (t *Task) Start(ctx context.Context) error {
	t.scheduler = gocron.NewScheduler(time.UTC)
	_, err := t.scheduler.CronWithSeconds("0/3 * * * * *").Do(func() {
		t.log.Info("I'm a Task1.")
	})
	if err != nil {
		t.log.Error("Task1 error" + err.Error())
	}

	_, err = t.scheduler.Every("3s").Do(func() {
		t.log.Info("I'm a Task2.")
	})
	if err != nil {
		t.log.Error("Task2 error" + err.Error())
	}

	t.scheduler.StartBlocking()
	return nil
}

func (t *Task) Stop(ctx context.Context) error {
	t.scheduler.Stop()
	t.log.Info("Task stop...")
	return nil
}
