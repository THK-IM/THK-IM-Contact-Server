package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	baseMiddleware "github.com/thk-im/thk-im-base-server/middleware"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/logic"
)

func queryContactList(appCtx *app.Context) gin.HandlerFunc {
	contactLogic := logic.NewContactLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		req := &dto.ContactListReq{}
		err := ctx.BindQuery(req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("queryContactList %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := contactLogic.QueryContactList(req)
		if errReq != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("queryContactList %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Info("queryContactList %v %v", req, resp)
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}
