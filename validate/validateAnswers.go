package validate

import (
	"strings"
)

type Section struct {
	Title         string
	ReplySent     string
	CorrectAnswer string
	Qualification float32
	Section       map[string]interface{}
}

func normalize(letter string) string {
	letter = strings.TrimSpace(letter)
	letter = strings.ToLower(letter)

	return letter
}

func (s *Section) SectionQuestion() (bool, map[string]interface{}) {
	s.Section = make(map[string]interface{})

	s.Section["title"] = s.Title
	s.Section["replySent"] = s.ReplySent
	s.Section["correctAnswer"] = s.CorrectAnswer

	s.ReplySent = normalize(s.ReplySent)
	s.CorrectAnswer = normalize(s.CorrectAnswer)

	if strings.Compare(s.ReplySent, s.CorrectAnswer) == 0 {
		s.Section["msgCorrect"] = "Correct answer"
		s.Qualification = 10.0
		s.Section["qualification"] = s.Qualification
	} else {
		s.Section["msgIncorrect"] = "Wrong answer"
		s.Qualification = 0.0
		s.Section["qualification"] = s.Qualification
	}
	return len(s.ReplySent) == 0, s.Section
}
