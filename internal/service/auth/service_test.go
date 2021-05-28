package auth

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/pkg/log/logrus"
	"github.com/kianooshaz/bookstore-api/pkg/random"
	"github.com/kianooshaz/bookstore-api/pkg/translate/i18n"
	"math/rand"
	"testing"
)

var (
	serviceTest *service
)

func setupTest(t *testing.T) {
	cfg := &config.Auth{
		AccessExpirationInMinute:  15,
		RefreshExpirationInMinute: 525600,
		JWTSecret:                 "secret",
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

	serviceTest = &service{
		cfg:        cfg,
		logger:     logger,
		translator: translator,
	}
}

func teardownTest() {
	serviceTest = nil
}

func newUserTest() *models.User {
	return &models.User{
		ID:          uint(rand.Uint32()),
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