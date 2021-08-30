package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	if err:= godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func  main()  {

	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)

	r.HandleFunc("/login/github", githubLoginHandler)
	r.HandleFunc("/login/github/callback/client/{client_id}", githubCallbackHandler)
	r.HandleFunc("/loggedin", func(w http.ResponseWriter, r *http.Request) {
		loggedinHandler(w, r, "")
	})

	fmt.Println("[ UP ON PORT 3000 ]")
	log.Panic(
		http.ListenAndServe(":3000", r),
	)
}

func getGithubClientID() string {
	githubClientID, exists := os.LookupEnv("CLIENT_ID")

	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}

	return githubClientID
}

func getGithubClientSecret() string {
	githubClientSecret, exists := os.LookupEnv("CLIENT_SECRET")

	if !exists {
		log.Fatal("Github Client Secret not defined in .env file")
	}

	return githubClientSecret
}

func getGithubAccessToken(code string) string {
	clientID := getGithubClientID()
	clientSecret := getGithubClientSecret()

	requestBodyMap := map[string]string{"client_id": clientID, "client_secret": clientSecret, "code": code}

	requestJSON, _ := json.Marshal(requestBodyMap)

	req, reqerr := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))
	if reqerr != nil {
		log.Panic("Request creation failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, resperr := http.DefaultClient.Do(req)

	if resperr != nil {
		log.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType     string `json:"token_type"`
		Scope         string `json:"scope"`
	}

	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	fmt.Println(ghresp)

	return ghresp.AccessToken
}

func getGithubData(accessToken string) string {
	req, reqerr := http.NewRequest("GET", "https://api.github.com/user", nil)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)
	return string(respbody)
}

func rootHandler(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprintf(w, `<a href="/login/github/">LOGIN</a>`)
	http.RedirectHandler("/login/github", 301)
	//http.Redirect(w, r, "/login/github", 301)
}

func githubLoginHandler(w http.ResponseWriter, r *http.Request)  {
	githubClientID := getGithubClientID()
	redirectURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		githubClientID, "http://localhost:3000/login/github/callback/client/sdfgasdf2q")

	fmt.Println(redirectURL)
	http.RedirectHandler(redirectURL, 301)
	//http.Redirect(w, r, redirectURL, 301)
}

func githubCallbackHandler(w http.ResponseWriter, r *http.Request)  {
	code := r.URL.Query().Get("code")
	vars := mux.Vars(r)
	client_id := vars["client_id"]
	fmt.Println(client_id)

	githubAccessToken := getGithubAccessToken(code)

	githubData := getGithubData(githubAccessToken)

	loggedinHandler(w, r, githubData)
}

func loggedinHandler(w http.ResponseWriter, r *http.Request, githubData string)  {
	if githubData == "" {
		fmt.Fprint(w, "UNAUTHORIZED!")
		return
	}

	w.Header().Set("Content-type", "application/json")

	var prettyJSON bytes.Buffer

	parserr := json.Indent(&prettyJSON, []byte(githubData), "", "\t")
	if parserr != nil {
		log.Panic("JSON parse error")
	}

	fmt.Fprint(w, string(prettyJSON.Bytes()))
}