package gogam

import (
	"image"
	"math"
	"time"

	//"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/f32"
	s "golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

type Engine struct {
	screenWidth  int
	screenHeight int
	startTime    time.Time
	lastUpdate   clock.Time
	eng          s.Engine
	scene        *s.Node
	arranger     *arrangerFunc
	sprites      map[*s.Node]*Sprite
}

func NewEngine(screenWidth, screenHeight int) *Engine {
	e := new(Engine)
	e.screenWidth = screenWidth
	e.screenHeight = screenHeight
	e.startTime = time.Now()
	e.eng = glsprite.Engine()
	e.scene = &s.Node{}

	e.eng.Register(e.scene)
	e.eng.SetTransform(e.scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})
	e.arranger = new(arrangerFunc)
	e.arranger.eng = e
	e.sprites = make(map[*s.Node]*Sprite)
	return e
}

func (e *Engine) LoadTexture(a image.Image) (s.Texture, error) {
	return e.eng.LoadTexture(a)
}

func (e *Engine) Draw(spr *Sprite) {
	e.sprites[spr.Node] = spr
	spr.Node.Arranger = e.arranger
	e.eng.Register(spr.Node)
	e.scene.AppendChild(spr.Node)
}

func (e *Engine) Render(sz size.Event) {
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	now := clock.Time(time.Since(e.startTime) * 60 / time.Second)
	e.arranger.sz = &sz
	e.eng.Render(e.scene, now, sz)
}

type arrangerFunc struct {
	eng *Engine
	sz  *size.Event
}

func (a *arrangerFunc) Arrange(e s.Engine, n *s.Node, t clock.Time) {
	sprite, _ := a.eng.sprites[n]
	frameTime := float32(t - a.eng.lastUpdate)
	updatePosition(sprite, frameTime)

	screenWidthScaler := float32(a.sz.WidthPx) / float32(a.eng.screenWidth)
	screenHeightScaler := float32(a.sz.HeightPx) / float32(a.eng.screenHeight)
	actualScaleX := screenWidthScaler * sprite.ScaleX
	actualScaleY := screenHeightScaler * sprite.ScaleY
	actualPositionX := screenWidthScaler * sprite.X
	actualPositionY := screenHeightScaler * sprite.Y

	e.SetSubTex(n, *sprite.GetCurrentFrame().Texture)

	r := sprite.Rotation * math.Pi / 180
	matrix := f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	}

	matrix.Translate(&matrix, actualPositionX, actualPositionY)
	matrix.Rotate(&matrix, r)
	matrix.Scale(&matrix, actualScaleX, actualScaleY)
	e.SetTransform(n, matrix)

	a.eng.lastUpdate = t
}

func updatePosition(sprite *Sprite, frameTime float32) {
	sprite.X += frameTime * sprite.Velocity[0]
	sprite.Y += frameTime * sprite.Velocity[1]
}
