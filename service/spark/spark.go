package spark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"paper-airplane/config"
	"paper-airplane/logging"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var Spark struct {
	Model struct {
		SparkLite string
	}
	AppId     string
	ApiSecret string
	ApiKey    string
	Client    *websocket.Conn
	Logger    *zap.Logger
}

func Send(msg string) error {
	data := genParamsSparkLite(Spark.AppId, msg)
	return Spark.Client.WriteJSON(data)
}

var Answer = make(chan string, 5)

func Init() {
	conf := config.App.Spark
	Spark.AppId = conf.AppId
	Spark.ApiSecret = conf.ApiSecret
	Spark.ApiKey = conf.ApiKey
	Spark.Model.SparkLite = "wss://spark-api.xf-yun.com/v1.1/chat"
	Spark.Logger = logging.Logger().Named("Spark")
	log := Spark.Logger

	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	var resp *http.Response
	var err error
	Spark.Client, resp, err = d.Dial(assembleAuthUrl1(Spark.Model.SparkLite, Spark.ApiKey, Spark.ApiSecret), nil)
	if err != nil {
		log.Error("websocket to spark failed", zap.Error(err), zap.String("reason", readResp(resp)))
	} else if resp.StatusCode != 101 {
		log.Error("websocket to spark failed, status code not 101", zap.Error(err), zap.String("reason", readResp(resp)))
	}

	var answer = ""
	//获取返回的数据
	go func() {
		for {
			_, msg, err := Spark.Client.ReadMessage()
			if err != nil {
				log.Error("read message error", zap.Error(err))
				continue
			}

			data := map[string]any{}
			if err := json.Unmarshal(msg, &data); err != nil {
				log.Error("failed unmarshal", zap.Error(err))
			}

			//解析数据
			payload := data["payload"].(map[string]any)
			choices := payload["choices"].(map[string]any)
			header := data["header"].(map[string]any)
			code := header["code"].(float64)

			if code != 0 {
				fmt.Println(data["payload"])
				return
			}
			status := choices["status"].(float64)
			fmt.Println(status)
			text := choices["text"].([]any)
			content := text[0].(map[string]any)["content"].(string)

			if status != 2 {
				answer += content
			} else {
				fmt.Println("收到最终结果")
				answer += content
				usage := payload["usage"].(map[string]any)
				temp := usage["text"].(map[string]any)
				totalTokens := temp["total_tokens"].(float64)
				fmt.Println("total_tokens:", totalTokens)
				Spark.Client.Close()
				break
			}
			if answer != "" {
				Answer <- answer
				answer = ""
			}
		}
	}()
}

// 生成参数
func genParamsSparkLite(appid, question string) map[string]any {
	messages := []Message{
		{Role: "user", Content: question},
	}

	data := map[string]any{
		"header": map[string]any{
			"app_id": appid,
		},
		"parameter": map[string]any{
			"chat": map[string]any{
				"domain":      "general",
				"temperature": float64(0.8),
				"top_k":       int64(6),
				"max_tokens":  int64(2048),
				"auditing":    "default",
			},
		},
		"payload": map[string]any{
			"message": map[string]any{
				"text": messages,
			},
		},
	}
	return data
}

// assembleAuthUrl1 创建鉴权 url apikey 即 hmac username
func assembleAuthUrl1(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		logging.Error("failed to parse Spark host url", zap.Error(err))
		return ""
	}
	//签名时间 "Tue, 28 May 2019 09:10:42 MST"
	date := time.Now().UTC().Format(time.RFC1123)
	//参与签名的字段 host ,date, request-line
	signFields := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sign := strings.Join(signFields, "\n")
	logging.Debug("generate spark signature string", zap.String("sign", sign))
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sign, apiSecret)
	logging.Debug("hmac with sha to base64", zap.String("base64", sha))
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey, "hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
