package model

import (
	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	RelationTypeBlack = iota + 1
	RelationTypeBeBlack
	RelationTypeFollow
	RelationTypeBeFollow
	RelationTypeFriend

	RelationBlack    = 1 << RelationTypeBlack
	RelationBeBlack  = 1 << RelationTypeBeBlack
	RelationFollow   = 1 << RelationTypeFollow
	RelationBeFollow = 1 << RelationTypeBeFollow
	RelationFriend   = 1 << RelationTypeFriend
)

const (
	ApplyInit = iota
	ApplyPassed
	ApplyRejected
)

type (
	UserContact struct {
		UserId     int64 `json:"user_id"`
		ContactId  int64 `json:"contact_id"`
		Relation   int64 `json:"relation"`
		CreateTime int64 `json:"create_time"`
		UpdateTime int64 `json:"update_time"`
	}

	UserContactApply struct {
		Id          int64 `json:"id"`
		UserId      int64 `json:"user_id"`
		ContactId   int64 `json:"contact_id"`
		ApplyType   int8  `json:"apply_type"`
		ApplyStatus int8  `json:"apply_status"`
		CreateTime  int64 `json:"create_time"`
		UpdateTime  int64 `json:"update_time"`
	}

	UserContactModel interface {
		FindOneByContactId(uId, contactId int64) (*UserContact, error)
		CreateUserRelation(uId, contactId, relation int64) (*UserContact, error)
		RemoveUserRelation(uId, contactId, relation int64) (*UserContact, error)
		CreateContactApply(uId, contactId int64, applyType int8) error
		ReviewContactApply(uId, contactId int64, applyType int8, status int8) error
	}

	defaultUserContactModel struct {
		shards        int64
		logger        *logrus.Entry
		db            *gorm.DB
		snowflakeNode *snowflake.Node
	}
)

func (d defaultUserContactModel) FindOneByContactId(uId, contactId int64) (*UserContact, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultUserContactModel) CreateUserRelation(uId, contactId, relation int64) (*UserContact, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultUserContactModel) RemoveUserRelation(uId, contactId, relation int64) (*UserContact, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultUserContactModel) CreateContactApply(uId, contactId int64, applyType int8) error {
	// TODO implement me
	panic("implement me")
}

func (d defaultUserContactModel) ReviewContactApply(uId, contactId int64, applyType int8, status int8) error {
	// TODO implement me
	panic("implement me")
}

func NewUserContactModel(db *gorm.DB, logger *logrus.Entry, snowflakeNode *snowflake.Node, shards int64) UserContactModel {
	return defaultUserContactModel{
		shards:        shards,
		logger:        logger,
		db:            db,
		snowflakeNode: snowflakeNode,
	}
}
