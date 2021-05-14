package postgres

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/internal/db/postgres/schema"
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/pkg/log/logrus"
	"github.com/kianooshaz/bookstore-api/pkg/random"
	"github.com/kianooshaz/bookstore-api/pkg/translate/i18n"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"testing"
)

var (
	repoTest *repository
)

func TestMain(m *testing.M) {
	setupTest()
	code := m.Run()
	tearDownTest()
	os.Exit(code)
}

func setupTest() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		"127.0.0.1",
		"postgres",
		"12345",
		"bookstore_test",
		"5432",
		"disable",
		"Asia/Tehran",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)

	}

	if err := db.Migrator().DropTable(
		&schema.User{},
		&schema.Book{},
		&schema.Wallet{},
		&schema.Category{},
		&schema.Comment{},
		&schema.Picture{},
	); err != nil {
		log.Fatalln(err)
	}

	if err := db.Migrator().CreateTable(
		&schema.User{},
		&schema.Book{},
		&schema.Wallet{},
		&schema.Category{},
		&schema.Comment{},
		&schema.Picture{},
	); err != nil {
		log.Fatalln(err)
	}

	translator, err := i18n.New("./build/i18n/")
	if err != nil {
		log.Fatalln(err)
	}

	logger, err := logrus.New(&logrus.Option{
		Path:         "../../../logs/test",
		Pattern:      "%Y-%m-%dT%H:%M",
		MaxAge:       "720h",
		RotationTime: "24h",
		RotationSize: "20MB",
	})

	repoTest = &repository{
		db:         db,
		translator: translator,
		logger:     logger,
	}
}

func tearDownTest() {
	repoTest = nil
}

func newUserTest() *models.User {
	return &models.User{
		Username:    random.String(8),
		FirstName:   random.String(8),
		LastName:    random.String(8),
		Email:       random.String(5) + "@" + random.String(3) + "." + random.String(3),
		PhoneNumber: "0912" + random.StringWithCharset(7, "0123456789"),
		Gender:      types.Male,
		Role:        types.Basic,
		Wallet: models.Wallet{
			Balance: types.Price(rand.Uint32()),
			Status:  types.WalletOpen,
		},
	}
}
