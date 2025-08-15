package contract

const IDKey = "ming:id"

type IDService interface {
	NewID() string
}
