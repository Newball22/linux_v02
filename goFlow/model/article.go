package model

import (
	"github.com/jinzhu/gorm"
	"goFlow/utils/errmsg"
)

//文章
type Article struct {
	gorm.Model
	Category    Category `gorm:"foreignkey:Cid"`
	Title       string   `gorm:"type:varchar(180);not null" json:"title"`
	Cid         int      `gorm:"type:int;not null" json:"cid"`
	Description string   `gorm:"type:varchar(200)" json:"description"`
	Content     string   `gorm:"type:longtext;not null" json:"content"`
	Img         string   `gorm:"type:varchar(180);not null" json:"img"`
}

//添加文章
func AddArt(data *Article) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询文章
func FindArt(id int) (Article, int) {
	var art Article
	err := DB.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR
	}
	return art, errmsg.SUCCESS
}

//查询分类下的所有文章
func FindAllCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int) {
	artList := []Article{}
	var total int
	err := DB.Preload("Category").Where("cid=?", cid).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, total
	}
	return artList, errmsg.SUCCESS, total
}

//查询所有文章列表
func FindAllArt(pageSize int, pageNum int) ([]Article, int, int) {
	artList := []Article{}
	var total int
	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, total
	}
	return artList, errmsg.SUCCESS, total
}

//编辑文章
func EditArt(id int, data *Article) int {
	var (
		art  Article
		maps = make(map[string]interface{})
	)
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["description"] = data.Description
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := DB.Model(&art).Where("id=?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//删除文章
func DeleteArt(id int) int {
	var art Article
	err := DB.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
