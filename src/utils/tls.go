package utils

import (
	rand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

//生成 SSL证书、RSA 私钥
func GetTLS(certPath , keyPath string)  {
	//生成序列号
	max := new(big.Int).Lsh(big.NewInt(1),128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	// x509识别
	subject := pkix.Name{
		Organization: []string{"github.com/L1ng14"},
		OrganizationalUnit: []string{"l1ng14"},
		CommonName: "github.com",
	}

	// x509证书
	tempalte := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: subject,
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(365*24*time.Hour),
		KeyUsage: x509.KeyUsageDataEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},  // 证书只允许在 127.0.0.1上运行
	}

	//生成 RSA 私钥
	pk ,_ := rsa.GenerateKey(rand.Reader ,2048)

	// 公钥，生成 SSL 证书时需要用到
	darBtyes,_ := x509.CreateCertificate(rand.Reader,&tempalte,&tempalte,&pk.PublicKey,pk)

	//结合公钥，写入证书到相应路径
	certOut , _ := os.Create(certPath)
	defer certOut.Close()
	pem.Encode(certOut, &pem.Block{
		Type: "CERTIFICATE",
		Bytes: darBtyes,
	})


	//写入私钥,写入到相应路径
	keyOut ,_:= os.Create(keyPath)
	defer keyOut.Close()
	pem.Encode(keyOut,&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk),
	})


}
