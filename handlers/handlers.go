package handlers

import (
	"log"
	"net/http"

	"github.com/Proyecto/proyecto_ingles/database"
	"github.com/Proyecto/proyecto_ingles/model"
	"github.com/Proyecto/proyecto_ingles/validate"

	"github.com/labstack/echo/v4"
)

type Answer struct {
	answer []string
}

var (
	sliceStudents model.Estudents
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.gohtml", map[string]interface{}{
		"title": "HOME",
		"msg":   "Hello",
	})
}

func FormHandler(c echo.Context) error {
	sliceForm := model.AllQuestions()
	return c.Render(http.StatusOK, "form.gohtml", map[string]interface{}{
		"question0": sliceForm[0].Title,
		"answers0":  sliceForm[0].Answers,
		"question1": sliceForm[1].Title,
		"answers1":  sliceForm[1].Answers,
		"question2": sliceForm[2].Title,
		"answers2":  sliceForm[2].Answers,
		"question3": sliceForm[3].Title,
		"answers3":  sliceForm[3].Answers,
		"question4": sliceForm[4].Title,
		"answers4":  sliceForm[4].Answers,
		"question5": sliceForm[5].Title,
		"answers5":  sliceForm[5].Answers,
		"question6": sliceForm[6].Title,
		"answers6":  sliceForm[6].Answers,
		"question7": sliceForm[7].Title,
		"answers7":  sliceForm[7].Answers,
		"question8": sliceForm[8].Title,
		"answers8":  sliceForm[8].Answers,
		"question9": sliceForm[9].Title,
		"answers9":  sliceForm[9].Answers,
	})
}

func CreateHandler(c echo.Context) error {
	name, email := c.FormValue("name"), c.FormValue("email")
	sliceQuestions := model.AllQuestions()

	model.Msg = &validate.Message{
		Email:   email,
		Content: name,
	}
	if model.Msg.Validate() == false {
		return model.RenderForm(c)
	}
	e := &model.Estudent{
		Name:  name,
		Email: email,
	}
	answer1 := &validate.Section{
		Title:         sliceQuestions[0].Title,
		ReplySent:     c.FormValue("answer0"),
		CorrectAnswer: sliceQuestions[0].Answers[1],
	}
	valueNull, one := answer1.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer1.Qualification)
	answer2 := &validate.Section{
		Title:         sliceQuestions[1].Title,
		ReplySent:     c.FormValue("answer1"),
		CorrectAnswer: sliceQuestions[1].Answers[2],
	}
	valueNull, two := answer2.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer2.Qualification)
	answer3 := &validate.Section{
		Title:         sliceQuestions[2].Title,
		ReplySent:     c.FormValue("list1"),
		CorrectAnswer: sliceQuestions[2].Answers[0],
	}
	valueNull, tree := answer3.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer3.Qualification)
	answer4 := &validate.Section{
		Title:         sliceQuestions[3].Title,
		ReplySent:     c.FormValue("list"),
		CorrectAnswer: sliceQuestions[3].Answers[1],
	}
	valueNull, four := answer4.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer4.Qualification)
	answer5 := &validate.Section{
		Title:         sliceQuestions[4].Title,
		ReplySent:     c.FormValue("answer4"),
		CorrectAnswer: sliceQuestions[4].Answers[1],
	}
	valueNull, five := answer5.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer5.Qualification)
	answer6 := &validate.Section{
		Title:         sliceQuestions[5].Title,
		ReplySent:     c.FormValue("answer5"),
		CorrectAnswer: sliceQuestions[5].Answers[2],
	}
	valueNull, six := answer6.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer6.Qualification)
	answer7 := &validate.Section{
		Title:         sliceQuestions[6].Title,
		ReplySent:     c.FormValue("answer6"),
		CorrectAnswer: sliceQuestions[6].Answers[1],
	}
	valueNull, seven := answer7.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer7.Qualification)
	answer8 := &validate.Section{
		Title:         sliceQuestions[7].Title,
		ReplySent:     c.FormValue("list7"),
		CorrectAnswer: sliceQuestions[7].Answers[1],
	}
	valueNull, eight := answer8.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer8.Qualification)
	answer9 := &validate.Section{
		Title:         sliceQuestions[8].Title,
		ReplySent:     c.FormValue("list8"),
		CorrectAnswer: sliceQuestions[8].Answers[2],
	}
	valueNull, nine := answer9.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer9.Qualification)
	answer10 := &validate.Section{
		Title:         sliceQuestions[9].Title,
		ReplySent:     c.FormValue("answer9"),
		CorrectAnswer: sliceQuestions[9].Answers[1],
	}
	valueNull, ten := answer10.SectionQuestion()
	if valueNull {
		model.Warning = "Por favor contesta todas las preguntas"
		return model.RenderForm(c)
	}
	e.Percentage = append(e.Percentage, answer10.Qualification)
	total := e.TotalQualification()

	err := database.CreateEstudent(e)
	if err != nil {
		log.Println("Ocurrio un error al insertar los datos")
	}

	return c.Render(http.StatusOK, "answers.gohtml", map[string]interface{}{
		"Name":          name,
		"Email":         email,
		"Qualification": total,
		"answer1":       one,
		"answer2":       two,
		"answer3":       tree,
		"answer4":       four,
		"answer5":       five,
		"answer6":       six,
		"answer7":       seven,
		"answer8":       eight,
		"answer9":       nine,
		"answer10":      ten,
	})
}

func GetAllHandler(c echo.Context) error {
	es, err := database.GetAllt()
	if err != nil {
		log.Print("Ocurrio un error al obtener todos los estudiantes")
	}
	return c.Render(http.StatusOK, "list.gohtml", map[string]interface{}{
		"list": es,
	})
}
