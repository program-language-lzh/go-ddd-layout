package file

import (
	"server/domain/file/entity"
	"server/domain/file/service"
)

type FileServiceImpl struct {
	fd service.FileDomain
}

type Service interface {
	CreateFile(file *entity.File) (*entity.File, error)
	GetFileById(id int64) (*entity.File, error)
}

func NewFileServiceImpl(frv service.FileDomain) Service {
	return &FileServiceImpl{fd: frv}
}

func (frv *FileServiceImpl) CreateFile(file *entity.File) (*entity.File, error) {
	return frv.fd.CreateFile(file)
}

func (frv *FileServiceImpl) GetFileById(id int64) (*entity.File, error) {
	return frv.fd.GetFileById(id)
}
