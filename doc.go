/*
Package basicauth makes it easy to lock pages behind Basic Authentication.

An example of how to use:

    package main

    import (
    	"log"
    	"net/http"
    	"os"

      "github.com/briangershon/basicauth"
    )

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

    func main() {
    	http.HandleFunc("/admin", adminHandler)
    	http.HandleFunc("/logout", logoutHandler)
    	err = http.ListenAndServe(":"+os.Getenv("PORT"), http.DefaultServeMux)
    	if err != nil {
    		log.Fatal("ListenAndServe: ", err)
    	}
    }

*/
package basicauth
