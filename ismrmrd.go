package ismrmrd

const (
	ISMRMRD_VERSION_MAJOR = 1
	ISMRMRD_VERSION_MINOR = 1
	ISMRMRD_VERSION_PATCH = 0

	ISMRMRD_POSITION_LENGTH  = 3
	ISMRMRD_DIRECTION_LENGTH = 3
	ISMRMRD_USER_INTS        = 8
	ISMRMRD_USER_FLOATS      = 8
	ISMRMRD_PHYS_STAMPS      = 3
	ISMRMRD_CHANNEL_MASKS    = 16

	ACQ_FIRST_IN_SLICE = 1 << 6
	ACQ_LAST_IN_SLICE  = 1 << 7
)

type EncodingCounters struct {
	KspaceEncodeStep1 uint16
	KspaceEncodeStep2 uint16
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
	Measurement_uid      uint32
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
	Data []float32
}

type ImageHeader struct {
	Version              uint16
	Flags                uint64
	MeasurementUid       uint32
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
	ImageDataype         uint16
	ImageType            uint16
	ImageIndex           uint16
	ImageSeriesIndex     uint16
	UserInt              [ISMRMRD_USER_INTS]int32
	UserFloat            [ISMRMRD_USER_FLOATS]float32
}
