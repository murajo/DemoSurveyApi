package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type SurveyItem struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	SurveyId int    `json:"surveyId"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

func GetAllSurveyItems() ([]*SurveyItem, error) {
	surveyItems := []*SurveyItem{}
	rows, err := db.Query("SELECT * FROM survey_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		surveyItem := &SurveyItem{}
		if err := rows.Scan(&surveyItem.ID, &surveyItem.Text, &surveyItem.SurveyId, &surveyItem.Created, &surveyItem.Updated); err != nil {
			return nil, err
		}
		surveyItems = append(surveyItems, surveyItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return surveyItems, nil
}

func GeSurveyItemBySurveyId(surveyId string) ([]*SurveyItem, error) {
	surveyItems := []*SurveyItem{}
	query := "SELECT * FROM survey_items"
	if surveyId != "" {
		query = query + " WHERE survey_id = " + surveyId
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		surveyItem := &SurveyItem{}
		if err := rows.Scan(&surveyItem.ID, &surveyItem.Text, &surveyItem.SurveyId, &surveyItem.Created, &surveyItem.Updated); err != nil {
			return nil, err
		}
		surveyItems = append(surveyItems, surveyItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return surveyItems, nil
}

func AddSurveyItem(surveyItem *SurveyItem) error {
	stmt, err := db.Prepare("INSERT INTO survey_items(text,survey_id) VALUES(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(surveyItem.Text, surveyItem.SurveyId); err != nil {
		return err
	}
	return nil
}

func UpdateSurveyItem(surveyItem *SurveyItem) error {
	stmt, err := db.Prepare("UPDATE survey_items SET text = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(surveyItem.Text, surveyItem.ID); err != nil {
		return err
	}
	return nil
}

func DeleteSurveyItem(id string) error {
	stmt, err := db.Prepare("DELETE FROM survey_items WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(id); err != nil {
		return err
	}
	return nil
}
