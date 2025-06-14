package repository

import "server/domain/file/entity"

type FileRepository interface {
	CreateFile(file *entity.File) (*entity.File, error)
	GetFileById(id int64) (*entity.File, error)
}
