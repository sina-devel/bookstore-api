package config

type (
	Config struct {
		Database Database `yaml:"database"`
		I18n     I18n     `yaml:"i18n"`
		Log      Log      `yaml:"log"`
	}

	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
		TimeZone string `yaml:"time_zone"`
		Charset  string `yaml:"charset"`
	}

	I18n struct {
		BundlePath string `yaml:"bundle_path"`
	}

	Log struct {
		MaxAge          string `yaml:"max_age"`
		MaxSize         string `yaml:"max_size"`
		FilenamePattern string `yaml:"filename_pattern"`
		RotationTime    string `yaml:"rotation_time"`
		InternalPath    string `yaml:"internal_path"`
		RequestPath     string `yaml:"request_path"`
	}
)
