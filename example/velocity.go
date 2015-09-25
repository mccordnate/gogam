package main

import (
	"github.com/mccordnate/gogam"
	"golang.org/x/image/math/f32"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	sproot "golang.org/x/mobile/exp/sprite"
	"image"
	_ "image/png"
	"log"
)

var eng *gogam.Engine
var cat *gogam.Sprite

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
			case key.Event:
				move(e)
			}
		}
	})
}

func onPaint(sz size.Event) {
	if eng == nil {
		load()
	}
	eng.Render(sz)
}

func load() {
	eng = gogam.NewEngine(1920, 1080)

	cat = gogam.NewSprite(20, 20)
	ass, err := asset.Open("cat.png")
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
	sub := &sproot.SubTex{t, image.Rect(0, 0, 139, 119)}
	stillAf1 := gogam.NewAnimationFrame(sub, 1000)
	a := gogam.NewAnimation([]*gogam.AnimationFrame{stillAf1})
	cat.AddAnimation("still", a)
	cat.SetAnimation("still")

	eng.Draw(cat)
}

func move(k key.Event) {
	if k.Code == key.CodeLeftArrow || k.Code == key.CodeA {
		if k.Direction == key.DirPress {
			cat.AddVelocity(f32.Vec2{-5, 0})
		} else if k.Direction == key.DirRelease {
			cat.AddVelocity(f32.Vec2{5, 0})
		}
	}
	if k.Code == key.CodeRightArrow || k.Code == key.CodeD {
		if k.Direction == key.DirPress {
			cat.AddVelocity(f32.Vec2{5, 0})
		} else if k.Direction == key.DirRelease {
			cat.AddVelocity(f32.Vec2{-5, 0})
		}
	}
	if k.Code == key.CodeUpArrow || k.Code == key.CodeW {
		if k.Direction == key.DirPress {
			cat.AddVelocity(f32.Vec2{0, -5})
		} else if k.Direction == key.DirRelease {
			cat.AddVelocity(f32.Vec2{0, 5})
		}
	}
	if k.Code == key.CodeDownArrow || k.Code == key.CodeS {
		if k.Direction == key.DirPress {
			cat.AddVelocity(f32.Vec2{0, 5})
		} else if k.Direction == key.DirRelease {
			cat.AddVelocity(f32.Vec2{0, -5})
		}
	}
}
