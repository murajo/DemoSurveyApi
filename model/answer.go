package model

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Answer struct {
	ID           int    `json:"id"`
	SurveyId     int    `json:"surveyId"`
	SurveyItemId int    `json:"surveyItemId"`
	Created      string `json:"created"`
	Updated      string `json:"updated"`
}

func GetAllAnswers() ([]*Answer, error) {
	rows, err := db.Query("SELECT * FROM answers")
	if err != nil {
		return nil, fmt.Errorf("failed to get all answer: %w", err)
	}
	defer rows.Close()

	answers := []*Answer{}
	for rows.Next() {
		answer := Answer{}
		err = rows.Scan(&answer.ID, &answer.SurveyId, &answer.SurveyItemId, &answer.Created, &answer.Updated)
		if err != nil {
			return nil, fmt.Errorf("failed to scan answer: %w", err)
		}
		answers = append(answers, &answer)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all answers: %w", err)
	}
	return answers, nil
}

func GeAnswersBySurveyId(surveyId string) ([]*Answer, error) {
	rows, err := db.Query("SELECT * FROM answers WHERE survey_id = ?", surveyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	answers := []*Answer{}
	for rows.Next() {
		answer := &Answer{}
		err = rows.Scan(&answer.ID, &answer.SurveyId, &answer.SurveyItemId, &answer.Created, &answer.Updated)
		if err != nil {
			return nil, fmt.Errorf("failed to scan answer: %w", err)
		}
		answers = append(answers, answer)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get answers by survey_id: %w", err)
	}
	return answers, nil
}

func AddAnswer(answer *Answer) (int, error) {
	result, err := db.Exec("INSERT INTO answers(survey_id,survey_item_id) VALUES(?,?)", answer.SurveyId, answer.SurveyItemId)
	if err != nil {
		return 0, fmt.Errorf("failed to add answer: %w", err)
	}
	answerId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get inserted answer id: %w", err)
	}
	return int(answerId), nil
}

func UpdateAnswer(answer *Answer) error {
	_, err := db.Exec("UPDATE answers SET survey_id = ?, survey_item_id = ? WHERE id = ?", answer.SurveyId, answer.SurveyItemId, answer.ID)
	if err != nil {
		return fmt.Errorf("failed to update answer: %w", err)
	}
	return nil
}

func DeleteAnswer(answerId string) error {
	if answerId == "" {
		return errors.New("empty survey id")
	}
	result, err := db.Exec("DELETE FROM answers WHERE id = ?", answerId)
	if err != nil {
		return fmt.Errorf("failed to delete answer: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete answer: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no answer found with id %s", answerId)
	}
	return nil
}
