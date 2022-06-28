package Leetcode

func replaceWords(dictionary []string, sentence string) string {
	//将字典放进map中
	dicMap := map[string]string{}
	//继承词最长
	maxLen := 0
	for _, successor := range dictionary {
		dicMap[successor] = successor
		if len(successor) > maxLen {
			maxLen = len(successor)
		}
	}
	//字符串转数组
	var words []string
	for i, j := 0, 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			words = append(words, sentence[j:i])
			j = i + 1
		}
		if i == len(sentence)-1 {
			words = append(words, sentence[j:])
		}
	}
	//循环数组，每一个word根据循环判断是否在map中，如果在map中，用map中successor代替
	for index, word := range words {
		for i := 0; i < len(word) && i <= maxLen; i++ {
			if successor, ok := dicMap[word[:i]]; ok {
				words[index] = successor
				break
			}
		}
	}
	//words 变成字符串
	wordStr := ""
	for _, word := range words {
		wordStr += word
		wordStr += " "
	}
	return wordStr[0 : len(wordStr)-1]
}
