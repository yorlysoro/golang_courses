package ransomnote

// canConstruct devuelve true si ransomNote puede construirse usando las letras de magazine.
// Cada letra en magazine solo puede usarse una vez.
func canConstruct(ransomNote string, magazine string) bool {
	// Contador de letras en magazine
	counts := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		counts[magazine[i]]++
	}

	// Verificar si ransomNote puede formarse
	for i := 0; i < len(ransomNote); i++ {
		c := ransomNote[i]
		if counts[c] == 0 {
			return false
		}
		counts[c]--
	}

	return true
}

/*func main() {
	tests := []struct {
		ransomNote string
		magazine   string
		expect     bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"hello", "lloehxlo", true},
		{"abc", "ab", false},
	}

	for _, t := range tests {
		result := canConstruct(t.ransomNote, t.magazine)
		fmt.Printf("ransomNote=%q magazine=%q -> %v (esperado: %v)\n",
			t.ransomNote, t.magazine, result, t.expect)
	}
}*/
