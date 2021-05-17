package model

type Question struct {
	Title   string
	Answers []string
}

type Header struct {
	Header string
}

func AllQuestions() []Question {
	return []Question{
		{
			Title: "I don´t need to go to work.",
			Answers: []string{
				"I don´t work either",
				"I don´t either",
				"I wasn´t either",
				"Neither need I",
			},
		},
		{
			Title: "They bought a new house. _ _ _ _ _  Too.",
			Answers: []string{
				"You do",
				"I was",
				"I did",
				"I do",
			},
		},
		{
			Title: "We reunite with my high school friends once a year.We do _ _ _ _.",
			Answers: []string{
				"Too",
				"Still",
				"Either",
				"So",
			},
		},
		{
			Title: "I really like Mexican food.",
			Answers: []string{
				"I don´t either",
				"So do I",
				"Neither do I",
			},
		},
		{
			Title: "I really miss Kate. _ _ _ _ _ do I.",
			Answers: []string{
				"Either",
				"So",
				"So Too",
				"Too",
			},
		}, {
			Title: "I haven´t tried Chinese food before. I haven´t _ _ _ _ _.",
			Answers: []string{
				"Yet",
				"Not",
				"Either",
				"Neither",
			},
		},
		{
			Title: "I haven´t received any warnings yet. _ _ _ _ _    _ _ _  my brother.",
			Answers: []string{
				"Neither did",
				"Neither has",
				"Neither will",
				"Neither is",
			},
		},
		{
			Title: "I don´t want to lose my job. _ _ _ Do I",
			Answers: []string{
				"Do",
				"Neither",
				"Either",
				"So",
			},
		},
		{
			Title: "I was so tired after the game. _ _ _ _ _ was I.",
			Answers: []string{
				"Either",
				"Did",
				"So",
				"Did",
			},
		},
		{
			Title: "I don´t like swimming.",
			Answers: []string{
				"I do either",
				"I don´t either",
				"So do I",
				"Neither I do",
			},
		},
	}
}
