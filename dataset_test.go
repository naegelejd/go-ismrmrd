package ismrmrd

import (
	"github.com/sbinet/go-hdf5"
	"os"
	"testing"
)

const (
	filename  = "testdata.h5"
	groupname = "dataset"
)

func TestCreateDataset(t *testing.T) {
	dset, err := Create(filename, groupname)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	if err := dset.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestOpenDataset(t *testing.T) {
	dset, err := Create(filename, groupname)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)
	if err := dset.Close(); err != nil {
		t.Fatal(err)
	}

	dset, err = Open(filename, groupname)
	if err != nil {
		t.Fatal(err)
	}
	if err := dset.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestWriteXMLHeader(t *testing.T) {
	hdf5.DisplayErrors(true)

	dset, err := Create(filename, groupname)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dset.Close()
		os.Remove(filename)
	}()

	header := "This is the test XML header"
	if err := dset.WriteXMLHeader(header); err != nil {
		t.Fatal(err)
	}
}

func TestReadXMLHeader(t *testing.T) {
	hdf5.DisplayErrors(true)

	dset, err := Create(filename, groupname)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		dset.Close()
		os.Remove(filename)
	}()

	header := "This is the test XML header"
	if err := dset.WriteXMLHeader(header); err != nil {
		t.Fatal(err)
	}

	xml, err := dset.ReadXMLHeader()
	if err != nil {
		t.Fatal(err)
	}

	if xml != header {
		t.Fatalf("XML header does not match what was written (%s)", xml)
	}
}
