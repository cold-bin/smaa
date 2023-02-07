// @author cold bin
// @date 2023/2/5

package service

import (
	"eliminate-male-appearance-anxiety/_util"
	"eliminate-male-appearance-anxiety/dao/_redis"
	"eliminate-male-appearance-anxiety/global"
	"eliminate-male-appearance-anxiety/model"
	"eliminate-male-appearance-anxiety/model/_req"
	"eliminate-male-appearance-anxiety/model/_res"
	"eliminate-male-appearance-anxiety/pkg/_jwt"
	"eliminate-male-appearance-anxiety/pkg/num"
	"eliminate-male-appearance-anxiety/pkg/sms"
	"errors"
	"github.com/cold-bin/goutil/general/mlog"
	"time"
)

func GetCode(param *_req.GetCode) (err error) {
	code := _util.Rand6Code()
	if err = sms.Send(param.Phone, code); err != nil {
		return
	}
	// 存储redis
	return _redis.SetEx(global.KeyCodePrefix+param.Phone, code, time.Minute*5)
}

var ErrPhone = errors.New("phone number is wrong")
var ErrSmsCode = errors.New("sms code is wrong or overdue")

func Register(param *_req.Register) (res _res.TwoToken, err error) {
	var (
		ok   bool
		code string
	)
	ok, err = _util.JudgePhone(param.Phone)
	if err != nil {
		return
	}
	// 手机号不合法
	if !ok {
		err = ErrPhone
		return
	}

	// 取验证码
	code, err = _redis.GetEx(global.KeyCodePrefix + param.Phone)
	if err != nil {
		return
	}

	// 比较验证码
	if param.Code != code {
		err = ErrSmsCode
		return
	}

	// 加密
	salt := _util.Rand6Str()
	// 入库
	u := &model.FUser{
		Phone:        param.Phone,
		Password:     _util.HM256(param.Password + salt),
		Salt:         salt,
		SerialNumber: num.GetUniqueId(),
		NickName:     param.NickName}
	err = u.Create()
	if err != nil {
		return
	}

	// 验证通过，签发token
	res.AccessToken, res.RefreshToken, err = _jwt.CreateToken(u.ID)
	return res, err
}

func RegretPassword(param *_req.RegretPassword) (err error) {
	// 先校验用户是否存在
	u := model.FUser{Phone: param.Phone}
	err = u.FindFieldByPhone("id")
	if err != nil {
		return
	}
	// 不存在用户
	if u.ID == 0 {
		err = ErrUserPC
		return
	}

	// 查询验证码
	var code string
	code, err = _redis.GetEx(global.KeyCodePrefix + param.Phone)
	if err != nil {
		return
	}

	if code != param.Code {
		err = ErrUserPC
		return
	}

	// 更新密码
	salt := _util.Rand6Str()
	u2 := model.FUser{Phone: param.Phone, Password: _util.HM256(param.Password + salt), Salt: salt}
	return u2.UpdateFieldByPhone("password")
}

var ErrPP = errors.New("password or phone is wrong")

func LoginByPassword(param *_req.LoginByPassword) (res _res.TwoToken, err error) {
	// 校验密码
	u := model.FUser{Phone: param.Phone}
	err = u.FindFieldByPhone("id", "phone", "password", "salt")
	if err != nil {
		return
	}

	mlog.Debugf("用户登录：%v", u)
	// 没有这个用户
	if u.ID == 0 {
		err = ErrPP
		return
	}

	// 比较密码
	if _util.HM256(param.Password+u.Salt) != u.Password {
		err = ErrPP
	}

	res.AccessToken, res.RefreshToken, err = _jwt.CreateToken(u.ID)
	return
}

var ErrUserPC = errors.New("phone or code is wrong")

func LoginByCode(param *_req.LoginByCode) (res _res.TwoToken, err error) {
	// 先校验用户是否存在
	u := model.FUser{Phone: param.Phone}
	err = u.FindFieldByPhone("id")
	if err != nil {
		return
	}
	// 不存在用户
	if u.ID == 0 {
		err = ErrUserPC
		return
	}

	// 查询验证码
	var code string
	code, err = _redis.GetEx(global.KeyCodePrefix + param.Phone)
	if err != nil {
		return
	}

	if code != param.Code {
		err = ErrUserPC
		return
	}

	res.AccessToken, res.RefreshToken, err = _jwt.CreateToken(u.ID)
	return
}
