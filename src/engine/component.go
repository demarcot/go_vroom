package engine;

type InputComponent interface {
  HandleInput(e *Entity);
};

type RenderComponent interface {
  Render(e *Entity);
};
