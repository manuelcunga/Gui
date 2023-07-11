package contract

type IGPTGenerator interface {
	GenerateText(query string) (string, error)
}
