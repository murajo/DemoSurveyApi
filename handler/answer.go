package handler

import (
	"net/http"

	"main/model"

	"github.com/labstack/echo"
)

func GetAnswers(c echo.Context) error {
	answers, err := model.GetAllAnswers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, answers)
}

func GetAnswer(c echo.Context) error {
	surveyId := c.Param("id")
	answer, err := model.GeAnswerBySurveyId(surveyId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, answer)
}

func AddAnswer(c echo.Context) error {
	answer := &model.Answer{}
	if err := c.Bind(&answer); err != nil {
		return err
	}
	if err := model.AddAnswer(answer); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, answer)
}

func UpdateAnswer(c echo.Context) error {
	answer := &model.Answer{}
	if err := c.Bind(answer); err != nil {
		return err
	}
	if err := model.UpdateAnswer(answer); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, answer)
}

func DeleteAnswer(c echo.Context) error {
	id := c.Param("id")
	if err := model.DeleteAnswer(id); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
