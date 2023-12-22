package logic

import (
	"github.com/sirupsen/logrus"
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/model"
	msgDto "github.com/thk-im/thk-im-msgapi-server/pkg/dto"
	msgModel "github.com/thk-im/thk-im-msgapi-server/pkg/model"
)

type BlackLogic struct {
	appCtx *app.Context
}

func NewBlackLogic(appCtx *app.Context) *BlackLogic {
	return &BlackLogic{appCtx: appCtx}
}

func (l *BlackLogic) AddBlackContact(req *dto.AddBlackReq, claims baseDto.ThkClaims) error {
	err := l.appCtx.UserContactModel().CreateUserRelation(req.UId, req.ContactId, model.RelationBlack)
	if err == nil {
		queryUserSessionReq := &msgDto.QueryUserSessionReq{
			UId:      req.UId,
			EntityId: req.ContactId,
			Type:     msgModel.SingleSessionType,
		}
		userSession, errQuery := l.appCtx.MsgApi().QueryUserSession(queryUserSessionReq, claims)
		if errQuery == nil {
			if userSession.SId > 0 && userSession.Status&msgModel.RejectBitInUserSessionStatus == 0 {
				newStatus := userSession.Status | msgModel.RejectBitInUserSessionStatus
				updateUserSessionReq := &msgDto.UpdateUserSessionReq{
					UId:    req.UId,
					SId:    userSession.SId,
					Status: &newStatus,
				}
				errUpdate := l.appCtx.MsgApi().UpdateUserSession(updateUserSessionReq, claims)
				if errUpdate != nil {
					l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("UpdateUserSession err %v", errUpdate)
				}
			}
		} else {
			l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("QueryUserSession err %v", errQuery)
		}

		errSend := SendBeBlackedMessage(l.appCtx, req.ContactId, req.UId, claims)
		if errSend != nil {
			l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("SendBeBlackedMessage err %v", errSend)
		}
	}
	return err
}

func (l *BlackLogic) RemoveBlackContact(req *dto.RemBlackReq, claims baseDto.ThkClaims) error {
	err := l.appCtx.UserContactModel().RemoveUserRelation(req.UId, req.ContactId, model.RelationBlack)
	if err == nil {
		queryUserSessionReq := &msgDto.QueryUserSessionReq{
			UId:      req.UId,
			EntityId: req.ContactId,
			Type:     msgModel.SingleSessionType,
		}
		userSession, errQuery := l.appCtx.MsgApi().QueryUserSession(queryUserSessionReq, claims)
		if errQuery == nil {
			if userSession.SId > 0 && userSession.Status&msgModel.RejectBitInUserSessionStatus > 0 {
				newStatus := userSession.Status & (userSession.Status ^ msgModel.RejectBitInUserSessionStatus)
				updateUserSessionReq := &msgDto.UpdateUserSessionReq{
					UId:    req.UId,
					SId:    userSession.SId,
					Status: &newStatus,
				}
				errUpdate := l.appCtx.MsgApi().UpdateUserSession(updateUserSessionReq, claims)
				if errUpdate != nil {
					l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("UpdateUserSession err %v", errUpdate)
				}
			}
		} else {
			l.appCtx.Logger().WithFields(logrus.Fields(claims)).Errorf("QueryUserSession err %v", errQuery)
		}
	}
	return err
}
