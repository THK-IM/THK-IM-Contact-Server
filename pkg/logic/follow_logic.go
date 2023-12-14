package logic

import (
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
)

type FollowLogic struct {
	appCtx *app.Context
}

func NewFollowLogic(appCtx *app.Context) *FollowLogic {
	return &FollowLogic{appCtx: appCtx}
}

func (f FollowLogic) Follow(req *dto.FollowReq) error {
	return nil
}

func (f FollowLogic) UnFollow(req *dto.UnFollowReq) error {
	return nil
}
