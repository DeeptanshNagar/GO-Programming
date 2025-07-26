package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// WebInfo holds extracted information from a website
type WebInfo struct {
	URL            string
	StatusCode     int
	ContentLength  int64
	ContentType    string
	Title          string
	MetaDesc       string
	Links          []string
	Images         []string
	EmailAddresses []string
	PhoneNumbers   []string
	Headers        map[string]string
	ResponseTime   time.Duration
}

const url = "https://upescsi.in/"

func main() {
	fmt.Println("🚀 Starting enhanced web scraper...")

	// Extract comprehensive information
	info, err := extractWebInfo(url)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	// Display the extracted information
	displayWebInfo(info)
}

// extractWebInfo fetches and extracts comprehensive information from a website
func extractWebInfo(targetURL string) (*WebInfo, error) {
	startTime := time.Now()

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	fmt.Printf("📡 Fetching data from: %s\n", targetURL)

	// Make the request
	response, err := client.Get(targetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer response.Body.Close()

	responseTime := time.Since(startTime)

	// Read the response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	content := string(bodyBytes)

	// Extract information
	info := &WebInfo{
		URL:           targetURL,
		StatusCode:    response.StatusCode,
		ContentLength: response.ContentLength,
		ContentType:   response.Header.Get("Content-Type"),
		ResponseTime:  responseTime,
		Headers:       make(map[string]string),
	}

	// Extract important headers
	importantHeaders := []string{"Server", "X-Powered-By", "Last-Modified", "Cache-Control"}
	for _, header := range importantHeaders {
		if value := response.Header.Get(header); value != "" {
			info.Headers[header] = value
		}
	}

	// Extract HTML elements
	info.Title = extractTitle(content)
	info.MetaDesc = extractMetaDescription(content)
	info.Links = extractLinks(content)
	info.Images = extractImages(content)
	info.EmailAddresses = extractEmails(content)
	info.PhoneNumbers = extractPhoneNumbers(content)

	return info, nil
}

// extractTitle extracts the page title
func extractTitle(content string) string {
	re := regexp.MustCompile(`<title[^>]*>([^<]+)</title>`)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return "No title found"
}

// extractMetaDescription extracts meta description
func extractMetaDescription(content string) string {
	re := regexp.MustCompile(`<meta[^>]*name=["\']description["\'][^>]*content=["\']([^"\']+)["\'][^>]*>`)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return "No description found"
}

// extractLinks extracts all href links
func extractLinks(content string) []string {
	re := regexp.MustCompile(`href=["\']([^"\']+)["\']`)
	matches := re.FindAllStringSubmatch(content, -1)

	var links []string
	seen := make(map[string]bool)

	for _, match := range matches {
		if len(match) > 1 {
			link := strings.TrimSpace(match[1])
			if link != "" && link != "#" && !seen[link] {
				links = append(links, link)
				seen[link] = true
			}
		}
	}
	return links
}

// extractImages extracts all image sources
func extractImages(content string) []string {
	re := regexp.MustCompile(`<img[^>]*src=["\']([^"\']+)["\'][^>]*>`)
	matches := re.FindAllStringSubmatch(content, -1)

	var images []string
	seen := make(map[string]bool)

	for _, match := range matches {
		if len(match) > 1 {
			img := strings.TrimSpace(match[1])
			if img != "" && !seen[img] {
				images = append(images, img)
				seen[img] = true
			}
		}
	}
	return images
}

// extractEmails extracts email addresses
func extractEmails(content string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	matches := re.FindAllString(content, -1)

	var emails []string
	seen := make(map[string]bool)

	for _, email := range matches {
		if !seen[email] {
			emails = append(emails, email)
			seen[email] = true
		}
	}
	return emails
}

// extractPhoneNumbers extracts phone numbers (basic patterns)
func extractPhoneNumbers(content string) []string {
	// Matches various phone number formats
	re := regexp.MustCompile(`(\+?\d{1,3}[-.\s]?)?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)
	matches := re.FindAllString(content, -1)

	var phones []string
	seen := make(map[string]bool)

	for _, phone := range matches {
		cleaned := strings.TrimSpace(phone)
		if cleaned != "" && !seen[cleaned] {
			phones = append(phones, cleaned)
			seen[cleaned] = true
		}
	}
	return phones
}

// displayWebInfo prints all extracted information in a formatted way
func displayWebInfo(info *WebInfo) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎯 WEB SCRAPING RESULTS")
	fmt.Println(strings.Repeat("=", 60))

	// Basic info
	fmt.Printf("🌐 URL: %s\n", info.URL)
	fmt.Printf("📊 Status Code: %d\n", info.StatusCode)
	fmt.Printf("⏱️  Response Time: %v\n", info.ResponseTime)
	fmt.Printf("📏 Content Length: %s\n", formatBytes(info.ContentLength))
	fmt.Printf("📄 Content Type: %s\n", info.ContentType)

	// Page info
	fmt.Printf("\n📰 Title: %s\n", info.Title)
	fmt.Printf("📝 Description: %s\n", info.MetaDesc)

	// Headers
	if len(info.Headers) > 0 {
		fmt.Printf("\n🔧 Important Headers:\n")
		for key, value := range info.Headers {
			fmt.Printf("   %s: %s\n", key, value)
		}
	}

	// Links
	if len(info.Links) > 0 {
		fmt.Printf("\n🔗 Links Found (%d):\n", len(info.Links))
		for i, link := range info.Links {
			if i < 10 { // Show first 10 links
				fmt.Printf("   %s\n", link)
			}
		}
		if len(info.Links) > 10 {
			fmt.Printf("   ... and %d more links\n", len(info.Links)-10)
		}
	}

	// Images
	if len(info.Images) > 0 {
		fmt.Printf("\n🖼️  Images Found (%d):\n", len(info.Images))
		for i, img := range info.Images {
			if i < 5 { // Show first 5 images
				fmt.Printf("   %s\n", img)
			}
		}
		if len(info.Images) > 5 {
			fmt.Printf("   ... and %d more images\n", len(info.Images)-5)
		}
	}

	// Contact info
	if len(info.EmailAddresses) > 0 {
		fmt.Printf("\n📧 Email Addresses Found:\n")
		for _, email := range info.EmailAddresses {
			fmt.Printf("   %s\n", email)
		}
	}

	if len(info.PhoneNumbers) > 0 {
		fmt.Printf("\n📞 Phone Numbers Found:\n")
		for _, phone := range info.PhoneNumbers {
			fmt.Printf("   %s\n", phone)
		}
	}

	fmt.Println("\n✅ Scraping completed successfully!")
}

// formatBytes converts bytes to human readable format
func formatBytes(bytes int64) string {
	if bytes < 1024 {
		return strconv.FormatInt(bytes, 10) + " B"
	}
	if bytes < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(bytes)/1024)
	}
	return fmt.Sprintf("%.1f MB", float64(bytes)/(1024*1024))
}
