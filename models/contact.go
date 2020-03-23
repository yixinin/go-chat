package models

import (
	"fmt"
	"go-lib/db"
	"time"
)

const (
	NotifyLevelMute    uint8 = 0 //静音 不提示消息
	NotifyLevelDisplay uint8 = 1 //仅显示， 不提示
	NotifyLevelNormal  uint8 = 2 //打开提示消息
	NotifyLevelPush    uint8 = 3 //推送消息
)

const (
	ContactStatusWaiting int32 = 1 + iota
	ContactStatusRejected
	ContactStatusApproved
	ContactStatusExpired
)

type Contact struct {
	Id       int64 `xorm:"pk autoincr"`
	UserId   int64 //x用户的联系人列表
	ToUserId int64
	Remarks  string //预设备注备注
	Status   int32  `xorm:"default(1)"` //状态 1=待通过 2=已拒绝 3=已通过 4=已过期
	Message  string //验证消息

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func (Contact) TableName() string {
	return "contact"
}

func AddContact(uid, toUid int64, msg, remarks string) (int64, error) {
	var contact = &Contact{
		UserId:   uid,
		ToUserId: toUid,
		Remarks:  remarks,
		Message:  msg,
	}
	n, err := db.Mysql.Insert(contact)
	return n, err
}

func ApproveContact(id int64, add bool, remarks string) (bool, error) {
	sess := db.Mysql.NewSession()
	defer sess.Close()
	err := sess.Begin()
	if err != nil {
		return false, err
	}
	//通过验证
	var contact = new(Contact)
	contact.Status = ContactStatusApproved
	if !add {
		contact.Status = ContactStatusRejected
	}
	n, err := sess.ID(id).Update(contact)
	if err != nil || n != 1 {
		sess.Rollback()
		return false, err
	}
	if !add { //拒绝
		err := sess.Commit()
		return err == nil, err
	}

	//查找联系人
	ok, err := sess.ID(id).Get(contact)
	if err != nil {
		sess.Rollback()
		return false, err
	}
	if !ok {
		sess.Rollback()
		return false, err
	}
	if contact.Status != ContactStatusWaiting {
		return false, nil
	}

	// 通过
	// 添加A的联系人B
	var userContactA = &UserContact{
		Uid: contact.UserId,

		UserId:  contact.ToUserId,
		Remarks: contact.Remarks,
		BeAdded: false,
	}

	ucaId, err := sess.Insert(userContactA)
	if err != nil {
		sess.Rollback()
		return false, err
	}
	if ucaId == 0 {
		sess.Rollback()
		return false, nil
	}

	// 添加B的联系人 A
	var userContactB = &UserContact{
		Uid: contact.ToUserId,

		UserId:  contact.UserId,
		Remarks: remarks,
		BeAdded: true,
	}
	ucbId, err := sess.Insert(userContactB)
	if err != nil {
		sess.Rollback()
		return false, err
	}
	if ucbId == 0 {
		sess.Rollback()
		return false, nil
	}
	err = sess.Commit()
	return err == nil, err
}

func DeleteContact(uid, id int64) (bool, error) {
	var sess = db.Mysql.NewSession()
	defer sess.Close()
	err := sess.Begin()
	defer sess.Rollback()
	if err != nil {
		return false, err
	}
	var addAId, addBId int64
	//查找A联系人
	var userContactA = &UserContact{
		Id:  id,
		Uid: uid,
	}
	ok, err := sess.ID(id).Get(userContactA)
	if err != nil {
		return false, err
	}
	if !ok {
		return ok, nil
	}
	userBId := userContactA.UserId
	if userContactA.BeAdded { //被添加
		addBId = userContactA.Uid
		addAId = userContactA.UserId
	} else {
		addBId = userContactA.UserId
		addAId = userContactA.Uid
	}

	// 删除A的联系人列表
	n, err := sess.ID(id).Delete(&UserContact{Uid: uid})
	if err != nil {
		return false, err
	}
	if n != 1 {
		return false, nil
	}

	//查找B联系人
	var userContactB = &UserContact{
		Uid:    userBId,
		UserId: uid,
	}
	ok, err = sess.Get(userContactB)
	if err != nil {
		return false, err
	}
	if !ok {
		return ok, nil
	}
	n, err = sess.ID(userContactB).Delete(&UserContact{Uid: userBId})
	if err != nil {
		return false, err
	}
	if n != 1 {
		return false, nil
	}

	//删除联系人关系
	var contact = &Contact{
		UserId:   addAId,
		ToUserId: addBId,
	}
	n, err = sess.Delete(contact)
	if err != nil {
		return false, err
	}
	if n != 1 {
		return false, nil
	}
	err = sess.Commit()
	return err == nil, err
}

//集合名：user_contact_userId
type UserContact struct {
	Uid         int64 `xorm:"-"`
	Id          int64 `xorm:"pk autoincr"`
	UserId      int64 `xorm:"unique"` //联系人
	Sort        int   `xorm:"default(1)"`
	IsFavorites bool
	Remarks     string //备注 默认为用户nickname
	NotifyLevel uint8
	BeAdded     bool //是否被动添加

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func (m UserContact) TableName() string {
	if m.Uid == 0 {
		return ""
	}
	return fmt.Sprintf("user_contact_%d", m.Uid)
}

func GetContactByUserId(uid, toUid int64) (*UserContact, bool, error) {
	var m = &UserContact{
		Uid:    uid,
		UserId: toUid,
	}
	ok, err := db.Mysql.Get(m)
	if !ok {
		return m, ok, err
	}
	if err != nil {
		return m, ok, err
	}
	return m, ok, err
}

func UpdateContact(uid, id int64, remarks string) (bool, error) {
	var userContact = &UserContact{
		Uid:     uid,
		Remarks: remarks,
	}
	n, err := db.Mysql.ID(id).Update(userContact)
	return n == 1, err
}
