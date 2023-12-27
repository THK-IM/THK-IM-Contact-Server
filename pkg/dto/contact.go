package dto

type ContactListReq struct {
	UId          int64 `form:"u_id"`
	RelationType int   `form:"relation_type"`
	Count        int   `form:"count"`
	Offset       int   `form:"offset"`
}

type LatestContactListReq struct {
	UId    int64 `form:"u_id"`
	Count  int   `form:"count"`
	Offset int   `form:"offset"`
	MTime  int64 `json:"m_time"`
}

type Contact struct {
	Id         int64  `json:"id"`
	Relation   int64  `json:"relation"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Sex        int8   `json:"sex"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type ContactPageListResp struct {
	Total int64      `json:"total"`
	Data  []*Contact `json:"data"`
}

type ContactListResp struct {
	Data []*Contact `json:"data"`
}
