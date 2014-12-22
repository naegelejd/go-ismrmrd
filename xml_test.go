package ismrmrd

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	version = 110

	patientName      = "Joe Naegele"
	patientWeight    = 72.57
	patientID        = "123abc45de"
	patientBirthdate = "1988-01-01"
	patientGender    = "M"

	studyDate              = "2014-12-31"
	studyTime              = "08:45:00"
	studyID                = "4242"
	accessionNumber        = 4242
	referringPhysicianName = "Dr. Oz"
	studyDescription       = "just a test"
	studyInstanceUID       = "1.2.345.678901.2.345.6.789012345.6789.0123456789.0123"

	measurementID         = "1"
	seriesDate            = "2014-12-31"
	seriesTime            = "08:46:00"
	patientPosition       = "HFS"
	initialSeriesNumber   = 2
	protocolName          = "JN template"
	seriesDescription     = "2D B0 Field Map"
	seriesInstanceUIDRoot = "1.2.345.678901.2.345.6.789"
	frameOfReferenceUID   = "1.2.345.678901.2.345.6.789012345.6789.0123456789.0123"

	referencedImageSequence0 = "1.2.1.2."
	referencedImageSequence1 = "1.2."
	referencedImageSequence2 = ""
	referencedImageSequence3 = ""

	systemVendor                   = "SECRET"
	systemModel                    = "SPECIAL"
	systemFieldStrengthT           = 3.0
	relativeReceiverNoiseBandwidth = 1.0
	receiverChannels               = 1
	institutionName                = "GOV"
	stationName                    = "controlroom"

	h1ResonanceFrequencyHz = 123136640

	matrixX = 128
	matrixY = 128
	matrixZ = 1

	fovX = 256
	fovY = 256
	fovZ = 256

	minY = 0
	maxY = 127
	cenY = 64

	minZ = 0
	maxZ = 4
	cenZ = 2

	tr0        = 150.0
	te0        = 4.8
	te1        = 0.0
	ti0        = 0.0
	flipAngle0 = 9.0

	trajectory = "cartesian"
)

var userParam0 = UserParameterString{"imageType", "ORIGINAL//PRIMARY//OTHER"}
var userParam1 = UserParameterString{"scanningSequence", "RM"}
var userParam2 = UserParameterString{"sequenceVariant", "NONE"}
var userParam3 = UserParameterString{"scanOptions", "NONE"}
var userParam4 = UserParameterString{"mrAcquisitionType", "2D"}
var userParam5 = UserParameterString{"freqEncodingDirection", "COL"}
var userParam6 = UserParameterDouble{"triggerTime", 0.0}

var testXML string
var testHeader *IsmrmrdHeader

func TestMain(m *testing.M) {
	f, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	testXML = string(b)

	testHeader = &IsmrmrdHeader{
		Version: version,
		SubjectInformation: &SubjectInformation{
			patientName, patientWeight, patientID, patientBirthdate, patientGender,
		},
		StudyInformation: &StudyInformation{
			studyDate, studyTime, studyID, accessionNumber,
			referringPhysicianName, studyDescription, studyInstanceUID,
		},
		MeasurementInformation: &MeasurementInformation{
			measurementID, seriesDate, seriesTime, patientPosition,
			initialSeriesNumber, protocolName, seriesDescription, nil,
			seriesInstanceUIDRoot, frameOfReferenceUID,
			ReferencedImageSequence{[]string{
				referencedImageSequence0, referencedImageSequence1,
				referencedImageSequence2, referencedImageSequence3,
			}},
		},
		AcquisitionSystemInformation: &AcquisitionSystemInformation{
			systemVendor, systemModel, systemFieldStrengthT,
			relativeReceiverNoiseBandwidth, receiverChannels,
			institutionName, stationName,
		},
		ExperimentalConditions: ExperimentalConditions{h1ResonanceFrequencyHz},
	}

	espace := EncodingSpace{
		MatrixSize{matrixX, matrixY, matrixZ},
		FieldOfView{fovX, fovY, fovZ},
	}
	rspace := EncodingSpace{
		MatrixSize{matrixX, matrixY, matrixZ},
		FieldOfView{fovX, fovY, fovZ},
	}
	elimits := EncodingLimits{
		KSpaceEncodingStep1: &Limit{minY, maxY, cenY},
		Slice:               &Limit{minZ, maxZ, cenZ},
	}

	e := Encoding{
		EncodedSpace:   espace,
		ReconSpace:     rspace,
		EncodingLimits: elimits,
		Trajectory:     trajectory,
	}

	testHeader.Encoding = append(testHeader.Encoding, e)

	testHeader.SequenceParameters = &SequenceParameters{
		[]float32{tr0}, []float32{te0, te1}, []float32{ti0}, []float32{flipAngle0}}

	testHeader.UserParameters = &UserParameters{
		UserParameterString: []UserParameterString{
			userParam0, userParam1, userParam2, userParam3, userParam4, userParam5,
		},
		UserParameterDouble: []UserParameterDouble{userParam6},
	}

	os.Exit(m.Run())
}

