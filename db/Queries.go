package db

var RegisterUser string = `INSERT INTO authorized_users (username) VALUES ("%s")`
var StoreUserDetails string = `INSERT INTO user_credentials (username, password) VALUES ("%s", "%s")`
var CheckUser string = `SELECT * FROM authorized_users WHERE username = "%s"`
var VerifyUserCredentials string = `SELECT * FROM user_credentials WHERE username = "%s" AND password = "%s"`