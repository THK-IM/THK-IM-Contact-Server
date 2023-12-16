package handler

import (
	"github.com/gin-gonic/gin"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/logic"
)

func appFriendApply(appCtx *app.Context) gin.HandlerFunc {
	friendLogic := logic.NewFriendLogic(appCtx)
	return func(ctx *gin.Context) {
		var req dto.AddFriendReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().Errorf("appFriendApply %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := friendLogic.AddFriendApply(&req)
		if errReq != nil {
			appCtx.Logger().Errorf("appFriendApply %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().Infof("appFriendApply %v %v", req, resp)
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
			appCtx.Logger().Errorf("reviewFriendApply %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := friendLogic.ReviewFriendApply(&req)
		if errReq != nil {
			appCtx.Logger().Errorf("reviewFriendApply %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().Infof("reviewFriendApply %v %v", req, resp)
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}
