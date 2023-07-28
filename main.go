package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	phoneId := os.Getenv("PHONE_ID")
	phoneNumber := os.Getenv("PHONE_NUMBER")
	accessToken := os.Getenv("TOKEN")

	// APIエンドポイントURL
	baseURL := "https://graph.facebook.com"
	version := "v17.0"
	url := fmt.Sprintf("%s/%s/%s/messages", baseURL, version, phoneId)

	// JSONデータを定義
	jsonData := `{
		"messaging_product": "whatsapp",
		"recipient_type": "individual",
		"to": "` + phoneNumber + `",
		"type": "template",
		"template": {
			"name": "hello_world",
			"language": {
				"code": "en_US"
			}
		}
	}`

	// POSTリクエストを作成
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
	}

	// ヘッダーを設定
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	// HTTPクライアントを作成
	client := &http.Client{}

	// リクエストを送信してレスポンスを取得
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %s", err)
	}
	defer resp.Body.Close()

	// レスポンスを表示
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Response Body:")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}
	fmt.Println(string(body))
}
