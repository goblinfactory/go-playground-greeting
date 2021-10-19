package dependencyinjection

// SimpleDataStore ...
type SimpleDataStore struct {
	userData map[string]string
}

// UserNameForID ...
func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// NewSimpleDataStore ...
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"19984": "Simon Stewart",
			"2":     "Sally",
			"3":     "Catherine",
			"4":     "Markus",
		},
	}
}
