package handler

import (
	"main/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllSurveyItems(c echo.Context) error {
	surveys, err := model.GetAllSurveyItems()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, surveys)
}

func GetSurveyItems(c echo.Context) error {
	surveyId := c.Param("id")
	survey, err := model.GeSurveyItemBySurveyId(surveyId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, survey)
}

func AddSurveyItem(c echo.Context) error {
	surveyItems := &model.SurveyItem{}
	if err := c.Bind(&surveyItems); err != nil {
		return err
	}
	// for _, surveyItem := range surveyItems {
	if err := model.AddSurveyItem(surveyItems); err != nil {
		return err
		// }
	}
	return c.JSON(http.StatusCreated, surveyItems)
}

func UpdateSurveyItem(c echo.Context) error {
	surveyItem := &model.SurveyItem{}
	if err := c.Bind(surveyItem); err != nil {
		return err
	}
	if err := model.UpdateSurveyItem(surveyItem); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, surveyItem)
}

func DeleteSurveyItem(c echo.Context) error {
	id := c.Param("id")
	if err := model.DeleteSurveyItem(id); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
