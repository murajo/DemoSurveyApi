package handler

import (
	"net/http"

	"main/model"

	"github.com/labstack/echo"
)

func GetAnswers(c echo.Context) error {
	answers, err := model.GetAllAnswers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, answers)
}

func GetAnswer(c echo.Context) error {
	surveyId := c.Param("id")
	answer, err := model.GeAnswersBySurveyId(surveyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, answer)
}

func AddAnswer(c echo.Context) error {
	answer := &model.Answer{}
	if err := c.Bind(&answer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id, err := model.AddAnswer(answer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	answer.ID = id
	return c.JSON(http.StatusCreated, answer)
}

func UpdateAnswer(c echo.Context) error {
	answer := &model.Answer{}
	if err := c.Bind(answer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := model.UpdateAnswer(answer); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, answer)
}

func DeleteAnswer(c echo.Context) error {
	id := c.Param("id")
	if err := model.DeleteAnswer(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"success": id})
}
