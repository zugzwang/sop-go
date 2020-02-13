package main

import "golang.org/x/crypto/openpgp"
import "bytes"
import "fmt"

// GenerateKey  creates a single default OpenPGP certificate with zero or more
// User IDs.
func GenerateKey(userID ...string) error {
	if len(userID) == 0 {
		return Err19
	}
	// FIXME: Multiple userIDs
	// FIXME: x/crypto breaks arbitrary userIDs. It expects name, comment, email.
	ent, err := openpgp.NewEntity(userID[0], SOPGOComment, "", nil)
	if err != nil {
		return Err99(err)
	}
	w := bytes.NewBuffer(nil)
	if err := ent.SerializePrivate(w, nil); err != nil {
		return Err99(err)
	}
	skSerialized := w.Bytes()
	sk, err := ArmorWithType(skSerialized, "BEGIN PRIVATE KEY BLOCK")
	if err != nil {
		return Err99(err)
	}
	fmt.Println(sk)

	return nil
}
