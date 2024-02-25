package tasks

import (
	"context"
	"github.com/samber/do"
	"gocourse22/pkg/scheduler"
)

func CalculateEmployysBonuses(_ *do.Injector) *EnqueueSodaResultParsing {
	return &EnqueueSodaResultParsing{}
}

type EnqueueSodaResultParsing struct {
}

func (r *EnqueueSodaResultParsing) TimeType() scheduler.TimeType {
	return scheduler.Every
}

func (r *EnqueueSodaResultParsing) Expression() string {
	return `1m`
}

func (r *EnqueueSodaResultParsing) Operation(_ context.Context) func() {
	return func() {
	}
}
