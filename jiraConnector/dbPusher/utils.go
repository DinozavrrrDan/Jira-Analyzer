package dbPusher

import (
	"database/sql"
	"fmt"
	"time"
)

func (databasePusher *DatabasePusher) insertIssue(projectId, authorId, assigneeId int, key, summary, description, Type, priority, status string, createdTime, closedTime, updatedTime time.Time, timeSpent int) error {

	_, err := databasePusher.database.Exec(
		"INSERT INTO issues "+
			"(projectId, authorId, assigneeId,"+
			" key, summary, description, type, priority, status,"+
			" createdTime, closedTime, updatedTime, timeSpent)"+
			" values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
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
		timeSpent)

	if err != nil {
		return fmt.Errorf("insertIssue: %v", err)
	}

	return nil
}

// updateIssue обвновляет данные issue заданного key в таблицк issues
func (databasePusher *DatabasePusher) updateIssue(projectID, authorId, assigneeId int, key, summary, description, Type, priority, status string, createdTime, closedTime, updatedTime time.Time, timespent int) error {

	_, err := databasePusher.database.Exec("UPDATE issues set"+
		" projectid = ?, authorid = ?, assigneeid = ?,"+
		" summary = ?, description = ?, type = ?, priority = ?, status = ?,"+
		" createdtime = ?, closedtime = ?, updatedtime = ?, timespent = ? where key = ?",
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
		key)

	if err != nil {
		return fmt.Errorf("updateIssue: %v", err)
	}

	return nil
}

// getIssueId получает id по ключу задачи из таблицы issues
func (databasePusher *DatabasePusher) getIssueId(issueKey string) (int64, error) {
	var issueID int64
	row := databasePusher.database.QueryRow("SELECT id FROM issues where key = ?", issueKey)

	if err := row.Scan(&issueID); err != nil {
		if err == sql.ErrNoRows {
			return issueID, fmt.Errorf("getIssueId %d: no issue", issueID)
		}
		return issueID, fmt.Errorf("getIssueId %d: %v", issueID, err)
	}

	return issueID, nil
}

// getProjectId получает id по названию проекта из таблицы project
func (databasePusher *DatabasePusher) getProjectId(projectTitle string) (int64, error) {
	var projectId int64
	row := databasePusher.database.QueryRow("SELECT id FROM project where title = ?", projectTitle)

	if err := row.Scan(&projectId); err != nil {
		if err == sql.ErrNoRows {
			return projectId, fmt.Errorf("getProjectId %d: no project", projectId)
		}
		return projectId, fmt.Errorf("getProjectId %d: %v", projectId, err)
	}

	if projectId == 0 {
		result, err := databasePusher.database.Exec("INSERT INTO project (title) VALUES(?)", projectTitle)
		if err != nil {
			return projectId, fmt.Errorf("getProjectId: %v", err.Error())
		}

		projectId, err := result.LastInsertId()
		if err != nil {
			return projectId, fmt.Errorf("getProjectId: %v", err.Error())
		}
	}

	return projectId, nil
}

// getAuthorId получает id по имени автора из таблицы author
func (databasePusher *DatabasePusher) getAuthorId(authorName string) (int64, error) {
	var authorId int64
	row := databasePusher.database.QueryRow("SELECT id FROM author where name = ?", authorName)

	if err := row.Scan(&authorId); err != nil {
		if err == sql.ErrNoRows {
			return authorId, fmt.Errorf("getAuthorId %d: no author", authorId)
		}
		return authorId, fmt.Errorf("getAuthorId %d: %v", authorId, err)
	}

	if authorId == 0 {
		result, err := databasePusher.database.Exec("INSERT INTO author (name) VALUES(?)", authorName)
		if err != nil {
			return authorId, fmt.Errorf("getAuthorId: %v", err.Error())
		}

		authorId, err := result.LastInsertId()
		if err != nil {
			return authorId, fmt.Errorf("getAuthorId: LastInsertId: %v", err.Error())
		}
	}

	return authorId, nil
}

// getAssigneeId получает id по имени assignee из таблицы author
func (databasePusher *DatabasePusher) getAssigneeId(assigneeName string) (int64, error) {
	var assigneeId int64
	row := databasePusher.database.QueryRow("SELECT id FROM author where name = ?", assigneeName)

	if err := row.Scan(&assigneeId); err != nil {
		if err == sql.ErrNoRows {
			return assigneeId, fmt.Errorf("getAssigneeId %d: no assignee", assigneeId)
		}
		return assigneeId, fmt.Errorf("getAssigneeId %d: %v", assigneeId, err)
	}

	if assigneeId == 0 {
		result, err := databasePusher.database.Exec("INSERT INTO author (name) VALUES(?)", assigneeName)
		if err != nil {
			return assigneeId, fmt.Errorf("getAssigneeId: %v", err.Error())
		}

		authorId, err := result.LastInsertId()
		if err != nil {
			return authorId, fmt.Errorf("getAssigneeId: LastInsertId: %v", err.Error())
		}
	}

	return assigneeId, nil
}

// checkIssueExists проверяет наличие issue заданного issueKey
func (databasePusher *DatabasePusher) checkIssueExists(issueKey string) (bool, error) {
	var issueId int64
	row := databasePusher.database.QueryRow("SELECT id FROM issues where key = ?", issueKey)

	if err := row.Scan(&issueId); err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("getAssigneeId %d: no assignee", issueId)
		}
		return false, fmt.Errorf("getAssigneeId %d: %v", issueId, err)
	}

	return issueId != 0, nil
}
