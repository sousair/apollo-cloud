package gormmodels

import (
	"time"

	"gorm.io/gorm"
)

type AlbumModel struct {
	gorm.Model
	ID                   string    `gorm:"type:uuid;primaryKey"`
	Name                 string    `gorm:"not null"`
	ReleaseDate          time.Time `gorm:"not null"`
	OwnerID              string    `gorm:"type:uuid;not null"`
	CoverImageLocationID string    `gorm:"type:uuid;not null"`

	Owner              *OwnerModel        `gorm:"foreignKey:OwnerID"`
	Musics             []*MusicModel      `gorm:"references:ID;foreignKey:AlbumID"`
	CoverImageLocation *FileLocationModel `gorm:"foreignKey:CoverImageLocationID"`
}
