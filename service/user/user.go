package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/huahuayu/go-gin-app/common/redis"
	"github.com/huahuayu/go-gin-app/common/session"
	. "github.com/huahuayu/go-gin-app/global"
	. "github.com/huahuayu/go-gin-app/model"
	"github.com/huahuayu/go-gin-app/view"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func Register(c *gin.Context, req *view.RegisterReq) {
	// set a lock to prevent concurrent insert
	lock := fmt.Sprintf(redis.KEY_INSERT_USER_LOCK, req.Email, req.Username)
	err := redis.Cli.ObtainLock(lock, 1*time.Second)
	defer redis.Cli.ReleaseLock(lock)
	if err != nil {
		view.ResponseErr(c, ErrTryLater, "")
		return
	}

	// check existence before register
	log.Debug("user register info: %s", req)
	_, exist, err := GetUserByEmail(req.Email)
	if err != nil {
		view.ResponseErr(c, ErrInternal, "")
		return
	}
	if exist {
		view.ResponseErr(c, ErrEmailAlreadyExist, "")
		return
	}

	_, exist, err = GetUserByUsername(req.Username)
	if err != nil {
		view.ResponseErr(c, ErrInternal, "")
		return
	}
	if exist {
		view.ResponseErr(c, ErrUsernameAlreadyExist, "")
		return
	}

	// register user
	user := new(TUser)
	user.Id = strings.Replace(uuid.New().String(), "-", "", -1)
	user.Email = req.Email
	user.Username = req.Username
	user.Pass = hashAndSalt(req.Password)
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	_, err = InsertUser(user)
	if err == nil {
		view.ResponseOK(c, "")
	} else {
		view.ResponseErr(c, ErrInternal, "")
		return
	}
}

func Login(c *gin.Context, req *view.LoginReq) {
	user, exist, err := GetUserByEmail(req.Email)
	if err != nil {
		view.ResponseErr(c, ErrInternal, "")
		return
	}

	if !exist || !comparePasswords(user.Pass, req.Password) {
		view.ResponseErr(c, ErrLoginFailed, "")
		return
	}

	token := uuid.New().String()
	session.Set(token, user)

	var res = &view.LoginRes{Token: token}
	view.ResponseOK(c, res)
}

func UpdatePassword(c *gin.Context, req *view.UpdatePasswordReq) {
	token := c.Request.Header.Get("Authorization")
	user := session.Get(token)

	if !comparePasswords(user.Pass, req.OldPassword) {
		view.ResponseErr(c, ErrOldPasswordNotRight, "")
		return
	}

	affected, err := UpdateUserPassword(user.Id, hashAndSalt(req.NewPassword))
	if err != nil {
		view.ResponseErr(c, ErrInternal, "")
		return
	}

	if affected == 0 {
		view.ResponseErr(c, ErrDataNotExist, "")
		return
	}

	// delete token, user need login again
	session.Del(token)

	view.ResponseOK(c, "")
}

func UpdateUsername(c *gin.Context, req *view.UpdateUsernameReq) {
	token := c.Request.Header.Get("Authorization")
	user := session.Get(token)
	_, exist, err := GetUserByUsername(req.NewUsername)
	if err != nil {
		view.ResponseErr(c, ErrInternal, "")
		return
	}
	if exist {
		view.ResponseErr(c, ErrUsernameAlreadyExist, "")
		return
	}

	affected, err := UpdateUserUsername(user.Username, req.NewUsername)
	if err != nil {
		view.ResponseErr(c, ErrInternal, "")
		return
	}

	// if username updated, reset session info
	if affected != 0 {
		user.Username = req.NewUsername
		session.Del(token)
		session.Set(token, user)
	}
	view.ResponseOK(c, "")
}

func Delete(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	user := session.Get(token)
	affected, err := DeleteUser(user.Id)
	if err != nil {
		view.ResponseErr(c, ErrInternal, "")
		return
	}
	if affected != 1 {
		view.ResponseErr(c, ErrDataNotExist, "")
		return
	}
	session.Del(token)
	view.ResponseOK(c, "")
}

func GetUser(c *gin.Context) *TUser {
	token := c.Request.Header.Get("Authorization")
	return session.Get(token)
}

func Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	session.Del(token)
	view.ResponseOK(c, "")
}

func Info(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	user := session.Get(token)

	var res = new(view.InfoRes)
	_ = copier.Copy(res, &user)
	view.ResponseOK(c, res)
}

func hashAndSalt(plainPwd string) string {
	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	if err != nil {
		log.Info(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Info(err)
		return false
	}
	return true
}
