package simplefactory

import "fmt"

const (
	puzzle = iota
	marble
	unsupported
)

type Toy interface {
	Play()
}

type ToyFactory struct{}

type Puzzle struct{}

func (*Puzzle) Play() {
	fmt.Println("palying puzzle")
}

type Marble struct{}

func (*Marble) Play() {
	fmt.Println("playing marble")
}

func (t *ToyFactory) New(typ int) Toy {
	switch typ {
	case puzzle:
		fmt.Println("puzzle")
		return new(Puzzle)
	case marble:
		fmt.Println("marble")
		return new(Marble)
	default:
		fmt.Println("unsupported type of toy")
		return nil
	}
}
