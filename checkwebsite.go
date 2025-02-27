package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-co-op/gocron"
)

// Configuration
const (
	websiteURL          = "00"            // آدرس وب‌سایت خود را وارد کنید
	checkInterval       = 5 * time.Minute // بازه زمانی بررسی (مثلاً ۵ دقیقه)
	asanakUsername      = "0"             // نام کاربری آسانک
	asanakPassword      = "0"
	source              = "ه"
	recieverPhoneNumber = "0" // شماره تلفن گیرنده پیامک
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	// Define the job
	s.Every(checkInterval).Do(checkWebsiteAndSendSMS)

	// Start the scheduler
	s.StartBlocking()
}

func checkWebsiteAndSendSMS() {
	fmt.Println("Checking website status...")
	err := checkWebsite("0")
	if err != nil {
		fmt.Printf("Website is down: %s\n", err)
		sendSMS("وب‌سایت خارج از دسترس است! لطفاً فوراً بررسی کنید.")
	} else {
		fmt.Println("Website is up and running.")
	}
}

func checkWebsite(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	return nil
}

func sendSMS(message string) {
	fmt.Println("Sending SMS...")

	data := url.Values{}
	data.Set("0", asanakUsername)
	data.Set("0", asanakPassword)
	data.Set("0", recieverPhoneNumber)
	data.Set("وب‌سایت خارج از دسترس است! لطفاً فوراً بررسی کنید.", message)

	// ارسال درخواست POST به API آسانک
	resp, err := http.PostForm("0", data)
	if err != nil {
		fmt.Printf("Error sending SMS: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("SMS sent successfully!")
	} else {
		fmt.Printf("Failed to send SMS. Status code: %d\n", resp.StatusCode)
	}
}
