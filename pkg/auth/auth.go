package auth

import (
	"cms/app/models/clerk"
	"cms/pkg/session"
	"errors"
)

func getUID() string {
	uid := session.Get("uid")
	if uid, ok := uid.(string); ok && len(uid) > 0 {
		return uid
	}
	return ""
}

func User() clerk.Clerk {
	uid := getUID()
	user, err := clerk.Find(uid)
	if err == nil {
		return user
	}
	return clerk.Clerk{}
}

func Attempt()  {
	
}

func Login()  {
	
}

func LoginUsingId()  {
	
}

func Logout()  {
	session.Forget("uid")
}

func Check() (bool, error) {

	if uid := getUID(); len(uid) <= 0 {
		return false, errors.New("登录信息过期")
	}
	user := User()
	if user.ID <= 0 {
		return false, errors.New("用户不存在")
	}

	if user.IsActive < 1 {
		return false, errors.New("用户被禁用")
	}

	return true, nil
}

