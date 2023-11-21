package database

var connection string

func init() { //otomatis di eksekusi menggunakan init
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}