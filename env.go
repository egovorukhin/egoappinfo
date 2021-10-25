package egoappinfo

type Environment string

const (
	None Environment = "none"
	Test             = "test"
	Prod             = "production"
)

func (e Environment) String() string {
	return string(e)
}

func IsTest() bool {
	if application.Environment == Test {
		return true
	}
	return false
}

func IsProd() bool {
	if application.Environment == Prod {
		return true
	}
	return false
}
