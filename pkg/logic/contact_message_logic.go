package logic

import (
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/model"
	msgDto "github.com/thk-im/thk-im-msgapi-server/pkg/dto"
	"time"
)

func SendBeBlackedMessage(appCtx *app.Context, uId, blackUId int64, claims baseDto.ThkClaims) error {
	body, errBody := dto.NewBeBlackedMsgBody(uId, blackUId).ToJson()
	if errBody != nil {
		return errBody
	}

	sendMsgReq := &msgDto.SendSysMessageReq{
		Type:      dto.SysMsgTypeBeBlacked,
		CTime:     time.Now().UnixMilli(),
		Body:      body,
		Receivers: []int64{uId, blackUId},
	}
	_, errSend := appCtx.MsgApi().SendSysMessage(sendMsgReq, claims)
	return errSend
}

func SendBeFollowMessage(appCtx *app.Context, uId, followId int64, claims baseDto.ThkClaims) error {
	body, errBody := dto.NewBeFollowMsgBody(uId, followId).ToJson()
	if errBody != nil {
		return errBody
	}

	sendMsgReq := &msgDto.SendSysMessageReq{
		Type:      dto.SysMsgTypeBeFollowed,
		CTime:     time.Now().UnixMilli(),
		Body:      body,
		Receivers: []int64{uId, followId},
	}
	_, errSend := appCtx.MsgApi().SendSysMessage(sendMsgReq, claims)
	return errSend
}

func SendFriendApplyMsg(appCtx *app.Context, apply *model.UserContactApply, msg string, claims baseDto.ThkClaims) error {
	body, errBody := dto.NewFriendApplyMsgBody(apply.ApplyId, apply.ApplyUserId, apply.ToUserId, msg).ToJson()
	if errBody != nil {
		return errBody
	}

	sendMsgReq := &msgDto.SendSysMessageReq{
		Type:      dto.SysMsgTypeFriendApply,
		CTime:     time.Now().UnixMilli(),
		Body:      body,
		Receivers: []int64{apply.ToUserId, apply.ApplyUserId},
	}
	_, errSend := appCtx.MsgApi().SendSysMessage(sendMsgReq, claims)
	return errSend
}

func SendFriendReviewMsg(appCtx *app.Context, apply *model.UserContactApply, msg string, claims baseDto.ThkClaims) error {
	body, errBody := dto.NewFriendReviewMsgBody(apply.ApplyId, apply.ApplyUserId, apply.ToUserId, msg, apply.ApplyStatus).ToJson()
	if errBody != nil {
		return errBody
	}

	sendMsgReq := &msgDto.SendSysMessageReq{
		Type:      dto.SysMsgTypeFriendReview,
		CTime:     time.Now().UnixMilli(),
		Body:      body,
		Receivers: []int64{apply.ApplyUserId, apply.ToUserId},
	}
	_, errSend := appCtx.MsgApi().SendSysMessage(sendMsgReq, claims)
	return errSend
}
