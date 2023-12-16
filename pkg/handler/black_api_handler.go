package handler

import (
	"github.com/gin-gonic/gin"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/logic"
)

func addBlack(appCtx *app.Context) gin.HandlerFunc {
	blackLogic := logic.NewBlackLogic(appCtx)
	return func(ctx *gin.Context) {
		var req dto.AddBlackReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().Errorf("addBlack %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		err = blackLogic.AddBlackContact(&req)
		if err != nil {
			appCtx.Logger().Errorf("addBlack %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, err)
		} else {
			appCtx.Logger().Infof("addBlack %v", req)
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}

func removeBlack(appCtx *app.Context) gin.HandlerFunc {
	blackLogic := logic.NewBlackLogic(appCtx)
	return func(ctx *gin.Context) {
		var req dto.RemBlackReq
		err := ctx.BindJSON(&req)
		if err != nil {
			appCtx.Logger().Errorf("removeBlack %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		err = blackLogic.RemoveBlackContact(&req)
		if err != nil {
			appCtx.Logger().Errorf("removeBlack %v %v", req, err.Error())
			appCtx.Logger().Error(err.Error())
			baseDto.ResponseInternalServerError(ctx, err)
		} else {
			appCtx.Logger().Infof("removeBlack %v", req)
			baseDto.ResponseSuccess(ctx, nil)
		}
	}
}
