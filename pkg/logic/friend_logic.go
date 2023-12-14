package logic

import (
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
)

type FriendLogic struct {
	appCtx *app.Context
}

func NewFriendLogic(appCtx *app.Context) *FriendLogic {
	return &FriendLogic{appCtx: appCtx}
}

func (f FriendLogic) QueryFriendList(id int64, count, offset int) (*dto.FriendListResp, error) {
	f.appCtx.Logger().Info(id, count, offset)
	return nil, nil
}

func (f FriendLogic) AddFriendApply(req *dto.AddFriendReq) (*dto.AddFriendResp, error) {
	f.appCtx.Logger().Info(req.UserId)
	return nil, nil
}

func (f FriendLogic) ReviewFriendApply(req *dto.ReviewFriendApplyReq) error {
	f.appCtx.Logger().Info(req.ApplyId, req.Pass)
	return nil
}
