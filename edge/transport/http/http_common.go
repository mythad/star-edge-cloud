package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

// Get function
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			return nil, err
		}

		fmt.Println(string(body))
		return body, nil
	}

	return nil, fmt.Errorf(string(resp.StatusCode))
}

// PostValue function.
func PostValue(url string, values map[string]string) ([]byte, error) {
	// postValues := url.Values{}
	// postValues.Add("publicKey", `-----BEGIN PUBLIC KEY-----MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB-----END PUBLIC KEY-----`)
	// postValues.Add("message", "")
	// resp, err := client.PostForm("http://192.168.1.35:9091", postValues)
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}

		fmt.Println(string(body))

		return body, nil
	}

	return nil, fmt.Errorf(string(resp.StatusCode))
}

// PostJSON function
func PostJSON(url string, json []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		// panic(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			return nil, err
		}
		fmt.Println(string(body))
		return body, nil
	}

	return nil, fmt.Errorf(string(resp.StatusCode))
}

// PostData function
func PostData(url string, data []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := client.Do(req)
	if err != nil {
		// panic(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			return nil, err
		}
		// fmt.Println(string(body))
		return body, nil
	}

	return nil, fmt.Errorf(string(resp.StatusCode))
}

// postForm function.
func postForm() {
	resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

// do function
func do() {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
