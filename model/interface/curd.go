package _interface

type CUDR interface {
	Create() (msg string, err error)
	Retrieve() (msg string, err error)
	Update() (msg string, err error)
	Delect() (msg string, err error)
}

