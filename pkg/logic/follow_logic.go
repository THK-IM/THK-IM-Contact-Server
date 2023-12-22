package logic

import (
	"github.com/sirupsen/logrus"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
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

func (l *FollowLogic) AddFollowContact(req *dto.FollowReq, claims baseDto.ThkClaims) error {
	err := l.appCtx.UserContactModel().CreateUserRelation(req.UId, req.ContactId, model.RelationFollow)
	if err == nil {
		errSend := SendBeFollowMessage(l.appCtx, req.ContactId, req.UId, claims)
		if errSend != nil {
			l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("SendBeFollowMessage err %v", errSend)
		}
	}
	return err
}

func (l *FollowLogic) RemoveFollowContact(req *dto.UnFollowReq) error {
	err := l.appCtx.UserContactModel().RemoveUserRelation(req.UId, req.ContactId, model.RelationFollow)
	return err
}
