package dao

import (
	"awesomeProject3/internal/db"
	"awesomeProject3/internal/entity"
	"fmt"
	"gorm.io/gorm"
)

var LR = LogRepository{db.DB}

type LogRepository struct {
	DB *gorm.DB
}

func (lr *LogRepository) CountByParam(p entity.LogParam) {
	fmt.Println(p)

}
func (lr *LogRepository) DeleteByParam(p entity.LogParam) {

}
func (lr *LogRepository) DeleteByPrimaryKey(id int) {

}
func (lr *LogRepository) Insert(e entity.LogEntity) {
	lr.DB.Create(&e)
}
func (lr *LogRepository) InsertSelective(e entity.LogEntity) {

}
func (lr *LogRepository) SelectByParam(p entity.LogParam) {

}
func (lr *LogRepository) SelectByPrimaryKey(id int) entity.LogEntity {
	var le entity.LogEntity
	lr.DB.First(&le, id)
	return le
}

func (lr *LogRepository) UpdateByParamSelective(e entity.LogEntity, p entity.LogParam) {

}
func (lr *LogRepository) UpdateByParam(e entity.LogEntity, p entity.LogParam) {

}
func (lr *LogRepository) UpdateByPrimaryKeySelective(e entity.LogEntity) {

}
func (lr *LogRepository) UpdateByPrimaryKey(e entity.LogEntity) {

}
func (lr *LogRepository) BatchInsert(es []entity.LogEntity) {

}
