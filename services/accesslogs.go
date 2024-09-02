package services

// func (a *AppDrawingAccessLogs) RecordAccess(userID string, drawingID string) error {
// 	query := `
// 		INSERT INTO DrawingAccessLog (ID, UserID, DrawingID, AccessedAt)
// 		VALUES (?, ?, ?, ?)
// 	`

// 	id := generateID()
// 	accessedAt := time.Now().Unix()

// 	_, err := a.DB.Exec(query, id, userID, drawingID, accessedAt)
// 	if err != nil {
// 		return fmt.Errorf("failed to record access: %w", err)
// 	}

// 	return nil
// }

// func (a *AppDrawingAccessLogs) GetUserLogs(userID string, count int) ([]AccessLog, error) {
// 	query := `
// 		SELECT ID, UserID, DrawingID, AccessedAt
// 		FROM DrawingAccessLog
// 		WHERE UserID = ?
// 		ORDER BY AccessedAt DESC
// 		LIMIT ?
// 	`

// 	rows, err := a.DB.Query(query, userID, count)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to query user logs: %w", err)
// 	}
// 	defer rows.Close()

// 	var logs []AccessLog
// 	for rows.Next() {
// 		var log AccessLog
// 		var accessedAt int64
// 		err := rows.Scan(&log.ID, &log.UserID, &log.DrawingID, &accessedAt)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to scan row: %w", err)
// 		}
// 		log.AccessedAt = time.Unix(accessedAt, 0)
// 		logs = append(logs, log)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("error iterating rows: %w", err)
// 	}

// 	return logs, nil
// }
