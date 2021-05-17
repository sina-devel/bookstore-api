package user

import (
	"github.com/golang/mock/gomock"
	"github.com/kianooshaz/bookstore-api/internal/config"
	"github.com/kianooshaz/bookstore-api/internal/db/mock/main_mock"
	"github.com/kianooshaz/bookstore-api/internal/models"
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/internal/params"
	"github.com/kianooshaz/bookstore-api/pkg/log/logrus"
	"github.com/kianooshaz/bookstore-api/pkg/random"
	"github.com/kianooshaz/bookstore-api/pkg/translate/i18n"
	"math/rand"
	"testing"
)

var (
	mockCtrl     *gomock.Controller
	mockMainRepo *main_mock.MockMainRepository
	serviceTest  *service
)

func teardownTest() {
	mockCtrl.Finish()
	mockCtrl = nil
	serviceTest = nil
}

func setupTest(t *testing.T) {
	cfg := &config.User{
		UsernameMinLength:  3,
		UsernameMaxLength:  50,
		PasswordMinLetters: 4,
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

	mockCtrl = gomock.NewController(t)
	mockMainRepo = main_mock.NewMockMainRepository(mockCtrl)

	serviceTest = &service{
		userCfg:    cfg,
		userRepo:   mockMainRepo,
		logger:     logger,
		translator: translator,
	}
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

func newCreateUserRequestTest() *params.CreateUserRequest {
	return &params.CreateUserRequest{
		Username:              random.String(8),
		Password:              random.String(10), // todo random password
		FirstName:             random.String(8),
		LastName:              random.String(8),
		Email:                 random.String(5) + "@" + random.String(3) + "." + random.String(3),
		IsEmailVerified:       true,
		PhoneNumber:           "0912" + random.StringWithCharset(7, "0123456789"),
		IsPhoneNumberVerified: true,
		Gender:                types.Male,
		Role:                  types.Basic,
	}
}

func newUpdateUserRequestTest() *params.UpdateUserRequest {
	return &params.UpdateUserRequest{
		FirstName:             random.String(8),
		LastName:              random.String(8),
		Email:                 random.String(5) + "@" + random.String(3) + "." + random.String(3),
		IsEmailVerified:       true,
		PhoneNumber:           "0912" + random.StringWithCharset(7, "0123456789"),
		IsPhoneNumberVerified: true,
		Gender:                types.Male,
		Role:                  types.Basic,
		Avatar:                random.String(5),
	}
}
