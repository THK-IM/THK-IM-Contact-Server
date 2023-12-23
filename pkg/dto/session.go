package dto

type CreateSessionReq struct {
	UId       int64 `json:"u_id"`
	ContactId int64 `json:"contact_id"`
}

type CreateSessionResp struct {
	SId      int64   `json:"s_id"`
	ParentId int64   `json:"parent_id"`
	EntityId int64   `json:"entity_id"`
	Type     int     `json:"type"`
	Name     string  `json:"name"`
	Remark   string  `json:"remark"`
	ExtData  *string `json:"ext_data"`
	Mute     int     `json:"mute"`
	Role     int     `json:"role"`
	CTime    int64   `json:"c_time"`
	MTime    int64   `json:"m_time"`
	Top      int64   `json:"top"`
	Status   int     `json:"status"`
	IsNew    bool    `json:"is_new"` // 如果之前已经创建，false
}
