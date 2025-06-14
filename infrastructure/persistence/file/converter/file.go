package converter

import (
	"server/domain/file/entity"
	"server/infrastructure/persistence/file/po"
)

func FileEntityToPO(file *entity.File) *po.File {
	return &po.File{
		ID:       file.ID,
		FileName: file.Name,
		FileType: file.Type,
	}
}

func FilePOToEntity(file *po.File) *entity.File {
	return &entity.File{
		ID:   file.ID,
		Name: file.FileName,
		Type: file.FileType,
	}
}
