package kitten

import (
	"github.com/nfrazier08/cat"
)

func Meow() string {
	return "Meow!"
}

func Meows() string {
	return "Meow! Meow! Meow!"
}

func BigMeow() string {
	return cat.WhenGrownUp(Meow())
}

func BigMeows() string {
	return cat.WhenGrownUp(Meows())
}

func Sleep() string {
	return "ZzzzZzzzZz"
}

func Play() string {
	return "The kitten loves to play with her ball of yarn!"
}
