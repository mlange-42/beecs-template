package sys

import (
	"fmt"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-template/params"
)

type PrintSystem struct {
	params *params.TestParams
}

func (s *PrintSystem) Initialize(w *ecs.World) {
	s.params = ecs.GetResource[params.TestParams](w)
}

func (s *PrintSystem) Update(w *ecs.World) {
	s.params.TestValue++
}

func (s *PrintSystem) Finalize(w *ecs.World) {
	fmt.Println(s.params.TestValue)
}
