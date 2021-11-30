package utils

import "abdukhashimov/mybron.uz/config"

func FirstN(s string) string {
	input := s

	if len(input) > config.SlugLength {
		return input[:config.SlugLength]
	}
	return input
}
