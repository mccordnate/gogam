package main

import (
	"github.com/mccordnate/gogam"
	"golang.org/x/image/math/f32"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"image"
	_ "image/png"
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
	// Initialize the engine with a working resolution of 1920x1080
	eng = gogam.NewEngine(1920, 1080)

	// Initialize a cat sprite at position (20,20)
	cat = gogam.NewSprite(20, 20)

	// Grab the cat texture
	t := gogam.NewTexture("cat.png", eng)

	// Get the cat's frame from the texture
	stillAf1 := gogam.NewAnimationFrame(t, image.Rect(0, 0, 139, 119), 1000)

	// Create an animation from the frame
	a := gogam.NewAnimation([]*gogam.AnimationFrame{stillAf1})

	// Add animation as a possible animation for the cat called "still"
	cat.AddAnimation("still", a)

	// Set the "still" animation as currently active animation
	cat.SetAnimation("still")

	// Tell the engine to draw the cat
	eng.Draw(cat)
}

// Move based on key press
// Creates smooth movement by changing the velocity of the cat
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