func TestSerialize(t *testing.T) {
	head := &IsmrmrdHeader{
		Version: version,
		SubjectInformation: &SubjectInformation{
			patientName, patientWeight, patientID, patientBirthdate, patientGender,
		},
		StudyInformation: &StudyInformation{
			studyDate, studyTime, studyID, accessionNumber,
			referringPhysicianName, studyDescription, studyInstanceUID,
		},
		MeasurementInformation: &MeasurementInformation{
			measurementID, seriesDate, seriesTime, patientPosition,
			initialSeriesNumber, protocolName, seriesDescription, nil,
			seriesInstanceUIDRoot, frameOfReferenceUID,
			ReferencedImageSequence{[]string{
				referencedImageSequence0, referencedImageSequence1,
				referencedImageSequence2, referencedImageSequence3,
			}},
		},
		AcquisitionSystemInformation: &AcquisitionSystemInformation{
			systemVendor, systemModel, systemFieldStrengthT,
			relativeReceiverNoiseBandwidth, receiverChannels,
			institutionName, stationName,
		},
		ExperimentalConditions: ExperimentalConditions{h1ResonanceFrequencyHz},
	}

	espace := EncodingSpace{
		MatrixSize{matrixX, matrixY, matrixZ},
		FieldOfView{fovX, fovY, fovZ},
	}
	rspace := EncodingSpace{
		MatrixSize{matrixX, matrixY, matrixZ},
		FieldOfView{fovX, fovY, fovZ},
	}
	elimits := EncodingLimits{
		KSpaceEncodingStep1: &Limit{minY, maxY, cenY},
		Slice:               &Limit{minZ, maxZ, cenZ},
	}

	e := Encoding{
		EncodedSpace:   espace,
		ReconSpace:     rspace,
		EncodingLimits: elimits,
		Trajectory:     trajectory,
	}

	head.Encoding = append(head.Encoding, e)

	head.SequenceParameters = &SequenceParameters{
		[]float32{tr0}, []float32{te0, te1}, []float32{ti0}, []float32{flipAngle0}}

	head.UserParameters = &UserParameters{
		UserParameterString: []UserParameterString{
			userParam0, userParam1, userParam2, userParam3, userParam4, userParam5,
		},
		UserParameterDouble: []UserParameterDouble{userParam6},
	}

	b, err := Serialize(head)
	if err != nil {
		t.Error(err)
	}
	print(string(b))

	// The test file contains a trailing newline :/
	if string(b) != testXML[:len(testXML)-1] {
		t.Fail()
	}
}

func TestDeserialize(t *testing.T) {
	b := []byte(testXML)
	head, err := Deserialize(b)
	if err != nil {
		t.Error(err)
	}

	if !equal(head, testHeader) {
		t.Fail()
	}
}

