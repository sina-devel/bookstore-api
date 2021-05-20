package config

type (
	Config struct {
		Database Database `yaml:"database"`
		Server   Server   `yaml:"server"`
		I18n     I18n     `yaml:"i18n"`
		Logger   Logger   `yaml:"logger"`
		User     User     `yaml:"user"`
	}

	Database struct {
		Postgres Postgres `yaml:"postgres"`
	}

	Postgres struct {
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		DBName    string `yaml:"db_name"`
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		SSLMode   string `yaml:"ssl_mode"`
		TimeZone  string `yaml:"time_zone"`
		Charset   string `yaml:"charset"`
		Migration bool   `yaml:"migration"`
	}

	Server struct {
		Port int `yaml:"port"`
	}

	I18n struct {
		BundlePath string `yaml:"bundle_path"`
	}

	Logger struct {
		MaxAge          string `yaml:"max_age"`
		MaxSize         string `yaml:"max_size"`
		FilenamePattern string `yaml:"filename_pattern"`
		RotationTime    string `yaml:"rotation_time"`
		InternalPath    string `yaml:"internal_path"`
	}

	User struct {
		UsernameMinLength  int `yaml:"username_min_length"`
		UsernameMaxLength  int `yaml:"username_max_length"`
		PasswordMinLetters int `yaml:"password_min_letters"`
	}
)
