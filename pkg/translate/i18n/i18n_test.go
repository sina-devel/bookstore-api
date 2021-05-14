package i18n

import (
	"testing"

	"github.com/kianooshaz/bookstore-api/pkg/translate"
	"github.com/kianooshaz/bookstore-api/pkg/translate/messages"
)

func TestMessageBundle_Translate(t *testing.T) {
	type args struct {
		message  string
		language translate.Language
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "translate farsi",
			args: args{
				message:  messages.DBError,
				language: translate.GetLanguage("fa"),
			},
			want: "خطایی وجود دارد",
		},
		{
			name: "translate english",
			args: args{
				message:  messages.UserNotFound,
				language: translate.GetLanguage("en"),
			},
			want: "user not found",
		},
		{
			name: "message key not found",
			args: args{
				message:  "NoKeyFound",
				language: translate.GetLanguage("en"),
			},
			want: "NoKeyFound",
		},
	}

	translator, err := New("testdata")
	if err != nil {
		t.Errorf("New() error : %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translator.Translate(tt.args.language, tt.args.message)
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
				message: messages.UserNotFound,
			},
			want: "user not found",
		},
		{
			name: "message key not found",
			args: args{
				message: "NoKeyFound",
			},
			want: "NoKeyFound",
		},
	}

	translator, err := New("testdata")
	if err != nil {
		t.Errorf("New() error : %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translator.TranslateEn(tt.args.message)
			if got != tt.want {
				t.Errorf("TranslateEn() got = %v, want %v", got, tt.want)
			}

		})
	}
}
