package handler

import (
	"main/model"

	"net/http"

	"github.com/labstack/echo"
)

func GetAllSurveyItems(c echo.Context) error {
	surveys, err := model.GetAllSurveyItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, surveys)
}

func GetSurveyItems(c echo.Context) error {
	surveyId := c.Param("id")
	survey, err := model.GetSurveyItemBySurveyId(surveyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, survey)
}

func AddSurveyItem(c echo.Context) error {
	surveyItems := &model.SurveyItem{}
	if err := c.Bind(&surveyItems); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id, err := model.AddSurveyItem(surveyItems)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	surveyItems.ID = id
	return c.JSON(http.StatusCreated, surveyItems)
}

func UpdateSurveyItem(c echo.Context) error {
	surveyItem := &model.SurveyItem{}
	if err := c.Bind(surveyItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := model.UpdateSurveyItem(surveyItem); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, surveyItem)
}

func DeleteSurveyItem(c echo.Context) error {
	id := c.Param("id")
	if err := model.DeleteSurveyItem(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"success": id})
}
