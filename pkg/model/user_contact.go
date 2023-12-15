package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thk-im/thk-im-base-server/snowflake"
	"gorm.io/gorm"
	"time"
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
	ApplyInit = iota + 1
	ApplyPassed
	ApplyRejected
)

type (
	UserContact struct {
		UserId     int64 `gorm:"user_id"`
		ContactId  int64 `gorm:"contact_id"`
		Relation   int64 `gorm:"relation"`
		CreateTime int64 `gorm:"create_time"`
		UpdateTime int64 `gorm:"update_time"`
	}

	UserContactApply struct {
		UserId       int64 `gorm:"user_id"`
		ApplyUserId  int64 `gorm:"apply_user_id"` // 申请人id
		ToUserId     int8  `gorm:"to_user_id"`    // 被申请人id
		RelationType int8  `gorm:"relation_type"`
		ApplyStatus  int8  `gorm:"apply_status"`
		CreateTime   int64 `gorm:"create_time"`
		UpdateTime   int64 `gorm:"update_time"`
	}

	UserContactModel interface {
		FindOneByContactId(uId, contactId int64) (*UserContact, error)
		CreateUserRelation(uId, contactId, relation int64) (err error)
		RemoveUserRelation(uId, contactId, relation int64) (err error)
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

func (d defaultUserContactModel) CreateUserRelation(uId, contactId, relation int64) (err error) {
	tx := d.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	return d.createUserRelation(tx, uId, contactId, relation)
}

func (d defaultUserContactModel) RemoveUserRelation(uId, contactId, relation int64) (err error) {
	tx := d.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	return d.removeUserRelation(tx, uId, contactId, relation)
}
func (d defaultUserContactModel) CreateContactApply(uId, contactId int64, applyType int8) error {
	return nil
}

func (d defaultUserContactModel) ReviewContactApply(uId, contactId int64, applyType int8, status int8) error {
	return nil
}

func (d defaultUserContactModel) createUserRelation(tx *gorm.DB, uId, contactId, relation int64) (err error) {
	userTable := d.genUserContactTableName(uId)
	contactTable := d.genUserContactTableName(contactId)
	sql := "insert into %s (user_id, contact_id, relation, create_time, update_time) values (?, ?, ?, ?, ?)  on duplicate key update relation = relation | ?, update_time = ? "
	now := time.Now().UnixMilli()
	err = tx.Exec(fmt.Sprintf(sql, userTable), uId, contactId, relation, now, now, relation, now).Error
	if err != nil {
		return err
	}
	reverseRelation := relation << 1
	if relation == RelationFriend {
		reverseRelation = relation
	}
	err = tx.Exec(fmt.Sprintf(sql, contactTable), contactId, uId, reverseRelation, now, now, reverseRelation, now).Error
	return err
}

func (d defaultUserContactModel) removeUserRelation(tx *gorm.DB, uId, contactId, relation int64) (err error) {
	userTable := d.genUserContactTableName(uId)
	contactTable := d.genUserContactTableName(contactId)
	sql := "update %s set relation = relation & (relation ^ ?), update = ? where user_id = ? and contact_id = ? "
	now := time.Now().UnixMilli()
	err = tx.Exec(fmt.Sprintf(sql, userTable), relation, now, uId, contactId).Error
	if err != nil {
		return err
	}
	reverseRelation := relation << 1
	if relation == RelationFriend {
		reverseRelation = relation
	}
	err = tx.Exec(fmt.Sprintf(sql, contactTable), reverseRelation, now, uId, contactId).Error
	return err
}

func (d defaultUserContactModel) genUserContactTableName(uId int64) string {
	return fmt.Sprintf("user_contact_%d", uId%d.shards)
}

func (d defaultUserContactModel) genUserContactApplyTableName(uId int64) string {
	return fmt.Sprintf("user_contact_apply_%d", uId%d.shards)
}

func NewUserContactModel(db *gorm.DB, logger *logrus.Entry, snowflakeNode *snowflake.Node, shards int64) UserContactModel {
	return defaultUserContactModel{
		shards:        shards,
		logger:        logger,
		db:            db,
		snowflakeNode: snowflakeNode,
	}
}
