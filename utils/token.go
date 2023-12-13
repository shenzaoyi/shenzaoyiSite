package utils

//	此包包含鉴权的各种函数
//	生成token,根据tokenString校验，获取Claims

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"os"
)

// 自定义的Claims，继承jwt.RegisteredClaims
type MyClaims struct {
	UId                  int `json:"UId"`
	jwt.RegisteredClaims     // 结构体嵌套
}

// 生成公钥和私钥,写入文件中，使用的时候直接读取就OK

func Genekey() {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("无法生成RSA私钥：", err)
		return
	}

	// 保存私钥到文件
	privateKeyFile, err := os.Create("./config/private.key")
	if err != nil {
		fmt.Println("无法创建私钥文件：", err)
		return
	}
	defer privateKeyFile.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	if err := pem.Encode(privateKeyFile, privateKeyBlock); err != nil {
		fmt.Println("无法写入私钥文件：", err)
		return
	}

	fmt.Println("私钥已保存到 private.key 文件")

	// 获取公钥
	publicKey := &privateKey.PublicKey

	// 保存公钥到文件
	publicKeyFile, err := os.Create("./config/public.key")
	if err != nil {
		fmt.Println("无法创建公钥文件：", err)
		return
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	if err := pem.Encode(publicKeyFile, publicKeyBlock); err != nil {
		fmt.Println("无法写入公钥文件：", err)
		return
	}

	fmt.Println("公钥已保存到 public.key 文件")
}

func readPrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("无法解码PEM数据")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func readPublicKeyFromFile(filePath string) (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("无法解码PEM数据")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("无法解析RSA公钥")
	}

	return rsaPublicKey, nil
}

// 使用Jwt,生成token
func GeneToken(m MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, m)
	privateKey, err := readPrivateKeyFromFile("./config/private.key")
	if err != nil {

		return "", fmt.Errorf("生成token出错: %v", err)
	}
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("生成token出错: %v", err)
	}
	return tokenString, nil
}

func CheckToken(tokenString string) (*MyClaims, error) {
	//	获取公钥
	publicKey, err := readPublicKeyFromFile("./config/public.key")
	if err != nil {
		return nil, fmt.Errorf("生成token出错：%v", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("签名方法不匹配")
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token验证失败：%v", err)
	}
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, fmt.Errorf("token claims 格式错误")
	}
	return claims, nil
}
