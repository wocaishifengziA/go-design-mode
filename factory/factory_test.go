package factory

import "testing"

func TestFactoryMethod(t *testing.T) {
	var factory ToyFactory

	factory = &PuzzleFactory{}
	toy1 := factory.New()
	toy1.Play()

	factory = &MarbleFactory{}
	toy2 := factory.New()
	toy2.Play()
}
