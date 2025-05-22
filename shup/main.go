package shup

import (
	"fmt"
)

type CollectorWriter struct {
	data []byte
}

func (cw *CollectorWriter) Write(data []byte) (int, error) {
	cw.data = append(cw.data, data...)
	return 0, nil
}

func (cw CollectorWriter) Collect() string {
	return string(cw.data[:])
}

func Main(filename string) {
	fmt.Printf("Collating %s\n", filename)
	holder := CollectorWriter{}
	Collate(0, filename, &holder, &FileRegistry{})

	fmt.Printf("%s\n", holder.Collect())
}