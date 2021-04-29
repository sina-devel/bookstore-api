package i18n

import (
	"github.com/kianooshaz/bookstore-api/internal/models/types"
	"github.com/kianooshaz/bookstore-api/pkg/translator/messages"
	"testing"
)

func TestMessageBundle_Translate(t *testing.T) {

	type args struct {
		message  string
		language types.Language
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "translate farsi",
			args: args{
				message:  messages.NoUserFound,
				language: types.GetLanguage("fa"),
			},
			want: "کاربر مورد نظر یافت نشد",
		},
		{
			name: "translate english",
			args: args{
				message:  messages.NoUserFound,
				language: types.GetLanguage("en"),
			},
			want: "no user found",
		},
		{
			name: "message key not found",
			args: args{
				message:  "NoKeyFound",
				language: types.GetLanguage("en"),
			},
			want: "NoKeyFound",
		},
	}

	translate, err := New("../../../build/i18n/")
	if err != nil {
		t.Errorf("New() error : %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translate.Translate(tt.args.language, tt.args.message)
			if got != tt.want {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestMessageBundle_TranslateEn(t *testing.T) {

	type args struct {
		message string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "translate",
			args: args{
				message: messages.NoUserFound,
			},
			want: "no user found",
		},
		{
			name: "message key not found",
			args: args{
				message: "NoKeyFound",
			},
			want: "NoKeyFound",
		},
	}

	translate, err := New("./testData/")
	if err != nil {
		t.Errorf("New() error : %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translate.TranslateEn(tt.args.message)
			if got != tt.want {
				t.Errorf("TranslateEn() got = %v, want %v", got, tt.want)
			}

		})
	}
}
