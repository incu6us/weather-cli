package client

type Result interface {
	ClientName() string
	Data() any
}
