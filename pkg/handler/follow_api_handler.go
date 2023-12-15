package handler

import (
	"github.com/gin-gonic/gin"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/logic"
)

func followUser(appCtx *app.Context) gin.HandlerFunc {
	followLogic := logic.NewFollowLogic(appCtx)
	return func(ctx *gin.Context) {
		var req dto.FollowReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().Error(err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		errReq := followLogic.AddFollowContact(&req)
		if errReq != nil {
			appCtx.Logger().Error(errReq.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}

func unFollowUser(appCtx *app.Context) gin.HandlerFunc {
	followLogic := logic.NewFollowLogic(appCtx)
	return func(ctx *gin.Context) {
		var req dto.UnFollowReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().Error(err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		err = followLogic.RemoveFollowContact(&req)
		if err != nil {
			appCtx.Logger().Error(err.Error())
			baseDto.ResponseInternalServerError(ctx, err)
		} else {
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}
