package endpoint

type Endpoint struct {
	Alphabet string
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		Alphabet: "abcdefghijklmnopqrstuvwxyz1234567890",
	}
}
func (e *Endpoint) MoveText(shift int) string {
	runeAlphabet := []rune(e.Alphabet)
	var result []rune
	for i := 0; i < shift; i++ {
		result = append(result, runeAlphabet[i+shift])
	}
	return string(result)
}
