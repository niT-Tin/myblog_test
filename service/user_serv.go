package service

import (
	"myblog/repositories"
)

type IUserService interface {
	Register(m repositories.Merge) repositories.Merge
	Login(m repositories.Merge) repositories.Merge
	GetUserById(int64) repositories.Merge
	GetAllUsers() repositories.Merge
	GetUserByUserName(string) repositories.Merge
	GetUserByNickName(string) repositories.Merge
	DeleteUser(m repositories.Merge) repositories.Merge
}

type UserService struct {
	UserRepo repositories.IUser
}

func NewUserService(iu repositories.IUser) IUserService {
	return &UserService{UserRepo: iu}
}

// Register 用户注册
func (us *UserService) Register(m repositories.Merge) repositories.Merge {
	// TODO 将用户密码进行再次加密
	encryptedPwd := ""
	m.U.Password = encryptedPwd
	if !us.UserRepo.IsExist(&m.U) {
		return repositories.Merge{IsSuccess: false}
	} else {
		// TODO 产生Token
		// TODO 将token存至redis 并且设置过期时间
		return us.Login(m)
	}
}

// Login 用户登陆
func (us *UserService) Login(m repositories.Merge) repositories.Merge {
	var tmp repositories.Merge
	// TODO 验证帐号
	// TODO 将token存至redis 并且设置过期时间
	return tmp
}

// GetUserById 根据用户id查询用户数据
func (us *UserService) GetUserById(id int64) repositories.Merge {
	byId, err := us.UserRepo.SelectById(id)
	if err != nil {
		return NoErrorRepeat(err, "get user by id error")
	}
	return repositories.Merge{U: *byId}
}

// GetAllUsers 获取所有用户数据
func (us *UserService) GetAllUsers() repositories.Merge {
	all, err := us.UserRepo.SelectAll()
	if err != nil {
		return NoErrorRepeat(err, "get all user error")
	}
	return repositories.Merge{US: all}
}

// GetUserByUserName 根据用户名获取用户数据
func (us *UserService) GetUserByUserName(userName string) repositories.Merge {
	name, err := us.UserRepo.SelectByUserName(userName)
	if err != nil {
		return NoErrorRepeat(err, "get user by user name error")
	}
	return repositories.Merge{U: *name}
}

// GetUserByNickName 根据用户昵称获取用户数据
func (us *UserService) GetUserByNickName(nickName string) repositories.Merge {
	user, err := us.UserRepo.SelectByNickName(nickName)
	if err != nil {
		return NoErrorRepeat(err, "get user by nick name error")
	}
	return repositories.Merge{U: *user}
}

// DeleteUser 删除用户数据
func (us *UserService) DeleteUser(m repositories.Merge) repositories.Merge {
	return repositories.Merge{IsSuccess: us.UserRepo.Delete(&m.U)}
}
