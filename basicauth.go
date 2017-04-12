package basicauth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// IsAuthenticated returns true if Basic Authentication username and password match.
func IsAuthenticated(r *http.Request, username, password string) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == username && pair[1] == password
}

// TriggerLogin shows Basic Authentication dialog.
func TriggerLogin(w http.ResponseWriter, realm string) {
	w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}
