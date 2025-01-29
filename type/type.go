package _type

import "fmt"

type MinhaString string

type User struct {
	ID   uint64
	Name string
}

func (u *User) Update(newName string) {
	u.Name = newName
}

func main() {
	user := User{24, "Gustavo"}
	user.Update("Henrique")
	fmt.Println(user)
}
