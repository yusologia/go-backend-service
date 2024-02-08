package Console

import (
	"github.com/go-co-op/gocron"
	"github.com/yusologia/go-core/console"
)

func Schedules(sch *gocron.Scheduler) {
	//addSchedule(sch.Every(1).Day().At("00:05"), Command.TestCommand{})
}

func addSchedule(schedule *gocron.Scheduler, command console.BaseInterface) {
	schedule.Do(command.Handle)
}
