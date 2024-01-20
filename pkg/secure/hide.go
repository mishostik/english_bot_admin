package secure

import (
	"github.com/fernet/fernet-go"
	"log"
	"time"
)

type Shield struct {
	EncryptionKey *fernet.Key
}

func NewShield(key string) *Shield {
	k := fernet.MustDecodeKeys(key)
	return &Shield{
		EncryptionKey: k[0],
	}
}

func (s *Shield) DecryptMessage(cipher string) string {
	tok := fernet.VerifyAndDecrypt([]byte(cipher), 0*time.Second, []*fernet.Key{s.EncryptionKey})
	return string(tok)
}

func (s *Shield) EncryptMessage(message string) (string, error) {
	tok, err := fernet.EncryptAndSign([]byte(message), s.EncryptionKey)
	return string(tok), err
}

func HidePassword(password string) string {
	token, err := shield.EncryptMessage(password)
	if err != nil {
		log.Fatalf("Error encrypting password: %v", err)
	}

	return token
}
