package main

import (
	"fmt"
	"os"
	"time"

	"github.com/maxnetish/algo-go/src/bubble-sort"
	"github.com/maxnetish/algo-go/src/merge-sort"
	"github.com/maxnetish/algo-go/src/quick-sort"
	"github.com/maxnetish/algo-go/src/shell-sort"
	"github.com/maxnetish/algo-go/src/utils"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("Arguments: ", os.Args)

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

	sorters := []Sorter{&bubble.Sorter, &quick.Sorter, &merge.Sorter, &shell.Sorter}

	for _, s := range sorters {
		numberOfCasesToRun := len(cases)
		if s.Slow() {
			numberOfCasesToRun = 5
		}
		fmt.Println("**************************************")
		fmt.Println(s)
		for caseInd := 0; caseInd < numberOfCasesToRun; caseInd++ {
			probeSlice := utils.MakeBigArray(cases[caseInd].args)
			dur = utils.ElapsedTime(func() {
				s.Sort(probeSlice)
			})
			fmt.Println("randoms:", cases[caseInd].args.NumberOfElements, "duration:", dur)
			dur = utils.ElapsedTime(func() {
				s.Sort(probeSlice)
			})
			fmt.Println("sorted:", cases[caseInd].args.NumberOfElements, "duration:", dur)
			probeSlice[len(probeSlice)-3] = 0
			dur = utils.ElapsedTime(func() {
				s.Sort(probeSlice)
			})
			fmt.Println("near sorted:", cases[caseInd].args.NumberOfElements, "duration:", dur)
		}
	}
}
