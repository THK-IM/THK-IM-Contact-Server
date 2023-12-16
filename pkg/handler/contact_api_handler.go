package handler

import (
	"github.com/gin-gonic/gin"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/logic"
)

func queryContactList(appCtx *app.Context) gin.HandlerFunc {
	contactLogic := logic.NewContactLogic(appCtx)
	return func(ctx *gin.Context) {
		req := &dto.ContactListReq{}
		err := ctx.BindQuery(req)
		if err != nil {
			appCtx.Logger().Errorf("queryContactList %v", err.Error())
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := contactLogic.QueryContactList(req)
		if errReq != nil {
			appCtx.Logger().Errorf("queryContactList %v %v", req, err.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			appCtx.Logger().Info("queryContactList %v %v", req, resp)
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}
