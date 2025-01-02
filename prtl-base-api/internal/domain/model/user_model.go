package model

type User struct {
    ID        int    `json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"password"`
}

func NewUser(id int, username, password string) *User {
    return &User{
        ID:        id,
        Username: username,
        Password: password,
    }
}
