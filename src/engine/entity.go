package engine;

import (
  math "demarcot/vroom/src/engine/math"
)

type Entity struct {
  Pos math.Vec2;
  Input InputComponent;
};

func (e *Entity) Update(timeElapsedMs int) {
  e.Input.HandleInput(e, timeElapsedMs);
}
