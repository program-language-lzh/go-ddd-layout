package service

import (
	"fmt"
	"server/domain/file/entity"
	"server/domain/file/repository"
	"strings"
)

type FileDomain interface {
	CreateFile(file *entity.File) (*entity.File, error)
	GetFileById(id int64) (*entity.File, error)
}

type FileDomainImpl struct {
	fr repository.FileRepository
}

func NewFileDomainImpl(repo repository.FileRepository) FileDomain {
	return &FileDomainImpl{fr: repo}
}

func (fd *FileDomainImpl) CreateFile(file *entity.File) (*entity.File, error) {
	f, err := fd.fr.CreateFile(file)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, fmt.Errorf("file %s already exists", file.Name)
		}
		return nil, err
	}
	return f, nil
}

func (fd *FileDomainImpl) GetFileById(id int64) (*entity.File, error) {
	return fd.fr.GetFileById(id)
}
