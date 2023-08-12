package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"

	eng "demarcot/vroom/src/engine"
	comp "demarcot/vroom/src/engine/components"
	"demarcot/vroom/src/engine/math"
)

var inputManager comp.InputManager;

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	// Initialize GLFW
  err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

  // Load configs
  configSvc := eng.ConfigService{}
  configSvc.LoadConfig("./config.tcd")

  // Create/Configure Game Window
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

  // Initialize keyboard map
  inputManager = comp.InputManager{
    Keys: [348]bool{false},
  };

  // Create Player
  // Just Placeholder until I make a smarter entity manager
  player := eng.Entity{
    Pos: math.NewVec2Default(),
    Input: comp.PlayerInputComponent {
      Manager: &inputManager,
    },
  };

  fmt.Println("Press WASD to move. Press Escape to quit.");
  fmt.Println("Note: There is no rendering code yet, so the window will be empty.");

  // Game Loop
  lastTime := time.Now();
  var currentTime time.Time;
	for !window.ShouldClose() {
    // handle input
		glfw.PollEvents();

    currentTime = time.Now();
    // update

    player.Update(int(currentTime.Sub(lastTime).Milliseconds()));
    lastTime = currentTime;

    // render
		window.SwapBuffers();
    time.Sleep(10 * time.Millisecond); // Forcing a short sleep. Otherwise it's too fast and game time elapsed in ms is 0 which throws off calculations
	}
}

func keyCallback (w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
  // fmt.Printf("key %v, action %v\n", key, action)
  if (key == glfw.KeyEscape) {
    w.SetShouldClose(true);
  } else {
    inputManager.SetKey(key, action != glfw.Release);
  }
}
