package tabio


type RecordUnmarshaler interface {
	UnmarshalRecord([]byte) error
}
