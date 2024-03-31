package logic

import (
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	baseErrorx "github.com/thk-im/thk-im-base-server/errorx"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	msgDto "github.com/thk-im/thk-im-msgapi-server/pkg/dto"
	msgModel "github.com/thk-im/thk-im-msgapi-server/pkg/model"
)

type SessionLogic struct {
	appCtx *app.Context
}

func NewSessionLogic(appCtx *app.Context) *SessionLogic {
	return &SessionLogic{appCtx: appCtx}
}

func (l SessionLogic) CreateSession(req *dto.CreateSessionReq, claims baseDto.ThkClaims) (*dto.CreateSessionResp, error) {
	createSessionReq := &msgDto.CreateSessionReq{
		UId:          req.UId,
		Type:         msgModel.SingleSessionType,
		EntityId:     req.ContactId,
		FunctionFlag: msgDto.FuncAll,
	}
	resp, errCreate := l.appCtx.MsgApi().CreateSession(createSessionReq, claims)
	if errCreate != nil {
		return nil, errCreate
	}
	if resp == nil || resp.SId == 0 {
		return nil, baseErrorx.ErrInternalServerError
	}

	err := l.appCtx.UserContactModel().SetSessionId(req.UId, req.ContactId, resp.SId)
	if err != nil {
		return nil, err
	}

	return &dto.CreateSessionResp{
		SId:          resp.SId,
		ParentId:     resp.ParentId,
		EntityId:     resp.EntityId,
		Type:         resp.Type,
		Name:         resp.Name,
		Remark:       resp.Remark,
		ExtData:      resp.ExtData,
		Mute:         resp.Mute,
		Role:         resp.Role,
		CTime:        resp.CTime,
		MTime:        resp.MTime,
		Top:          resp.Top,
		FunctionFlag: resp.FunctionFlag,
		Status:       resp.Status,
		IsNew:        resp.IsNew,
	}, nil
}
