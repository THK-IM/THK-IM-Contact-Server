package logic

import (
	baseDto "github.com/thk-im/thk-im-base-server/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/dto"
	"github.com/thk-im/thk-im-contact-server/pkg/model"
)

type ContactLogic struct {
	appCtx *app.Context
}

func NewContactLogic(appCtx *app.Context) *ContactLogic {
	return &ContactLogic{appCtx: appCtx}
}

func (l *ContactLogic) QueryContact(userId, toUserId int64) (*dto.Contact, error) {
	contact, err := l.appCtx.UserContactModel().FindOneByContactId(userId, toUserId)
	if err != nil {
		return nil, err
	}
	if contact.ContactId == 0 {
		contact.ContactId = toUserId
		contact.UserId = userId
		contact.Relation = 0
	}
	dtoContact := l.contactModel2Dto(contact)
	return dtoContact, nil
}

func (l *ContactLogic) QueryContactList(req *dto.ContactListReq) (*dto.ContactPageListResp, error) {
	userContacts, total, err := l.appCtx.UserContactModel().FindContacts(req.UId, req.RelationType, req.Count, req.Offset)
	if err != nil {
		return nil, err
	}
	res := &dto.ContactPageListResp{
		Total: total,
		Data:  make([]*dto.Contact, 0),
	}
	if len(userContacts) > 0 {
		res.Data = make([]*dto.Contact, 0)
		for _, uc := range userContacts {
			dtoContact := l.contactModel2Dto(uc)
			res.Data = append(res.Data, dtoContact)
		}
	}
	return res, nil
}

func (l *ContactLogic) QueryLatestContactList(req *dto.LatestContactListReq) (*dto.ContactListResp, error) {
	userContacts, err := l.appCtx.UserContactModel().FindLatestContacts(req.UId, req.MTime, req.Count, req.Offset)
	if err != nil {
		return nil, err
	}
	res := &dto.ContactListResp{
		Data: make([]*dto.Contact, 0),
	}
	if len(userContacts) > 0 {
		res.Data = make([]*dto.Contact, 0)
		for _, uc := range userContacts {
			dtoContact := l.contactModel2Dto(uc)
			res.Data = append(res.Data, dtoContact)
		}
	}
	return res, nil
}
func (l *ContactLogic) UpdateContactName(req *dto.UpdateContactNotName, claims baseDto.ThkClaims) error {
	return l.appCtx.UserContactModel().SetNoteName(req.UId, req.ContactId, req.NoteName)
}

func (l *ContactLogic) contactModel2Dto(contact *model.UserContact) *dto.Contact {
	return &dto.Contact{
		Id:         contact.ContactId,
		Relation:   contact.Relation,
		NoteName:   contact.NoteName,
		CreateTime: contact.CreateTime,
		UpdateTime: contact.UpdateTime,
	}
}
