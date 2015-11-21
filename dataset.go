package ismrmrd

import (
	"fmt"
	"strings"

	"github.com/sbinet/go-hdf5"
)

type Dataset struct {
	file      *hdf5.File
	groupname string
	group     *hdf5.Group
}

func Open(filename, groupname string) (dset *Dataset, err error) {
	var file *hdf5.File
	file, err = hdf5.OpenFile(filename, hdf5.F_ACC_RDWR)
	if err != nil {
		return
	}

	var group *hdf5.Group
	group, err = file.OpenGroup(groupname)
	if err != nil {
		return
	}

	dset = &Dataset{file, groupname, group}
	return
}

func Create(filename, groupname string) (dset *Dataset, err error) {
	var file *hdf5.File
	file, err = hdf5.CreateFile(filename, hdf5.F_ACC_TRUNC)
	if err != nil {
		return
	}

	var group *hdf5.Group
	group, err = file.CreateGroup(groupname)
	if err != nil {
		return
	}

	dset = &Dataset{file, groupname, group}
	return
}

func (d *Dataset) Close() error {
	if err := d.group.Close(); err != nil {
		return err
	}
	if err := d.file.Close(); err != nil {
		return err
	}
	return nil
}

func (d *Dataset) ReadXMLHeader() (xml string, err error) {
	var dataset *hdf5.Dataset
	dataset, err = d.file.OpenDataset(d.makePath("xml"))
	if err != nil {
		return
	}
	defer dataset.Close()

	// TODO: delete this block... just checking that datatypes
	// are automatically loaded correctly
	var dtype *hdf5.Datatype
	dtype, err = dataset.Datatype()
	if err != nil {
		return
	}
	defer dtype.Close()
	if !dtype.Equal(hdf5.T_GO_STRING) {
		err = fmt.Errorf("invalid datatype")
		return
	}

	wrapper := make([]string, 1)
	if err = dataset.Read(&wrapper); err != nil {
		return
	}
	fmt.Println("read: ", wrapper[0])

	xml = wrapper[0]

	return
}

func (d *Dataset) WriteXMLHeader(header string) error {
	dataspace, err := hdf5.CreateSimpleDataspace([]uint{1}, nil)
	if err != nil {
		return err
	}
	defer dataspace.Close()

	datatype := hdf5.T_GO_STRING

	dataset, err := d.file.CreateDataset(d.makePath("xml"), datatype, dataspace)
	if err != nil {
		return err
	}
	defer dataset.Close()

	wrapper := []string{header}
	if err := dataset.Write(&wrapper); err != nil {
		return err
	}

	return nil
}

func (d *Dataset) numberOfElements(path string) uint {
	dataset, err := d.file.OpenDataset(path)
	if err != nil {
		return 0
	}
	dataspace := dataset.Space()
	dims, _, err := dataspace.SimpleExtentDims()
	if err != nil {
		return 0
	}

	if len(dims) < 1 {
		return 0
	}

	return dims[0]
}

func (d *Dataset) NumberOfAcquisitions() int {
	return int(d.numberOfElements(d.makePath("data")))
}

// func (d *Dataset) ReadAcquisition(acqNum int) (*Acquisition, error) {

// }

// func (d *Dataset) AppendAcquisition(acq *Acquisition) error {

// }

func (d *Dataset) NumberOfImages(imgPath string) int {
	return int(d.numberOfElements(d.makePath(imgPath, "header")))
}

// func (d *Dataset) ReadImage(imgPath string, imgNum int) (*Image, error) {

// }

// func (d *Dataset) AppendImage(imgPath string, img *Image) error {

// }

func (d *Dataset) makePath(components ...string) string {
	return strings.Join(append([]string{d.groupname}, components...), "/")
}

func (d *Dataset) NumberOfArrays(arrPath string) int {
	return int(d.numberOfElements(d.makePath(arrPath)))
}

// func (d *Dataset) ReadArray(arrPath string, arrNum int) (*Array, error) {

// }

// func (d *Dataset) AppendArray(arrPath string, arr *Array) error {

// }
