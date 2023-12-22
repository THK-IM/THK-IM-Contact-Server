package dto

type FollowReq struct {
	UId       int64 `json:"u_id" binding:"required"`
	ContactId int64 `json:"contact_id" binding:"required"` // 对方id
}

type UnFollowReq struct {
	UId       int64 `json:"u_id" binding:"required"`
	ContactId int64 `json:"contact_id" binding:"required"` // 对方id
}
