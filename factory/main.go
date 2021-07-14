package factory

import "fmt"

type Toy interface {
	Play()
}

type ToyFactory interface {
	New() Toy
}

type Puzzle struct{}

func (*Puzzle) Play() {
	fmt.Println("palying puzzle")
}

type Marble struct{}

func (*Marble) Play() {
	fmt.Println("playing marble")
}

type PuzzleFactory struct{}

func (p *PuzzleFactory) New() Toy {
	return new(Puzzle)
}

type MarbleFactory struct{}

func (p *MarbleFactory) New() Toy {
	return new(Marble)
}
