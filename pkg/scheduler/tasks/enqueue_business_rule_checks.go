package tasks

import (
	"context"
	"github.com/samber/do"
	"gocourse22/pkg/scheduler"
)

func CalculateFinances(_ *do.Injector) *EnqueueBusinessRuleTask {
	return &EnqueueBusinessRuleTask{}
}

type EnqueueBusinessRuleTask struct {
}

func (r *EnqueueBusinessRuleTask) TimeType() scheduler.TimeType {
	return scheduler.Every
}

func (r *EnqueueBusinessRuleTask) Expression() string {
	return `1m`
}

func (r *EnqueueBusinessRuleTask) Operation(_ context.Context) func() {
	return func() {}
}
