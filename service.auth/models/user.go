package AuthModels

import (
	"errors"
	"net/mail"
	"time"

	"avirtan/work_craft/config"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// основная модель пользователя
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Login        string    `gorm:"unique; not null"`
	Email        string    `gorm:"unique; not null"`
	Nickname     string    `gorm:"unique; not null" json:"nickname"`
	Password     string    `gorm:"not null"`
	Verification bool
	CreatedAt    time.Time
	IconID       *uint
}

// метод на создание пользователя в бд
func (u *User) CreateUser(login string, password string, email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 9)
	u.Login = login
	u.Password = string(hash[:])
	u.Email = email
	u.Nickname = login
	result := config.DB.Create(u)
	return result.Error
}

//метод на поиск пользователя
func (u *User) FindUser(login string, password string) error {
	result := config.DB.Where("Login = ? OR Email = ?", login, login).Find(u)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if result.Error != nil || result.RowsAffected == 0 || err != nil {
		return errors.New("not currect login/email or password")
	}
	return nil
}

func (u *User) DeleteUser() {
	config.DB.Where("login = ?", u.Login).Delete(u)
}

// метод на получение всех пользователей
//! мб не нужен будет
func GetAllUSer() {

}
