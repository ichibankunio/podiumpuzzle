package mygame

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type sprite struct {
	img   *ebiten.Image
	x     float64
	y     float64
	alpha float64
	touchID ebiten.TouchID
	afterLatestTouch int
}

const (
	IMG_ALOUD = iota
	IMG_BRONZE
	IMG_COLOR_OK
	IMG_EDIT_COLOR
	IMG_EDIT_PLACE
	IMG_EDIT_TEXT
	IMG_FILTER_BRONZE
	IMG_FILTER_GOLD
	IMG_FILTER_SILVER
	IMG_FIVE
	IMG_GOLD
	IMG_INSTRUCTION
	IMG_PARSLEY
	IMG_PODIUM_OK
	IMG_POSITION_OK
	IMG_RECORD
	IMG_RESTART
	IMG_RETRY
	IMG_ROULETTE
	IMG_SAMPLE
	IMG_SEVENRACE
	IMG_SHARE
	IMG_SILENT
	IMG_SILVER
	IMG_STOP
	IMG_TEN
	IMG_TOTITLE
)

var (
	editColorButtons []*sprite
	editTextButtons []*sprite
	editPlaceButtons []*sprite
	
	sample *sprite
	instruction *sprite

	cutinBg *sprite

	shareButton *sprite
	totitleButton *sprite
	retryButton *sprite
	stopButton *sprite
	restartButton *sprite
	recordButton *sprite

	fiveButton *sprite
	tenButton *sprite

	soundIcon *sprite

	sevenraceIcon *sprite
	rouletteIcon *sprite
	parsleyIcon *sprite


)

func initSprite() {
	editTextButtons = []*sprite{
		newSprite(images[IMG_EDIT_TEXT], 100, 500),
		newSprite(images[IMG_EDIT_TEXT], SCREEN_WIDTH - 220, 500),
	}

	editColorButtons = []*sprite{
		newSprite(images[IMG_EDIT_COLOR], 100, 620),
		newSprite(images[IMG_EDIT_COLOR], SCREEN_WIDTH - 220, 620),
	}

	editPlaceButtons = []*sprite{
		newSprite(images[IMG_EDIT_PLACE], 100, 740),
		newSprite(images[IMG_EDIT_PLACE], SCREEN_WIDTH - 220, 740),
	}

	sample = newSprite(images[IMG_SAMPLE], centerX(images[IMG_SAMPLE], SCREEN_WIDTH / 2), 850)
	instruction = newSprite(images[IMG_INSTRUCTION], 10, 810)

	cutinBg = newSprite(ebiten.NewImage(SCREEN_WIDTH, 200), 0, 520)
	cutinBg.img.Fill(color.RGBA{127, 255, 212, 255})
	cutinBg.alpha = 0.95

	shareButton = newSprite(images[IMG_SHARE], 90, 550)
	totitleButton = newSprite(images[IMG_TOTITLE], SCREEN_WIDTH - 90 - 150, 550)
	retryButton = newSprite(images[IMG_RETRY], centerX(images[IMG_RETRY], SCREEN_WIDTH / 2), 690)
	stopButton = newSprite(images[IMG_STOP], 10, 50)
	restartButton = newSprite(images[IMG_RESTART], 90, 550)

	fiveButton = newSprite(images[IMG_FIVE], 90, 600)
	tenButton = newSprite(images[IMG_TEN], SCREEN_WIDTH - 90 - 150, 600)
	recordButton = newSprite(images[IMG_RECORD], centerX(images[IMG_RECORD], SCREEN_WIDTH / 2), 720)

	soundIcon = newSprite(images[IMG_ALOUD], 50, 210)

	sevenraceIcon = newSprite(images[IMG_SEVENRACE], 100, 24)
	rouletteIcon = newSprite(images[IMG_ROULETTE], 228, 24)
	parsleyIcon = newSprite(images[IMG_PARSLEY], SCREEN_WIDTH - 100 - 84, 24)
}

func (s *sprite) isJustTouched3() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.x) && x <= int(s.x)+s.img.Bounds().Dx() && y >= int(s.y) && y <= int(s.y)+s.img.Bounds().Dy() {
			return true
		}
	}

	t := inpututil.JustPressedTouchIDs()
	if len(t) > 0 {
		x, y := ebiten.TouchPosition(t[0])
		if x >= int(s.x) && x <= int(s.x)+s.img.Bounds().Dx() && y >= int(s.y) && y <= int(s.y)+s.img.Bounds().Dy() {
			return true
		}
	}

	return false
}
 
func (s *sprite) isJustTouched() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.x) && x <= int(s.x)+s.img.Bounds().Dx() && y >= int(s.y) && y <= int(s.y)+s.img.Bounds().Dy() {
			return true
		}
	}

	touch := inpututil.JustPressedTouchIDs()
	if len(touch) > 0 {
		for _, t := range touch {
			x, y := ebiten.TouchPosition(t)
			if x >= int(s.x) && x <= int(s.x)+s.img.Bounds().Dx() && y >= int(s.y) && y <= int(s.y)+s.img.Bounds().Dy() {
				s.touchID = t
				s.afterLatestTouch = 0
				return true
			}
		}
		
	}

	return false
}

func (s *sprite) isJustReleased() bool {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.x) && x <= int(s.x)+s.img.Bounds().Dx() && y >= int(s.y) && y <= int(s.y)+s.img.Bounds().Dy() {
			return true
		}
	}

	if inpututil.IsTouchJustReleased(s.touchID) {
		return true
	}

	return false
}

func (s *sprite) isTouched() bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || ebiten.IsKeyPressed(ebiten.KeySpace) {
		x, y := ebiten.CursorPosition()

		if x >= int(s.x) && x <= int(s.x)+s.img.Bounds().Dx() && y >= int(s.y) && y <= int(s.y)+s.img.Bounds().Dy() {
			return true
		}
	}

	t := ebiten.TouchIDs()
	if len(t) > 0 {
		x, y := ebiten.TouchPosition(t[0])
		if x >= int(s.x) && x <= int(s.x)+s.img.Bounds().Dx() && y >= int(s.y) && y <= int(s.y)+s.img.Bounds().Dy() {
			return true
		}
	}

	return false
}

func (s *sprite) updateFrame() {
	s.afterLatestTouch ++
}

func (s *sprite) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	op.ColorM.Scale(1, 1, 1, s.alpha)
	screen.DrawImage(s.img, op)
}

func centerX(img *ebiten.Image, center int) float64 {
	return float64(center - img.Bounds().Dx()/2)
}

func newSprite(img *ebiten.Image, x, y float64) *sprite {
	s := &sprite{
		img:   img,
		x:     x,
		y:     y,
		alpha: 1,
		afterLatestTouch: 10000,
	}

	return s
}
