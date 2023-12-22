package dto

import "encoding/json"

const (
	SysMsgTypeBeBlacked    = -20 // 被人拉黑消息
	SysMsgTypeBeFollowed   = -21 // 被人关注消息
	SysMsgTypeFriendApply  = -22 // 好友申请消息
	SysMsgTypeFriendReview = -23 // 好友审核消息
)

type BeBlackedMsgBody struct {
	UId      int64 `json:"u_id"`       // 本人id
	BlackUId int64 `json:"black_u_id"` // 设置本人为黑名单的用户id
}

func NewBeBlackedMsgBody(uId, blackId int64) *BeBlackedMsgBody {
	return &BeBlackedMsgBody{
		UId:      uId,
		BlackUId: blackId,
	}
}

func (g *BeBlackedMsgBody) ToJson() (string, error) {
	d, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

type BeFollowMsgBody struct {
	UId      int64 `json:"u_id"`      // 本人id
	FollowId int64 `json:"follow_id"` // 关注人id
}

func NewBeFollowMsgBody(uId, followId int64) *BeFollowMsgBody {
	return &BeFollowMsgBody{
		UId:      uId,
		FollowId: followId,
	}
}

func (g *BeFollowMsgBody) ToJson() (string, error) {
	d, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

type FriendApplyMsgBody struct {
	ApplyId  int64  `json:"apply_id"`
	ApplyUId int64  `json:"apply_u_id"`
	ToUId    int64  `json:"to_u_id"`
	Msg      string `json:"msg"`
}

func NewFriendApplyMsgBody(applyId, applyUId, toUId int64, msg string) *FriendApplyMsgBody {
	return &FriendApplyMsgBody{
		ApplyId:  applyId,
		ApplyUId: applyUId,
		ToUId:    toUId,
		Msg:      msg,
	}
}

func (g *FriendApplyMsgBody) ToJson() (string, error) {
	d, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

type FriendReviewMsgBody struct {
	ApplyId  int64  `json:"apply_id"`
	ApplyUId int64  `json:"apply_u_id"`
	ToUId    int64  `json:"to_u_id"`
	Msg      string `json:"msg"`
	Passed   int8   `json:"passed"`
}

func NewFriendReviewMsgBody(applyId, applyUId, toUId int64, msg string, passed int8) *FriendReviewMsgBody {
	return &FriendReviewMsgBody{
		ApplyId:  applyId,
		ApplyUId: applyUId,
		ToUId:    toUId,
		Msg:      msg,
		Passed:   passed,
	}
}

func (g *FriendReviewMsgBody) ToJson() (string, error) {
	d, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
