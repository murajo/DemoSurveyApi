package model

import (
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
	answers := []*Answer{}
	rows, err := db.Query("SELECT * FROM answers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		answer := &Answer{}
		if err := rows.Scan(&answer.ID, &answer.SurveyId, &answer.SurveyItemId, &answer.Created, &answer.Updated); err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return answers, nil
}

func GeAnswerBySurveyId(surveyId string) ([]*Answer, error) {
	answers := []*Answer{}
	query := "SELECT * FROM answers"
	if surveyId != "" {
		query = query + " WHERE survey_id = " + surveyId
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		answer := &Answer{}
		if err := rows.Scan(&answer.ID, &answer.SurveyId, &answer.SurveyItemId, &answer.Created, &answer.Updated); err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return answers, nil
}

func AddAnswer(answer *Answer) error {
	stmt, err := db.Prepare("INSERT INTO answers(survey_id,survey_item_id) VALUES(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(answer.SurveyId, answer.SurveyItemId); err != nil {
		return err
	}
	return nil
}

func UpdateAnswer(answer *Answer) error {
	stmt, err := db.Prepare("UPDATE answers SET title = ?, question = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(answer.SurveyId, answer.SurveyItemId, answer.ID); err != nil {
		return err
	}
	return nil
}

func DeleteAnswer(id string) error {
	stmt, err := db.Prepare("DELETE FROM answers WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(id); err != nil {
		return err
	}
	return nil
}
