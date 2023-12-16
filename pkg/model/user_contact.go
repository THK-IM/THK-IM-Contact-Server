package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thk-im/thk-im-base-server/errorx"
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
		ApplyId      int64 `gorm:"apply_id"`
		ApplyUserId  int64 `gorm:"apply_user_id"` // 申请人id
		ToUserId     int64 `gorm:"to_user_id"`    // 被申请人id
		RelationType int8  `gorm:"relation_type"`
		ApplyStatus  int8  `gorm:"apply_status"`
		CreateTime   int64 `gorm:"create_time"`
		UpdateTime   int64 `gorm:"update_time"`
	}

	UserContactModel interface {
		FindContacts(uId int64, contactType, count, offset int) ([]*UserContact, int64, error)
		FindOneByContactId(uId, contactId int64) (*UserContact, error)
		CreateUserRelation(uId, contactId, relation int64) (err error)
		RemoveUserRelation(uId, contactId, relation int64) (err error)
		FindOneByContactApplyId(uId, applyId int64) (apply *UserContactApply, err error)
		CreateContactApply(uId, contactId int64, relationType int8) (applyId int64, err error)
		ReviewContactApply(uId, applyId int64, passed int8) (toUserApply *UserContactApply, err error)
	}

	defaultUserContactModel struct {
		shards        int64
		logger        *logrus.Entry
		db            *gorm.DB
		snowflakeNode *snowflake.Node
	}
)

func (d defaultUserContactModel) FindContacts(uId int64, contactType, count, offset int) ([]*UserContact, int64, error) {
	tableName := d.genUserContactTableName(uId)
	total := int64(0)
	userContacts := make([]*UserContact, 0)
	relation := 1 << contactType
	err := d.db.Table(tableName).Where("user_id = ? and relation & ? > 0", uId, relation).Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = d.db.Table(tableName).Where("user_id = ? and relation & ? > 0", uId, relation).Offset(offset).Limit(count).Scan(&userContacts).Error
	return userContacts, total, err
}

