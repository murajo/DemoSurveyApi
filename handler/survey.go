package handler

import (
	"net/http"

	"main/model"

	"github.com/labstack/echo"
)

func GetSurveys(c echo.Context) error {
	surveys, err := model.GetAllSurveys()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, surveys)
}

func GetSurvey(c echo.Context) error {
	id := c.Param("id")
	survey, err := model.GeSurveyById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, survey)
}

func AddSurvey(c echo.Context) error {
	survey := &model.Survey{}
	if err := c.Bind(survey); err != nil {
		return err
	}
	surveyId, err := model.AddSurvey(survey)
	if err != nil {
		return err
	}
	survey.ID = surveyId
	return c.JSON(http.StatusCreated, survey)
}

func UpdateSurvey(c echo.Context) error {
	survey := &model.Survey{}
	if err := c.Bind(survey); err != nil {
		return err
	}
	if err := model.UpdateSurvey(survey); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, survey)
}

func DeleteSurvey(c echo.Context) error {
	id := c.Param("id")
	if err := model.DeleteSurvey(id); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
