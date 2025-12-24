package seeders

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/DavidAfdal/purchasing-systeam/internal/models"
)

func Seedrs(db *gorm.DB) {

	users := []models.User{
		{
			ID:       uuid.New(),
			Username: "admin",
			Role:     "Admin",
			Password: "password123",
		},
		{
			ID:       uuid.New(),
			Username: "staff1",
			Role:     "Staff",
			Password: "password123",
		},
		{
			ID:       uuid.New(),
			Username: "staff2",
			Role:     "Staff",
			Password: "password123",
		},
	}

	for i := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users[i].Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Failed to hash password:", err)
			continue
		}
		users[i].Password = string(hashedPassword)

		if err := db.FirstOrCreate(&users[i], models.User{Username: users[i].Username}).Error; err != nil {
			log.Println("Failed to seed user:", err)
		}
	}

	items := []models.Item{
		{ID: uuid.New(), Name: "Buku Tulis", Stock: 100, Price: 5000},
		{ID: uuid.New(), Name: "Pulpen Biru", Stock: 200, Price: 2000},
		{ID: uuid.New(), Name: "Pensil HB", Stock: 150, Price: 1500},
		{ID: uuid.New(), Name: "Penghapus", Stock: 120, Price: 1000},
		{ID: uuid.New(), Name: "Spidol", Stock: 80, Price: 4000},
	}

	for _, i := range items {
		if err := db.FirstOrCreate(&i, models.Item{Name: i.Name}).Error; err != nil {
			log.Println("Failed to seed item:", err)
		}
	}

	suppliers := []models.Supplier{
		{
			ID:      uuid.New(),
			Name:    "PT Alat Tulis Nusantara",
			Email:   "contact@alattulis.co.id",
			Address: "Jl. Sudirman No.10, Jakarta",
		},
		{
			ID:      uuid.New(),
			Name:    "CV Stationery Sejahtera",
			Email:   "info@stationerysejahtera.com",
			Address: "Jl. Merdeka No.5, Bandung",
		},
		{
			ID:      uuid.New(),
			Name:    "PT Buku Kita",
			Email:   "sales@bukukita.id",
			Address: "Jl. Mangga Dua No.20, Surabaya",
		},
	}

	for _, s := range suppliers {
		if err := db.FirstOrCreate(&s, models.Supplier{Email: s.Email}).Error; err != nil {
			log.Println("Failed to seed supplier:", err)
		}
	}

	log.Println("Seeding Users, Items, and Suppliers finished successfully!")
}
