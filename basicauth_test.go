package basicauth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoAuthorizationHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	assert.Equal(t, false, IsAuthenticated(req, "username", "password"))
}

func TestValidAuth(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("username", "password")
	assert.Equal(t, true, IsAuthenticated(req, "username", "password"))
}

func TestInvalidAuthBadUsername(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("badname", "password")
	assert.Equal(t, false, IsAuthenticated(req, "username", "password"))
}

func TestInvalidAuthBadPassword(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("username", "badpassword")
	assert.Equal(t, false, IsAuthenticated(req, "username", "password"))
}

func TestInvalidAuthBadUsernameAndPassword(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("badusername", "badpassword")
	assert.Equal(t, false, IsAuthenticated(req, "username", "password"))
}

func TestInvalidAuthDecodeError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Basic with spaces is not valid base64")
	assert.Equal(t, false, IsAuthenticated(req, "username", "password"))
}

func TestInvalidAuthNotEnoughParts(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Basic QmFzaWMgdXNlcm5hbWU=") // base64 encoded this string: "Basic username"
	assert.Equal(t, false, IsAuthenticated(req, "username", "password"))
}

func TestTriggerLogin(t *testing.T) {
	rr := httptest.NewRecorder()
	TriggerLogin(rr, "Danger")
	assert.Equal(t, 401, rr.Code)
	assert.Equal(t, `Basic realm="Danger"`, rr.Header().Get("Www-Authenticate"))
	assert.Equal(t, "401 Unauthorized\n", rr.Body.String())
}
