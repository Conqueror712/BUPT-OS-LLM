package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func calculateMD5(inputString string) string {
	hash := md5.New()
	hash.Write([]byte(inputString))
	encrypted := hex.EncodeToString(hash.Sum(nil))
	return encrypted
}

func doRequest(c *gin.Context) {
	url := "https://api.baichuan-ai.com/v1/chat"
	apiKey := "your_api_key"
	secretKey := "your_secret_key"

	data := map[string]interface{}{
		"model": "Baichuan2-53B",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": "世界第一高峰是",
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.String(http.StatusInternalServerError, "请求失败，无法解析JSON数据")
		return
	}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	signature := calculateMD5(secretKey + string(jsonData) + timeStamp)

	headers := map[string]string{
		"Content-Type":     "application/json",
		"Authorization":    "Bearer " + apiKey,
		"X-BC-Request-Id":  "your requestId",
		"X-BC-Timestamp":   timeStamp,
		"X-BC-Signature":   signature,
		"X-BC-Sign-Algo":   "MD5",
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.String(http.StatusInternalServerError, "请求失败，无法创建请求")
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "请求失败，无法发送请求")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "请求失败，无法读取响应体")
		return
	}

	if resp.StatusCode == http.StatusOK {
		c.String(http.StatusOK, fmt.Sprintf("请求成功！\n响应header: %+v\n响应body: %s", resp.Header, string(body)))
	} else {
		c.String(http.StatusInternalServerError, fmt.Sprintf("请求失败，状态码: %d", resp.StatusCode))
	}
}

func main() {
	r := gin.Default()

	r.POST("/chat", doRequest)

	r.Run(":8080")
}