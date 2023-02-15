package model

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Survey struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Question string `json:"question"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

func GetAllSurveys() ([]*Survey, error) {
	rows, err := db.Query("SELECT * FROM surveys")
	if err != nil {
		return nil, fmt.Errorf("failed to get all survey: %w", err)
	}
	defer rows.Close()

	surveys := []*Survey{}
	for rows.Next() {
		survey := Survey{}
		err = rows.Scan(&survey.ID, &survey.Title, &survey.Question, &survey.Created, &survey.Updated)
		if err != nil {
			return nil, fmt.Errorf("failed to scan survey: %w", err)
		}
		surveys = append(surveys, &survey)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all surveys: %w", err)
	}
	return surveys, nil
}

func GetSurveyById(surveyId string) (*Survey, error) {
	if surveyId == "" {
		return nil, errors.New("empty survey id")
	}
	survey := Survey{}
	err := db.QueryRow("SELECT * FROM surveys WHERE id=?", surveyId).
		Scan(&survey.ID, &survey.Title, &survey.Question, &survey.Created, &survey.Updated)
	if err != nil {
		return nil, fmt.Errorf("failed to get survey: %w", err)
	}
	return &survey, nil
}

func AddSurvey(survey *Survey) (int, error) {
	result, err := db.Exec("INSERT INTO surveys(title, question) VALUES(?,?)", survey.Title, survey.Question)
	if err != nil {
		return 0, fmt.Errorf("failed to add survey: %w", err)
	}
	surveyId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get insert survey id: %w", err)
	}
	return int(surveyId), nil
}

func UpdateSurvey(survey *Survey) error {
	result, err := db.Exec("UPDATE surveys SET title=?, question=? WHERE id=?", survey.Title, survey.Question, survey.ID)
	if err != nil {
		return fmt.Errorf("failed to update survey: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to update survey: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no update with id %d", survey.ID)
	}
	return nil
}

func DeleteSurvey(surveyId string) error {
	if surveyId == "" {
		return errors.New("empty survey id")
	}
	result, err := db.Exec("DELETE FROM surveys WHERE id=?", surveyId)
	if err != nil {
		return fmt.Errorf("failed to delete survey: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete survey: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no survey found with id %s", surveyId)
	}
	return nil
}
