package model

type AccessToken struct {
	UUID   string `gorm:"primarykey;"`
	Token  string
	UserID uint `gorm:"foreignKey;"`
}

func (a *AccessToken) String() string {
	return a.Token
}
