package gormmodels

import (
	"time"

	"gorm.io/gorm"
)

type MusicModel struct {
	gorm.Model
	ID               string    `gorm:"type:uuid;primaryKey"`
	Name             string    `gorm:"not null"`
	DurationInMs     int       `gorm:"not null"`
	ReleaseDate      time.Time `gorm:"not null"`
	AlbumID          string    `gorm:"default:null"`
	OwnerID          string    `gorm:"not null"`
	CoverImageURL    string    `gorm:"default:null"`
	SongDataLocation string    `gorm:"not null"`

	Album AlbumModel `gorm:"foreignKey:AlbumID"`
	Owner OwnerModel `gorm:"foreignKey:OwnerID"`
}
