package model

import (
	"net/http"

	"github.com/Proyecto/proyecto_ingles/validate"
	"github.com/labstack/echo/v4"
)

var (
	Msg     *validate.Message
	Warning string
)

func RenderForm(c echo.Context) error {
	sliceForm := AllQuestions()
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
		"error":     Msg.Errors,
		"warning":   Warning,
	})
}
