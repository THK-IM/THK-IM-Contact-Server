package dto

type FollowReq struct {
	UserId int64 `json:"user_id" binding:"required"` // 对方id
}

type UnFollowReq struct {
	UserId int64 `json:"user_id" binding:"required"` // 对方id
}
