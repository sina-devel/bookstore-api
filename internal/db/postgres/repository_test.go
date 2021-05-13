package postgres

import (
	"fmt"
	"github.com/kianooshaz/bookstore-api/internal/db/postgres/schema"
	"github.com/kianooshaz/bookstore-api/pkg/log/logrus"
	"github.com/kianooshaz/bookstore-api/pkg/translate/i18n"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
