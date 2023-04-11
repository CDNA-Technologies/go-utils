package testing

func IsErrorEqual(a, b error) bool {
	// If both are nil then true
	if a == nil && b == nil {
		return true
	}

	// If either one in nil then false
	if a == nil || b == nil {
		return false
	}

	return a.Error() == b.Error()
}
