package ddget

import (
	"flag"
	"fmt"
)

type DdgetFlags struct {
	Table         string
	Key           string
	ValueAttrNmae string
	NoNewLine     bool
}

func ParseFlag() (*DdgetFlags, error) {
	ddgf := &DdgetFlags{}

	flag.StringVar(&ddgf.Table, "t", "", "Table name")
	flag.StringVar(&ddgf.Key, "k", "", "Item key")
	flag.StringVar(&ddgf.ValueAttrNmae, "v", "", "Value attribute name")
	flag.BoolVar(&ddgf.NoNewLine, "n", false, "Do not print newline")
	flag.Parse()

	if ddgf.Table == "" {
		err := fmt.Errorf("Table name is required")
		return nil, err
	}

	if ddgf.Key == "" {
		err := fmt.Errorf("Item key is required")
		return nil, err
	}

	return ddgf, nil
}
