package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSendEmail(t *testing.T) {
	body := strings.NewReader(`{"email": "matheuscoppi22@gmail.com"}`)
	server, c, rec := createServer(t, "post", "/forgot-password", body, "")

	err := server.HandleRequestForgotPassword(c)

	jsonResponse := rec.Body.String()

	fmt.Println(jsonResponse)

	assert.NoError(t, err)
}
