package testing

/**
	Compare the given errors and return true if equal, or else returns false.
**/
func IsErrorEqual(wantErr, err error) bool {
	// If both are nil then true
	if wantErr == nil && err == nil {
		return true
	}

	// If either one in nil then false
	if wantErr == nil || err == nil {
		return false
	}

	return wantErr.Error() == err.Error()
}
