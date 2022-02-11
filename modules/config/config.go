package config

import (
	"io/ioutil"
	"log"
	"os"

	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"encoding/json"
	"encoding/pem"
)

var Config Configuration

type SSH struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	HostKeyFile string `json:"hostKeyFile"`
	TextDisplay string `json:"textDisplay"`
	ClientAuth  bool   `json:"clientAuth"`
}

type Configuration struct {
	SSH SSH `json:"ssh"`
}

func LoadConfig(configFile string) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal([]byte(file), &Config)
	generateKey() //generate key RSA
}

func generateKey() error {
	_, err := os.Stat(Config.SSH.HostKeyFile)
	if os.IsNotExist(err) {
		log.Printf("generateKey NewRSA\n")
		key, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			return err
		}
		if err = key.Validate(); err != nil {
			return err
		}

		priv := x509.MarshalPKCS1PrivateKey(key)
		privBlock := pem.Block{
			Type:    "RSA PRIVATE KEY",
			Headers: nil,
			Bytes:   priv,
		}
		privatePEM := pem.EncodeToMemory(&privBlock)
		if err := ioutil.WriteFile(Config.SSH.HostKeyFile, privatePEM, 0644); err != nil {
			return err
		}
	}
	return nil
}
