package utils

import "abdukhashimov/mybron.uz/config"

func FirstN(s string) string {
	if len(s) > config.SlugLength {
		return s[:config.SlugLength]
	}
	return s
}
