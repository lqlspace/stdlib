package cryptox

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	Data struct {
		Orig string
		Signature string
	}
)

func TestClientCertificate(t *testing.T) {
	// 1. 创建秘钥对
	err := generateRsaKey(1024, "client_pub_rsa.pem", "client_rsa.pem")
	assert.Nil(t, err)

	// 2. 使用公钥进行加密
	src := []byte("just a certificate test")
	cipherText, err := RSAEncrypt(src, "server_pub_rsa.pem")
	assert.Nil(t, err)

	// 3. 数字签名
	signText, err := RSASign(cipherText, "client_rsa.pem")
	assert.Nil(t, err)

	var data Data
	data.Orig = hex.EncodeToString(cipherText)
	data.Signature = hex.EncodeToString(signText)

	dataBytes, err := json.Marshal(&data)
	assert.Nil(t, err)

	// 4. 发送给服务端
	urlPath := `http://127.0.0.1:8080/str/post`
	req, err := http.NewRequest("POST", urlPath, bytes.NewReader(dataBytes))
	assert.Nil(t, err)

	// 5. 接收来自服务器端的响应
	rsp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	b, err := ioutil.ReadAll(rsp.Body)
	assert.Nil(t, err)
	defer rsp.Body.Close()

	// 6. 反序列化
	var da Data
	err  = json.Unmarshal(b, &da)
	assert.Nil(t, err)

	// 6. 验证数字签名
	sig, _ := hex.DecodeString(da.Signature)
	orig, _ := hex.DecodeString(da.Orig)
	err = RSAVerify(orig, sig, "server_pub_rsa.pem")
	assert.Nil(t, err)

	t.Logf("data = %s\n", string(orig))
}


func TestServerCertificate(t *testing.T) {
	// 1. 创建服务器端秘钥对
	err := generateRsaKey(1024, "server_pub_rsa.pem", "server_rsa.pem")
	assert.Nil(t, err)

	// 2. 绑定路由
	http.HandleFunc("/str/post", HandleCertificate)

	// 3. 启动服务
	fmt.Printf("conn with  params startup: ...\n")
	http.ListenAndServe(":8080", nil)
}


func HandleCertificate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "can not parse body")
	}
	defer r.Body.Close()

	var data Data
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Fprintf(w, "unmarshal data failed\n")
	}

	orig, err := hex.DecodeString(data.Orig)
	if err != nil {
		fmt.Fprintf(w, "hex decode failed\n")
	}
	sig, err := hex.DecodeString(data.Signature)
	if err != nil {
		fmt.Fprintf(w, "hex sig decode failed\n")
	}
	//1. 使用客户端公钥verify是否来自客户端
	err = RSAVerify(orig, sig, "client_pub_rsa.pem")
	if err != nil {
		fmt.Fprintf(w, "verify client identity failed")
		return
	}

	// 4. 使用私钥进行解密
	plainText, err := RSADecrypt(orig, "server_rsa.pem")
	if err != nil {
		fmt.Fprintf(w, "decrypt data failed:%s\n", err.Error())
		return
	}

	fmt.Printf("data = %s\n", string(plainText))

	// 5. 将接收到的数据签名后返回
	signedText, err := RSASign(plainText, "server_rsa.pem")
	if err != nil {
		fmt.Fprintf(w, "server sign failed: %s\n", err.Error())
		return
	}
	data = Data{
		Orig: hex.EncodeToString(plainText),
		Signature: hex.EncodeToString(signedText),
	}
	dataBytes, err := json.Marshal(&data)
	if err != nil {
		fmt.Fprintf(w, "server json failed: %s\n", err)
		return
	}

	fmt.Fprintf(w, "%s", dataBytes)
}


func generateRsaKey(keySize int, pubName, privName string) error {
	//1. 创建私钥
	priv, err := generateRSAPriv(keySize, privName)
	if err != nil {
		return err
	}

	//2. 创建公钥
	if err := generateRSAPub(priv, pubName); err != nil {
		return err
	}

	return nil
}

func generateRSAPriv(keySize int, privName string) (*rsa.PrivateKey, error) {
	privateKey,err:=rsa.GenerateKey(rand.Reader,keySize)
	if err != nil{
		return nil, err
	}

	derText := x509.MarshalPKCS1PrivateKey(privateKey)

	block:=pem.Block{
		Type: "rsa private key",
		Bytes:derText,
	}
	fo,err:=os.Create(privName)
	if err!=nil{
		panic(err)
	}
	defer fo.Close()
	pem.Encode(fo,&block)

	return privateKey, nil
}


func generateRSAPub(priv *rsa.PrivateKey, privName string) error {
	publicKey:=priv.PublicKey

	derStream,err:=x509.MarshalPKIXPublicKey(&publicKey)
	if err!=nil{
		panic(err)
	}
	block := pem.Block{
		Type:"rsa public key",
		Bytes:derStream,
	}

	file, err := os.Create(privName)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	pem.Encode(file,&block)

	return nil
}


func RSAEncrypt(plainText []byte,fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buf)

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pubKey := pubInterface.(*rsa.PublicKey)
	cipherText,err := rsa.EncryptPKCS1v15(rand.Reader,pubKey,plainText)
	if err != nil{
		return nil, err
	}
	return cipherText, nil
}

func RSADecrypt(cipherText []byte, fileName string) ([]byte, error) {
	file,err:=os.Open(fileName)
	if err!=nil{
		return nil, err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block,_ := pem.Decode(buf)
	privKey,err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		return nil, err
	}

	plainText,err := rsa.DecryptPKCS1v15(rand.Reader,privKey,cipherText)
	return plainText, nil
}


func RSASign(cipherText []byte, fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buf)

	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write(cipherText)
	digest := h.Sum(nil)

	signText, err := rsa.SignPKCS1v15(nil, privKey, crypto.SHA256, digest)
	if err != nil {
		return nil, err
	}

	return signText, nil
}


func RSAVerify(text, sig []byte,fileName string) error {
	file,err:=os.Open(fileName)
	if err!=nil{
		return err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	block,_ := pem.Decode(buf)
	pubKey,err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		return err
	}

	h := sha256.New()
	h.Write(text)
	digest := h.Sum(nil)

	pk := pubKey.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pk, crypto.SHA256, digest, sig)
	if err != nil {
		return err
	}

	return nil
}




