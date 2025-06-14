package file

import (
	"errors"
	"server/domain/file/entity"
	"server/domain/file/repository"
	"server/infrastructure/persistence/file/converter"
	"server/infrastructure/persistence/file/po"

	"gorm.io/gorm"
)

type FileRepositoryImpl struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) repository.FileRepository {
	return &FileRepositoryImpl{db: db}
}

// createFileInfo implements FileRepository.
func (fr *FileRepositoryImpl) CreateFile(file *entity.File) (*entity.File, error) {
	poFile := converter.FileEntityToPO(file)
	result := fr.db.Create(poFile)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("unable to save file")
	}

	entityFile := converter.FilePOToEntity(poFile)

	return entityFile, nil
}

// getFileInfoById implements FileRepository.
func (fr *FileRepositoryImpl) GetFileById(id int64) (*entity.File, error) {
	file := &po.File{}

	result := fr.db.First(file, id)
	if result.Error != nil {
		return nil, result.Error
	}

	entityFile := converter.FilePOToEntity(file)

	return entityFile, nil
}
