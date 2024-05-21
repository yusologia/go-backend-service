package Console

import (
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/console"
	"service/app/Console/Command"
)

func RegisterCommand(cobraCmd *cobra.Command) {
	console.Commands(cobraCmd, []console.BaseInterface{
		&Command.TestCommand{},
	})
}

func RegisterSchedule(sch *gocron.Scheduler) {
	//addSchedule(sch.Every(1).Minute(), &Command.TestCommand{})
}

func addSchedule(schedule *gocron.Scheduler, command console.BaseInterface) {
	schedule.Do(command.Handle)
}
