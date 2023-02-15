package model

import (
	"errors"
	"fmt"

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
	rows, err := db.Query("SELECT * FROM survey_items")
	if err != nil {
		return nil, fmt.Errorf("failed to get all survey_items: %w", err)
	}
	defer rows.Close()

	surveyItems := []*SurveyItem{}
	for rows.Next() {
		surveyItem := SurveyItem{}
		err = rows.Scan(&surveyItem.ID, &surveyItem.Text, &surveyItem.SurveyId, &surveyItem.Created, &surveyItem.Updated)
		if err != nil {
			return nil, fmt.Errorf("failed to scan survey_item: %w", err)
		}
		surveyItems = append(surveyItems, &surveyItem)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all survey_items: %w", err)
	}
	return surveyItems, nil
}

func GetSurveyItemBySurveyId(surveyId string) ([]*SurveyItem, error) {
	rows, err := db.Query("SELECT * FROM survey_items WHERE survey_id = ?", surveyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	surveyItems := []*SurveyItem{}
	for rows.Next() {
		surveyItem := &SurveyItem{}
		err = rows.Scan(&surveyItem.ID, &surveyItem.Text, &surveyItem.SurveyId, &surveyItem.Created, &surveyItem.Updated)
		if err != nil {
			return nil, fmt.Errorf("failed to scan survey_item: %w", err)
		}
		surveyItems = append(surveyItems, surveyItem)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get survey_items by survey_id: %w", err)
	}
	return surveyItems, nil
}

func AddSurveyItem(surveyItem *SurveyItem) (int, error) {
	res, err := db.Exec("INSERT INTO survey_items(text,survey_id) VALUES(?,?)", surveyItem.Text, surveyItem.SurveyId)
	if err != nil {
		return 0, fmt.Errorf("failed to add survey_item: %w", err)
	}
	surveyItemId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get insert survey_item id: %w", err)
	}
	return int(surveyItemId), nil
}

func UpdateSurveyItem(surveyItem *SurveyItem) error {
	_, err := db.Exec("UPDATE survey_items SET text = ? WHERE id = ?", surveyItem.Text, surveyItem.ID)
	if err != nil {
		return fmt.Errorf("failed to update survey_item: %w", err)
	}
	return nil
}

func DeleteSurveyItem(surveyItemId string) error {
	if surveyItemId == "" {
		return errors.New("empty survey id")
	}
	result, err := db.Exec("DELETE FROM survey_items WHERE id=?", surveyItemId)
	if err != nil {
		return fmt.Errorf("failed to delete survey_item: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to delete survey_item: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no survey_item found with id %s", surveyItemId)
	}
	return nil
}
