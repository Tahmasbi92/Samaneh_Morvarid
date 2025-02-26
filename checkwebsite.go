package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-co-op/gocron"
)

// Configuration
const (
	websiteURL    = "https://example.com" // آدرس وب‌سایت خود را وارد کنید
	checkInterval = 5 * time.Minute       // بازه زمانی بررسی (مثلاً ۵ دقیقه)

	// Asanak credentials (جایگزین اطلاعات درست کنید)
	asanakAPIURL         = "https://api.asanak.ir/sms/send" // URL API آسانک
	asanakAPIKey         = "YOUR_API_KEY"                   // کلید API آسانک شما
	asanakPhoneNumber    = "YOUR_ASANAK_NUMBER"             // شماره تلفن آسانک شما
	recipientPhoneNumber = "+989123456789"                  // شماره تلفن گیرنده پیامک
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
	err := checkWebsite(websiteURL)
	if err != nil {
		fmt.Printf("Website is down: %s\n", err)
		sendSMS("Website is down! Please check immediately.")
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

	// Create the request body
	data := url.Values{}
	data.Set("apikey", asanakAPIKey)
	data.Set("receptor", recipientPhoneNumber)
	data.Set("message", message)
	data.Set("originator", asanakPhoneNumber)

	// Create the HTTP request
	resp, err := http.PostForm(asanakAPIURL, data)
	if err != nil {
		fmt.Printf("Asanak error: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	// Print the response from Asanak
	fmt.Printf("Asanak response: %s\n", string(body))

	// You might want to parse the response to check for success
	// and handle any errors reported by Asanak.
}
