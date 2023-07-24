package main

import (
	"fmt"
	"runtime"
	"github.com/go-gl/glfw/v3.3/glfw"

  eng "demarcot/vroom/src/engine"
  comp "demarcot/vroom/src/engine/components"
)

var inputManager comp.InputManager;

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

  configSvc := eng.ConfigService{}
  configSvc.LoadConfig("./config.tcd")

  inputManager = comp.InputManager{
    Keys: [348]bool{false},
  };

  inputComponent := comp.PlayerInputComponent{
    Manager: &inputManager,
  };

  // create entity manager
  player := eng.Entity{
    X: 0,
    Y: 0,
    Input: inputComponent,
  };

	window, err := glfw.CreateWindow(configSvc.GetIntVal("WINDOW_WIDTH", 50),
    configSvc.GetIntVal("WINDOW_HEIGHT", 50),
    configSvc.GetStrVal("WINDOW_TITLE", "default title"),
    nil,
    nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
  window.SetKeyCallback(keyCallback)

	for !window.ShouldClose() {
    // handle input
		glfw.PollEvents()

    // update
    player.Update()

    // render
		window.SwapBuffers()
	}
}

func keyCallback (w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
  fmt.Printf("key %v, action %v\n", key, action)
  if (key == glfw.KeyEscape) {
    w.SetShouldClose(true);
  } else {
    inputManager.SetKey(key, action != glfw.Release);
  }
}
