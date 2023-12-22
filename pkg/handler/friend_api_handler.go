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

func appFriendApply(appCtx *app.Context) gin.HandlerFunc {
	friendLogic := logic.NewFriendLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		var req dto.AddFriendReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().Errorf("appFriendApply %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := friendLogic.AddFriendApply(&req)
		if errReq != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("appFriendApply %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Infof("appFriendApply %v %v", req, resp)
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}

func reviewFriendApply(appCtx *app.Context) gin.HandlerFunc {
	friendLogic := logic.NewFriendLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		var req dto.ReviewFriendApplyReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("reviewFriendApply %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := friendLogic.ReviewFriendApply(&req)
		if errReq != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("reviewFriendApply %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Infof("reviewFriendApply %v %v", req, resp)
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}
