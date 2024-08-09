package spark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"
)

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

func hmacWithShaTobase64(data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}


// buildAuthUrl 创建需要的鉴权 URL
func buildAuthUrl(hosturl, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		log.Error("failed to parse Spark host url", zap.Error(err))
		return ""
	}
	date := time.Now().UTC().Format(time.RFC1123)

	signFields := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}

	sign := strings.Join(signFields, "\n")
	log.Debug("generate spark signature string", zap.String("sign", sign))

	sha := hmacWithShaTobase64(sign, apiSecret)

	log.Debug("hmac with sha to base64", zap.String("base64", sha))

	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey, "hmac-sha256", "host date request-line", sha)

	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)

	callurl := hosturl + "?" + v.Encode()
	return callurl
}
