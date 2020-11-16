package model

import (
	. "github.com/huahuayu/go-gin-app/common/db"
	log "github.com/sirupsen/logrus"
	"time"
)

// can be generated by xorm reverse: https://gitea.com/xorm/reverse
type TUser struct {
	Id       string    `xorm:"not null pk VARCHAR(32)"`
	Email    string    `xorm:"VARCHAR(50)"`
	Username string    `xorm:"VARCHAR(32)"`
	Pass     string    `xorm:"VARCHAR(100)"`
	CreateAt time.Time `xorm:"created"` // auto populate create_at field, refer https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-04/1.created.html
	UpdateAt time.Time `xorm:"updated"` // auto populate update_At field, refer https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-06/2.updated.html
	DeleteAt time.Time `xorm:"deleted"` // soft delete enabled, refer https://gobook.io/read/gitea.com/xorm/manual-zh-CN/chapter-07/1.deleted.html
}

func InsertUser(user *TUser) (lastInsertId string, err error) {
	if _, err = DB.Insert(user); err != nil {
		log.Warn(err)
		return "", err
	}
	return user.Id, nil
}

func GetUserById(id string) (user *TUser, exist bool, err error) {
	user = new(TUser)
	exist, err = DB.Where("id = ?", id).Get(user)
	if err != nil {
		log.Warn(err.Error())
		return nil, exist, err
	}

	if exist {
		return user, exist, nil
	}

	return nil, exist, nil
}

func GetUserByUsername(username string) (user *TUser, exist bool, err error) {
	user = new(TUser)
	// SELECT `id`, `email`, `username`, `pass`, `create_at`, `update_at`, `delete_at` FROM `t_user` WHERE (username = ?) AND (delete_at=? OR delete_at IS NULL) LIMIT 1
	// query DO NOT include soft deleted rows, if you want get soft delete row, use DB.Where("email = ?", email).Unscoped().Get(user)
	exist, err = DB.Where("username = ?", username).Get(user)
	if err != nil {
		log.Warn(err.Error())
		return nil, exist, err
	}

	if exist {
		return user, exist, nil
	}

	return nil, exist, nil
}

func GetUserByEmail(email string) (user *TUser, exist bool, err error) {
	user = new(TUser)
	exist, err = DB.Where("email = ?", email).Get(user)
	if err != nil {
		log.Warn(err.Error())
		return nil, exist, err
	}

	if exist {
		return user, exist, nil
	}

	return nil, exist, nil
}

func DeleteUser(id string) (affected int64, err error) {
	user := new(TUser)
	// soft delete, UPDATE `t_user` SET `delete_at` = ? WHERE (delete_at=? OR delete_at IS NULL) AND `id`=?
	// if you want to delete in database, use DB.ID(id).Unscoped().Delete(user)
	return DB.ID(id).Delete(user)
}

func UpdateUserPassword(id string, pass string) (affected int64, err error) {
	result, err := DB.Exec("update `t_user` set pass = ? , update_at = ? where id = ?", pass, time.Now(), id)
	if err != nil {
		return 0, err
	} else {
		affected, err = result.RowsAffected()
		if err != nil {
			log.Error(err)
			return 0, err
		}
		return affected, nil
	}
}

func UpdateUserUsername(oldUsername string, newUsername string) (affected int64, err error) {
	result, err := DB.Exec("update `t_user` set username = ? , update_at = ? where username = ? and not exists (select 1 from (select 1 from t_user where username = ?) a)", newUsername, time.Now(), oldUsername, newUsername)
	if err != nil {
		return 0, err
	} else {
		affected, err = result.RowsAffected()
		if err != nil {
			log.Error(err)
			return 0, err
		}
		return affected, nil
	}
}
