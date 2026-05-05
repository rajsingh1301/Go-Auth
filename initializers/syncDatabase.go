package initializers

import "auth-go-test/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}
