package profile

type Achivement struct {
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"not null"`
	State string `gorm:"not null"`
	Value int32  `gorm:"not null"`
	Icon  *Icon
}
