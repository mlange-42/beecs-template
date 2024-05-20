package obs

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/beecs-template/params"
)

type TestObserver struct {
	testParams *params.TestParams
}

func (o *TestObserver) Initialize(w *ecs.World) {
	o.testParams = ecs.GetResource[params.TestParams](w)
}

func (o *TestObserver) Update(w *ecs.World) {}

func (o *TestObserver) Header() []string {
	return []string{"TestValue"}
}

func (o *TestObserver) Values(w *ecs.World) []float64 {
	return []float64{float64(o.testParams.TestValue)}
}
