package dto

type AddFriendReq struct {
	UId       int64  `json:"u_id" binding:"required"`       // 用户id
	ContactId int64  `json:"contact_id" binding:"required"` // 对方id
	Channel   int8   `json:"channel"`                       // 渠道
	Msg       string `json:"msg"`                           // 申请留言
}

type AddFriendResp struct {
	ToUserId    int64  `json:"to_user_id"`             // 对方id
	Relation    int64  `json:"relation"`               // 好友关系
	ApplyId     *int64 `json:"apply_id,omitempty"`     // 好友申请状态
	ApplyStatus *int8  `json:"apply_status,omitempty"` // 好友申请状态
}

type ReviewFriendApplyReq struct {
	UId     int64  `json:"u_id" binding:"required"`     // 用户id
	ApplyId int64  `json:"apply_id" binding:"required"` // 好友申请 申请id
	Pass    int8   `json:"pass" binding:"required"`     // 是否通过 1 待审核 2通过 3驳回
	Msg     string `json:"msg"`
}

type ReviewFriendResp struct {
	ToUserId    int64  `json:"to_user_id"`             // 对方id
	Relation    int64  `json:"relation"`               // 好友关系
	ApplyId     *int64 `json:"apply_id,omitempty"`     // 好友申请状态
	ApplyStatus *int8  `json:"apply_status,omitempty"` // 好友申请状态
}
