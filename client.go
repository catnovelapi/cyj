package ciyuanjiAPI

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"github.com/ciyjbo/buildHttps"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

type CiyuanjiClient struct {
	host       string
	contentKey string
	paramKey   string
	token      string
}

func InitCiyuanjiClient() *CiyuanjiClient {
	return &CiyuanjiClient{
		host:       "https://api.hwnovel.com/api/ciyuanji/client",
		contentKey: "ZUreQN0E",
		paramKey:   "NpkTYvpvhJjEog8Y051gQDHmReY54z5t3F0zSd9QEFuxWGqfC8g8Y4GPuabq0KPdxArlji4dSnnHCARHnkqYBLu7iIw55ibTo18",
	}
}
func (book *CiyuanjiClient) NewAndroidAppToken(token string) *CiyuanjiClient {
	book.token = token
	return book
}

func (book *CiyuanjiClient) getHeaders() map[string]string {
	return map[string]string{
		"channel":      "25",
		"Targetmodel":  "SM-N9700",
		"Platform":     "1",
		"oaid":         "",
		"User-Agent":   "Mozilla/5.0 (Linux; Android 11; Pixel 4 XL Build/RP1A.200720.009; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.115 Mobile Safari/537.36",
		"deviceno":     "d0b7cef20c3c6b5f",
		"version":      "3.3.2",
		"token":        book.token,
		"Content-Type": "application/json"}
}

type PostDataStruct struct {
	Param     string `json:"param"`
	RequestId string `json:"requestId"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
}

func (book *CiyuanjiClient) get(url string, params map[string]any) gjson.Result {
	return buildHttps.Get(book.host+url, book.param(params), book.getHeaders()).Debug().HttpClient().Gjson()
}

func (book *CiyuanjiClient) post(path string, params map[string]any) gjson.Result {
	formData := book.param(params)
	return buildHttps.Post(book.host+path, &PostDataStruct{
		Param:     formData["param"],
		RequestId: formData["requestId"],
		Sign:      formData["sign"],
		Timestamp: formData["timestamp"],
	}, book.getHeaders()).HttpClient().Gjson()
}

func (book *CiyuanjiClient) decryptDESECB(d, key []byte) string {
	data, _ := base64.StdEncoding.DecodeString(string(d))
	if len(key) > 8 {
		key = key[:8]
	}
	block, _ := des.NewCipher(key)
	bs := block.BlockSize()
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = book.pkcs5UnPadding(out)
	return string(out)
}

func (book *CiyuanjiClient) pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func (book *CiyuanjiClient) byteBase64(bArr []byte) []byte {
	encoded := base64.StdEncoding.EncodeToString(bArr)
	return []byte(encoded)
}

func (book *CiyuanjiClient) byteMd5(bArr []byte) []byte {
	encoded := md5.Sum(bArr)
	return encoded[:]
}

func (book *CiyuanjiClient) encodeHex(bArr []byte) []byte {
	length := len(bArr)
	cArr := make([]byte, length*2)
	i := 0
	for i2 := 0; i2 < length; i2++ {
		i3 := i + 1
		cArr[i] = "0123456789ABCDEF"[(bArr[i2]&240)>>4]
		i = i3 + 1
		cArr[i3] = "0123456789ABCDEF"[bArr[i2]&15]
	}
	return cArr
}

func (book *CiyuanjiClient) encrypt(data, key []byte) string {
	if len(key) > 8 {
		key = key[:8]
	}
	block, _ := des.NewCipher(key)
	bs := block.BlockSize()
	data = book.pkcs5Padding(data, bs)
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out)
}

func (book *CiyuanjiClient) pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (book *CiyuanjiClient) signParam(param, requestId, timestamp string) string {
	signStr := "param=" + param + "&requestId=" + requestId + "&timestamp=" + timestamp + "&key=" + book.paramKey
	return strings.ToUpper(string(book.encodeHex(book.byteMd5(book.byteBase64([]byte(signStr))))))
}

func (book *CiyuanjiClient) param(params map[string]any) map[string]string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	requestId := strings.ReplaceAll(uuid.New().String(), "-", "")
	if params == nil {
		params = make(map[string]any)
	}
	params["Timestamp"] = timestamp
	jsonBytes, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	encryptParamInfo := book.encrypt(jsonBytes, []byte(book.contentKey))
	m := map[string]string{
		"timestamp": strconv.FormatInt(timestamp, 10),
		"requestId": requestId,
		"sign":      book.signParam(encryptParamInfo, requestId, strconv.FormatInt(timestamp, 10)),
		"param":     encryptParamInfo,
	}
	return m
}
