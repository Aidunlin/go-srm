package db

import (
	"fmt"
	"strings"

	"github.com/Aidunlin/go-srm/model"
	"github.com/go-mysql-org/go-mysql/client"
)

func AuthenticateAdmin(admin model.AdminRecord) (model.AdminRecord, bool) {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return admin, false
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %v WHERE email = ? AND password = ? LIMIT 1", adminTableName)
	result, execErr := db.Execute(sql, admin.Email, admin.Password)
	if execErr != nil || len(result.Values) < 1 {
		return admin, false
	}

	return model.GetAdminRecordFromResult(result, 0), true
}

func SelectAdminWithId(id int) (model.AdminRecord, bool) {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return model.AdminRecord{}, false
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %v WHERE id = ? LIMIT 1", adminTableName)
	result, execErr := db.Execute(sql, id)
	if execErr != nil || len(result.Values) < 1 {
		return model.AdminRecord{}, false
	}

	return model.GetAdminRecordFromResult(result, 0), true
}

func SelectAdminWithEmail(email string) (model.AdminRecord, bool) {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return model.AdminRecord{}, false
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %v WHERE email = ? LIMIT 1", adminTableName)
	result, execErr := db.Execute(sql, email)
	if execErr != nil || len(result.Values) < 1 {
		return model.AdminRecord{}, false
	}

	return model.GetAdminRecordFromResult(result, 0), true
}

func InsertAdmin(admin model.AdminRecord) bool {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	paramsMap := admin.ToMap()

	names := []string{}
	placeholders := []string{}
	values := make([]interface{}, len(paramsMap))

	i := 0
	for name, value := range paramsMap {
		names = append(names, name)
		placeholders = append(placeholders, "?")
		values[i] = value
		i++
	}

	namesString := strings.Join(names, ", ")
	placeholdersString := strings.Join(placeholders, ", ")

	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", adminTableName, namesString, placeholdersString)
	_, execErr := db.Execute(sql, values...)
	return execErr == nil
}
