package dto

type AddFriendReq struct {
	UserId    int64 `json:"user_id" binding:"required"`
	ContactId int64 `json:"contact_id" binding:"required"` // 对方id
}

type AddFriendResp struct {
	ToUserId    int64  `json:"to_user_id"`             // 对方id
	Relation    int64  `json:"relation"`               // 好友关系
	ApplyId     *int64 `json:"apply_id,omitempty"`     // 好友申请状态
	ApplyStatus *int8  `json:"apply_status,omitempty"` // 好友申请状态
}

type ReviewFriendApplyReq struct {
	UserId  int64 `json:"user_id" binding:"required"`  // 用户id
	ApplyId int64 `json:"apply_id" binding:"required"` // 好友申请 申请id
	Pass    int8  `json:"pass" binding:"required"`     // 是否通过 2通过 3驳回
}

type ReviewFriendResp struct {
	ToUserId    int64  `json:"to_user_id"`             // 对方id
	Relation    int64  `json:"relation"`               // 好友关系
	ApplyId     *int64 `json:"apply_id,omitempty"`     // 好友申请状态
	ApplyStatus *int8  `json:"apply_status,omitempty"` // 好友申请状态
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
