package auth

import "os"

func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWTSecret"))
}
