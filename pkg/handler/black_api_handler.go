package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	baseMiddleware "github.com/thk-im/thk-im-base-server/middleware"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/logic"
	userSdk "github.com/thk-im/thk-im-user-server/pkg/sdk"
)

func addBlack(appCtx *app.Context) gin.HandlerFunc {
	blackLogic := logic.NewBlackLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		var req dto.AddBlackReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("addBlack %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}

		requestUid := ctx.GetInt64(userSdk.UidKey)
		if requestUid > 0 && requestUid != req.UId {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("addBlack %d, %d", requestUid, req.UId)
			baseDto.ResponseForbidden(ctx)
			return
		}

		err = blackLogic.AddBlackContact(&req, claims)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("addBlack %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, err)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Infof("addBlack %v", req)
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}

func removeBlack(appCtx *app.Context) gin.HandlerFunc {
	blackLogic := logic.NewBlackLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		var req dto.RemBlackReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("removeBlack %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		requestUid := ctx.GetInt64(userSdk.UidKey)
		if requestUid > 0 && requestUid != req.UId {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("removeBlack %d, %d", requestUid, req.UId)
			baseDto.ResponseForbidden(ctx)
			return
		}
		err = blackLogic.RemoveBlackContact(&req, claims)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("removeBlack %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, err)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Infof("removeBlack %v", req)
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}
