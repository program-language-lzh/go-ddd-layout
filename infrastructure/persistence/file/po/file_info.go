package po

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID       int64  `gorm:"column:id;	type:int;	not null;	primaryKey;	autoIncrement;	comment: 'Primary key ID'"`
	FileName string `gorm:"column:filename; type:varchar(255);	 not null; comment: 'File Name'"`
	FileType string `gorm:"column:filetype; type: varchar(50); not null; comment: 'File Type'"`
}

func (f *File) TableName() string {
	return "t_file_info"
}
