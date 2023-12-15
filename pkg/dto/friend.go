package dto

type AddFriendReq struct {
	UserId    int64 `json:"user_id" binding:"required"`
	ContactId int64 `json:"contact_id" binding:"required"` // 对方id
}

type AddFriendResp struct {
	ToUserId     int64  `json:"to_user_id"`    // 对方id
	ApplyId      *int64 `json:"apply_id"`      // 好友申请 申请id
	FriendStatus int8   `json:"friend_status"` // 好友状态 1 好友 0 非好友
	BlackStatus  int8   `json:"black_status"`  // 黑名单状态 1 你是对方的黑名单，2 对方在你的黑名单中 0 非黑名单
}

type ReviewFriendApplyReq struct {
	ApplyId int64 `json:"apply_id" binding:"required"` // 好友申请 申请id
	Pass    int8  `json:"pass" binding:"required"`     // 是否通过 1通过 0驳回
}

type Friend struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Sex      int8   `json:"sex"`
}

type FriendListResp struct {
	Total int      `json:"total"`
	Data  []Friend `json:"data"`
}
