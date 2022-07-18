package db

type Action string

type Register struct {
	Username Action `json: "username"`
	Password Action `json: "password"`
}