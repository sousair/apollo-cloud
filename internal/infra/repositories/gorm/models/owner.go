package gormmodels

import (
	"time"

	"gorm.io/gorm"
)

type OwnerModel struct {
	gorm.Model
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Albums []*AlbumModel `gorm:"foreignKey:OwnerID;references:ID"`
	Musics []*MusicModel `gorm:"foreignKey:OwnerID;references:ID"`
}
