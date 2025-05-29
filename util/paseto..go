package util

import "fmt"

// download PASETO PACKAGE

func VerifyPasswprdPaseto(rawPW, encryptedPW string) (bool, error) {

	const checkPW string = "check if password is true"

	// for now false dulu
	return false, fmt.Errorf("for now salah dulu %s", checkPW)
}
