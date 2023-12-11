package users

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func (login *Login) PasswordMatches(user User) (bool, error) {
// 	fmt.Println(user.Password, login.Password)
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
// 	// fmt.Println(err)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
// 			return false, err
// 		default:
// 			return false, err
// 		}
// 	}
// 	return true, nil
// }
