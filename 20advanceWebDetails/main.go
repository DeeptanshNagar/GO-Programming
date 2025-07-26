package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// 🌐 Website address we want to visit
const url = "https://upesacm.org/"

func main() {
	fmt.Println("📡 Connecting to:", url)

	// 🚀 Step 1: Make an HTTP client that follows redirects and captures TLS info
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return fmt.Errorf("❌ too many redirects")
			}
			fmt.Println("🔀 Redirected to:", req.URL)
			return nil
		},
		Timeout: 10 * time.Second,
	}

	// 🚀 Step 2: Send request
	response, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// ✅ Step 3: Print basic response info
	fmt.Println("\n🌐 --- Website Info ---")
	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Status      :", response.Status)
	fmt.Println("Protocol    :", response.Proto)

	// ✅ Step 4: Print headers
	fmt.Println("\n📦 --- Response Headers ---")
	for key, value := range response.Header {
		fmt.Printf("%s: %s\n", key, value)
	}

	// ✅ Step 5: SSL Certificate Details (only for HTTPS)
	if response.TLS != nil && len(response.TLS.PeerCertificates) > 0 {
		cert := response.TLS.PeerCertificates[0]
		fmt.Println("\n🔐 --- SSL Certificate Info ---")
		fmt.Println("Issuer        :", cert.Issuer)
		fmt.Println("Subject       :", cert.Subject)
		fmt.Println("Valid From    :", cert.NotBefore)
		fmt.Println("Valid Until   :", cert.NotAfter)
		fmt.Println("DNS Names     :", cert.DNSNames)
	} else {
		fmt.Println("\n⚠️ No SSL certificate information found (site might be HTTP only)")
	}

	// ✅ Step 6: Read HTML Content
	fmt.Println("\n📜 --- Reading HTML Content ---")
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Printf("Total Bytes Received: %d bytes\n", len(databytes))

	// ✅ Step 7: Print a preview of the HTML
	fmt.Println("\n📝 --- HTML Preview (First 500 chars) ---")
	if len(content) > 500 {
		fmt.Println(content[:500] + "...")
	} else {
		fmt.Println(content)
	}

	// ✅ Step 8: Save the entire HTML to a file
	fileName := "upesacm.html"
	err = os.WriteFile(fileName, databytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n💾 Full HTML saved to:", fileName)
}
