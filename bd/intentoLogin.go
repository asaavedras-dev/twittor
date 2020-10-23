package bd

import (
	"twittor/models"

	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin el chequeo de login a la BD*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBd := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBd, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
