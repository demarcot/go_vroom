package engine;

type InputComponent interface {
  HandleInput(e *Entity, timeElapsedMs int);
};

type RenderComponent interface {
  Render(e *Entity, timeElapsedMs int);
};
