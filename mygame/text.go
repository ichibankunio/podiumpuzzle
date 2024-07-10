package mygame

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type txtObj struct {
	txt     string
	spr *sprite
	clr     color.Color
	padUp   int
	padLeft int
	font    font.Face
	isVert  bool
	hidden bool
}

var (
	timerText *txtObj
	timerDeltaText *txtObj

	infoText *txtObj
	cutinText *txtObj
	editCounterText *txtObj

	resultTimeTitle *txtObj
	resultTimeText *txtObj
	resultScoreTitle *txtObj
	resultScoreText *txtObj

	resultBonusText *txtObj

	resultRank *txtObj

	praiseText *txtObj

	recommendText *txtObj

	titleText []*txtObj

	recordTitle1 *txtObj
	recordTitle2 *txtObj
	recordTime1 *txtObj
	recordScore1 *txtObj
	recordTime2 *txtObj
	recordScore2 *txtObj
	recordBestText *txtObj
)

func initText() {

	timerText = newText("00.00", 0, 132, color.White, color.Transparent, smallFont, 0, 0, false)
	timerText.setCenter(SCREEN_WIDTH / 2)

	timerDeltaText = newText("+0.00", 0, 132, color.White, color.Transparent, smallFont, 0, 0, false)
	timerDeltaText.setCenter(SCREEN_WIDTH / 2 + 122)
	timerDeltaText.hidden = true

	infoText = newText("1/5 最短5手", 0, 180, color.White, color.Transparent, smallFont, 0, 0, false)
	infoText.setCenter(SCREEN_WIDTH / 2)

	cutinText = newText("1/5 最短5手", 0, 596, color.Black, color.Transparent, middleFont, 0, 0, false)
	cutinText.setCenter(SCREEN_WIDTH / 2 + 60)

	editCounterText = newText("1", 0, 530, color.RGBA{64, 64, 64, 255}, color.Transparent, bigFont, 0, 0, false)
	editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)
	editCounterText.hidden = true

	recordBestText = newText(lang[LANG_RECORD], 0, 150, color.White, color.Transparent, smallFont, 0, 0, false)
	recordBestText.setCenter(SCREEN_WIDTH / 2)

	resultTimeTitle = newText(lang[LANG_TIME], 0, 130, color.White, color.Transparent, middleFont, 0, 0, false)
	resultTimeTitle.setCenter(SCREEN_WIDTH / 2)
	resultTimeText = newText("00.00", 0, 210, color.White, color.Transparent, middleFont, 0, 0, false)
	resultTimeText.setCenter(SCREEN_WIDTH / 2)
	resultScoreTitle = newText(lang[LANG_SCORE], 0, 320, color.White, color.Transparent, middleFont, 0, 0, false)
	resultScoreTitle.setCenter(SCREEN_WIDTH / 2)
	resultScoreText = newText("0", 0, 400, color.White, color.Transparent, middleFont, 0, 0, false)
	resultScoreText.setCenter(SCREEN_WIDTH / 2)

	resultBonusText = newText(lang[LANG_BONUS], 0, 470, color.White, color.Transparent, smallFont, 0, 0, false)
	resultBonusText.setCenter(SCREEN_WIDTH / 2)
	resultBonusText.hidden = true

	resultRank = newText(lang[LANG_GOLD], 330, 400, color.White, color.Transparent, middleFont, 0, 0, false)
	resultRank.hidden = true

	praiseText = newText("EXCELLENT!", 0, 60, color.RGBA{255, 255, 0, 255}, color.Transparent, middleFont, 0, 0, false)
	praiseText.setCenter(SCREEN_WIDTH / 2)

	recommendText = newText(lang[LANG_APP], 0, 118, color.White, color.Transparent, smallFont, 0, 0, false)
	recommendText.setCenter(SCREEN_WIDTH / 2)

	titleText = []*txtObj{
		newText(lang[LANG_PO], 112, 320, dataToColor(medalClrData[0]), color.Transparent, mainFont, 0, 0, false),
		newText(lang[LANG_DI], 222, 300, dataToColor(medalClrData[1]), color.Transparent, mainFont, 0, 0, false),
		newText(lang[LANG_UM], SCREEN_WIDTH - 112 - 96, 340, dataToColor(medalClrData[2]), color.Transparent, mainFont, 0, 0, false),
		newText(lang[LANG_PU], 106, 320 + 110, color.RGBA{255, 0, 0, 255}, color.White, mainFont, 10, 4, false),
		newText(lang[LANG_ZZ], 222, 300 + 110, color.RGBA{255, 0, 0, 255}, color.White, mainFont, 20, 4, false),
		newText(lang[LANG_LE], SCREEN_WIDTH - 112 - 96, 340 + 110, color.RGBA{255, 0, 0, 255}, color.White, mainFont, 0, 4, false),
	}

	recordTitle1 = newText(lang[LANG_FIVE], 0, 240, color.White, color.Transparent, middleFont, 0, 0, false)
	recordTitle2 = newText(lang[LANG_TEN], 0, 490, color.White, color.Transparent, middleFont, 0, 0, false)
	recordTitle1.setCenter(SCREEN_WIDTH / 2)
	recordTitle2.setCenter(SCREEN_WIDTH / 2)

	recordTime1 = newText(lang[LANG_TIME], 0, 310, color.White, color.Transparent, smallFont, 0, 0, false)
	recordScore1 = newText(lang[LANG_TIME], 0, 370, color.White, color.Transparent, smallFont, 0, 0, false)
	recordTime1.setCenter(SCREEN_WIDTH / 2)
	recordScore1.setCenter(SCREEN_WIDTH / 2)

	recordTime2 = newText(lang[LANG_TIME], 0, 560, color.White, color.Transparent, smallFont, 0, 0, false)
	recordScore2 = newText(lang[LANG_TIME], 0, 620, color.White, color.Transparent, smallFont, 0, 0, false)
	recordTime2.setCenter(SCREEN_WIDTH / 2)
	recordScore2.setCenter(SCREEN_WIDTH / 2)
}



