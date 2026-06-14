package endpoint

type Endpoint struct {
	Alphabet string
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		Alphabet: "abcdefghijklmnopqrstuvwxyz1234567890",
	}
}

func (e *Endpoint) MoveText(shift int, text string) string {
	runeAlphabet := []rune(e.Alphabet)
	alphaLen := len(runeAlphabet)
	idx := make(map[rune]int, alphaLen)
	for i, r := range runeAlphabet {
		idx[r] = i
	}
	var result []rune
	for _, r := range []rune(text) {
		if i, ok := idx[r]; ok {
			ni := (i + shift) % alphaLen
			if ni < 0 {
				ni += alphaLen
			}
			result = append(result, runeAlphabet[ni])
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
