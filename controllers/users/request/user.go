package request

import (
	_userDomain "acp14/business/users"
	"time"
)

type UserRequest struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	// Token     string    `json:"token"`
}

func ToDomain(user UserRequest) _userDomain.Domain {
	return _userDomain.Domain{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
