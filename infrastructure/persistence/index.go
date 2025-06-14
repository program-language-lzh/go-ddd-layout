package persistence

import (
	filepo "server/domain/file/repository"
	userpo "server/domain/user/repository"
	"server/infrastructure/persistence/file"
	"server/infrastructure/persistence/user"

	"gorm.io/gorm"
)

type Repositories struct {
	User userpo.UserRepository
	File filepo.FileRepository
	// other repositories...
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewUserRepository(db),
		File: file.NewFileRepository(db),
	}
}
