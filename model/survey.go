package model

import (
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
	surveys := []*Survey{}
	rows, err := db.Query("SELECT * FROM surveys")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		survey := &Survey{}
		if err := rows.Scan(&survey.ID, &survey.Title, &survey.Question, &survey.Created, &survey.Updated); err != nil {
			return nil, err
		}
		surveys = append(surveys, survey)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return surveys, nil
}

func GeSurveyById(surveyId string) ([]*Survey, error) {
	surveys := []*Survey{}
	query := "SELECT * FROM surveys"
	if surveyId != "" {
		query = query + " WHERE id = " + surveyId
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		survey := &Survey{}
		if err := rows.Scan(&survey.ID, &survey.Title, &survey.Question, &survey.Created, &survey.Updated); err != nil {
			return nil, err
		}
		surveys = append(surveys, survey)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return surveys, nil
}

func AddSurvey(survey *Survey) (int, error) {
	stmt, err := db.Prepare("INSERT INTO surveys(title,question) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rows, err := stmt.Exec(survey.Title, survey.Question)
	if err != nil {
		return 0, err
	}
	id, err := rows.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func UpdateSurvey(survey *Survey) error {
	stmt, err := db.Prepare("UPDATE surveys SET title = ?, question = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(survey.Title, survey.Question, survey.ID); err != nil {
		return err
	}
	return nil
}

func DeleteSurvey(id string) error {
	stmt, err := db.Prepare("DELETE FROM surveys WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(id); err != nil {
		return err
	}
	return nil
}
