package db

type User struct {
	Id        string `bson:"_id,omitempty"`
	Username  string `bson:"username,omitempty"`
	Password  string `bson:"password,omitempty"`
	CreatedAt int64  `bson:"created_at,omitempty"`
	UpdatedAt int64  `bson:"updated_at,omitempty"`
}
