package internal

// ThisIsInternalOnly can only be seen by packages in this module but not externally as part of the API
func ThisIsInternalOnly() int {
	return 42
}
