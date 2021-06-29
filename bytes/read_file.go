package bytes

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(fname string) (nums []int64, err error) {

	b, err := ioutil.ReadFile(fname)

	if err != nil { return nil, err }
	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int64, 0, len(lines))
	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 { continue }

		// Atoi better suits the job when we know exactly what we're dealing

		n, err := strconv.ParseInt(l,10,64)
		if err != nil { return nil, err }
		nums = append(nums, n)
	}
	return nums, nil

}
