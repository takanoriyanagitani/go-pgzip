package pgzip

type EncodeLevel int

const (
	EncodeLevelUnspecified EncodeLevel = iota
	EncodeLevelStore
	EncodeLevelFast
	EncodeLevelBest
	EncodeLevelDefault
	EncodeLevelHuffmanOnly
	EncodeLevelConstant
)

var EncodeLevelMap map[string]EncodeLevel = map[string]EncodeLevel{
	"Store":       EncodeLevelStore,
	"Fast":        EncodeLevelFast,
	"Best":        EncodeLevelBest,
	"Default":     EncodeLevelDefault,
	"HuffmanOnly": EncodeLevelHuffmanOnly,
	"Constant":    EncodeLevelConstant,
}

func EncodeLevelFromStr(s string) EncodeLevel {
	val, found := EncodeLevelMap[s]
	switch found {
	case true:
		return val
	default:
		return EncodeLevelUnspecified
	}
}

type EncodeConfig struct {
	EncodeLevel
}

var EncodeConfigDefault EncodeConfig = EncodeConfig{
	EncodeLevel: EncodeLevelDefault,
}