func (d defaultUserContactModel) FindOneByContactId(uId, contactId int64) (*UserContact, error) {
	tableName := d.genUserContactTableName(uId)
	sql := fmt.Sprintf("select * from %s where user_id = ? and contact_id = ?", tableName)
	userContact := &UserContact{}
	err := d.db.Raw(sql, uId, contactId).Scan(userContact).Error
	return userContact, err
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

func (d defaultUserContactModel) FindOneByContactApplyId(uId, applyId int64) (apply *UserContactApply, err error) {
	userTable := d.genUserContactApplyTableName(uId)
	toUserApply := &UserContactApply{}
	err = d.db.Table(userTable).Find(toUserApply).Where("user_id = ? and apply_id = ?", uId, applyId).Error
	if err != nil {
		return
	}
	apply = toUserApply
	return
}

func (d defaultUserContactModel) CreateContactApply(uId, contactId int64, relationType int8) (applyId int64, err error) {
	userTable := d.genUserContactApplyTableName(uId)
	toUserTable := d.genUserContactApplyTableName(contactId)
	now := time.Now().UnixMilli()
	applyId = d.snowflakeNode.Generate().Int64()
	tx := d.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	userContactApply := &UserContactApply{
		UserId:       uId,
		ApplyId:      applyId,
		ApplyUserId:  uId,
		ToUserId:     contactId,
		RelationType: relationType,
		ApplyStatus:  ApplyInit,
		CreateTime:   now,
		UpdateTime:   now,
	}
	err = tx.Table(userTable).Create(userContactApply).Error
	if err != nil {
		return
	}

	toUserContactApply := &UserContactApply{
		UserId:       contactId,
		ApplyId:      applyId,
		ApplyUserId:  uId,
		ToUserId:     contactId,
		RelationType: relationType,
		ApplyStatus:  ApplyInit,
		CreateTime:   now,
		UpdateTime:   now,
	}
	err = tx.Table(toUserTable).Create(toUserContactApply).Error
	return
}

func (d defaultUserContactModel) ReviewContactApply(uId, applyId int64, passed int8) (toUserApply *UserContactApply, err error) {
	userTable := d.genUserContactApplyTableName(uId)
	now := time.Now().UnixMilli()

	tx := d.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	toUserApply = &UserContactApply{}
	err = tx.Table(userTable).Where("user_id = ? and apply_id = ?", uId, applyId).Find(toUserApply).Error
	if err != nil {
		return
	}

	if toUserApply.UserId != toUserApply.ToUserId {
		err = errorx.ErrPermission
		return
	}

	if toUserApply.ApplyStatus != ApplyInit {
		return
	}

	updateMap := make(map[string]interface{})
	updateMap["apply_status"] = passed
	updateMap["update_time"] = now

	err = tx.Table(userTable).Where("user_id = ? and apply_id = ?", uId, applyId).Updates(updateMap).Error
	if err != nil {
		return
	}

	applyUserTable := d.genUserContactApplyTableName(toUserApply.ApplyUserId)
	err = tx.Table(applyUserTable).Where("user_id = ? and apply_id = ?", toUserApply.ApplyUserId, applyId).Updates(updateMap).Error
	if err != nil {
		return
	}

	if passed == ApplyPassed {
		relation := int64(1 << toUserApply.RelationType)
		err = d.createUserRelation(tx, toUserApply.ApplyUserId, toUserApply.ToUserId, relation)
	}
	return
}

func (d defaultUserContactModel) createUserRelation(tx *gorm.DB, uId, contactId, relation int64) (err error) {
	userTable := d.genUserContactTableName(uId)
	contactTable := d.genUserContactTableName(contactId)

	reverseRelation := relation << 1
	if relation == RelationFriend {
		reverseRelation = relation
	}

	relationSql := ""
	reverseRelationSql := ""
	if relation == RelationBlack {
		relationSql = fmt.Sprintf("%d", RelationBlack)
		reverseRelationSql = fmt.Sprintf("%d", RelationBeBlack)
	} else {
		relationSql = fmt.Sprintf("relation & %d | %d", 0, relation)
		reverseRelationSql = fmt.Sprintf("relation | %d", reverseRelation)
	}

	sql := "insert into %s (user_id, contact_id, relation, create_time, update_time) values (?, ?, ?, ?, ?)  on duplicate key update relation = %s, update_time = ? "
	now := time.Now().UnixMilli()
	err = tx.Exec(fmt.Sprintf(sql, userTable, relationSql), uId, contactId, relation, now, now, now).Error
	if err != nil {
		return err
	}
	err = tx.Exec(fmt.Sprintf(sql, contactTable, reverseRelationSql), contactId, uId, reverseRelation, now, now, now).Error
	return err
}

func (d defaultUserContactModel) removeUserRelation(tx *gorm.DB, uId, contactId, relation int64) (err error) {
	userTable := d.genUserContactTableName(uId)
	contactTable := d.genUserContactTableName(contactId)

	reverseRelation := relation << 1
	if relation == RelationFriend {
		reverseRelation = relation
	}

	relationSql := ""
	reverseRelationSql := ""
	if relation == RelationBlack {
		relationSql = "0"
		reverseRelationSql = "0"
	} else {
		relationSql = fmt.Sprintf("relation & (relation ^ %d)", relation)
		reverseRelationSql = fmt.Sprintf("relation & (relation ^ %d)", reverseRelation)
	}

	sql := "update %s set relation = %s, update_time = ? where user_id = ? and contact_id = ? "
	now := time.Now().UnixMilli()
	err = tx.Exec(fmt.Sprintf(sql, userTable, relationSql), now, uId, contactId).Error
	if err != nil {
		return err
	}
	err = tx.Exec(fmt.Sprintf(sql, contactTable, reverseRelationSql), now, contactId, uId).Error
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
