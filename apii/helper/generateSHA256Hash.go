package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func GenerateSHA256Hash(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(jsonData)
	return hex.EncodeToString(hash[:]), nil
}
