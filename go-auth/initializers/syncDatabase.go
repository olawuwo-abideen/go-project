package initializers

import (
	"github.com/olawuwo-abideen/go-project/go-auth/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
