package ocpp

import (
	"time"

	"github.com/evcc-io/evcc/util"
	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
	"github.com/lorenzodonini/ocpp-go/ocppj"
)

var instance *CS

func Instance() *CS {
	if instance == nil {
		cs := ocpp16.NewCentralSystem(nil, nil)

		instance = &CS{
			log: util.NewLogger("ocpp"),
			cps: make(map[string]*CP),
			cs:  cs,
		}

		ocppj.SetLogger(instance)

		cs.SetCoreHandler(instance)
		cs.SetNewChargePointHandler(instance.NewChargePoint)
		cs.SetChargePointDisconnectedHandler(instance.ChargePointDisconnected)
		cs.SetFirmwareManagementHandler(instance)

		go Instance().errorHandler(cs.Errors())
		go cs.Start(8887, "/{ws}")

		time.Sleep(time.Second)
	}

	return instance
}
