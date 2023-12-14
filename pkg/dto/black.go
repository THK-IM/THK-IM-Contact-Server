package dto

type AddBlackReq struct {
	UserId int64 `json:"user_id" binding:"required"` // 对方id
}

type RemBlackReq struct {
	UserId int64 `json:"user_id" binding:"required"` // 对方id
}
