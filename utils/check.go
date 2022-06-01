package utils

const (
	maxUsernameLen = 32
	maxPasswordLen = 32
	randomSaltLen  = 32
)

func CheckNameValid(username string) bool {
	return len(username) != 0 && len(username) <= maxUsernameLen
}

func CheckPasswordValid(password string) bool {
	return len(password) != 0 && len(password) <= maxUsernameLen
}
