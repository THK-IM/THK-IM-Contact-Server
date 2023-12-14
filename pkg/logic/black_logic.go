package logic

import (
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
)

type BlackLogic struct {
	appCtx *app.Context
}

func NewBlackLogic(appCtx *app.Context) *BlackLogic {
	return &BlackLogic{appCtx: appCtx}
}

func (l BlackLogic) AddBlack(req *dto.AddBlackReq) error {
	return nil
}

func (l BlackLogic) RemoveBlack(req *dto.RemBlackReq) error {
	return nil
}
