package logic

import (
	"github.com/sirupsen/logrus"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-base-server/errorx"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/model"
)

type FriendLogic struct {
	appCtx *app.Context
}

func NewFriendLogic(appCtx *app.Context) *FriendLogic {
	return &FriendLogic{appCtx: appCtx}
}

func (l *FriendLogic) AddFriendApply(req *dto.AddFriendReq, claims baseDto.ThkClaims) (*dto.AddFriendResp, error) {
	userContact, errQuery := l.appCtx.UserContactModel().FindOneByContactId(req.UId, req.ContactId)
	if errQuery != nil {
		return nil, errQuery
	}

	if userContact.Relation&model.RelationFriend > 0 {
		return &dto.AddFriendResp{
			ToUserId: req.ContactId,
			Relation: userContact.Relation,
		}, nil
	}
	apply, err := l.appCtx.UserContactModel().CreateContactApply(req.UId, req.ContactId, model.RelationTypeFriend, req.Channel)
	if err != nil {
		return nil, err
	}
	errSend := SendFriendApplyMsg(l.appCtx, apply, req.Msg, claims)
	if errSend != nil {
		l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("SendFriendApplyMsg err %v", errSend)
	}
	return &dto.AddFriendResp{
		ToUserId:    req.ContactId,
		Relation:    userContact.Relation,
		ApplyId:     &apply.ApplyId,
		ApplyStatus: &apply.ApplyStatus,
	}, nil
}

func (l *FriendLogic) ReviewFriendApply(req *dto.ReviewFriendApplyReq, claims baseDto.ThkClaims) (*dto.ReviewFriendResp, error) {
	apply, errDb := l.appCtx.UserContactModel().ReviewContactApply(req.UId, req.ApplyId, req.Pass)
	if errDb != nil {
		return nil, errDb
	}
	if apply.ApplyId != req.ApplyId {
		return nil, errorx.ErrParamsError
	}

	errSend := SendFriendReviewMsg(l.appCtx, apply, req.Msg, claims)
	if errSend != nil {
		l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("SendFriendReviewMsg err %v", errSend)
	}

	if apply.ApplyStatus == model.ApplyPassed {
		errChat := StartChat(l.appCtx, apply.ToUserId, apply.ApplyUserId, "你们已经是好友，可以开始聊天了", claims)
		if errChat != nil {
			l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("StartChat err %v", errChat)
		}
	}

	resp := &dto.ReviewFriendResp{
		ToUserId:    apply.ApplyUserId,
		ApplyId:     &apply.ApplyId,
		ApplyStatus: &apply.ApplyStatus,
	}
	userContact, errQuery := l.appCtx.UserContactModel().FindOneByContactId(req.UId, apply.ApplyUserId)
	if errQuery == nil {
		resp.Relation = userContact.Relation
	}
	return resp, nil
}
