package core

import (
    "golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    HandleError(err)
    return string(hashedPassword)
}

func CompareHashAndPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err == nil {
        return true
    } else {
        return false
    }
}
