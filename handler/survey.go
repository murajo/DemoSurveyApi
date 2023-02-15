package handler

import (
	"net/http"

	"main/model"

	"github.com/labstack/echo"
)

func GetSurveys(c echo.Context) error {
	surveys, err := model.GetAllSurveys()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, surveys)
}

func GetSurvey(c echo.Context) error {
	id := c.Param("id")
	survey, err := model.GetSurveyById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, survey)
}

func AddSurvey(c echo.Context) error {
	survey := &model.Survey{}
	if err := c.Bind(survey); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id, err := model.AddSurvey(survey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	survey.ID = id
	return c.JSON(http.StatusCreated, survey)
}

func UpdateSurvey(c echo.Context) error {
	survey := &model.Survey{}
	if err := c.Bind(survey); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := model.UpdateSurvey(survey); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, survey)
}

func DeleteSurvey(c echo.Context) error {
	id := c.Param("id")
	if err := model.DeleteSurvey(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"success": id})
}
