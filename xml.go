package ismrmrd

import (
	"encoding/xml"
)

const Namespace = "http://www.ismrm.org/ISMRMRD"

// xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xs="http://www.w3.org/2001/XMLSchema" xsi:schemaLocation="http://www.ismrm.org/ISMRMRD ismrmrd.xsd"`

type IsmrmrdHeader struct {
	XMLName                      xml.Name
	Version                      int64                         `xml:"version"`
	SubjectInformation           *SubjectInformation           `xml:"subjectInformation"`
	StudyInformation             *StudyInformation             `xml:"studyInformation"`
	MeasurementInformation       *MeasurementInformation       `xml:"measurementInformation"`
	AcquisitionSystemInformation *AcquisitionSystemInformation `xml:"acquisitionSystemInformation"`
	ExperimentalConditions       ExperimentalConditions        `xml:"experimentalConditions"`
	Encoding                     []Encoding                    `xml:"encoding"`
	SequenceParameters           *SequenceParameters           `xml:"sequenceParameters"`
	UserParameters               *UserParameters               `xml:"userParameters"`
}

type SubjectInformation struct {
	PatientName      string  `xml:"patientName"`
	PatientWeightKg  float32 `xml:"patientWeight_kg"`
	PatientID        string  `xml:"patientID"`
	PatientBirthdate string  `xml:"patientBirthdate"`
	PatientGender    string  `xml:"patientGender"`
}

type StudyInformation struct {
	StudyDate              string `xml:"studyDate"`
	StudyTime              string `xml:"studyTime"`
	StudyID                string `xml:"studyID"`
	AccessionNumber        int64  `xml:"accessionNumber"`
	ReferringPhysicianName string `xml:"referringPhysicianName"`
	StudyDescription       string `xml:"studyDescription"`
	StudyInstanceUID       string `xml:"studyInstanceUID"`
}

type MeasurementInformation struct {
	MeasurementID           string                  `xml:"measurementID"`
	SeriesDate              string                  `xml:"seriesDate"`
	SeriesTime              string                  `xml:"seriesTime"`
	PatientPosition         string                  `xml:"patientPosition"`
	InitialSeriesNumber     int64                   `xml:"initialSeriesNumber"`
	ProtocolName            string                  `xml:"protocolName"`
	SeriesDescription       string                  `xml:"seriesDescription"`
	MeasurementDependency   []MeasurementDependency `xml:"measurementDependency"`
	SeriesInstanceUIDRoot   string                  `xml:"seriesInstanceUIDRoot"`
	FrameOfReferenceUID     string                  `xml:"frameOfReferenceUID"`
	ReferencedImageSequence ReferencedImageSequence `xml:"referencedImageSequence"`
}

type MeasurementDependency struct {
	DependencyType string `xml:"dependencyType"`
	MeasurementID  string `xml:"measurementID"`
}

type ReferencedImageSequence struct {
	ReferencedSOPInstanceUID []string `xml:"referencedSOPInstanceUID"`
}

type AcquisitionSystemInformation struct {
	SystemVendor                  string  `xml:"systemVendor"`
	SystemModel                   string  `xml:"systemModel"`
	SystemFieldStrengthT          float32 `xml:"systemFieldStrength_T"`
	RelativeReceiverNoiseBandwith float32 `xml:"relativeReceiverNoiseBandwidth"`
	ReceiverChannels              uint16  `xml:"receiverChannels"`
	InstitutionName               string  `xml:"institutionName"`
	StationName                   string  `xml:"stationName"`
}

type ExperimentalConditions struct {
	H1ResonanceFrequencyHz int64 `xml:"H1resonanceFrequency_Hz"`
}

type Encoding struct {
	EncodedSpace          EncodingSpace          `xml:"encodedSpace"`
	ReconSpace            EncodingSpace          `xml:"reconSpace"`
	EncodingLimits        EncodingLimits         `xml:"encodingLimits"`
	Trajectory            string                 `xml:"trajectory"`
	TrajectoryDescription *TrajectoryDescription `xml:"trajectoryDescription"`
	ParallelImaging       *ParallelImaging       `xml:"parallelImaging"`
}

type EncodingSpace struct {
	MatrixSize    MatrixSize  `xml:"matrixSize"`
	FieldOfViewMM FieldOfView `xml:"fieldOfView_mm"`
}

type MatrixSize struct {
	X uint16 `xml:"x"`
	Y uint16 `xml:"y"`
	Z uint16 `xml:"z"`
}

type FieldOfView struct {
	X float32 `xml:"x"`
	Y float32 `xml:"y"`
	Z float32 `xml:"z"`
}

type EncodingLimits struct {
	KSpaceEncodingStep0 *Limit `xml:"kspace_encoding_step_0"`
	KSpaceEncodingStep1 *Limit `xml:"kspace_encoding_step_1"`
	KSpaceEncodingStep2 *Limit `xml:"kspace_encoding_step_2"`
	Average             *Limit `xml:"average"`
	Slice               *Limit `xml:"slice"`
	Contrast            *Limit `xml:"contrast"`
	Phase               *Limit `xml:"phase"`
	Repetition          *Limit `xml:"repetition"`
	Set                 *Limit `xml:"set"`
	Segment             *Limit `xml:"segment"`
}

type Limit struct {
	Minimum uint16 `xml:"minimum"`
	Maximum uint16 `xml:"maximum"`
	Center  uint16 `xml:"center"`
}

type TrajectoryDescription struct {
	Name                string                `xml:"name"`
	UserParameterLong   []UserParameterLong   `xml:"userParameterLong"`
	UserParameterDouble []UserParameterDouble `xml:"userParameterDouble"`
	Comment             string                `xml:"comment"`
}

type ParallelImaging struct {
	AccelerationFactor    AccelerationFactor `xml:"accelerationFactor"`
	CalibrationMode       string             `xml:"calibrationMode"`
	InterleavingDimension string             `xml:"interleavingDimension"`
}

type AccelerationFactor struct {
	KSpaceEncodingStep1 uint16 `xml:"kspace_encoding_step_1"`
	KSpaceEncodingStep2 uint16 `xml:"kspace_encoding_step_2"`
}

type SequenceParameters struct {
	TR           []float32
	TE           []float32
	TI           []float32
	FlipAngleDeg []float32 `xml:"flipAngle_deg"`
}

type UserParameters struct {
	UserParameterLong   []UserParameterLong   `xml:"userParameterLong"`
	UserParameterDouble []UserParameterDouble `xml:"userParameterDouble"`
	UserParameterString []UserParameterString `xml:"userParameterString"`
	UserParameterBase64 []UserParameterBase64 `xml:"userParameterBase64"`
}

type UserParameterLong struct {
	Name  string `xml:"name"`
	Value int64  `xml:"value"`
}

type UserParameterDouble struct {
	Name  string  `xml:"name"`
	Value float64 `xml:"value"`
}

type UserParameterString struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type UserParameterBase64 struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func Serialize(head *IsmrmrdHeader) ([]byte, error) {
	head.XMLName = xml.Name{Namespace, "ismrmrdHeader"}
	return xml.MarshalIndent(head, "", "  ")
}

func Deserialize(data []byte) (*IsmrmrdHeader, error) {
	var head IsmrmrdHeader
	err := xml.Unmarshal(data, &head)
	return &head, err
}
