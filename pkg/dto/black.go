package dto

type AddBlackReq struct {
	UId       int64 `json:"u_id" binding:"required"`
	ContactId int64 `json:"contact_id" binding:"required"` // 对方id
}

type RemBlackReq struct {
	UId       int64 `json:"u_id" binding:"required"`
	ContactId int64 `json:"contact_id" binding:"required"` // 对方id
}
