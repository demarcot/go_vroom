package components

import (
	eng "demarcot/vroom/src/engine"
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type InputManager struct {
  Keys [348]bool;
};

func (i *InputManager) SetKey (key glfw.Key, isPressed bool) {
  i.Keys[key] = isPressed;
}

type PlayerInputComponent struct {
  Manager *InputManager;
};

func (c PlayerInputComponent) HandleInput (p *eng.Entity) {
  if (c.Manager.Keys[glfw.KeyJ]) {
    p.X++;
    p.Y++;
    fmt.Printf("Player input handler pos: %d, %d\n", p.X, p.Y);
  }
}
