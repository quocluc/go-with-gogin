package models

const DRIVER_NAME = "mysql"
const USER_NAME = "root"
const PASSWORD = "mysql"
const DATABASE_NAME = "loglag_gps"
const IP_DATABASE = "127.0.0.1"
const PORT = "3306"
const IP_CONNECT_DATABASE = "tcp(" + IP_DATABASE + ":" + PORT + ")"
const DATABASE_CONNECT = USER_NAME + ":" + PASSWORD + "@" + IP_CONNECT_DATABASE + "/" + DATABASE_NAME

