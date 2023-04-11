package testing

func IsErrorEqual(a, b error) bool {
	isANil := a == nil
	isBNil := b == nil

	// If both are nil then true
	if isANil && isBNil {
		return true
	}

	// If either one in nil then false
	if isANil || isBNil {
		return false
	}

	// If error message differ than true
	if a.Error() != b.Error() {
		return false
	}

	return true
}
