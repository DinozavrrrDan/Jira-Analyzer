package service

import (
	"fmt"
	"time"
)

func (databasePusher *DatabasePusherService) insertInfoIntoIssues(projectId, authorId, assigneeId int, key, summary, description, Type, priority, status string, createdTime, closedTime, updatedTime time.Time, timeSpent int) error {

	err := databasePusher.database.QueryRow("INSERT INTO issues (projectId, authorId, assigneeId, key, summary, description, type, priority, status, createdTime, closedTime, updatedTime, timeSpent) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
		projectId,
		authorId,
		assigneeId,
		key,
		summary,
		description,
		Type,
		priority,
		status,
		createdTime,
		closedTime,
		updatedTime,
		timeSpent).Err()

	if err != nil {
		return fmt.Errorf("ERROR: %v", err.Error())
	}

	return nil
}

// updateIssue обвновляет данные issue заданного key в таблицк issues
func (databasePusher *DatabasePusherService) updateIssue(projectID, authorId, assigneeId int, key, summary, description, Type, priority, status string, createdTime, closedTime, updatedTime time.Time, timespent int) error {
	err := databasePusher.database.QueryRow("UPDATE issues set projectid = $1, authorid = $2, assigneeid = $3, summary = $4, description = $5, type = $6, priority = $7, status = $8, createdtime = $9, closedtime = $10, updatedtime = $11, timespent = $12 where key = $13",
		projectID,
		authorId,
		assigneeId,
		summary,
		description,
		Type,
		priority,
		status,
		createdTime,
		closedTime,
		updatedTime,
		timespent,
		key).Err()

	if err != nil {
		return fmt.Errorf("ERROR: %v", err.Error())
	}

	return nil
}

// getIssueId получает id по ключу задачи из таблицы issues
func (databasePusher *DatabasePusherService) getIssueId(issueKey string) (int, error) {
	var issueID int
	_ = databasePusher.database.QueryRow("SELECT id FROM issues where key = $1", issueKey).Scan(&issueID)

	return issueID, nil
}

// getProjectId получает id по названию проекта из таблицы project
func (databasePusher *DatabasePusherService) getProjectId(projectTitle string) (int, error) {
	var projectId int
	_ = databasePusher.database.QueryRow("SELECT id FROM project where title = $1", projectTitle).Scan(&projectId)

	if projectId == 0 {
		err := databasePusher.database.QueryRow("INSERT INTO project (title) VALUES($1) RETURNING id", projectTitle).
			Scan(&projectId)
		if err != nil {
			return projectId, fmt.Errorf("ERROR: %v", err.Error())
		}
	}

	return projectId, nil
}

// getAuthorId получает id по имени автора из таблицы author
func (databasePusher *DatabasePusherService) getAuthorId(authorName string) (int, error) {
	var authorId int
	_ = databasePusher.database.QueryRow("SELECT id FROM author where name = $1", authorName).Scan(&authorId)

	if authorId == 0 {
		err := databasePusher.database.QueryRow("INSERT INTO author (name) VALUES($1) RETURNING id", authorName).
			Scan(&authorId)

		if err != nil {
			return authorId, fmt.Errorf("ERROR: %v", err.Error())
		}
	}

	return authorId, nil
}

// getAssigneeId получает id по имени assignee из таблицы author
func (databasePusher *DatabasePusherService) getAssigneeId(assignee string) (int, error) {
	var assigneeId int
	_ = databasePusher.database.QueryRow("SELECT id FROM author where name = $1", assignee).
		Scan(&assigneeId)

	if assigneeId == 0 {
		err := databasePusher.database.QueryRow("INSERT INTO author (name) VALUES($1) RETURNING id",
			assignee).Scan(&assigneeId)
		if err != nil {
			return assigneeId, fmt.Errorf("ERROR: %v", err.Error())
		}
	}

	return assigneeId, nil
}

// checkIssueExists проверяет наличие issue заданного issueKey
func (databasePusher *DatabasePusherService) checkIssueExists(issueKey string) bool {
	var issueId int

	_ = databasePusher.database.QueryRow("SELECT id FROM issues where key = $1", issueKey).Scan(&issueId)

	return issueId != 0
}