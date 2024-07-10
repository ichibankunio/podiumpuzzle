package mygame

import (
	"image/color"
	"math/rand"
	"reflect"

	// "reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	LEFT_CENTER = iota
	RIGHT_CENTER
)

var medals []*medal

// var medalTxtData = []string{"Au", "Ag", "Cu"}
var medalTxtData = []string{"銀", "金", "銅"}
var medalClrData = [][]uint8{
	{192, 192, 192, 255},
	// {255, 215, 0, 255},
	{255, 255, 0, 255},

	{196, 112, 34, 255},
}
var medalBgData = []int{IMG_GOLD, IMG_SILVER,IMG_BRONZE}
var medalFilterData = []int{IMG_FILTER_SILVER, IMG_FILTER_GOLD, IMG_FILTER_BRONZE}

var medalPlaceData = []float64{30, 60, 0}

type medal struct {
	txt *txtObj
	bg *sprite
	podium *sprite
	filter *sprite

	clrIndex int
	txtIndex int
	placeIndex int
}

func initMedal() {
	// newMedals()
}

func editColor(target int) {
	l := medals[0].clrIndex
	c := medals[1].clrIndex
	r := medals[2].clrIndex
	switch target {
	case LEFT_CENTER:
		medals[0].clrIndex = c
		medals[1].clrIndex = l
	case RIGHT_CENTER:
		medals[1].clrIndex = r
		medals[2].clrIndex = c
	}

	for i, _ := range medals {
		medals[i].txt.clr = dataToColor(medalClrData[medals[i].clrIndex])
	}
}

func editPlace(target int) {
	l := medals[0].placeIndex
	c := medals[1].placeIndex
	r := medals[2].placeIndex
	switch target {
	case LEFT_CENTER:
		medals[0].placeIndex = c
		medals[1].placeIndex = l
	case RIGHT_CENTER:
		medals[1].placeIndex = r
		medals[2].placeIndex = c
	}
	for i, _ := range medals {
		medals[i].podium = newSprite(ebiten.NewImage(132, 50 + int(medalPlaceData[medals[i].placeIndex])), centerX(medals[i].podium.img, SCREEN_WIDTH * (1 + i) / 4), 300 - medalPlaceData[medals[i].placeIndex] + 96)
		medals[i].podium.img.Fill(color.White)
		medals[i].txt.spr.y = 300 - medalPlaceData[medals[i].placeIndex]
		medals[i].bg.y = medals[i].txt.spr.y
		medals[i].filter.y = medals[i].txt.spr.y
	}
}

func editTextEffect(target int) {
	switch target {
	case LEFT_CENTER:
		medals[0].txt.clr = color.RGBA{228, 0, 127, 255}
		medals[1].txt.clr = color.RGBA{228, 0, 127, 255}
	case RIGHT_CENTER:
		medals[1].txt.clr = color.RGBA{228, 0, 127, 255}
		medals[2].txt.clr = color.RGBA{228, 0, 127, 255}
	}
}

func resetEditTextEffect() {
	for i, _ := range medals {
		medals[i].txt.clr = dataToColor(medalClrData[medals[i].clrIndex])
	}
}

func editColorEffect(target int) {
	switch target {
	case LEFT_CENTER:
		medals[0].bg.img.Fill(dataToColor(medalClrData[medals[0].clrIndex]))
		medals[1].bg.img.Fill(dataToColor(medalClrData[medals[1].clrIndex]))
	case RIGHT_CENTER:
		medals[1].bg.img.Fill(dataToColor(medalClrData[medals[1].clrIndex]))
		medals[2].bg.img.Fill(dataToColor(medalClrData[medals[2].clrIndex]))
	}
}

func resetEditColorEffect() {
	for i, _ := range medals {
		medals[i].bg.img.Fill(color.Black)
	}
}

func editPlaceEffect(target int) {
	switch target {
	case LEFT_CENTER:
		medals[0].podium.img.Fill(color.RGBA{0, 161, 233, 255})
		medals[1].podium.img.Fill(color.RGBA{0, 161, 233, 255})
	case RIGHT_CENTER:
		medals[1].podium.img.Fill(color.RGBA{0, 161, 233, 255})
		medals[2].podium.img.Fill(color.RGBA{0, 161, 233, 255})
	}
}

func resetEditPlaceEffect() {
	for i, _ := range medals {
		medals[i].podium.img.Fill(color.White)
	}
}

func editText(target int) {
	l := medals[0].txtIndex
	c := medals[1].txtIndex
	r := medals[2].txtIndex
	switch target {
	case LEFT_CENTER:
		medals[0].txtIndex = c
		medals[1].txtIndex = l
	case RIGHT_CENTER:
		medals[1].txtIndex = r
		medals[2].txtIndex = c
	}
	for i, _ := range medals {
		medals[i].txt.txt = medalTxtData[medals[i].txtIndex]
		medals[i].filter.img = images[medalFilterData[medals[i].txtIndex]]
	}
}



func dataToColor(data []uint8) color.Color {
	return color.RGBA{data[0], data[1], data[2], data[3]}
}



func newMedals() {
	medals = nil
	fastest = 0

	base := [][]int{
		{0, 1, 2},//txt
		{0, 1, 2},//clr
		{0, 1, 2},//pos
	}
	for _, b := range base {
		rand.Shuffle(3, func(i, j int) {b[i], b[j] = b[j], b[i]})
	}

	for i:=0; i<3; i++ {
		// txt := newText(medalTxtData[i], 0, 300 - medalPlaceData[i], color.Black, color.Transparent, mainFont, 0, 0, false)
		t := newText(medalTxtData[base[0][i]], 0, 300 - medalPlaceData[base[2][i]], dataToColor(medalClrData[base[1][i]]), color.Transparent, mainFont, 0, 0, false)
		t.setCenter(SCREEN_WIDTH * (i+1) / 4)

		bg := newSprite(ebiten.NewImage(96, 96), t.spr.x, t.spr.y)
		bg.img.Fill(color.Black)

		podImg := ebiten.NewImage(132, 50 + int(medalPlaceData[base[2][i]]))
		pod := newSprite(podImg, centerX(podImg, SCREEN_WIDTH * (i + 1) / 4), 300 - medalPlaceData[base[2][i]] + 96)
		pod.img.Fill(color.White)

		fil := newSprite(images[medalFilterData[base[0][i]]], t.spr.x, t.spr.y)
		medals = append(medals, &medal{
			txt: t,
			bg: bg,
			podium: pod,
			filter: fil,
			txtIndex: base[0][i],
			clrIndex: base[1][i],
			placeIndex: base[2][i],
		})
	}
	for _, v := range sortData {
		if reflect.DeepEqual(base[0], v[:3]) {
			fastest += v[3]
			// println("txt",v[3])	
		}
		if reflect.DeepEqual(base[1], v[:3]) {
			fastest += v[3]
			// println("clr",v[3])
		}
		if reflect.DeepEqual(base[2], v[:3]) {
			fastest += v[3]
			// println("place",v[3])
		}
	}
	if fastest < 4 {
		newMedals()
	}
}

func (m *medal) draw(screen *ebiten.Image, lang int) {
	m.bg.draw(screen)
	m.txt.draw(screen)
	if lang == 0 {
		m.filter.draw(screen)
	}
	m.podium.draw(screen)
}