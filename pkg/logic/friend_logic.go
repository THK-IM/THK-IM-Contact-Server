package logic

import (
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

func (f FriendLogic) AddFriendApply(req *dto.AddFriendReq) (*dto.AddFriendResp, error) {
	addFriendResp := &dto.AddFriendResp{
		ToUserId: req.ContactId,
	}
	userContact, errQuery := f.appCtx.UserContactModel().FindOneByContactId(req.UserId, req.ContactId)
	if errQuery != nil {
		return nil, errQuery
	}
	addFriendResp.Relation = userContact.Relation

	if userContact.Relation&model.RelationFriend > 0 {
		return addFriendResp, nil
	}
	id, err := f.appCtx.UserContactModel().CreateContactApply(req.UserId, req.ContactId, model.RelationTypeFriend)
	if err != nil {
		return nil, err
	}
	addFriendResp.ApplyId = &id
	status := int8(model.ApplyInit)
	addFriendResp.ApplyStatus = &status

	return addFriendResp, nil
}

func (f FriendLogic) ReviewFriendApply(req *dto.ReviewFriendApplyReq) (*dto.ReviewFriendResp, error) {
	apply, errDb := f.appCtx.UserContactModel().ReviewContactApply(req.UserId, req.ApplyId, req.Pass)
	if errDb != nil {
		return nil, errDb
	}
	if apply.ApplyId != req.ApplyId {
		return nil, errorx.ErrParamsError
	}
	resp := &dto.ReviewFriendResp{
		ToUserId:    apply.ApplyUserId,
		ApplyId:     &apply.ApplyId,
		ApplyStatus: &apply.ApplyStatus,
	}
	userContact, errQuery := f.appCtx.UserContactModel().FindOneByContactId(req.UserId, apply.ApplyUserId)
	if errQuery == nil {
		resp.Relation = userContact.Relation
	}
	return resp, nil
}
