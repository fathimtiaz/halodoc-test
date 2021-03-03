package config

func Get() (Config, error) {
	return getJson()
}
