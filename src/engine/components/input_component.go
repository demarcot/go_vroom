package components

import (
	eng "demarcot/vroom/src/engine"
	engMath "demarcot/vroom/src/engine/math"
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var playerSpeed = 5.0; // temporary, should either be on player or in a physics component

type InputManager struct {
  Keys [348]bool;
};

func (i *InputManager) SetKey (key glfw.Key, isPressed bool) {
  i.Keys[key] = isPressed;
}

type PlayerInputComponent struct {
  Manager *InputManager;
};

func (c PlayerInputComponent) HandleInput (p *eng.Entity, timeElapsedMs int) {
  movementVec := engMath.NewVec2Default();

  // Determine the direction the player wants
  if (c.Manager.Keys[glfw.KeyW]) {
    movementVec.Y++;
  }

  if (c.Manager.Keys[glfw.KeyS]) {
    movementVec.Y--;
  }

  if (c.Manager.Keys[glfw.KeyD]) {
    movementVec.X++;
  }

  if (c.Manager.Keys[glfw.KeyA]) {
    movementVec.X--;
  }

  if (movementVec.Mag() > 0) {
    // Normalize the chosen direction
    movementVec.Norm();

    // Scale the direction vector by speed and the amount of time that has elapsed
    interval := float64(timeElapsedMs) / 1000;
    movementVec.Scale(playerSpeed * interval);

    // Update player position
    p.Pos.MutAdd(movementVec);

    fmt.Printf("pos: %f, %f\n", p.Pos.X, p.Pos.Y)
  }
}
