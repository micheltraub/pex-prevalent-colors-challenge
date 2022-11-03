package models

import "image"

type PrevalentColor interface {
	FetchImage() (image.Image, error)
	CalculatePrevalentColors(image.Image) error
	GetCalculatedPrevalentColors() (string, string, string)
	GetUrl() string
}
