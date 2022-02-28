package entity_test

import (
	"shortener/src/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenerUrlOriginal(t *testing.T) {
	sut := entity.Shorteners{}
	sut, err := sut.Validate()

	assert.Equal(t, "não existe url orginal", err.Error())
}

func TestShortenerUrlMax(t *testing.T) {
	sut := entity.Shorteners{
		UrlOriginal: "123das12sda213sad132asd132asd123asd123123asd123asd123asd123asd123asd132asd12313asd2132das132das132asd132asd123ad132sasd132123sda2",
	}
	sut, err := sut.Validate()
	assert.Equal(t, "url muito grande", err.Error())
}

func TestShortenerUser(t *testing.T) {
	sut := entity.Shorteners{
		UrlOriginal: "http://shortener.com",
	}
	sut, err := sut.Validate()
	assert.Equal(t, "o usuario é inválido", err.Error())
}

func TestShortenerCreated(t *testing.T) {
	sut := entity.Shorteners{
		UrlOriginal: "http://shortener.com",
		UserId:      "0000-00",
	}
	sut, _ = sut.Validate()

	var visits int16 = 0

	assert.Equal(t, "http://shortener.com", sut.UrlOriginal)
	assert.Equal(t, "0000-00", sut.UserId)
	assert.NotNil(t, sut.UrlShortened)
	assert.Equal(t, visits, sut.Visits)
}
