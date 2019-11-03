package common

import (
	"beego_admin/models"
	"beego_admin/utils"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/go-ozzo/ozzo-validation"
)

type User struct {
	models.Model
	Account string
	Password string
	Nickname string
	WechatOpenId string
	HeadImgUrl string
	Sex int
	Province string
	City string
	LoginIp string
	loginTimes int
	Salt string
}

func (user User)Validate() error {
	return validation.ValidateStruct(&user,
	)
}

// 添加用户信息
func (user *User)Create() (err error) {
	// 数据验证
	err = user.Validate()
	if err != nil {
		return err
	}
	// 判断用户名或者昵称未使用
	if !models.DB.Unscoped().Where("nickname = ?", user.Nickname).Find(&User{}).RecordNotFound(){
		return errors.New("昵称已经存在")
	}

	if err = models.DB.Create(&user).Error; err != nil {
		logs.Error("用户创建失败", err)
		return errors.New("信息保存失败")
	}
	return nil
}

func (user *User)FindUserByOpenId() (u User, err error) {
	if models.DB.Where("wechat_open_id = ? ", user.WechatOpenId).First(&u).RecordNotFound() {
		// 微信用户不存在,添加用户
		// 判断用户名是否重复,如果重复,增加字符串,知道不重复未知
		user.Nickname = user.GenerateNickname(user.Nickname)

		user.loginTimes = 1
		user.Salt = utils.RandString(8)
		err := user.Create()
		if err != nil{
			return u, err
		}
		_,_ = user.FindUserByOpenId()
	}
	return u, nil
}

func (user *User)GenerateNickname(nickname string) string {
	if !models.DB.Where("nickname = ? ", nickname).Find(&user).RecordNotFound() {
		nickname := utils.RandString(2)
		var tempUser = &User{}
		tempUser.GenerateNickname(nickname)
	}
	return nickname
}