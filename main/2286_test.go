package main

import (
	"testing"
)

func Test2286(t *testing.T) {
	b := Constructor(2, 5)
	b.Gather(4, 0)
	b.Gather(2, 0)
	b.Scatter(5, 1)
	b.Scatter(5, 1)
}
