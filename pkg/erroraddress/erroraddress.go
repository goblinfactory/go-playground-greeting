package erroraddress

// Field indicates the particular address field
type Field int

// Example of enum below
// ---------------------
// we start first item with iota + 1 so that iota (0) can be used to represent unassigned.

const (
	// PostCode of the address
	PostCode Field = iota + 1
	// City of the address
	City
	// Country of the address
	Country
)

// AddressError is an example of how to do a custom error object, in this case a address error.
type AddressError struct {
	Field   Field
	message string
	err     error
}

func (ae AddressError) Error() string {
	return ae.message
}

// NewAddressError ...
func NewAddressError(field Field, message string, err error) error {
	return AddressError{field, message, err}
}

// we implement Unwrap so that we can use this error to wrap an existing error,
// so that callers can get at the underlying error that we wrapped with a custom type.
// Normally you'd wrap an error with what seems like a basic string only, i.e.a non custom type using fmt.Errorf(" my message: %w", err)

func (ae AddressError) Unwrap() error {
	return ae.err
}
