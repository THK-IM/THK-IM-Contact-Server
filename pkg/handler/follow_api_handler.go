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

func followUser(appCtx *app.Context) gin.HandlerFunc {
	followLogic := logic.NewFollowLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		var req dto.FollowReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("followUser %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		errReq := followLogic.AddFollowContact(&req)
		if errReq != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("followUser %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Infof("followUser %v", req)
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}

func unFollowUser(appCtx *app.Context) gin.HandlerFunc {
	followLogic := logic.NewFollowLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		var req dto.UnFollowReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("unFollowUser %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		err = followLogic.RemoveFollowContact(&req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("unFollowUser %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, err)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Infof("unFollowUser %v", req)
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}
