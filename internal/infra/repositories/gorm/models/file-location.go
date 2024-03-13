package gormmodels

type FileLocationModel struct {
	ID        string `gorm:"uuid;primaryKey;default:uuid_generate_v4()"`
	URL       string `gorm:"not null"`
	Provider  string `gorm:"not null"`
	Extension string `gorm:"not null"`
}