func newText(txt string, x, y float64, clr color.Color, bgClr color.Color, font font.Face, padUp, padLeft int, isVert bool) *txtObj {
	var bgImg *ebiten.Image
	if isVert {
		height := font.Metrics().Height.Ceil() * len([]rune(txt))
		bgImg = ebiten.NewImage(font.Metrics().Height.Ceil()+padLeft*2, height+padUp*2)
	}else {
		width := text.BoundString(font, txt).Dx()
		bgImg = ebiten.NewImage(width+padLeft*2, font.Metrics().Height.Ceil()+padUp*2)
	}

	bgImg.Fill(bgClr)	

	t := &txtObj{
		txt: txt,
		spr: newSprite(bgImg, x, y),
		clr: clr,
		font: font,
		padUp: 0,
		padLeft: 0,
		isVert: false,
	}
	return t
}

func (t *txtObj) setCenter(center int) {
	if t.isVert {
		*&t.spr.x = float64(center - t.font.Metrics().Height.Ceil()/2 - t.padLeft)
	} else {
		width := text.BoundString(t.font, t.txt).Dx()
		*&t.spr.x = float64(center - width/2 - t.padLeft)
	}
}

func (t *txtObj) setRight(right int) {
	if t.isVert {
		*&t.spr.x = float64(right - t.font.Metrics().Height.Ceil() - t.padLeft)
	} else {
		width := text.BoundString(t.font, t.txt).Dx()
		*&t.spr.x = float64(right - width - t.padLeft)
	}
}

func (t *txtObj) draw(screen *ebiten.Image) {
	if !t.hidden {
		if t.isVert {
			t.spr.draw(screen)
		
			for i, v := range []rune(t.txt) {
				if string(v) == "、" || string(v) == "。" {
					// int(t.spr.x)は危険かもしれない
					text.Draw(screen, string(v), t.font, int(t.spr.x)+t.padLeft+t.font.Metrics().Height.Ceil()-text.BoundString(t.font, string(v)).Dx(), int(t.spr.y)-t.font.Metrics().Height.Ceil()/8+t.padUp+i*t.font.Metrics().Height.Ceil()+t.font.Metrics().Height.Ceil()-text.BoundString(t.font, string(v)).Dy(), t.clr)
				} else {
					text.Draw(screen, string(v), t.font, int(t.spr.x)+t.padLeft, int(t.spr.y)-t.font.Metrics().Height.Ceil()/8+t.padUp+t.font.Metrics().Height.Ceil()+i*t.font.Metrics().Height.Ceil(), t.clr)
	
				}
			}
		} else {
			t.spr.draw(screen)
	
			text.Draw(screen, t.txt, t.font, int(t.spr.x)+t.padLeft, int(t.spr.y)-t.font.Metrics().Height.Ceil()/8+t.font.Metrics().Height.Ceil()+t.padUp, t.clr)
		}
	}
	
}
