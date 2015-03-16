package ismrmrd

const (
	ISMRMRD_VERSION_MAJOR = 1
	ISMRMRD_VERSION_MINOR = 2
	ISMRMRD_VERSION_PATCH = 2

	ISMRMRD_USER_INTS        = 8
	ISMRMRD_USER_FLOATS      = 8
	ISMRMRD_PHYS_STAMPS      = 3
	ISMRMRD_CHANNEL_MASKS    = 16
	ISMRMRD_NDARRAY_MAXDIM   = 7
	ISMRMRD_POSITION_LENGTH  = 3
	ISMRMRD_DIRECTION_LENGTH = 3

	ISMRMRD_USHORT   = 1
	ISMRMRD_SHORT    = 2
	ISMRMRD_UINT     = 3
	ISMRMRD_INT      = 4
	ISMRMRD_FLOAT    = 5
	ISMRMRD_DOUBLE   = 6
	ISMRMRD_CXFLOAT  = 7
	ISMRMRD_CXDOUBLE = 8

	ISMRMRD_IMTYPE_MAGNITUDE = 1
	ISMRMRD_IMTYPE_PHASE     = 2
	ISMRMRD_IMTYPE_REAL      = 3
	ISMRMRD_IMTYPE_IMAG      = 4
	ISMRMRD_IMTYPE_COMPLEX   = 5
)

type EncodingCounters struct {
	KSpaceEncodeStep1 uint16
	KSpaceEncodeStep2 uint16
	Average           uint16
	Slice             uint16
	Contrast          uint16
	Phase             uint16
	Repetition        uint16
	Set               uint16
	Segment           uint16
	User              [ISMRMRD_USER_INTS]uint16
}

type AcquisitionHeader struct {
	Version              uint16
	Flags                uint64
	MeasurementUID       uint32
	ScanCounter          uint32
	AcquisitionTimeStamp uint32
	PhysiologyTimeStamp  [ISMRMRD_PHYS_STAMPS]uint32
	NumberOfSamples      uint16
	AvailableChannels    uint16
	ActiveChannels       uint16
	ChannelMask          [ISMRMRD_CHANNEL_MASKS]uint64
	DiscardPre           uint16
	DiscardPost          uint16
	CenterSample         uint16
	EncodingSpaceRef     uint16
	TrajectoryDimensions uint16
	SampleTimeUs         float32
	Position             [ISMRMRD_POSITION_LENGTH]float32
	ReadDirection        [ISMRMRD_DIRECTION_LENGTH]float32
	PhaseDirection       [ISMRMRD_DIRECTION_LENGTH]float32
	SliceDirection       [ISMRMRD_DIRECTION_LENGTH]float32
	PatientablePosition  [ISMRMRD_POSITION_LENGTH]float32
	Idx                  EncodingCounters
	UserInt              [ISMRMRD_USER_INTS]int32
	UserFloat32          [ISMRMRD_USER_FLOATS]float32
}

type Acquisition struct {
	Head AcquisitionHeader
	Traj []float32
	Data []complex64
}

type ImageHeader struct {
	Version              uint16
	DataType             uint16
	Flags                uint64
	MeasurementUID       uint32
	MatrixSize           [3]uint16
	FieldOfView          [3]float32
	Channels             uint16
	Position             [ISMRMRD_POSITION_LENGTH]float32
	ReadDirection        [ISMRMRD_DIRECTION_LENGTH]float32
	PhaseDirection       [ISMRMRD_DIRECTION_LENGTH]float32
	SliceDirection       [ISMRMRD_DIRECTION_LENGTH]float32
	PatientTablePosition [ISMRMRD_POSITION_LENGTH]float32
	Average              uint16
	Slice                uint16
	Contrast             uint16
	Phase                uint16
	Repetition           uint16
	Set                  uint16
	AcquisitionTimeStamp uint32
	PhysiologyTimeStamp  [ISMRMRD_PHYS_STAMPS]uint32
	ImageType            uint16
	ImageIndex           uint16
	ImageSeriesIndex     uint16
	UserInt              [ISMRMRD_USER_INTS]int32
	UserFloat            [ISMRMRD_USER_FLOATS]float32
	AttributeStringLen   uint32
}
