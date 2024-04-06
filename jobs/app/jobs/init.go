package jobs

import (
	"github.com/wiselike/revel"
)

var jobLog = revel.AppLog

func init() {
	revel.RegisterModuleInit(func(m *revel.Module) {
		jobLog = m.Log
	})
}
