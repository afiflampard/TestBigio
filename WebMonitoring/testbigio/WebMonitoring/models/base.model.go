package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	pgName := os.Getenv("PGNAME")
	pgPassword := os.Getenv("PGPASSWORD")
	pgDB := os.Getenv("PGDATABASE")
	pgHost := os.Getenv("PGHOST")
	pgPort := os.Getenv("PGPORT")

	postgresConname := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%v", pgHost, pgName, pgDB, pgPassword, pgPort)
	fmt.Println("canname is\t\t", postgresConname)

	db, err := gorm.Open(postgres.Open(postgresConname), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbMigrate := db.Debug().Migrator().DropTable(&User{}, &Role{}, &Raport{})
	if dbMigrate != nil {
		log.Fatal("Cannot drop Table")
	}
	db.AutoMigrate(&User{}, &Role{}, &Raport{})

	var roles = []Role{
		Role{
			Role: "Guru",
		},
		Role{
			Role: "Siswa",
		},
		Role{
			Role: "Admin",
		},
	}

	pass, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	var Users = []User{
		User{
			Email:    "afiflampard32@gmail.com",
			Name:     "Guru",
			Mobile:   "08576543434",
			Address:  "Jabon",
			Password: string(pass),
			RoleID:   1,
		},
		User{
			Email:    "afiflampard123@gmail.com",
			Name:     "Siswa",
			Mobile:   "08576543439",
			Address:  "Jabon",
			Password: string(pass),
			RoleID:   2,
		},
		User{
			Email:    "afiflampard09@gmail.com",
			Name:     "Admin",
			Mobile:   "08576543439",
			Address:  "Jabon",
			Password: string(pass),
			RoleID:   3,
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

func GetDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	pgName := os.Getenv("PGNAME")
	pgPassword := os.Getenv("PGPASSWORD")
	pgDB := os.Getenv("PGDATABASE")
	pgHost := os.Getenv("PGHOST")
	pgPort := os.Getenv("PGPORT")

	postgresConname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v ", pgHost, pgPort, pgName, pgDB, pgPassword)
	fmt.Println("canname is\t\t", postgresConname)

	db, err := gorm.Open(postgres.Open(postgresConname), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
