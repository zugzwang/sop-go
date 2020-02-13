package main

import (
	"bytes"
	"golang.org/x/crypto/openpgp/armor"
)

// ArmorWithType make bytes input to armor format
func ArmorWithType(input []byte, armorType string) (string, error) {
	var b bytes.Buffer
	w, err := armor.Encode(&b, armorType, nil)
	if err != nil {
		return "", err
	}
	if _, err = w.Write(input); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	return b.String(), nil
}

