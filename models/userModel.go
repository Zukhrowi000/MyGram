package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`                       // Kunci asing untuk menghubungkan dengan Role
	Role     Role   `json:"role" gorm:"foreignKey:RoleID"` // Penunjuk untuk GORM
}
