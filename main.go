package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Pokemon struct {
	Name string
}

func (p *Pokemon) IChooseU() {
	fmt.Println("I choose u,", p.Name)
}

func StartFight() {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			if err != nil {
				err = errors.Wrap(err, time.Now().String())
				fmt.Printf("Catch: %s", err)
			}
		}
	}()

	b := &Pokemon{"Bulbasaur"}
	b.IChooseU()

	p := &Pokemon{"Pikachu"}
	p = nil
	p.IChooseU()
}

func MillionFiles() {
	if err := os.Mkdir("files", 0777); err != nil {
		return
	}

	for i := 0; i < 1000000; i++ {
		f, err := os.Create("files/" + strconv.Itoa(i) + ".txt")
		if err != nil {
			fmt.Println(err)
		}
		func() {
			defer f.Close()
		}()
	}
}

func main() {
	StartFight()
	MillionFiles()
}
