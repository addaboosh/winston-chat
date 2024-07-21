package store

type ResourceNotFoundError struct {}

func (e *ResourceNotFoundError) Error() string {
	return "resource not found"
}

type UUIDCreateError struct {}

func (e *UUIDCreateError) Error() string {
	return "eeror creating new UUID"
}
