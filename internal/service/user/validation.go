package user

import (
	"unicode"

	"github.com/kianooshaz/bookstore-api/pkg/derrors"
	"github.com/kianooshaz/bookstore-api/pkg/log"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func (s *service) validateUsername(username string) error {

	if l := len(username); l <= s.userCfg.UsernameMinLength || l >= s.userCfg.UsernameMaxLength {
		s.logger.Error(&log.Field{
			Section:  "user.validation",
			Function: "validateUsername",
			Params:   map[string]interface{}{"username": username},
			Message:  s.translator.TranslateEn(messages.InvalidUsernameLength),
		})

		return derrors.New(derrors.KindInvalid, messages.InvalidUsernameLength)
	}

	return nil
}

func (s *service) validatePassword(password string) error {
	var number, upper, special bool
	var letters int

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			s.logger.Error(&log.Field{
				Section:  "user.validation",
				Function: "validatePassword",
				Params:   map[string]interface{}{"password": password},
				Message:  s.translator.TranslateEn(messages.InvalidPassword),
			})

			return derrors.New(derrors.KindInvalid, messages.InvalidPassword)
		}
	}

	if letters >= s.userCfg.PasswordMinLetters && number && upper && special {
		return nil
	}

	s.logger.Error(&log.Field{
		Section:  "user.validation",
		Function: "validatePassword",
		Params:   map[string]interface{}{"password": password},
		Message:  s.translator.TranslateEn(messages.InvalidPassword),
	})

	return derrors.New(derrors.KindInvalid, messages.InvalidPassword)
}
