package repositories

import (
	"gorm.io/gorm"
	"myblog/models"
)

type IUser interface {
	Insert(user *models.User) (int64, error)
	Update(user *models.User) (int64, error)
	SelectById(int64) (*models.User, error)
	SelectAll() ([]models.User, error)
	SelectByUserName(user string) (*models.User, error)
	SelectByNickName(user string) (*models.User, error)
	Delete(user *models.User) bool
	IsExist(user *models.User) bool
	IsMatch(user *models.User) bool
}

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) IUser {
	return &User{db: db}
}

// Insert 插入用户数据
func (u *User) Insert(user *models.User) (int64, error) {
	return NoRepeat1(u.db.Create(user), "插入用户数据错误")
}

// Update 更新用户数据
func (u *User) Update(user *models.User) (int64, error) {
	return NoRepeat1(u.db.Updates(user), "更新用户数据错误")
}

// SelectById 根据id查询用户数据
func (u *User) SelectById(userId int64) (*models.User, error) {
	var usr models.User
	_, err2 := NoRepeat1(u.db.First(&usr, userId), "根据id查询用户数据错误")
	if err2 != nil {
		return &models.User{}, err2
	}
	return &usr, nil
}

// SelectAll 查询所有用户数据
func (u *User) SelectAll() ([]models.User, error) {
	var usrs []models.User
	_, err2 := NoRepeat1(u.db.Find(&usrs), "查询所有用户错误")
	if err2 != nil {
		return usrs, err2
	}
	return usrs, nil
}

// SelectByUserName 根据用户名查询用户数据
func (u *User) SelectByUserName(userName string) (*models.User, error) {
	var ut models.User
	_, err2 := NoRepeat1(u.db.Where("user_name = ?", userName).First(&ut), "根据用户名查询用户数据错误")
	if err2 != nil {
		return &models.User{}, err2
	}
	return &ut, nil
}

// SelectByNickName 根据用户昵称查询用户数据
func (u *User) SelectByNickName(userNickName string) (*models.User, error) {
	var ut models.User
	_, err2 := NoRepeat1(u.db.Where("nick_name = ?", userNickName).First(&ut), "根据用户昵称查询用户数据错误")
	if err2 != nil {
		return &models.User{}, err2
	}
	return &ut, nil
}

// Delete 删除用户数据
func (u *User) Delete(user *models.User) bool {
	_, err2 := NoRepeat1(u.db.Delete(user), "删除用户数据错误")
	if err2 != nil {
		return false
	}
	return true
}

// IsExist 判断用户是否存在
func (u *User) IsExist(user *models.User) bool {
	var tu models.User
	_, err2 := NoRepeat1(u.db.Where("user_name = ?", user.UserName).First(&tu), "用户存在判断错误")
	if err2 != nil {
		return false
	}
	if tu.Password != "" && tu.ID != 0 {
		return true
	}
	return false
}

// IsMatch 判断帐号密码是否匹配
func (u *User) IsMatch(user *models.User) bool {
	var tu models.User
	_, err2 := NoRepeat1(u.db.Where("user_name = ? and password = ?", user.UserName, user.Password).First(&tu), "用户匹配过程错误")
	if err2 != nil {
		return false
	}
	if tu.UserName == user.UserName && tu.Password == user.Password {
		return true
	}
	return false
}
