package entity

import (
	"errors"
	"shortener/src/random"
)

type Shorteners struct {
	ID           string `db:"id" json:"id"`
	UrlShortened string `db:"url_shortened" json:"urlShortened"`
	UrlOriginal  string `db:"url_original" json:"urlOriginal"`
	UserId       string `db:"user_id" json:"userId"`
	Visits       int16  `db:"visits" json:"visits"`
}

func (s *Shorteners) Validate() (Shorteners, error) {
	if s.UrlOriginal == "" {
		return Shorteners{}, errors.New("não existe url orginal")
	} else if len(s.UrlOriginal) > 128 {
		return Shorteners{}, errors.New("url muito grande")
	} else if s.UserId == "" {
		return Shorteners{}, errors.New("o usuario é inválido")
	}
	return Shorteners{
		UrlShortened: random.RandStringRunes(5),
		UrlOriginal:  s.UrlOriginal,
		UserId:       s.UserId,
	}, nil
}
