package security
func Compare(value string, hashedValue string) bool {
	return Hash(value) == hashedValue
}