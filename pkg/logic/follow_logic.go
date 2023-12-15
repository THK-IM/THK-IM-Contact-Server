package logic

import (
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/model"
)

type FollowLogic struct {
	appCtx *app.Context
}

func NewFollowLogic(appCtx *app.Context) *FollowLogic {
	return &FollowLogic{appCtx: appCtx}
}

func (f FollowLogic) AddFollowContact(req *dto.FollowReq) error {
	err := f.appCtx.UserContactModel().CreateUserRelation(req.UserId, req.ContactId, model.RelationFollow)
	return err
}

func (f FollowLogic) RemoveFollowContact(req *dto.UnFollowReq) error {
	err := f.appCtx.UserContactModel().RemoveUserRelation(req.UserId, req.ContactId, model.RelationFollow)
	return err
}
