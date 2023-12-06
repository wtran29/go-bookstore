package users

import (
	"encoding/json"
	"time"
)

type PublicUser struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}
type PrivateUser struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

func (users Users) ReadJson(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for idx, user := range users {
		result[idx] = user.ReadJson(isPublic)
	}
	return result
}

func (user *User) ReadJson(isPublic bool) interface{} {
	if isPublic {

		return PublicUser{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Status:    user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}
	return privateUser
}
