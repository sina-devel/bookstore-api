package postgres

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/db/postgres/schema"
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/pkg/log/logrus"
	"github.com/kianooshaz/bookstore-api/pkg/random"
	"github.com/kianooshaz/bookstore-api/pkg/translate/i18n"
	"log"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setupTest(t *testing.T) *repository {
	cfg := &config.Postgres{
		Username:  "postgres",
		Password:  "12345",
		DBName:    "bookstore_test",
		Host:      "127.0.0.1",
		Port:      "5432",
		SSLMode:   "disable",
		TimeZone:  "Asia/Tehran",
		Charset:   "utf8mb4",
		Migration: true,
	}

	translator, err := i18n.New("../../../build/i18n/")
	if err != nil {
		t.Fatal(err)
	}

	logger, err := logrus.New(&logrus.Option{
		Path:         "../../../logs/test",
		Pattern:      "%Y-%m-%dT%H:%M",
		MaxAge:       "720h",
		RotationTime: "24h",
		RotationSize: "20MB",
	})
	if err != nil {
		t.Fatal(err)
	}

	repo := &repository{
		cfg:        cfg,
		translator: translator,
		logger:     logger,
	}

	if err := repo.connect(); err != nil {
		t.Fatal(err)
	}

	if err := repo.db.Migrator().DropTable(
		&schema.User{},
		&schema.Book{},
		&schema.Wallet{},
		&schema.Category{},
		&schema.Comment{},
		&schema.Picture{},
	); err != nil {
		log.Fatalln(err)
	}

	if err := repo.db.Migrator().CreateTable(
		&schema.User{},
		&schema.Book{},
		&schema.Wallet{},
		&schema.Category{},
		&schema.Comment{},
		&schema.Picture{},
	); err != nil {
		log.Fatalln(err)
	}

	return repo
}

func newUserTest() *models.User {

	return &models.User{
		Username:    random.String(8),
		Password:    random.String(25),
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
