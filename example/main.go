package main

import (
	"fmt"
	"log"

	"github.com/naegelejd/go-ismrmrd"
)

func main() {
	dset, err := ismrmrd.Open("testdata.h5", "dataset")
	if err != nil {
		log.Fatal(err)
	}

	xml, err := dset.ReadXMLHeader()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xml)

	fmt.Println(dset.NumberOfAcquisitions())
}
