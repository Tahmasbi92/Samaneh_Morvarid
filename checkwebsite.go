package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	asanakAPIURL   = "https://api.asanak.ir/v1/sms/send" // جایگزین کنید
	asanakUsername = "YOUR_ASANAK_USERNAME"              // جایگزین کنید
	asanakPassword = "YOUR_ASANAK_PASSWORD"              // جایگزین کنید
	receptor       = "YOUR_PHONE_NUMBER"                 // جایگزین کنید
)

func sendSMS(message string) error {
	// ایجاد پارامترهای درخواست
	data := url.Values{}
	data.Set("username", asanakUsername)
	data.Set("password", asanakPassword)
	data.Set("receptor", receptor)
	data.Set("message", message)

	// ایجاد درخواست POST
	resp, err := http.PostForm(asanakAPIURL, data)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// خواندن بدنه پاسخ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// بررسی وضعیت پاسخ
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// پردازش پاسخ (بر اساس فرمت پاسخ API آسانک)
	fmt.Println("SMS sent successfully, response:", string(body))

	return nil
}

func main() {
	message := "This is a test message from Go program"
	err := sendSMS(message)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
