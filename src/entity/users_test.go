package entity_test

import (
	"shortener/src/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameBlank(t *testing.T) {
	shortenersTest := entity.Users{}
	err := shortenersTest.Prepare()
	assert.Equal(t, err.Error(), "the name is required and cannot be blank")
}

func TestNameMin(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "user",
	}
	err := shortenersTest.Prepare()
	assert.Equal(t, err.Error(), "username is too weak please create a validate username")
}
func TestNameMax(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "useruseruseruseruseruseruseruseruseruseruseruseruseruseruseruseru",
	}
	err := shortenersTest.Prepare()
	assert.Equal(t, err.Error(), "this does not appear to be a valid username")
}

func TestEmailBlank(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "Username",
	}
	err := shortenersTest.Prepare()
	assert.Equal(t, err.Error(), "the Email is required and cannot be blank")
}

func TestPasswordBlank(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "Username",
		Email:    "email",
	}
	err := shortenersTest.Prepare()
	assert.Equal(t, err.Error(), "password is required and cannot be blank")
}

func TestPasswordMin(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "Username",
		Email:    "email",
		Password: "12345",
	}
	err := shortenersTest.Prepare()
	assert.Equal(t, err.Error(), "password is too weak please create a stronger password")
}
func TestPasswordMax(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "Username",
		Email:    "email",
		Password: "012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678",
	}
	err := shortenersTest.Prepare()
	assert.Equal(t, err.Error(), "this does not appear to be a valid password")
}
func TestEmailValidate(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "Username",
		Email:    "email",
		Password: "123456789",
	}
	err := shortenersTest.Prepare()
	assert.Equal(t, err.Error(), "the e-mail is invalid")
}
func TestValidUser(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "Username",
		Email:    "email@email.com",
		Password: "123456789",
	}
	err := shortenersTest.Prepare()

	assert.Equal(t, err, nil)
}

func TestFormated(t *testing.T) {
	shortenersTest := entity.Users{
		Username: "        Username             ",
		Email:    "      email@email.com    ",
		Password: "123456789",
	}

	shortenersTest.Prepare()

	assert.Equal(t, "Username", shortenersTest.Username)
	assert.Equal(t, "email@email.com", shortenersTest.Email)
}
