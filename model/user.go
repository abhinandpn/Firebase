package model

type User struct {
	Firstname string `firestore:"firstname"`
	Lastname  string `firestore:"lastname"`
	Username  string `firestore:"username"`
}
