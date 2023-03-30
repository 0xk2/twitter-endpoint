package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/0xk2/twitter-endpoint/utils"
)

type ResponseData struct {
	Success bool `json:"success"`
}

type FirstStepResponse struct {
	Url string `json:"url"`
}

var codeChallenge = utils.RandString(8)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("auth handler")
	state := r.URL.Query().Get("state")
	clientId := os.Getenv("TWITTER_CLIENT_ID")
	html := "<html><body>Error!</body></html>"
	baseUrl := "https://twitter-endpoint.herokuapp.com/"
	if state == "" { // no state, this is step 1
		state = utils.RandString(8)

		firstStepUrl := "https://twitter.com/i/oauth2/authorize?response_type=code&client_id=" + clientId + "&redirect_uri=" + baseUrl + "&scope=tweet.read%20users.read%20tweet.write&state=" + state + "&code_challenge=" + codeChallenge + "&code_challenge_method=plain"
		html = "<html><body><a href='" + firstStepUrl + "'>Click here</a></body></html>"
	} else { // if request has state then this is step 2
		code := r.URL.Query().Get("code")
		// create a post request to twitter
		urlStr := "https://api.twitter.com/2/oauth2/token"

		values := url.Values{"code": {"" + code + ""},
			"grant_type":   {"authorization_code"},
			"redirect_uri": {"" + baseUrl + ""},
			// "client_id":     {"" + clientId + ""},
			"code_verifier": {"challenge"}}
		request, err := http.NewRequest("POST", urlStr, strings.NewReader(values.Encode()))
		if err != nil {
			// handle error
			html = "<html><body>Error from new request!</body></html>"
		}
		// set Content-Type header to application/x-www-form-urlencoded
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Set("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAA6VmQEAAAAAQjuxPEE4EiMOKwthNn0H7wHWnss%3DVfBB4jB4mxIrQg1BXkjuFu9ClwrFFb103XfxEkpV2DdP3TH17V")

		// create a new HTTP client and send the request
		client := &http.Client{}
		response, err := client.Do(request)
		html = "<html><body>Request sent to twitter!</body></html>"
		// if err != nil {
		// 	// handle error
		// 	html = "<html><body>Error from response!</body></html>"
		// } else {
		// 	body, e := ioutil.ReadAll(response.Body)
		// 	if e == nil {
		// 		html = string(body)
		// 	}
		// }
		defer response.Body.Close()
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", html)
}
