package gogam

import (
	"time"

	"golang.org/x/image/math/f32"
	spr "golang.org/x/mobile/exp/sprite"
)

type Sprite struct {
	X             float32
	Y             float32
	ScaleX        float32
	ScaleY        float32
	Rotation      float32
	Velocity      f32.Vec2
	AnimStartTime time.Time
	Animations    map[string]*Animation
	CurrentAnim   string
	Node          *spr.Node
}

func NewSprite(x, y float32) *Sprite {
	s := new(Sprite)
	s.Node = &spr.Node{}
	s.X = x
	s.Y = y
	s.ScaleX = 100
	s.ScaleY = 100
	s.Rotation = 0
	s.Animations = make(map[string]*Animation)
	return s
}

func (s *Sprite) AddAnimation(name string, a *Animation) {
	s.Animations[name] = a
}

func (s *Sprite) SetAnimation(name string) {
	_, ok := s.Animations[name]
	if ok {
		s.CurrentAnim = name
		s.AnimStartTime = time.Now()
	}
}

func (s *Sprite) GetCurrentAnimation() *Animation {
	if s.CurrentAnim == "" {
		return nil
	} else {
		return s.Animations[s.CurrentAnim]
	}
}

func (s *Sprite) MoveTo(x, y float32) {
	s.X = x
	s.Y = y
}

func (s *Sprite) Translate(x, y float32) {
	s.X = s.X + x
	s.Y = s.Y + y
}

func (s *Sprite) Scale(x, y float32) {
	s.ScaleX = s.ScaleX * x
	s.ScaleY = s.ScaleY * y
}

func (s *Sprite) Rotate(d float32) {
	s.Rotation += d
	for s.Rotation < 0 {
		s.Rotation += 360
	}
	for s.Rotation > 360 {
		s.Rotation -= 360
	}
}

func (s *Sprite) SetVelocity(v f32.Vec2) {
	s.Velocity = v
}

func (s *Sprite) AddVelocity(v f32.Vec2) {
	s.Velocity[0] += v[0]
	s.Velocity[1] += v[1]
}

type Animation struct {
	Frames   []*AnimationFrame
	Duration time.Duration
}

func (s *Sprite) GetCurrentFrame() *AnimationFrame {
	a := s.GetCurrentAnimation()
	if a == nil {
		return nil
	}

	timeSinceStart := time.Since(s.AnimStartTime)

	curAnimTime := timeSinceStart % a.Duration
	animTime := time.Duration(0)
	for i, af := range a.Frames {
		animTime = animTime + af.Duration
		if curAnimTime < animTime {
			return a.Frames[i]
		}
	}

	return nil
}

func NewAnimation(afs []*AnimationFrame) *Animation {
	a := new(Animation)
	a.Frames = afs
	a.Duration = 0
	for _, af := range afs {
		a.Duration = a.Duration + af.Duration
	}
	return a
}

type AnimationFrame struct {
	Texture  *spr.SubTex
	Duration time.Duration
}

func NewAnimationFrame(t *spr.SubTex, duration int) *AnimationFrame {
	af := new(AnimationFrame)
	af.Texture = t
	af.Duration = time.Duration(duration)
	return af
}
