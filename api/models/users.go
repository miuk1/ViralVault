// model for users
package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username          string `json:"username"`
	Email             string `json:"email"`
	PasswordHash      string `json:"-"`
	PasswordSalt      string `json:"-"`
	TwoFactorEnabled  bool   `json:"two_factor_enabled"`
	SessionToken      string `json:"-"`
	SessionExpiration int64  `json:"-"`
}
