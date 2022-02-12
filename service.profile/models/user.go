package profile

import (
	"time"

	"avirtan/work_craft/config"

	"github.com/gofrs/uuid"
)

// основная модель пользователя
type User struct {
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Nickname     string     `gorm:"unique; not null" json:"nickname"`
	Login        string     `gorm:"unique; not null" json:"login"`
	Email        string     `gorm:"unique; not null" json:"email"`
	Password     string     `gorm:"not null" json:"-"`
	Level        uint8      `gorm:"default:0"`
	Verification bool       `json:"verification"`
	CreatedAt    time.Time  `json:"createAt"`
	IconID       *uint      `json:"icon"`
	Character    *Character `json:"character"`
	Statistic    *Statistic `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"statistic"`
	Roles        []*Role    `gorm:"many2many:user_role; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"roles"`
}

func (u *User) DeleteUser() {
	config.DB.Where("login = ?", u.Login).Delete(u)
}

//! Delete
func (u *User) SetUserRole(roles []*Role) error {
	err := config.DB.Model(u).Association("Roles").Append(roles)
	return err
}

// метод на получение всех пользователей
func GetAllUser() ([]User, error) {
	users := []User{}
	result := config.DB.Find(&users)
	return users, result.Error
}

func GetUserById(id string) (*User, error) {
	user := &User{}
	result := config.DB.First(user, "id = ?", id)
	return user, result.Error
}

func (u *User) UpdateUser(json *User) error {
	u.Email = json.Email
	u.Login = json.Login
	u.Verification = json.Verification
	u.IconID = json.IconID
	config.DB.Model(u).Association("Roles").Clear()
	u.Roles = json.Roles
	result := config.DB.Save(u)
	return result.Error
}

func (u *User) GetRoles() error {
	roles := []*Role{}
	err := config.DB.Model(u).Association("Roles").Find(&roles)
	u.Roles = roles
	return err
}

func (u *User) GetCharacter() error {
	character := &Character{}
	err := config.DB.Model(u).Association("Character").Find(&character)
	if err != nil {
		return err
	}
	u.Character = character
	return nil
}

func (u *User) GetStatistic() error {
	statistic := &Statistic{}
	err := config.DB.Model(u).Association("Statistic").Find(&statistic)
	if err != nil {
		return err
	}
	if statistic.ID == 0 {
		statistic.UserID = u.ID
		err = statistic.Create()
		if err != nil {
			return err
		}
	}
	u.Statistic = statistic
	return nil
}
