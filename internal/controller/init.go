package controller

import (
	"server/internal/service"
	"server/pkg/log"
)

func Initialization() {
	Api = &api{srv: service.Api, log: log.GetLog()}
}
