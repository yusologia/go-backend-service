package middleware

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"service/internal/pkg/config"
	"service/internal/pkg/encryption"
	error2 "service/internal/pkg/error"
	"time"
)

func AuthenticatePrivateAPI(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logiaId := r.Header.Get("Logia-ID")
		if logiaId == "" {
			error2.ErrXtremePrivateAPIAuthentication("logia ID is missing!")
		}

		logiaName := r.Header.Get("Logia-Name")
		if logiaName == "" {
			error2.ErrXtremePrivateAPIAuthentication("logia Name missing!")
		}

		logiaSecret := r.Header.Get("Logia-Secret")
		if logiaSecret == "" {
			error2.ErrXtremePrivateAPIAuthentication("logia Secret missing!")
		}

		key := fmt.Sprintf("%s-%s", logiaId, logiaName)
		if _, ok := config.XtremeCache.Get(key); !ok {
			ec := encryption.NewPrivateAPIEncryption(logiaId)

			decrypt, err := ec.Decrypt(logiaName, logiaSecret)
			if err != nil {
				error2.ErrXtremePrivateAPIAuthentication(err.Error())
			}

			err = bcrypt.CompareHashAndPassword([]byte(decrypt), []byte(config.PrivateAPICredential[logiaName]["key"].(string)))
			if err != nil {
				error2.ErrXtremePrivateAPIAuthentication("Your credential is incorrect!")
			}

			config.XtremeCache.Set(key, decrypt, time.Hour)
		}

		next.ServeHTTP(w, r)
	})
}
