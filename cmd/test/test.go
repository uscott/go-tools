package main

import (
	"fmt"

	"github.com/uscott/go-tools/mathx"
)

func main() {
	var (
		x, chunkSz float64
		prec       uint
	)
	x = 48443.9283202302930398
	prec = 10
	fmt.Println()
	fmt.Printf("x: %f, prec: %d => %s\n", x, prec, mathx.FtoS(x, prec))
	prec = 0
	fmt.Printf("x: %f, prec: %d => %s\n", x, prec, mathx.FtoS(x, prec))
	chunkSz = 20
	fmt.Printf(
		"round x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkRdToS(x, chunkSz))
	chunkSz = 0.5
	fmt.Printf(
		"round x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkRdToS(x, chunkSz))
	chunkSz = 0.0001
	fmt.Printf(
		"round x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkRdToS(x, chunkSz))
	chunkSz = 20
	fmt.Printf(
		"ceil x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkClToS(x, chunkSz))
	chunkSz = 0.5
	fmt.Printf(
		"ceil x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkClToS(x, chunkSz))
	chunkSz = 0.0001
	fmt.Printf(
		"ceil x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkClToS(x, chunkSz))
	chunkSz = 20
	fmt.Printf(
		"floor x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkFlToS(x, chunkSz))
	chunkSz = 0.5
	fmt.Printf(
		"floor x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkFlToS(x, chunkSz))
	chunkSz = 0.0001
	fmt.Printf(
		"floor x: %f, chunkSz: %f => %s\n", x, chunkSz, mathx.ChunkFlToS(x, chunkSz))
}
