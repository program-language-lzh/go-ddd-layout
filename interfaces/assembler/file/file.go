package file

import (
	"server/domain/file/entity"
	dto "server/interfaces/dto/file"
)

func DTOToEntity(file *dto.File) *entity.File {
	return &entity.File{
		ID:   file.ID,
		Name: file.Filename,
		Type: file.Filetype,
	}
}

func EntityToDTO(file *entity.File) *dto.File {
	return &dto.File{
		ID:       file.ID,
		Filename: file.Name,
		Filetype: file.Type,
	}
}

func ListEntityToDTO(files []entity.File) []dto.File {
	var dtoFiles []dto.File
	for _, file := range files {
		dtoFiles = append(dtoFiles, *EntityToDTO(&file))
	}
	return dtoFiles
}
