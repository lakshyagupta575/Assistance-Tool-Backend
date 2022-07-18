package db

var CheckUser string = `SELECT * FROM authorized_users WHERE username = "%s"`
var RegisterUser string = `INSERT INTO authorized_users (username) VALUES ("%s")`