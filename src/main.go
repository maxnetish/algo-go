package main

import (
	"fmt"
	"os"
	"time"

	"github.com/maxnetish/algo-go/src/bubble-sort"
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
				To:               100000,
			},
		},
		{
			args: utils.MakeBigArrayParams{
				NumberOfElements: 20000,
				To:               100000,
			},
		},
	}

	for _, oneCase := range cases {
		dur = utils.ElapsedTime(func() {
			swaps, passes = bubble.Sort(utils.MakeBigArray(oneCase.args))
		})
		fmt.Println("elements:", oneCase.args.NumberOfElements, "duration:", dur, "swaps:", swaps, "passes:", passes)
	}
}
