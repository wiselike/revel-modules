package app

import (
	"github.com/wiselike/revel"
)

func init() {
	revel.OnAppStart(func() {
		revel.AppLog.Info("Go to /@tests to run the tests.")
	})
}
