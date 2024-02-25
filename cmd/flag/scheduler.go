package flag

import "github.com/urfave/cli/v2"

const (
	SchedulerEnable = "scheduler-enable"
)

func SchedulerEnableFlag() cli.Flag {
	return &cli.BoolFlag{
		Category:    "scheduler",
		Name:        SchedulerEnable,
		Usage:       "Scheduler enable",
		EnvVars:     []string{"SCHEDULER_ENABLE"},
		Value:       false,
		DefaultText: "false",
	}
}
