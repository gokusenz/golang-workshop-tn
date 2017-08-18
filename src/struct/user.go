package main

type User struct {
	id       int
	username string
	email    string
	status   bool
}

func (u User) add() bool {
	panic("Under construction")
	return false
}

func main() {
	u := User{}
	u.add()
}
