package main

import (
	"github.com/teamyapp/experimental/yibolu/identity/app/idgen"
	"net/http"
)

/*
	1. Client library
		a. sign in
		b. callback
	2. Identity Backend
		a. clientID generator
		b. verify jwt
		c. sign in backend
			redirection to oauth provider
 			generate jwt
			show all sign in methods
	
	Notes:  1. generalize oauth code to support all kinds of oauth providers
			2. pull secrets from vault
			3. receive notification whenever vault config is updated
*/

func main() {
	// try to sign-in

	// check if we have jwt and if jwt is valid

	client := identity.NewClient()

	client.GetClientID()

	jwt := getJWT()
	clientID, isSignedIn := getClientID(jwt)

	if isSignedIn {
		// do the real work
		return
	}

	for {
		// retry when clientID is occupied
		err := SetupSignInCallback(clientID, handleSignInResult) // call api of client library to get the jwt
		// A websocket linked to identity service, where have a listener to mq
		if err == nil {
			break
		}
		// we need to get another clientID
		clientID, isSignedIn = getClientID(jwt)
	}

	// open frontend
	signInMethods := showSignInMethods()
	signInMethod := selectSignInMethod(signInMethods)

	// use clientID

	err = trySignIn(signInMethod, clientID) // pop web browser or a new page
}

func getClientID(jwt string)  {

}

func SetupSignInCallback(clientID string, handleSignInResult func()) {
	subscribeQueue(clientID, handleSignInResult) // 在服务端做验证 锁定 以及 rabbitmq的 subscribe
}

// client request with jwt -> backend -> get user id from jwt -> use user id to call other api

// messageQueue invoke this callback

func handleSignInResult(jwt string, err error)  {
	if err != nil {
		handleError(err)
		return
	}

	saveJWT(jwt)
	// other logic
}




