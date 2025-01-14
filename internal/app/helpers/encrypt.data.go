package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"siap_app/config"
	"siap_app/internal/app/entity"
)

func GetInputDataRequest(c *http.Request) ([]byte, error) {
	var (
		encryptedData entity.RequestData
		requestInput  entity.RequestInput
	)

	body, err := ioutil.ReadAll(c.Body)
	if err != nil {
		return nil, errors.New("failed to read request body")
	}
	defer c.Body.Close()

	err = json.Unmarshal(body, &encryptedData)
	if err != nil {
		return nil, errors.New("failed to parse JSON")
	}

	if encryptedData.Data == "" {
		return nil, errors.New("invalid input: empty data field")
	}

	data, err := DecryptData(encryptedData.Data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), &requestInput)
	if err != nil {
		return nil, errors.New("failed to parse decrypted data")
	}

	input, ok := requestInput.Request.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid request format")
	}

	dataBytes, err := json.Marshal(input)
	if err != nil {
		return nil, errors.New("error processing input data")
	}

	return dataBytes, nil
}

func DecryptData(input string) (string, error) {
	config := config.LoadDBConfig()
	ciphertext, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", errors.New("internal server error")
	}

	key := []byte(config.SecretDecrypt)
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.New("internal server error")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	plaintext := PKCS7Unpad(ciphertext)

	return string(plaintext), nil
}

func PKCS7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
