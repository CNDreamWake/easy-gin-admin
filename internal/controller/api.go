package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/internal/service"
	"server/internal/vo"
)

var Api = new(api)

type api struct {
	srv service.IApi
	log *zap.Logger
}

func (f *api) Test(c *gin.Context) {
	f.log.Info("info")
	err := f.srv.Test()
	if err != nil {
		vo.ResultErr(err, c)
		return
	}
	vo.OK(c)
}
