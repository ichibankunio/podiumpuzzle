package mygame

import (
	// "fmt"
	// "fmt"
	// "fmt"
	"image/color"
	"strings"

	// "math"
	"math/rand"
	"strconv"

	// "strings"
	"time"

	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	SCREEN_WIDTH  = 540
	SCREEN_HEIGHT = 960

	SAMPLE_RATE = 44100
)

type Game struct{
	State int
	ButtonData string
	Lang int

	// Time1 int
	// Time2 int
	// Score1 int
	// Score2 int
	Record []string
	Result []string
}

// ebitenmobile bind -target android -javapkg com.ku20298.devutil -o ebiten.aar ./mobile
const (
	STATE_TYTLE = iota
	STATE_CUTIN
	STATE_MAIN
	STATE_CUTIN_2
	STATE_RESULT
	STATE_STOP
	STATE_RECORD
)

var counter int
var fastest = 0
var round = 1
var lastRound = 5
var editCounter int

var isSilent bool

var lang []string

// var startTime time.Time
// // var stopTime time.Duration
// var current time.Duration
// var formerTime = []int{0, 0}
// var continuousTime = []int{0, 0}
var frameTime = 1.0 / 60.0
var	validTime float64
var formerTime float64

var visibleScore int
var score int
var editCountMin int
var editCountSum int
var bonus int

var isFirst = true

func init() {
	rand.Seed(time.Now().UnixNano())
	lang = langJa

	initResource()
	initSprite()
	initText()
	initMedal()
}

func isRightColor() bool {
	return medals[0].txtIndex == medals[0].clrIndex && medals[1].txtIndex == medals[1].clrIndex && medals[2].txtIndex == medals[2].clrIndex
}

func isRightPlace() bool {
	return medals[0].txtIndex == medals[0].placeIndex && medals[1].txtIndex == medals[1].placeIndex && medals[2].txtIndex == medals[2].placeIndex
}

func isEditFinished() bool {
	for i, v := range medals {
		if !(v.clrIndex == i && v.txtIndex == i && v.placeIndex == i) {
			return false
		}
	}
	return true
}


func formatTime(t float64) string {
	s := strconv.FormatFloat(t, 'f', 2, 64)
	if int(t) < 10 {
		s = "0" + s
	}
	return s
}


func computeScore(t float64) int {
	n := float64(100000)
	if lastRound == 10 {
		n = 200000
	}
	return int(n / t)
}

func computeBonus() int {
	bonus := 2222
	if lastRound == 10 {
		bonus = 4444
	}
	if editCountSum > editCountMin {
		bonus = bonus / (editCountSum - editCountMin)
	}
	return bonus
}

