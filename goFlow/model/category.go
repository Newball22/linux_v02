package model

import (
	"github.com/jinzhu/gorm"
	"goFlow/utils/errmsg"
)

//文章类型
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//检查分类是否存在
func CheckCate(name string) int {
	var cate Category
	DB.Select("id").Where("name=?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

//添加分类
func AddACate(data *Category) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类列表
func GetCateS(id int, pageSize int, pageNum int) ([]Category, int,int) {
	cateList := []Category{}
	var total int
	err := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cateList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR,total
	}
	return cateList, errmsg.SUCCESS,total
}

//编辑分类
func EditACate(id int, data *Category) int {
	var (
		cate Category
		maps = make(map[string]interface{})
	)
	maps["name"] = data.Name
	err := DB.Model(&cate).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//删除分类
func DeleteACate(id int) int {
	var cate Category
	err := DB.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
