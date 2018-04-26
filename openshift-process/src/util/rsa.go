package util

import (
	"crypto/rand"
	"crypto/rsa"
	"os"
	"encoding/pem"
	"crypto/x509"
	"golang.org/x/crypto/ssh"
	"bytes"
	"encoding/asn1"
	"net"
	"golang.org/x/crypto/ssh/agent"
	"fmt"
)

type KeyPair struct {
	RsaKey *rsa.PrivateKey
	PublicKey ssh.PublicKey
}

type Agent struct {
	Agent agent.Agent
}

func NewKeyPair() KeyPair {
	reader := rand.Reader
	bitSize := 4096

	key, err := rsa.GenerateKey(reader, bitSize)
	ExitOnError("Error while generating RSA key", err)

	pubKey, err := ssh.NewPublicKey(&key.PublicKey)
	ExitOnError("Cannot create public SSH key", err)

	return KeyPair{key, pubKey}
}

func (key KeyPair) WritePublicKey(filename string) {
	privateKeyFile, err := os.Create(filename)
	ExitOnError("Cannot create public key file", err)
	defer privateKeyFile.Close()

	_, err = privateKeyFile.WriteString(key.GetPublicKey())
	ExitOnError("Error during writing public key file", err)
}

func (key KeyPair) WritePrivateKey(filename string) {
	privateKeyFile, err := os.Create(filename)
	ExitOnError("Cannot create private key file", err)
	defer privateKeyFile.Close()

	_, err = privateKeyFile.WriteString(key.GetPrivateKey())
	ExitOnError("Error during writing private key file", err)
}

func (key KeyPair) WritePublicPem(filename string) {
	asn1Bytes, err := asn1.Marshal(key.RsaKey.PublicKey)
	ExitOnError("Cannot encode public key bytes", err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(filename)
	ExitOnError("Cannot create public key file", err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	ExitOnError("Error during writing public key file", err)
}

func (key KeyPair) GetPublicKey() string {
	pub := ssh.MarshalAuthorizedKey(key.PublicKey)
	return bytes.NewBuffer(pub).String()
}

func (key KeyPair) GetPrivateKey() string {
	buffer := bytes.NewBufferString("")

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key.RsaKey)}
	err := pem.Encode(buffer, privateKeyPEM)
	ExitOnError("Error when encoding private key", err)

	return buffer.String()
}

func NewSshAgentClient() Agent {
	sock := os.Getenv("SSH_AUTH_SOCK")
	fmt.Println(sock)
	socket, err := net.Dial("unix", sock)
	ExitOnError("Cannot connect to SSH Agent socket", err)

	return Agent{agent.NewClient(socket)}
}

func (sshAgent Agent) AddKey(key KeyPair) {
	addedKey := agent.AddedKey{}
	addedKey.PrivateKey = key.RsaKey
	addedKey.ConfirmBeforeUse = false
	addedKey.LifetimeSecs = 0

	err := sshAgent.Agent.Add(addedKey)
	ExitOnError("Cannot add key pair to ssh agent", err)
}