package main

import (
	"fmt"
	"os"
	"time"

	. "github.com/maxnetish/algo-go/src/bubble-sort"
	. "github.com/maxnetish/algo-go/src/merge-sort"
	"github.com/maxnetish/algo-go/src/quick-sort"
	"github.com/maxnetish/algo-go/src/utils"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("Arguments: ", os.Args)

	var swaps, passes int
	var dur time.Duration

	cases := []struct {
		args utils.MakeBigArrayParams
	}{
		{
			args: utils.MakeBigArrayParams{
				NumberOfElements: 10,
				To:               100,
			},
		},
		{
			args: utils.MakeBigArrayParams{
				NumberOfElements: 100,
				To:               1000,
			},
		},
		{
			args: utils.MakeBigArrayParams{
				NumberOfElements: 1000,
				To:               10000,
			},
		},
		{
			args: utils.MakeBigArrayParams{
				NumberOfElements: 10000,
				To:               1000,
			},
		},
		{
			args: utils.MakeBigArrayParams{
				NumberOfElements: 100000,
				To:               10000,
			},
		},
		{
			args: utils.MakeBigArrayParams{
				NumberOfElements: 1000000,
				To:               100000,
			},
		},
		{
			args: utils.MakeBigArrayParams{
				NumberOfElements: 10000000,
				To:               10000000,
			},
		},
	}

	fmt.Println("Bubble sort (cocktale sort)")
	for ind, oneCase := range cases {
		if ind < 4 {
			dur = utils.ElapsedTime(func() {
				BubbleSorter.Sort(utils.MakeBigArray(oneCase.args))
			})
			fmt.Println("elements:", oneCase.args.NumberOfElements, "duration:", dur)
		}
	}
	fmt.Println("We will not try to bubble sort out anymore. )")

	fmt.Println("Quick sort (Charles Antony Richard Hoare sort)")
	for _, oneCase := range cases {
		dur = utils.ElapsedTime(func() {
			swaps, passes = quick.Sort(utils.MakeBigArray(oneCase.args))
		})
		fmt.Println("elements:", oneCase.args.NumberOfElements, "duration:", dur, "swaps:", swaps, "passes:", passes)
	}

	fmt.Println("Merge sort (John von Neumann sort)")
	for _, oneCase := range cases {
		dur = utils.ElapsedTime(func() {
			MergeSorter.Sort(utils.MakeBigArray(oneCase.args))
		})
		fmt.Println("elements:", oneCase.args.NumberOfElements, "duration:", dur)
	}
}
