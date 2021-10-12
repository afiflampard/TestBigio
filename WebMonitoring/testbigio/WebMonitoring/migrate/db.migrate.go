package migrate

import (
	"log"
	"onboarding/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	tableExist := (db.Migrator().HasTable(&models.User{}) && db.Migrator().HasTable(&models.Role{}) && db.Migrator().HasTable(&models.Raport{}))
	if !tableExist {
		dbMigrate := db.Debug().Migrator().DropTable(&models.User{}, &models.Role{}, &models.Raport{})
		if dbMigrate != nil {
			log.Fatal("Cannot drop Table")
		}
		db.AutoMigrate(&models.User{}, &models.Raport{})

		var roles = []models.Role{
			models.Role{
				Role: "Guru",
			},
			models.Role{
				Role: "Siswa",
			},
			models.Role{
				Role: "Admin",
			},
		}

		pass, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.MinCost)
		if err != nil {
			log.Println(err)
		}
		var Users = []models.User{
			models.User{
				Email:    "afiflampard32@gmail.com",
				Name:     "Guru",
				Mobile:   "08576543434",
				Address:  "Jabon",
				Password: string(pass),
				RoleID:   1,
			},
			models.User{
				Email:    "afiflampard123@gmail.com",
				Name:     "Siswa",
				Mobile:   "08576543439",
				Address:  "Jabon",
				Password: string(pass),
				RoleID:   2,
			},
			models.User{
				Email:    "afiflampard09@gmail.com",
				Name:     "Admin",
				Mobile:   "08576543439",
				Address:  "Jabon",
				Password: string(pass),
				RoleID:   2,
			},
		}

		for _, role := range roles {
			err := db.Debug().Create(&role).Error
			if err != nil {
				log.Fatalf("Failed to create Role")
			}
		}

		for _, user := range Users {
			err := db.Debug().Create(&user).Error
			if err != nil {
				log.Fatalf("Failed to create User")
			}
		}

	}

}
