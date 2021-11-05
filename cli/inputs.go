package cli

import (
	"flag"
	"fmt"
	"strconv"
)

func GetUserInput(typSpecifier *string) ([]int64, error) {
	flag.StringVar(typSpecifier, "inputs", "", "Usage: grpcfactorial --inputs integers 16 464 100 ...")
	flag.Parse()
	lst := flag.Args()
	if len(lst) == 0 {
		return nil, fmt.Errorf("no input supplied")
	}
	readyInp, err := sanitizeInputs(lst)
	if err != nil {
		return nil, err
	}
	return readyInp, nil
}

func sanitizeInputs(args []string) ([]int64, error) {
	sanitized := make([]int64, 0)
	for _, value := range args {
		i64, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("non int64 value %+v supplied in input : %+v", value, err)
		}
		sanitized = append(sanitized, i64)
	}
	return sanitized, nil
}
