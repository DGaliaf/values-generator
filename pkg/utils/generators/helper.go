package generators

var (
	supportedTypes = []string{"string", "int", "uuid"}
)

func IsSupportedType(t string) bool {
	for i := 0; i < len(supportedTypes); i++ {
		if t == supportedTypes[i] {
			return true
		}
	}

	return false
}
