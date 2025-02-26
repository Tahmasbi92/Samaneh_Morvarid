package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Configuration
const (
	websiteURL    = "https://example.com" // آدرس وب‌سایت خود را وارد کنید
	checkInterval = 5 * time.Minute       // بازه زمانی بررسی (مثلاً ۵ دقیقه)

	// Twilio credentials
	twilioAccountSid     = "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // شناسه حساب Twilio
	twilioAuthToken      = "your_auth_token"                 // توکن احراز هویت Twilio
	twilioPhoneNumber    = "+1234567890"                     // شماره تلفن Twilio شما
	recipientPhoneNumber = "+989123456789"                   // شماره تلفن گیرنده پیامک
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

	client := twilio.NewRestClient(twilioAccountSid, twilioAuthToken)

	params := &twilioApi.CreateMessageParams{
		To:   twilio.String(recipientPhoneNumber),
		From: twilio.String(twilioPhoneNumber),
		Body: twilio.String(message),
	}

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Printf("Twilio error: %s\n", err)
	} else {
		if resp.Status != nil {
			fmt.Printf("SMS sent successfully! Status: %s\n", *resp.Status)
		} else {
			fmt.Println("SMS sent successfully!")
		}
	}
}
