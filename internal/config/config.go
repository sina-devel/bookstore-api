package config

type (
	Config struct {
		Database Database `yaml:"database"`
		I18n     I18n     `yaml:"i18n"`
		Logger   Logger   `yaml:"logger"`
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
)
