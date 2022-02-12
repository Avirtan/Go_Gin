package profile

import "avirtan/work_craft/config"

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique; not null"`
}

// добавить роль
func (r *Role) AddRole() error {
	result := config.DB.Create(&r)
	return result.Error
}

// метод на получение всех ролей
func GetAllRole() (*[]Role, error) {
	roles := &[]Role{}
	result := config.DB.Find(&roles)
	return roles, result.Error
}

func GetRoleById(id uint) (*Role, error) {
	role := &Role{}
	result := config.DB.First(role, id)
	return role, result.Error
}

func (r *Role) DeleteRole() error {
	result := config.DB.Delete(&r)
	return result.Error
}
