package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
func GetHHAuthURL() string {
	res := fmt.Sprintf(
		"https://hh.ru/oauth/authorize?response_type=code&client_id=%s&state=%s&redirect_uri=%s",
		os.Getenv("CLIENT_ID"),
		"hh_auth",
		os.Getenv("REDIRECT_URI"),
	)
	return res
}

func getCodeFromUrl(urlStr string) string {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(parsedUrl.RawQuery)
	return m["code"][0]
}
func getAccessToken(code string) string {
	log.Println("Code:", code)
	urlStr := "https://hh.ru/oauth/token"
	data := url.Values{}
	data.Set("client_id", os.Getenv("CLIENT_ID"))
	data.Add("client_secret", os.Getenv("CLIENT_SECRET"))
	data.Add("code", code)
	data.Add("grant_type", "authorization_code")
	data.Add("redirect_uri", os.Getenv("REDIRECT_URI"))
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, _ := client.Do(r)
	var jsonData map[string]string
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(body, &jsonData)
	return jsonData["access_token"]
}
func storeAccessToken(token string) {
	fh, err := os.OpenFile(".token", os.O_WRONLY+os.O_CREATE, 0700)
	if err != nil { // если возникла ошибка
		log.Panicln("Unable to create file:", err)
	}
	fh.Write([]byte(token))
	fh.Close()
}
func main() {
	log.Println("start")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(GetHHAuthURL())
	codeUrl := input("Input redirected url:")
	accessToken := getAccessToken(getCodeFromUrl(codeUrl))
	storeAccessToken(accessToken)
	log.Println("Access token:", accessToken)
	log.Println("finish")
}
