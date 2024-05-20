package sys

import (
	"fmt"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-template/params"
)

type TestSystem struct {
	params *params.TestParams
}

func (s *TestSystem) Initialize(w *ecs.World) {
	s.params = ecs.GetResource[params.TestParams](w)
}

func (s *TestSystem) Update(w *ecs.World) {
	s.params.TestValue++
}

func (s *TestSystem) Finalize(w *ecs.World) {
	fmt.Println(s.params.TestValue)
}
