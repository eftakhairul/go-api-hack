package models

import "errors"

// User Model
type User struct {
	Model
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Address   string `gorm:"column:address" json:"address"`
	Email     string `gorm:"column:email" json:"email"`
}

// Validate validates user input
func (u *User) Validate() error {
	if u.FirstName == "" || u.LastName == "" {
		return errors.New("Invalid input")
	}

	return nil
}
