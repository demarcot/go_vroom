package main

import (
	"fmt"
	"runtime"
  "unsafe"
	"time"

  "github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	eng "demarcot/vroom/src/engine"
	comp "demarcot/vroom/src/engine/components"
	"demarcot/vroom/src/engine/math"
)

var inputManager comp.InputManager;

var WIDTH int;
var HEIGHT int;

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
  fmt.Println("Starting...");

	// Initialize GLFW
  err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

  // Load configs
  configSvc := eng.ConfigService{}
  configSvc.LoadConfig("./config.tcd")

  WIDTH = configSvc.GetIntVal("WINDOW_WIDTH", 50);
  HEIGHT = configSvc.GetIntVal("WINDOW_HEIGHT", 50);

  // Create/Configure Game Window
  glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(WIDTH,
    HEIGHT,
    configSvc.GetStrVal("WINDOW_TITLE", "default title"),
    nil,
    nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
  window.SetKeyCallback(keyCallback)

  // Initialize Glow
  if err := gl.Init(); err != nil {
		panic(err)
	}

  // Initialize keyboard map
  inputManager = comp.InputManager{
    Keys: [348]bool{false},
  };

  // Create Player
  // Just Placeholder until I make a smarter entity manager
  player := eng.Entity{
    Pos: math.NewVec2(100, 100),
    Input: comp.PlayerInputComponent {
      Manager: &inputManager,
    },
  };

  fmt.Println("Press WASD to move. Press Escape to quit.");
  fmt.Println("Note: Rendering code in main file is temporary.");

  shaders := compileShaders();
  shaderProgram := linkShaders(shaders);
  VAO := createTriangleVAO(player.Pos);

  // Game Loop
  lastTime := time.Now();
  var currentTime time.Time;
	for !window.ShouldClose() {
    // handle input
		glfw.PollEvents();

    // update
    currentTime = time.Now();
    player.Update(int(currentTime.Sub(lastTime).Milliseconds()));
    lastTime = currentTime;

    // render
    gl.ClearColor(0.2, 0.5, 0.5, 1.0);
    gl.Clear(gl.COLOR_BUFFER_BIT);

    VAO = createTriangleVAO(player.Pos)
    gl.UseProgram(shaderProgram)
    gl.BindVertexArray(VAO)
    gl.DrawArrays(gl.TRIANGLES, 0, 3)
    gl.BindVertexArray(0)

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

func compileShaders() []uint32 {
	// create the vertex shader
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	shaderSourceChars, freeVertexShaderFunc := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, shaderSourceChars, nil)
	gl.CompileShader(vertexShader)
	checkShaderCompileErrors(vertexShader)

	// create the fragment shader
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	shaderSourceChars, freeFragmentShaderFunc := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, shaderSourceChars, nil)
	gl.CompileShader(fragmentShader)
	checkShaderCompileErrors(fragmentShader)

	defer freeFragmentShaderFunc()
	defer freeVertexShaderFunc()

	return []uint32{vertexShader, fragmentShader}
}

/*
 * Link the provided shaders in the order they were given and return the linked program.
 */
func linkShaders(shaders []uint32) uint32 {
	program := gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(program, shader)
	}
	gl.LinkProgram(program)
	checkProgramLinkErrors(program)

	// shader objects are not needed after they are linked into a program object
	for _, shader := range shaders {
		gl.DeleteShader(shader)
	}

	return program
}

/*
 * Creates the Vertex Array Object for a triangle.
 */
func createTriangleVAO(pos math.Vec2) uint32 {
  halfWidth := float32(WIDTH)/2
  halfHeight:= float32(HEIGHT)/2

  scaledPosX := (pos.X - float64(halfWidth)) / float64(halfWidth)
  scaledPosY := (pos.Y - float64(halfHeight)) / float64(halfHeight)

  x1 := float32(scaledPosX) + 25.0/float32(WIDTH)
  y1 := float32(scaledPosY) + 25.0/float32(HEIGHT)

  x2 := float32(scaledPosX) - 25.0/float32(WIDTH)
  y2 := float32(scaledPosY) + 25.0/float32(HEIGHT)

  x3 := float32(scaledPosX)
  y3 := float32(scaledPosY) - 25.0/float32(HEIGHT)

	vertices := []float32{
		x1, y1, 0.0,
		x2, y2, 0.0,
		x3, y3, 0.0,
	}

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)

	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// specify the format of our vertex input
	// (shader) input 0
	// vertex has size 3
	// vertex items are of type FLOAT
	// do not normalize (already done)
	// stride of 3 * sizeof(float) (separation of vertices)
	// offset of where the position data starts (0 for the beginning)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO
}

func checkShaderCompileErrors(shader uint32) {
	checkGlError(shader, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"ERROR::SHADER::COMPILE_FAILURE")
}

func checkProgramLinkErrors(program uint32) {
	checkGlError(program, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog,
		"ERROR::PROGRAM::LINKING_FAILURE")
}

type getGlParam func(uint32, uint32, *int32)
type getInfoLog func(uint32, int32, *int32, *uint8)

func checkGlError(glObject uint32, errorParam uint32, getParamFn getGlParam,
	getInfoLogFn getInfoLog, failMsg string) {

	var success int32
	getParamFn(glObject, errorParam, &success)
	if success != 1 {
		var infoLog [512]byte
		getInfoLogFn(glObject, 512, nil, (*uint8)(unsafe.Pointer(&infoLog)))
		fmt.Println(string(infoLog[:512]))
	}
}

var vertexShaderSource = `
#version 410 core

layout (location = 0) in vec3 position;

void main()
{
    gl_Position = vec4(position.x, position.y, position.z, 1.0);
}
`

var fragmentShaderSource = `
#version 410 core

out vec4 color;

void main()
{
    color = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
`