func updateRecordText(r []string) {
	recordTime1.txt = "タイム " + r[0]
	recordScore1.txt = "総合スコア " + r[1]
	recordTime2.txt = "タイム " + r[2]
	recordScore2.txt = "総合スコア " + r[3]
	recordTime1.setCenter(SCREEN_WIDTH / 2)
	recordTime2.setCenter(SCREEN_WIDTH / 2)
	recordScore1.setCenter(SCREEN_WIDTH / 2)
	recordScore2.setCenter(SCREEN_WIDTH / 2)
}

 
func (g *Game) Update() error {
	// g.Lang = 1
	// if isFirst {
	// 	if g.Lang == 1 {
	// 		lang = langEn
	// 		medalTxtData = []string{"Si", "Go", "Br"}
	// 		initText()
	// 	}else {
	// 		lang = langJa
	// 	}

	// 	isFirst = false
	// }
	
	switch g.State {
	case STATE_TYTLE:
		if fiveButton.isJustTouched() {
			lastRound = 5
			g.State = STATE_CUTIN
			newMedals()
			cutinText.txt = strconv.Itoa(round) + "/" + strconv.Itoa(lastRound) + " " + lang[LANG_EXCELLENT] + strconv.Itoa(fastest) + lang[LANG_SHIFTS]
			infoText.txt = cutinText.txt

			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}
		if tenButton.isJustTouched() {
			lastRound = 10
			g.State = STATE_CUTIN
			newMedals()
			cutinText.txt = strconv.Itoa(round) + "/" + strconv.Itoa(lastRound) + " " + lang[LANG_EXCELLENT] + strconv.Itoa(fastest) + lang[LANG_SHIFTS]

			infoText.txt = cutinText.txt
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}
		if recordButton.isJustTouched() {
			g.State = STATE_RECORD
			totitleButton.x = centerX(totitleButton.img, SCREEN_WIDTH / 2)
			totitleButton.y = 720

			updateRecordText(g.Record)

			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}

		if soundIcon.isJustTouched() {
			if isSilent {
				for _, v := range se {
					v.SetVolume(0.7)
				}
				soundIcon.img = images[IMG_ALOUD]
				isSilent = false
				se[SE_PIRON].Rewind()
				se[SE_PIRON].Play()
			}else {
				for _, v := range se {
					v.SetVolume(0)
				}
				soundIcon.img = images[IMG_SILENT]
				isSilent = true
			}
		}
	case STATE_CUTIN:
		if counter < 6 {
			cutinText.spr.x -= 10
		}
		if counter >= 180 && counter < 186 {
			cutinText.spr.x -= 10
			cutinBg.alpha -= 0.1
		}

		if counter == 186 {
			cutinText.hidden = true
			g.State = STATE_MAIN
			// startTime = time.Now().Add(-current)
			counter = 0

			totitleButton.alpha = 1
			retryButton.alpha = 1
		}else {
			counter ++
		}

		if stopButton.isJustTouched() {
			stopButton.alpha = 0.6
			g.State = STATE_STOP
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}

	case STATE_MAIN:
		validTime += frameTime
		// if rand.Intn(30) == 0 {
		// 	validTime += 0.01
		// }
		if timerText.spr.x != 200 {
			if len(formatTime(validTime)) > 5 {
				timerText.spr.x = 200
			}
		}
		timerText.txt = formatTime(validTime)

		// updateTime()
		// timerText.txt = formatTime(continuousTime)

		for i, v := range editColorButtons {
			if v.isJustTouched() {
				editColor(i)
				editColorEffect(i)
				editCounter++
				editCounterText.txt = strconv.Itoa(editCounter)
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)
				se[SE_EDIT].Rewind()
				se[SE_EDIT].Play()
			}
			if v.afterLatestTouch == 3 {
				v.alpha = 0.6
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 - 3)
			}
			if v.afterLatestTouch == 6 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 + 3)
			}
			if v.afterLatestTouch == 10 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)

				v.alpha = 0.95
				resetEditColorEffect()
			}
			v.updateFrame()
		}
		for i, v := range editPlaceButtons {
			if v.isJustTouched() {
				editPlace(i)
				
				editPlaceEffect(i)
				editCounter++
				editCounterText.txt = strconv.Itoa(editCounter)
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)
				se[SE_EDIT].Rewind()
				se[SE_EDIT].Play()
			}
			if v.afterLatestTouch == 3 {
				v.alpha = 0.6
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 - 3)
			}
			if v.afterLatestTouch == 6 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 + 3)
			}
			if v.afterLatestTouch == 10 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)

				v.alpha = 0.95
				resetEditPlaceEffect()
			}
			v.updateFrame()

		}
		for i, v := range editTextButtons {
			if v.isJustTouched() {
				editText(i)
				editTextEffect(i)
				editCounter++
				editCounterText.txt = strconv.Itoa(editCounter)
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)
				se[SE_EDIT].Rewind()
				se[SE_EDIT].Play()

			}
			if v.afterLatestTouch == 3 {
				v.alpha = 0.6
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 - 3)
			}
			if v.afterLatestTouch == 6 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 + 3)
			}
			if v.afterLatestTouch == 10 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)
				v.alpha = 0.95
				resetEditTextEffect()
			}
			v.updateFrame()
		}

		if isEditFinished() {
			g.State = STATE_CUTIN_2
			cutinBg.alpha = 0.95
			if editCounter == fastest {
				se[SE_GREAT].Rewind()
				se[SE_GREAT].Play()
			}else {
				se[SE_SUCCESS].Rewind()
				se[SE_SUCCESS].Play()
			}
			counter = 0
			timerDeltaText.hidden = false
			
			timerDeltaText.txt = "+" + formatTime(validTime - formerTime)

			formerTime = validTime

			if editCounter == fastest {
				praiseText.txt = "EXCELLENT!"
				praiseText.setCenter(SCREEN_WIDTH / 2)
				praiseText.clr = color.RGBA{255, 255, 0, 255}
	
			}else {	
				praiseText.txt = "GOOD!"
				praiseText.setCenter(SCREEN_WIDTH / 2)
				praiseText.clr = color.White
			}
			
		}
		
		if stopButton.isJustTouched() {
			stopButton.alpha = 0.6
			g.State = STATE_STOP
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}

		if editCounter == 0 {
			editCounterText.hidden = true
		}else {
			editCounterText.hidden = false
		}
	case STATE_CUTIN_2:
		vy := 3.0
		if editCounter == fastest {
			
			vy = 16.0
			if counter == 15 {
				editCounterText.clr = color.White
			}
			if counter == 30 {
				editCounterText.clr = color.RGBA{64, 64, 64, 255}
			}
			if counter == 45 {
				editCounterText.clr = color.White
			}
			if counter == 60 {
				editCounterText.clr = color.RGBA{64, 64, 64, 255}
			}
		}
		if counter < 4 {
			medals[0].txt.spr.y -= vy
			medals[0].bg.y -= vy
			medals[0].filter.y -= vy
			medals[0].podium.y -= vy
		}else if counter < 8 {
			medals[0].txt.spr.y += vy
			medals[0].bg.y += vy
			medals[0].filter.y += vy
			medals[0].podium.y += vy
			medals[1].txt.spr.y -= vy
			medals[1].bg.y -= vy
			medals[1].filter.y -= vy
			medals[1].podium.y -= vy
		}else if counter < 12 {
			medals[1].txt.spr.y += vy
			medals[1].bg.y += vy
			medals[1].filter.y += vy
			medals[1].podium.y += vy
			medals[2].txt.spr.y -= vy
			medals[2].bg.y -= vy
			medals[2].filter.y -= vy
			medals[2].podium.y -= vy
		}else if counter < 16 {
			medals[2].txt.spr.y += vy
			medals[2].bg.y += vy
			medals[2].filter.y += vy
			medals[2].podium.y += vy
		}
		
		if counter % 20 == 0 {
			timerDeltaText.hidden = false
		}else if counter % 10 == 0 {
			timerDeltaText.hidden = true
		}
		
		if counter == 150 {
			if round == lastRound {
				g.State = STATE_RESULT
				resultTimeText.txt = timerText.txt
				score = computeScore(validTime)
				bonus = computeBonus()
				score += bonus

				
				t := strings.Split(resultTimeText.txt, ".")
				s, _ := strconv.Atoi(t[0])
				ms, _ := strconv.Atoi(t[1])
			
				rt1 := strings.Split(g.Record[0], ".")
				rs1, _ := strconv.Atoi(rt1[0])
				rms1, _ := strconv.Atoi(rt1[1])
			
				rt2 := strings.Split(g.Record[2], ".")
				rs2, _ := strconv.Atoi(rt2[0])
				rms2, _ := strconv.Atoi(rt2[1])
			
				switch lastRound {
				case 5:
					if (rs1 == 0 && rms1 == 0) || s < rs1 || (s == rs1 && ms < rms1) {
						g.Record[0] = resultTimeText.txt
					}

					rscore, _ := strconv.Atoi(g.Record[1])
					if score > rscore {
						g.Record[1] = strconv.Itoa(score)
						println(score)
					}
				case 10:
					if (rs2 == 0 && rms2 == 0) || s < rs2 || (s == rs2 && ms < rms2) {
						g.Record[2] = resultTimeText.txt
					}

					rscore, _ := strconv.Atoi(g.Record[3])
					if score > rscore {
						g.Record[3] = strconv.Itoa(score)
					}
				}



				resultTimeText.setCenter(SCREEN_WIDTH / 2)
				resultScoreText.setCenter(SCREEN_WIDTH / 2)
				resultBonusText.txt += strconv.Itoa(bonus)
				resultBonusText.setCenter(SCREEN_WIDTH / 2)
				switch  {
				case score > 10000:
					resultRank.txt = medalTxtData[1]
					resultRank.clr = dataToColor(medalClrData[1])
				case score > 7000:
					resultRank.txt = medalTxtData[0]
					resultRank.clr = dataToColor(medalClrData[0])
				default:
					resultRank.txt = medalTxtData[2]
					resultRank.clr = dataToColor(medalClrData[2])
				}
				validTime = 0
				formerTime = 0
				timerText.txt = formatTime(validTime)
				timerText.setCenter(SCREEN_WIDTH / 2)


				g.Result[0] = resultTimeText.txt
				g.Result[1] = strconv.Itoa(score)
				g.Result[2] = resultRank.txt
				g.Result[3] = strconv.Itoa(lastRound)

			
				se[SE_SCORE_RISE].Rewind()
				se[SE_SCORE_RISE].Play()

			}else {
				round ++
				g.State = STATE_CUTIN
				editCountMin += fastest
				editCountSum += editCounter
			}
			
			newMedals()
			cutinText.txt = strconv.Itoa(round) + "/" + strconv.Itoa(lastRound) + " " + lang[LANG_EXCELLENT] + strconv.Itoa(fastest) + lang[LANG_SHIFTS]

			infoText.txt = cutinText.txt
			editCounter = 0
			editCounterText.txt = strconv.Itoa(editCounter)
			editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)
			cutinText.setCenter(SCREEN_WIDTH / 2 + 60)
			counter = 0
			cutinText.hidden = false
			cutinBg.alpha = 0.95
			timerDeltaText.hidden = true
		}else {
			counter ++
		}

		for _, v := range editColorButtons {
			if v.afterLatestTouch == 3 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 - 3)
			}
			if v.afterLatestTouch == 6 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 + 3)
			}
			if v.afterLatestTouch == 10 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)

				v.alpha = 1
				resetEditColorEffect()
			}
			v.updateFrame()
		}
		for _, v := range editPlaceButtons {
			if v.afterLatestTouch == 3 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 - 3)
			}
			if v.afterLatestTouch == 6 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 + 3)
			}
			if v.afterLatestTouch == 10 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)

				v.alpha = 1
				resetEditPlaceEffect()
			}
			v.updateFrame()
		}
		for _, v := range editTextButtons {
			if v.afterLatestTouch == 3 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 - 3)
			}
			if v.afterLatestTouch == 6 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10 + 3)
			}
			if v.afterLatestTouch == 10 {
				editCounterText.setCenter(SCREEN_WIDTH / 2 - 10)
				v.alpha = 1
				resetEditTextEffect()
			}
			v.updateFrame()
		}
	case STATE_RESULT:
		if visibleScore < score - bonus {
			visibleScore += 50
			resultScoreText.txt = strconv.Itoa(visibleScore - visibleScore % 100 + (counter + 1) % 10 * 10 + counter % 10)
		}

		if visibleScore % 1800 == 0 {
			if visibleScore < score - bonus {
				se[SE_SCORE_RISE].Rewind()
				se[SE_SCORE_RISE].Play()
			}
		}
		if visibleScore >= score - bonus {
			if se[SE_SCORE_RISE].IsPlaying() {
				resultScoreText.txt = strconv.Itoa(score - bonus)
			}else {
				// if visibleScore == score - bonus  {
				
				if visibleScore < score {
					if !se[SE_SCORE_RISE_2].IsPlaying() {
						se[SE_SCORE_RISE_2].Rewind()
						se[SE_SCORE_RISE_2].Play()
						resultBonusText.hidden = false
						counter = 0
					}
					visibleScore += 50
					resultScoreText.txt = strconv.Itoa(visibleScore - visibleScore % 100 + (counter + 1) % 10 * 10 + counter % 10)
				}
				if visibleScore >= score {
					resultScoreText.txt = strconv.Itoa(score)
					if counter == 110 && resultRank.hidden {
						resultRank.hidden = false
						resultScoreText.spr.x -= 25
						se[SE_METAL].Rewind()
						se[SE_METAL].Play()
					}
				}
			}
		}

		if visibleScore == 100 {
			resultScoreText.setCenter(SCREEN_WIDTH / 2)
		}
		if visibleScore == 1000 {
			resultScoreText.setCenter(SCREEN_WIDTH / 2)
		}
		if visibleScore == 10000 {
			resultScoreText.setCenter(SCREEN_WIDTH / 2)
		}

		counter ++
		
		if totitleButton.isJustTouched() {
			g.State = STATE_TYTLE
			round = 1
			score = 0
			visibleScore = 0
			counter = 0
			editCountMin = 0
			editCountSum = 0
			resultBonusText.hidden = true
			resultBonusText.txt = lang[LANG_BONUS]
			resultRank.hidden = true
			se[SE_SCORE_RISE].Pause()
			se[SE_SCORE_RISE_2].Pause()
			se[SE_METAL].Pause()
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()

		}

		if retryButton.isJustTouched() {
			g.State = STATE_CUTIN
			newMedals()
			round = 1
			cutinText.txt = strconv.Itoa(round) + "/" + strconv.Itoa(lastRound) + " " + lang[LANG_EXCELLENT] + strconv.Itoa(fastest) + lang[LANG_SHIFTS]

			infoText.txt = cutinText.txt
			score = 0
			visibleScore = 0
			counter = 0
			editCountMin = 0
			editCountSum = 0
			resultBonusText.hidden = true
			resultBonusText.txt = lang[LANG_BONUS]
			resultRank.hidden = true
			se[SE_SCORE_RISE].Pause()
			se[SE_SCORE_RISE_2].Pause()
			se[SE_METAL].Pause()

			cutinBg.alpha = 0.95

			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}
		if shareButton.isJustTouched() {
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}
	case STATE_STOP:
		if totitleButton.isJustTouched() {
			g.State = STATE_TYTLE
			round = 1
			score = 0
			visibleScore = 0
			counter = 0
			editCountMin = 0
			editCountSum = 0

			stopButton.alpha = 1

			cutinText.setCenter(SCREEN_WIDTH / 2 + 60)
			cutinBg.alpha = 0.95
			cutinText.hidden = false

			validTime = 0
			formerTime = 0
			timerText.txt = formatTime(validTime)
			timerText.setCenter(SCREEN_WIDTH / 2)
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}

		if retryButton.isJustTouched() {
			g.State = STATE_CUTIN
			newMedals()
			round = 1
			cutinText.txt = strconv.Itoa(round) + "/" + strconv.Itoa(lastRound) + " " + lang[LANG_EXCELLENT] + strconv.Itoa(fastest) + lang[LANG_SHIFTS]

			infoText.txt = cutinText.txt
			score = 0
			visibleScore = 0
			counter = 0
			editCountMin = 0
			editCountSum = 0

			stopButton.alpha = 1
			cutinText.setCenter(SCREEN_WIDTH / 2 + 60)
			counter = 0
			cutinBg.alpha = 0.95
			cutinText.hidden = false

			validTime = 0
			formerTime = 0
			timerText.txt = formatTime(validTime)
			timerText.setCenter(SCREEN_WIDTH / 2)
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}

		if restartButton.isJustTouched() {
			g.State = STATE_MAIN
			stopButton.alpha = 1
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}
	case STATE_RECORD:
		if totitleButton.isJustTouched() {
			g.State = STATE_TYTLE
			totitleButton.x = SCREEN_WIDTH - 90 - 150
			totitleButton.y = 550
			se[SE_PIRON].Rewind()
			se[SE_PIRON].Play()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))

	switch g.State {
	case STATE_TYTLE:
		fiveButton.draw(screen)
		tenButton.draw(screen)
		recordButton.draw(screen)
		soundIcon.draw(screen)
		sample.draw(screen)
		instruction.draw(screen)
		recommendText.draw(screen)
		sevenraceIcon.draw(screen)
		rouletteIcon.draw(screen)
		parsleyIcon.draw(screen)
		for _, v := range titleText {
			v.draw(screen)
		}
	case STATE_CUTIN:
		for _, medal := range medals {
			medal.draw(screen, g.Lang)
		}

		for _, v := range editColorButtons {
			v.draw(screen)
		}
		for _, v := range editPlaceButtons {
			v.draw(screen)
		}
		for _, v := range editTextButtons {
			v.draw(screen)
		}
		sample.draw(screen)
		instruction.draw(screen)

		timerText.draw(screen)
		infoText.draw(screen)

		cutinBg.draw(screen)
		cutinText.draw(screen)
		stopButton.draw(screen)

	case STATE_MAIN:
		for _, medal := range medals {
			medal.draw(screen, g.Lang)
		}

		editCounterText.draw(screen)

		for _, v := range editColorButtons {
			v.draw(screen)
		}
		for _, v := range editPlaceButtons {
			v.draw(screen)
		}
		for _, v := range editTextButtons {
			v.draw(screen)
		}
		sample.draw(screen)
		instruction.draw(screen)

		timerText.draw(screen)
		infoText.draw(screen)
		stopButton.draw(screen)

	case STATE_CUTIN_2:
		for _, medal := range medals {
			medal.draw(screen, g.Lang)
		}

		editCounterText.draw(screen)

		for _, v := range editColorButtons {
			v.draw(screen)
		}
		for _, v := range editPlaceButtons {
			v.draw(screen)
		}
		for _, v := range editTextButtons {
			v.draw(screen)
		}
		sample.draw(screen)
		instruction.draw(screen)

		timerText.draw(screen)
		timerDeltaText.draw(screen)
		infoText.draw(screen)

		praiseText.draw(screen)
		// cutinBg.draw(screen)
		// cutinText.draw(screen)

	case STATE_RESULT:
		resultTimeTitle.draw(screen)
		resultTimeText.draw(screen)
		resultScoreTitle.draw(screen)
		resultScoreText.draw(screen)
		resultBonusText.draw(screen)
		resultRank.draw(screen)

		sample.draw(screen)
		instruction.draw(screen)

		shareButton.draw(screen)
		totitleButton.draw(screen)
		retryButton.draw(screen)
		// cutinBg.draw(screen)
	case STATE_STOP:
		timerText.draw(screen)
		infoText.draw(screen)

		restartButton.draw(screen)
		totitleButton.draw(screen)
		retryButton.draw(screen)
		sample.draw(screen)
		instruction.draw(screen)
	case STATE_RECORD:
		recordBestText.draw(screen)
		totitleButton.draw(screen)
		recordTitle1.draw(screen)
		recordTitle2.draw(screen)
		
		recordTime1.draw(screen)
		recordScore1.draw(screen)
		recordTime2.draw(screen)
		recordScore2.draw(screen)
		sample.draw(screen)
		instruction.draw(screen)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

// qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM!01234567895連続金銀銅.+0.001/5 最短4手タップでタイマーがスタート最高記録タイム総合スコア手数ボーナスEXCELLENT!他のおすすめアプリ表彰台パズル5回/GOOD!

