package model

type RequestLevel struct {
	ID    uint8  `gorm:"primaryKey;column=id"`
	Label string `gorm:"column=label"`
}

func (RequestLevel) TableName() string {
	return "request_levels"
}
