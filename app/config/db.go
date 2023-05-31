package config

import "fmt"

const (
	DBName     = "reservator"
	DBPort     = "3306"
	DBUser     = "reservator_user"
	DBPassword = "reservator_user"
	DBHost     = "serverdb"
)

func GetConnectionString() string {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	return connection
}
