package go_gematria

import (
	"fmt"
	"regexp"
	"strings"
)

func characterScores(str string) (uint, uint, uint, error) {
	if len(str) != 1 {
		return 0, 0, 0, fmt.Errorf("string must only contain 1 character")
	}
	j, e, s := jewishCodes[str], englishCodes[str], simpleCodes[str]
	return j, e, s, nil
}

func NewGematria(data string) (Gematria, error) {
	re, err := regexp.Compile(`[^a-zA-Z\d.\s]`)
	if err != nil {
		return Gematria{}, fmt.Errorf("failed to compile gematria regex: %w", err)
	}
	data = re.ReplaceAllString(data, "")
	data = strings.TrimLeft(data, "")
	dataBytes := []byte(data)
	var letters []Gematria
	for i := 0; i < len(dataBytes); i++ {
		wat := strings.ToUpper(string(dataBytes[i]))
		if len(wat) == 1 && wat != "" && wat != " " {
			j, e, s, err := characterScores(wat)
			if err != nil {
				return Gematria{}, err
			}
			letters = append(letters, Gematria{
				Jewish:  j,
				English: e,
				Simple:  s,
			})
		}
	}
	var jf, ef, sf uint
	for _, gs := range letters {
		jf += gs.Jewish
		ef += gs.English
		sf += gs.Simple
	}
	return Gematria{Jewish: jf, English: ef, Simple: sf}, nil
}

func (s Gematria) String() string {
	output := "\t"
	output += fmt.Sprintf("%s = %d \t", "Jewish", s.Jewish)
	output += fmt.Sprintf("%s = %d \t", "English", s.English)
	output += fmt.Sprintf("%s = %d \t", "Simple", s.Simple)
	return output
}
