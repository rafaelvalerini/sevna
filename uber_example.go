/*

Example Usage:

1) Sign up for credentials at https://developer.uber.com/
2) Export credentials to environment

export UBER_API_SERVER_TOKEN="xxx"
export UBER_API_CLIENT_ID="xxx"
export UBER_API_CLIENT_SECRET="xxx"
export UBER_API_REDIRECT_URI="http://localhost:8000/authenticate/"
export UBER_API_SCOPES="profile request places all_trips"

3) $ go run uber_rides_client.go

4) $ open http://localhost:8000/

5) Follow link to trigger oauth flow for uber

6) Simulate webhook call via curl

curl -X POST -d '{"event_id": "3a3f3da4-14ac-4056-bbf2-d0b9cdcb0777","event_time": 1427343990,"event_type": "all_trips.status_changed","meta": {"user_id": "d13dff8b","resource_id": "2a2f3da4","resource_type": "request","status": "accepted"},"resource_href": "https://api.uber.com/v1/requests/2a2f3da4"}' --header "Content-Type:application/json"  --header "X-Uber-Signature:898ceee917f74f694a65ec8ca0ca45012f86cdfab852598f7411298aeaed3738" http://localhost:8000/webhook

*/

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	uber "github.com/r-medina/go-uber"
)

var (
	clientID          = os.Getenv("UBER_API_CLIENT_ID")
	clientSecret      = os.Getenv("UBER_API_CLIENT_SECRET")
	clientRedirectURI = os.Getenv("UBER_API_REDIRECT_URI")
	clientScopes      = os.Getenv("UBER_API_SCOPES")
	serverToken       = os.Getenv("UBER_API_SERVER_TOKEN")
)

var client = uber.NewClient(serverToken)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body><a href=\"/authorize/\">Login to take an Uber to your Nest...</a></body></html>")
}

func authorize(w http.ResponseWriter, r *http.Request) {

	url, _ := client.OAuth(clientID, clientSecret, clientRedirectURI, clientScopes)

	http.Redirect(w, r, url, http.StatusFound)
	return
}

func authenticate(w http.ResponseWriter, r *http.Request) {

	accessToken := r.FormValue("code")

	/*
		Save access token from authorization request and persist to data store (refresh as needed)
	*/
	client.SetAccessToken(accessToken)

	/*
		Get user profile from /v1/me endpoint + or make a request estimate call (https://developer.uber.com/docs/rides/api/v1-me)
		{
		  "first_name": "Uber",
		  "last_name": "Developer",
		  "email": "developer@uber.com",
		  "picture": "https://...",
		  "promo_code": "teypo",
		  "mobile_verified": true,
		  "uuid": "91d81273-45c2-4b57-8124-d0165f8240c0"
		}
	*/
	profile, err := client.GetUserProfile()
	if err != nil {
		fmt.Println(err)
	} else {
		io.WriteString(w, string(profile.Email)+" "+string(accessToken))
	}

	return
}

func webhook(w http.ResponseWriter, r *http.Request) {

	/*
		Listen for webhook notifications (https://developer.uber.com/docs/rides/webhooks)
	*/

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		fmt.Println(err)

	} else {

		hash := hmac.New(sha256.New, []byte(clientSecret))
		hash.Write([]byte(body))

		/*
			Validate uber signature and if valid act on webhook
		*/
		if r.Header.Get("X-Uber-Signature") == hex.EncodeToString(hash.Sum(nil)) {
			/*
			 3) Setup go api client with server token + user access token

			    var client = uber.NewClient(serverToken)
			 	  client.SetAccessToken(accessToken)
			*/
			io.WriteString(w, string(body))
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
	}

}

func main() {

	http.HandleFunc("/authorize/", authorize)
	http.HandleFunc("/authenticate/", authenticate)
	http.HandleFunc("/webhook/", webhook)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)

}
