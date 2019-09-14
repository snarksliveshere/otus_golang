package search

import (
	"regexp"
	"sort"
	"strings"
)

type word struct {
	word string
	freq int
}

func CommonWords(str string, top int) []string {
	s := normalizeStr(str)
	dict := getDicts(s)
	getSortSlice(dict)
	return getWordsByNum(top, dict)
}

func getWordsByNum(num int, word []word) []string {
	w := make([]string, 0, num)
	if len(word) < num {
		num = len(word)
	}
	for i := 0; i < num; i++ {
		w = append(w, word[i].word)
	}

	return w
}

func getDicts(s string) (w []word) {
	m := map[string]word{}
	for _, v := range strings.Fields(s) {
		if val, ok := m[v]; ok {
			val.freq += 1
			m[v] = val
		} else {
			m[v] = word{
				word: v,
				freq: 1,
			}
		}
	}
	for _, v := range m {
		w = append(w, v)
	}

	return
}

func getSortSlice(word []word) {
	sort.Slice(word, func(i, j int) bool {
		return word[i].freq > word[j].freq
	})
}

func normalizeStr(str string) (s string) {
	excessSymbols := `\.|,|;|:|\n|\r|\*|`
	extraSpaces := `\s{2,}|\t+`
	reg := regexp.MustCompile(excessSymbols)
	re := regexp.MustCompile(extraSpaces)
	s = reg.ReplaceAllString(str, "")
	s = re.ReplaceAllString(s, " ")
	s = strings.ToLower(s)

	return
}
