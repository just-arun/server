package userenum;

type UserEnum = int;

const (
	// Disabled refers to user disabled
	Disabled UserEnum = iota
	// Active refers to user disabled
	Active
	// ForgotPassword refers to user disabled
	ForgotPassword
	// UpdatePassword refers to user disabled
	UpdatePassword
)