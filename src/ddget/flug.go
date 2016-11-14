package ddget

import (
	"flag"
	"fmt"
	"os"
)

type DdgetFlags struct {
	Table         string
	Key           string
	ValueAttrNmae string
	NoNewLine     bool
	Profile       string
	Region        string
}

func ParseFlag() (*DdgetFlags, error) {
	ddgf := &DdgetFlags{}

	flag.StringVar(&ddgf.Table, "t", "", "Table name")
	flag.StringVar(&ddgf.Key, "k", "", "Item key")
	flag.StringVar(&ddgf.ValueAttrNmae, "v", "", "Value attribute name")
	flag.BoolVar(&ddgf.NoNewLine, "n", false, "Do not print newline")
	flag.StringVar(&ddgf.Profile, "p", "", "Profile name")
	flag.StringVar(&ddgf.Region, "r", "", "Region")
	flag.Parse()

	if ddgf.Table == "" {
		err := fmt.Errorf("Table name is required")
		return nil, err
	}

	if ddgf.Key == "" {
		err := fmt.Errorf("Item key is required")
		return nil, err
	}

	if ddgf.Profile == "" {
		accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
		secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

		if accessKeyId == "" || secretAccessKey == "" {
			ddgf.Profile = "default"
		}
	}

	return ddgf, nil
}
