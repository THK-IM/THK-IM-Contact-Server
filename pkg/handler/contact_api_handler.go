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

func createSession(appCtx *app.Context) gin.HandlerFunc {
	sessionLogic := logic.NewSessionLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		var req dto.CreateSessionReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("createSession %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		requestUid := ctx.GetInt64(userSdk.UidKey)
		if requestUid > 0 && requestUid != req.UId {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("createSession %d, %d", requestUid, req.UId)
			baseDto.ResponseForbidden(ctx)
			return
		}
		resp, errReq := sessionLogic.CreateSession(&req, claims)
		if errReq != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("createSession %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Infof("createSession %v %v", req, resp)
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}

func setNoteName(appCtx *app.Context) gin.HandlerFunc {
	contactLogic := logic.NewContactLogic(appCtx)
	return func(ctx *gin.Context) {
		claims := ctx.MustGet(baseMiddleware.ClaimsKey).(baseDto.ThkClaims)
		var req dto.UpdateContactNotName
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("setNoteName %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		requestUid := ctx.GetInt64(userSdk.UidKey)
		if requestUid > 0 && requestUid != req.UId {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("setNoteName %d, %d", requestUid, req.UId)
			baseDto.ResponseForbidden(ctx)
			return
		}
		errReq := contactLogic.UpdateContactName(&req, claims)
		if errReq != nil {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("setNoteName %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().WithFields(logrus.Fields(claims)).Infof("setNoteName %v", req)
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}
