package extcert

import (
	cryptoRand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	certFileSuffix = ".crt"
	keyFileSuffix  = ".key"
)

// names 定义Name
type names struct {
	Country          string // 国家
	Province         string // 省/州
	Locality         string // 地区
	Organization     string // 组织
	OrganizationUnit string // 组织单位
}

// config 定义配置
type config struct {
	CommonName string
	Names      names
	Host       []string
	Expire     uint64 // 小时
}

// createSignFile 根据rootCA rootKey生成签发证书文件
func createSignFile(rootCA *x509.Certificate, rootKey *rsa.PrivateKey, filenamePrefix string, cfg *config) error {
	cert, key, err := createSign(rootCA, rootKey, cfg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filenamePrefix+certFileSuffix, cert, 0755)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filenamePrefix+keyFileSuffix, key, 0755)
}

// createSign 根据rootCA rootKey生成签发证书
func createSign(rootCA *x509.Certificate, rootKey *rsa.PrivateKey, cfg *config) (cert, key []byte, err error) {
	tpl := &x509.Certificate{
		SerialNumber: big.NewInt(rand.Int63()),
		Subject: pkix.Name{
			CommonName:         cfg.CommonName,
			Country:            []string{cfg.Names.Country},
			Organization:       []string{cfg.Names.Organization},
			OrganizationalUnit: []string{cfg.Names.OrganizationUnit},
			Province:           []string{cfg.Names.Province},
			Locality:           []string{cfg.Names.Locality},
		},
		NotBefore: time.Now().Add(-time.Hour),
		NotAfter:  time.Now().Add(time.Hour * time.Duration(cfg.Expire)),

		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		EmailAddresses:        []string{},
		IPAddresses:           []net.IP{},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}
	for _, host := range cfg.Host {
		if ip := net.ParseIP(host); ip != nil {
			tpl.IPAddresses = append(tpl.IPAddresses, ip)
		} else {
			tpl.DNSNames = append(tpl.DNSNames, host)
		}
	}

	// 生成公钥私钥对
	var priKey *rsa.PrivateKey
	priKey, err = rsa.GenerateKey(cryptoRand.Reader, 2048)
	if err != nil {
		return
	}
	cert, err = x509.CreateCertificate(cryptoRand.Reader, tpl, rootCA, &priKey.PublicKey, rootKey)
	if err != nil {
		return
	}
	// Generate cert
	cert = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})
	// Generate key
	key = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priKey),
	})
	return cert, key, nil
}

// CreateCAFile 生成ca证书文件
func CreateCAFile(filenamePrefix string, cfg *config) error {
	ca, key, err := createCA(cfg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filenamePrefix+certFileSuffix, ca, 0755)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filenamePrefix+keyFileSuffix, key, 0755)
}

// createCA 生成ca证书
func createCA(cfg *config) (ca, key []byte, err error) {
	var privateKey *rsa.PrivateKey

	privateKey, err = rsa.GenerateKey(cryptoRand.Reader, 2048)
	if err != nil {
		return
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:         cfg.CommonName,
			Country:            []string{cfg.Names.Country},
			Organization:       []string{cfg.Names.Organization},
			OrganizationalUnit: []string{cfg.Names.OrganizationUnit},
			Province:           []string{cfg.Names.Province},
			Locality:           []string{cfg.Names.Locality},
		},
		NotBefore: time.Now().Add(-time.Hour),
		NotAfter:  time.Now().Add(time.Hour * time.Duration(cfg.Expire)),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	var crt []byte
	crt, err = x509.CreateCertificate(cryptoRand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return
	}
	// Generate cert
	ca = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: crt,
	})
	// Generate key
	key = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	return ca, key, nil
}

func TestGenerateCA(t *testing.T) {
	err := CreateCAFile("ca", &config{
		Names: names{
			Country: "CN",
		},
		Expire: 365 * 24,
	})
	require.NoError(t, err)

	// invalid ca file
	_, _, err = ParseCrtAndKeyFile("invalid.crt", "ca.key")
	require.Error(t, err)

	// invalid key file
	_, _, err = ParseCrtAndKeyFile("ca.crt", "invalid.key")
	require.Error(t, err)

	// ca key file
	ca, key, err := ParseCrtAndKeyFile("ca.crt", "ca.key")
	require.NoError(t, err)
	require.Equal(t, "CN", ca.Subject.Country[0])

	err = key.Validate()
	require.NoError(t, err)

	// invalid ca file
	_, _, err = LoadCrtAndKeyFile("invalid.crt", "ca.key")
	require.Error(t, err)

	// invalid key file
	_, _, err = LoadCrtAndKeyFile("ca.crt", "invalid.key")
	require.Error(t, err)

	// ca key file
	caBytes, keyBytes, err := LoadCrtAndKeyFile("ca.crt", "ca.key")
	require.NoError(t, err)

	ca, key, err = ParseCrtAndKey(caBytes, keyBytes)
	require.NoError(t, err)
	require.Equal(t, "CN", ca.Subject.Country[0])

	err = key.Validate()
	require.NoError(t, err)

	// file
	caBytes, keyBytes, err = LoadPair("ca.crt", "ca.key")
	require.NoError(t, err)

	_, _, err = ParseCrtAndKey(caBytes, keyBytes)
	require.NoError(t, err)
	require.Equal(t, "CN", ca.Subject.Country[0])

	// base64 string
	caStr := base64Prefix + base64.StdEncoding.EncodeToString(caBytes)
	keyStr := base64Prefix + base64.StdEncoding.EncodeToString(keyBytes)
	caBytes, keyBytes, err = LoadPair(caStr, keyStr)
	require.NoError(t, err)

	ca, key, err = ParseCrtAndKey(caBytes, keyBytes)
	require.NoError(t, err)
	require.Equal(t, "CN", ca.Subject.Country[0])

	// invalid base64 string
	_, _, err = LoadPair(base64Prefix+"invalidbase64", base64Prefix+"invalidbase64")
	require.Error(t, err)

	err = key.Validate()
	require.NoError(t, err)

	os.Remove("ca.crt")
	os.Remove("ca.key")
}

func TestSign(t *testing.T) {
	err := CreateCAFile("ca", &config{
		CommonName: "server",
		Names: names{
			Country:      "CN",
			Organization: "test",
		},
		Expire: 365 * 24,
	})
	require.NoError(t, err)

	ca, key, err := ParseCrtAndKeyFile("ca.crt", "ca.key")
	require.NoError(t, err)

	err = createSignFile(ca, key, "server", &config{
		CommonName: "server.com",
		Host:       []string{"server.com"},
		Names: names{
			Country:      "CN",
			Organization: "test",
		},
		Expire: 365 * 24,
	})
	require.NoError(t, err)

	srvCa, srvKey, err := ParseCrtAndKeyFile("server.crt", "server.key")
	require.NoError(t, err)
	require.Equal(t, "server.com", srvCa.Subject.CommonName)

	err = srvKey.Validate()
	require.NoError(t, err)
	os.Remove("ca.crt")
	os.Remove("ca.key")
	os.Remove("server.crt")
	os.Remove("server.key")
}

func TestInvalid(t *testing.T) {
	t.Run("invalid key", func(t *testing.T) {
		_, err := ParseKey([]byte{})
		require.Error(t, err)
	})

	t.Run("invalid crt", func(t *testing.T) {
		_, _, err := ParseCrtAndKey([]byte{}, []byte{})
		require.Error(t, err)
	})
}
