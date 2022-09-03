package config

func MigrateStrut(strut interface{}) {
	Database.AutoMigrate(strut)
}
