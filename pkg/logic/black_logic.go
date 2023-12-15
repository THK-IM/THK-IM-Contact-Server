package logic

import (
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/model"
)

type BlackLogic struct {
	appCtx *app.Context
}

func NewBlackLogic(appCtx *app.Context) *BlackLogic {
	return &BlackLogic{appCtx: appCtx}
}

func (l BlackLogic) AddBlackContact(req *dto.AddBlackReq) error {
	err := l.appCtx.UserContactModel().CreateUserRelation(req.UserId, req.ContactId, model.RelationBlack)
	return err
}

func (l BlackLogic) RemoveBlackContact(req *dto.RemBlackReq) error {
	err := l.appCtx.UserContactModel().RemoveUserRelation(req.UserId, req.ContactId, model.RelationBlack)
	return err
}
