package model

type AccessToken struct {
	UUID      string `gorm:"primarykey;"`
	Token     string `gorm:"not null"`
	UserID    uint   `gorm:"foreignKey;"`
	UserAgent string `gorm:"not null"`
}

func (a *AccessToken) String() string {
	return a.Token
}