func equal(h1, h2 *IsmrmrdHeader) bool {
	if h1.Version != h2.Version {
		print(h1.Version)
		print(h2.Version)
		print("Bad version")
		return false
	} else if *h1.SubjectInformation != *h2.SubjectInformation {
		print("Bad SubjectInformation")
		return false
	} else if *h1.StudyInformation != *h2.StudyInformation {
		print("Bad StudyInformation")
		return false
	} else if *h1.AcquisitionSystemInformation != *h2.AcquisitionSystemInformation {
		print("Bad AcquisitionSystemInformation")
		return false
	} else if h1.ExperimentalConditions != h2.ExperimentalConditions {
		print("Bad ExperimentalConditions")
		return false
	} else if !same_meas_info(h1.MeasurementInformation, h2.MeasurementInformation) {
		print("Bad MeasurementInformation")
		return false
	} else if !same_encoding(&h1.Encoding, &h2.Encoding) {
		print("Bad Encoding")
		return false
	} else if !same_seq_params(h1.SequenceParameters, h2.SequenceParameters) {
		print("Bad SequenceParameters")
		return false
	} else if !same_user_params(h1.UserParameters, h2.UserParameters) {
		print("Bad UserParameters")
		return false
	}

	return true
}

func same_meas_info(m1, m2 *MeasurementInformation) bool {
	if len(m1.MeasurementDependency) != len(m2.MeasurementDependency) {
		return false
	}
	for i := 0; i < len(m1.MeasurementDependency); i++ {
		if m1.MeasurementDependency[i] != m2.MeasurementDependency[i] {
			return false
		}
	}

	if len(m1.ReferencedImageSequence.ReferencedSOPInstanceUID) !=
		len(m2.ReferencedImageSequence.ReferencedSOPInstanceUID) {
		return false
	}
	for i := 0; i < len(m1.ReferencedImageSequence.ReferencedSOPInstanceUID); i++ {
		if m1.ReferencedImageSequence.ReferencedSOPInstanceUID[i] !=
			m2.ReferencedImageSequence.ReferencedSOPInstanceUID[i] {
			return false
		}
	}

	return m1.MeasurementID == m2.MeasurementID &&
		m1.SeriesDate == m2.SeriesDate &&
		m1.SeriesTime == m2.SeriesTime &&
		m1.PatientPosition == m2.PatientPosition &&
		m1.InitialSeriesNumber == m2.InitialSeriesNumber &&
		m1.ProtocolName == m2.ProtocolName &&
		m1.SeriesDescription == m2.SeriesDescription &&
		m1.SeriesInstanceUIDRoot == m2.SeriesInstanceUIDRoot &&
		m1.FrameOfReferenceUID == m2.FrameOfReferenceUID
}

func same_encoding(e1, e2 *[]Encoding) bool {
	if len(*e1) != len(*e2) {
		return false
	}

	for i := 0; i < len(*e1); i++ {
		if (*e1)[i] != (*e2)[i] {
			return false
		}
	}
	return true
}

func same_seq_params(p1, p2 *SequenceParameters) bool {
	if !(len(p1.TR) == len(p2.TR) && len(p1.TE) == len(p2.TE) && len(p1.TI) == len(p2.TI) && len(p1.FlipAngleDeg) == len(p2.FlipAngleDeg)) {
		return false
	}

	for i := 0; i < len(p1.TR); i++ {
		if p1.TR[i] != p2.TR[i] {
			return false
		}
	}

	for i := 0; i < len(p1.TE); i++ {
		if p1.TE[i] != p2.TE[i] {
			return false
		}
	}

	for i := 0; i < len(p1.TI); i++ {
		if p1.TI[i] != p2.TI[i] {
			return false
		}
	}

	for i := 0; i < len(p1.FlipAngleDeg); i++ {
		if p1.FlipAngleDeg[i] != p2.FlipAngleDeg[i] {
			return false
		}
	}

	return true
}

func same_user_params(p1, p2 *UserParameters) bool {
	return true
}
