package env

func GetString(key, fallback string) string {
	if key == "" {
		return fallback
	}

	return key
}
