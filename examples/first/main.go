package main

import (
	"fmt"

	"github.com/ed-henrique/frevo/pkg/duration"
	"github.com/ed-henrique/frevo/pkg/sim"
)

type a struct {
	b int
	s *sim.Simulation
}

func (aa a) Do() {
	fmt.Println(aa.b, aa.s.CurrentTime()+duration.Seconds(10))
	aa.b++

	aa.s.Schedule(aa.s.CurrentTime()+duration.Seconds(10), aa)
}

func main() {
	s := sim.New(sim.WithStopTime(duration.Seconds(30)), sim.WithWarnings())
	aa := a{s: s}

	s.Schedule(duration.Seconds(10), aa)
	s.Run()
}
