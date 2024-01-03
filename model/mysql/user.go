package mysql

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;"`
	Password string `gorm:"type:char(32);not null;"`
	Mobile   string `gorm:"type:varchar(11);unique"`
}

type Identity struct {
	gorm.Model
	UserId   int64
	RealName string
	CardNo   string
}

// 根据手机号查询用户
func GetUserByMobile(mobile string) (user *User, err error) {
	err = MyDB.Where("mobile = ?", mobile).First(&user).Error
	return
}

// 根据用户id修改密码
func UpdateUserPwd(password string, userId int64) (err error) {
	err = MyDB.Model(&User{}).Where("id = ?", userId).Update("password", password).Error
	return
}

// 根据用户名查询用户名
func GetUserByUsername(username string) (user *User, err error) {
	err = MyDB.Where("username = ?", username).First(&user).Error
	return
}

// 根据用户id修改手机号
func UpdateUserMobile(user *User) (err error) {
	err = MyDB.Where(user).Save(&user).Error
	return
}

// 根据用户id查看用户是否已经完成实名认证
func GetUserIdentityByUserId(userId int64) (ident *Identity, err error) {
	err = MyDB.Where("user_id = ?", userId).First(&ident).Error
	return
}

// 添加用户实名认证表
func AddIdentity(ident *Identity) (err error) {
	err = MyDB.Create(&ident).Error
	return
}
