package main

type Sorter interface {
	Sort(a []int)
	Slow() bool
}
