package accurateprevalent

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"pex-prevalent-colors-challenge/internal/app/helpers"
	"pex-prevalent-colors-challenge/internal/app/models"

	"github.com/nfnt/resize"
)

// - accuratePrevalentColor: Get the exact 3 pixels that appears the most
type AccuratePrevalentColor struct {
	URL    string
	Color1 string
	Color2 string
	Color3 string
}

func NewAccuratePrevalentColor(url string, color1 string, color2 string, color3 string) models.PrevalentColor {
	return &AccuratePrevalentColor{
		URL:    url,
		Color1: color1,
		Color2: color2,
		Color3: color3,
	}
}

func (accurate *AccuratePrevalentColor) FetchImage() (image.Image, error) {
	if len(accurate.URL) <= 0 {
		return nil, fmt.Errorf("URL is empty")
	}
	log.Println("📥 Fetching image from: ", accurate.URL)
	//Get the response bytes from the url
	response, err := http.Get(accurate.URL)
	if err != nil {
		return nil, err
	}

	img, filename, err := image.Decode(response.Body)
	if err != nil {
		log.Println("🚧🚨 Error decoding image: " + err.Error())
		return nil, err
	}
	log.Println("🖼️ Image type: " + filename)
	defer response.Body.Close()
	return img, nil
}

func (accurate *AccuratePrevalentColor) CalculatePrevalentColors(img image.Image) error {
	log.Println("🧑‍💻 Calculating prevalent color")
	//Reducing the size of the image with interpolation makes the pixels counting faster,
	//consuming less memory but less accurate
	img = resize.Resize(uint(img.Bounds().Dx())/5, 0, img, resize.Lanczos3)
	//count the pixels by hex code in map
	m := make(map[string]int)
	m["-"] = -1
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			hexPixel := helpers.RgbaToHex(img.At(x, y).RGBA())
			m[hexPixel] = m[hexPixel] + 1

			//use the same loop to set the 3 pixels that appears the most
			if (m[hexPixel] > m[accurate.Color1]) && (hexPixel != accurate.Color1) {
				if accurate.Color3 != accurate.Color2 {
					accurate.Color3 = accurate.Color2
				}
				if accurate.Color2 != accurate.Color1 {
					accurate.Color2 = accurate.Color1
				}
				accurate.Color1 = hexPixel
			} else if (m[hexPixel] > m[accurate.Color2]) && (hexPixel != accurate.Color2) && (hexPixel != accurate.Color1) {

				if accurate.Color3 != accurate.Color2 {
					accurate.Color3 = accurate.Color2
				}
				accurate.Color2 = hexPixel
			} else if m[hexPixel] > m[accurate.Color3] && (hexPixel != accurate.Color3) && (hexPixel != accurate.Color2) && (hexPixel != accurate.Color1) {
				accurate.Color3 = hexPixel
			}
		}
	}
	log.Println("🌈 Prevalent Colors: " + accurate.Color1 + ", " + accurate.Color2 + ", " + accurate.Color3)
	return nil
}

func (accurate *AccuratePrevalentColor) GetCalculatedPrevalentColors() (string, string, string) {
	return accurate.Color1, accurate.Color2, accurate.Color3
}
func (accurate *AccuratePrevalentColor) GetUrl() string {
	return accurate.URL
}
