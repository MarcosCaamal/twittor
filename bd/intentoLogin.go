package bd

import (
	"github.com/MarcosCaamal/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo de login a la BD*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encoontrado, _ := ChequeoYaExisteUsuario(email)
	if encoontrado == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true

}
