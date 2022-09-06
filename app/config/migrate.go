package config

// MigrateStrut : is used to migrate any struct to de database using Gorm
func MigrateStrut(strut interface{}) {
	Database.AutoMigrate(strut)
}
