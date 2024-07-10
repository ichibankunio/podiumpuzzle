package mygame

import (
	"bytes"
	"embed"
	"image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/image/font"
)

var audioContext *audio.Context

var (
	//go:embed images
	imagesDir embed.FS
	//go:embed fonts
	fontsDir embed.FS
	//go:embed se
	seDir embed.FS
)

const (
	SE_EDIT = iota
	SE_GREAT
	SE_METAL
	SE_PIRON
	SE_SCORE_RISE
	SE_SCORE_RISE_2
	SE_SUCCESS
)

var (
	images []*ebiten.Image
	mainFont font.Face
	middleFont font.Face
	smallFont font.Face
	bigFont font.Face

	se []*audio.Player
)

func initResource() {
	imagesFS, err := imagesDir.ReadDir("images")

    if err != nil {
        log.Fatal(err)
    }

    for _, fs := range imagesFS {
        if !fs.IsDir() {
            file, err := imagesDir.ReadFile("images/" + fs.Name())

            if err != nil {
                log.Fatal(err)
            }
			
			images = append(images, newImageFromBytes(file))
            
        }
    }

	audioContext = audio.NewContext(SAMPLE_RATE)

	seFS, err := seDir.ReadDir("se")

	if err != nil {
		log.Fatal(err)
	}

	for _, fs := range seFS {
		if !fs.IsDir() {
			file, err := seDir.ReadFile("se/" + fs.Name())

			if err != nil {
				log.Fatal(err)
			}

			se = append(se, newSEFromBytes(file))

		}
	}

	for _, v := range se {
		v.SetVolume(0.7)
	}

	
	fontByte, err := fontsDir.ReadFile("fonts/subsetMplus.ttf")
	if err != nil {
		log.Fatal(err)
	}

	mainFont = newFontFromBytes(fontByte, 96)
	middleFont = newFontFromBytes(fontByte, 48)
	smallFont = newFontFromBytes(fontByte, 36)
	bigFont = newFontFromBytes(fontByte, 256)
}

func newImageFromBytes(b []byte) *ebiten.Image {
	r := bytes.NewReader(b)
	p, _ := png.Decode(r)
	return ebiten.NewImageFromImage(p)
}

func newFontFromBytes(byteData []byte, size int) font.Face {
	tt, err := truetype.Parse(byteData)
	if err != nil {
		log.Fatal(err)
	}

	return truetype.NewFace(tt, &truetype.Options{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func newSEFromBytes(b []byte) *audio.Player {
	m, _ := mp3.DecodeWithSampleRate(SAMPLE_RATE, bytes.NewReader(b))
	p, _ := audio.NewPlayer(audioContext, m)
	return p
}
