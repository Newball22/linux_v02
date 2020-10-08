package model

import (
	"encoding/base64"
	"goFlow/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
	Avatar   string `gorm:"type:varchar(20)" json:"avatar"`
}

//检查用户是否存在
func CheckUser(name string) int {
	var user User
	DB.Select("id").Where("username=?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//添加用户
func AddAUser(data *User) int {
	data.Password = ScryptPassword(data.Password) //默认实现来gorm的钩子函数
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询单个用户
func FindAUser(id int) (User, int) {
	var user User
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

//查询用户列表
func FindAllUsers(pageSize int, pageNum int) ([]User, int, int) {
	var total int
	userList := []User{}
	err := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, total
	}
	return userList, errmsg.SUCCESS, total
}

//编辑用户信息(除了密码)
func EditAUser(id int, data *User) int {
	var (
		user User
		maps = make(map[string]interface{})
	)
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := DB.Model(&user).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//删除用户
func DeleteAUser(id int) int {
	var user User
	err := DB.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func ScryptPassword(password string) string {
	const keyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 15, 23, 6, 2, 28, 13, 3}
	hashPassword, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(hashPassword)
	return fpw
}
