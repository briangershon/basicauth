# basicauth

Easily lock web pages behind Basic Authentication.

Written in Golang.

# Sample usage

Here are two example http handlers -- a protected one and a logout one.

```
var basicAuthUsername = "mylogin"
var basicAuthPassword = "mypassword"

func adminHandler(w http.ResponseWriter, r *http.Request) {
	if !basicauth.IsAuthenticated(r, basicAuthUsername, basicAuthPassword) {
		basicauth.TriggerLogin(w, "Danger Ahead")
		return
	}

  // handler continues if user enters correct username and password

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
  basicauth.TriggerLogin(w, "Danger Ahead")
}
```
