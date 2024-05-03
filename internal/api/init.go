package api

import (
	"server/internal/service"
	"server/pkg/log"
)

func Initialization() {
	Api = &api{srv: service.Frontend, log: log.GetLog()}
}
