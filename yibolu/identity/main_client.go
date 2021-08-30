package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const clientID = "4da1b4a1f09b0ba7a81e"
const clientSecret = "7ebe0c784eaa7f836d373bb0ced17ee0bffda1dd"

const googleClientID = "893937988570-u5doflho5jj6169767q4svnl693f41eq.apps.googleusercontent.com"
const googleClientSecret = "u0Z-ldhCar1TYWqOByBPSCp6"

var (
	googleOauthConfig *oauth2.Config
)
func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
<body>
	<a href="/login">Google Log In</a>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
}

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", content)
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)
	http.ListenAndServe(":8080", nil)
}

//func main() {
//	fs := http.FileServer(http.Dir("public"))
//	http.Handle("/", fs)
//
//	// We will be using `httpClient` to make external HTTP requests later in our code
//	httpClient := http.Client{}
//
//	// Create a new redirect route route
//	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
//		// First, we need to get the value of the `code` query param
//		err := r.ParseForm()
//		if err != nil {
//			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
//			w.WriteHeader(http.StatusBadRequest)
//		}
//		code := r.FormValue("code")
//
//		// Next, lets for the HTTP request to call the github oauth enpoint
//		// to get our access token
//		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
//		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
//		if err != nil {
//			fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
//			w.WriteHeader(http.StatusBadRequest)
//			}
//		// We set this header since we want the response
//		// as JSON
//		req.Header.Set("accept", "application/json")
//
//		// Send out the HTTP request
//		res, err := httpClient.Do(req)
//		if err != nil {
//			fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
//			w.WriteHeader(http.StatusInternalServerError)
//		}
//		defer res.Body.Close()
//
//		// Parse the request body into the `OAuthAccessResponse` struct
//		var t OAuthAccessResponse
//		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
//			fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
//			w.WriteHeader(http.StatusBadRequest)
//		}
//
//		// Finally, send a response to redirect the user to the "welcome" page
//		// with the access token
//		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
//
//		w.WriteHeader(http.StatusFound)
//	})
//
//	http.ListenAndServe(":8080", nil)
//}
//
//type OAuthAccessResponse struct {
//	AccessToken string `json:"access_token"`
//}