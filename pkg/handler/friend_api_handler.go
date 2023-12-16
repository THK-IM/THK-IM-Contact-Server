package handler

import (
	"github.com/gin-gonic/gin"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/logic"
)

func queryFriendList(appCtx *app.Context) gin.HandlerFunc {
	friendLogic := logic.NewFriendLogic(appCtx)
	return func(ctx *gin.Context) {
		id := int64(0)
		count := int(20)
		offset := int(0)
		resp, errReq := friendLogic.QueryFriendList(id, count, offset)
		if errReq != nil {
			appCtx.Logger().Error(errReq.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}

func appFriendApply(appCtx *app.Context) gin.HandlerFunc {
	friendLogic := logic.NewFriendLogic(appCtx)
	return func(ctx *gin.Context) {
		var req dto.AddFriendReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().Error(err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := friendLogic.AddFriendApply(&req)
		if errReq != nil {
			appCtx.Logger().Error(errReq.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().Debug("appFriendApply", resp)
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}

func reviewFriendApply(appCtx *app.Context) gin.HandlerFunc {
	friendLogic := logic.NewFriendLogic(appCtx)
	return func(ctx *gin.Context) {
		var req dto.ReviewFriendApplyReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().Error(err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := friendLogic.ReviewFriendApply(&req)
		if errReq != nil {
			appCtx.Logger().Error(errReq.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().Debug("reviewFriendApply", resp)
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}
