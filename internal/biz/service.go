package biz

import (
	"awesomeProject3/internal/dao"
	"awesomeProject3/internal/dto"
	"awesomeProject3/internal/vo"
)

var LS = LogService{dao.LR}

type LogService struct {
	lr dao.LogRepository
}

func (s *LogService) CreateLog(dto dto.LogDto) {

}

func (s *LogService) DeleteLog(id int) {

}

func (s *LogService) UpdateLog(dto dto.LogDto) {

}

func (s *LogService) QueryLog(param dto.LogQueryParam) []vo.LogVo {
	var vs []vo.LogVo
	return vs
}

func (s *LogService) ViewLog(id int) vo.LogVo {
	logEntity := s.lr.SelectByPrimaryKey(id)
	v := vo.LogVo{ID: logEntity.ID, Content: logEntity.Content,
		Pictures: logEntity.Pictures, Date: logEntity.Date, Location: logEntity.Location}
	return v
}
