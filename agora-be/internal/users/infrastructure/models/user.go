package models

type UserModel struct {
	Id    string `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}
