package Leetcode

func uniqueMorseRepresentations(words []string) int {
	morseTable := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	uniqueMap := make(map[string]int)
	for _, word := range words {
		morseCode := ""
		for _, letter := range word {
			morseCode += morseTable[letter-97]
		}
		uniqueMap[morseCode] += 1
	}
	return len(uniqueMap)
}
