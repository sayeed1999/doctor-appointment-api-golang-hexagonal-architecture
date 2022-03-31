package constants

var ApplicationMessage constant

type constant struct {
	WrongPassword                          string
	EmailNotRegistered                     string
	InvalidEmail                           string
	ItemNotFound                           string
	ItemNotFoundByTheGivenPrimaryKeyOfItem string
}

func init() {
	ApplicationMessage.WrongPassword = "Your password does not match"
	ApplicationMessage.EmailNotRegistered = "This email is not yet registered"
	ApplicationMessage.InvalidEmail = "Email is not a valid email"
	ApplicationMessage.ItemNotFound = "%v not found"
	ApplicationMessage.ItemNotFoundByTheGivenPrimaryKeyOfItem = "%v not found by the given primary key ID of %v"
}
