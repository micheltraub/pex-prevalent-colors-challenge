package prevalentcolors

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
)

type PrevalentColor interface {
	FetchImage() (image.Image, error)
	CalculatePrevalentColors(image.Image) error
	GetCalculatedPrevalentColors() (string, string, string)
	GetUrl() string
	SortTopColors(map[string]int, string)
	ShouldDownscale(image.Rectangle) bool
}

func ProcessPrevalentColors(p PrevalentColor, csvCh chan []string) {
	// Todo: Handle invalid formats
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)

	img, err := p.FetchImage()
	if err != nil {
		//log error and return CSV line with error message
		log.Println("🚧🚨Error downloading file: " + err.Error())
		csvCh <- []string{p.GetUrl(), "-", "-", "-", err.Error()}
		return
	}

	err = p.CalculatePrevalentColors(img)
	if err != nil {
		//log error and return CSV line with error message
		log.Println("🚧🚨Error calculating prevalent colors: " + err.Error())
		csvCh <- []string{p.GetUrl(), "-", "-", "-", err.Error()}
		return
	}
	color1, color2, color3 := p.GetCalculatedPrevalentColors()

	csvCh <- []string{p.GetUrl(), color1, color2, color3, ""}
}
