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
		contactListReq := &dto.ContactListReq{}
		err := ctx.BindQuery(contactListReq)
		if err != nil {
			baseDto.ResponseBadRequest(ctx)
			return
		}
		resp, errReq := contactLogic.QueryContactList(contactListReq)
		if errReq != nil {
			appCtx.Logger().Error(errReq.Error())
			baseDto.ResponseInternalServerError(ctx, errReq)
		} else {
			baseDto.ResponseSuccess(ctx, resp)
		}
	}
}
