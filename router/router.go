package router

import (
	"main/handler"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.GET("/surveys", handler.GetSurveys)
	e.GET("/surveys/:id", handler.GetSurvey)
	e.POST("/surveys", handler.AddSurvey)
	e.PUT("/surveys", handler.UpdateSurvey)
	e.DELETE("/surveys/:id", handler.DeleteSurvey)

	e.GET("/surveyitems", handler.GetAllSurveyItems)
	e.GET("/surveyitems/:id", handler.GetSurveyItems)
	e.POST("/surveyitems", handler.AddSurveyItem)
	e.PUT("/surveyitems", handler.UpdateSurveyItem)
	e.DELETE("/surveyitems/:id", handler.DeleteSurveyItem)

	e.GET("/answers", handler.GetAnswers)
	e.GET("/answers/:id", handler.GetAnswer)
	e.POST("/answers", handler.AddAnswer)
	e.PUT("/answers", handler.UpdateAnswer)
	e.DELETE("/answers/:id", handler.DeleteAnswer)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE,
			http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}
