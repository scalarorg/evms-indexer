package db

func (db *DatabaseAdapter) SaveSingleValue(value any) error {
	result := db.PostgresClient.Save(value)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DatabaseAdapter) SaveMultipleValues(values []any) error {
	result := db.PostgresClient.Save(values)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
