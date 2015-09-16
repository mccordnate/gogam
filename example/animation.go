package main

import (
	"github.com/mccordnate/gogam"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	sproot "golang.org/x/mobile/exp/sprite"
	"image"
	_ "image/png"
	"log"
	"time"
)

var eng *gogam.Engine
var duck *gogam.Sprite

func main() {
	app.Main(func(a app.App) {
		var sz size.Event
		for e := range a.Events() {
			switch e := app.Filter(e).(type) {
			case size.Event:
				sz = e
			case paint.Event:
				onPaint(sz)
				a.EndPaint(e)
			case touch.Event:
				duck.MoveTo(e.X/sz.PixelsPerPt, e.Y/sz.PixelsPerPt)
			}
		}
	})
}

func onPaint(sz size.Event) {
	if eng == nil {
		load()
	}
	move()
	eng.Render(sz)
}

func load() {
	eng = gogam.NewEngine(1080, 1920)

	duck = gogam.NewSprite(20, 20)
	ass, err := asset.Open("duck.png")
	if err != nil {
		log.Fatal(err)
	}
	defer ass.Close()
	i, _, err := image.Decode(ass)
	if err != nil {
		log.Fatal(err)
	}
	t, err := eng.LoadTexture(i)
	if err != nil {
		log.Fatal(err)
	}
	sub2 := &sproot.SubTex{t, image.Rect(256, 0, 512, 288)}
	stillAf2 := gogam.NewAnimationFrame(sub2, int(time.Millisecond*300))
	sub3 := &sproot.SubTex{t, image.Rect(512, 0, 768, 288)}
	stillAf3 := gogam.NewAnimationFrame(sub3, int(time.Millisecond*300))
	a := gogam.NewAnimation([]*gogam.AnimationFrame{stillAf2, stillAf3})
	duck.AddAnimation("still", a)
	duck.SetAnimation("still")

	eng.Draw(duck)
}

func move() {
	duck.Rotate(1)
}
