package i18n

import (
	"testing"

	"github.com/kianooshaz/bookstore-api/pkg/translator"
	"github.com/kianooshaz/bookstore-api/pkg/translator/messages"
)

func TestMessageBundle_Translate(t *testing.T) {

	type args struct {
		message  string
		language translator.Language
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
				language: translator.GetLanguage("fa"),
			},
			want: "کاربر مورد نظر یافت نشد",
		},
		{
			name: "translate english",
			args: args{
				message:  messages.NoUserFound,
				language: translator.GetLanguage("en"),
			},
			want: "no user found",
		},
		{
			name: "message key not found",
			args: args{
				message:  "NoKeyFound",
				language: translator.GetLanguage("en"),
			},
			want: "NoKeyFound",
		},
	}

	translate, err := New("testdata")
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

	translate, err := New("testdata")
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
