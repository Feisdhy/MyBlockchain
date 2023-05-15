package trie

var indices = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "[17]"}

type node interface {
}

type (
	fullNode struct {
		Children [17]node
	}
	shortNode struct {
		Key []byte
		Val node
	}
	hashNode  []byte
	valueNode []byte
)
