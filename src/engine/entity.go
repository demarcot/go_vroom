package engine;

import ()

type Entity struct {
  X int;
  Y int;
  Input InputComponent;
};

func (e *Entity) Update() {
  e.Input.HandleInput(e);
}
