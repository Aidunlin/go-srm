package db

import (
	"fmt"
	"math"
	"strings"

	"github.com/Aidunlin/go-srm/model"
	"github.com/go-mysql-org/go-mysql/client"
)

const address = "localhost:3306"
const user = "root"
const password = ""
const databaseName = "srm"
const tableName = "student"
const paginateLimit = 10

func GetTotalPages(totalResults int64) int {
	return int(math.Ceil(float64(totalResults) / float64(paginateLimit)))
}

func SelectStudents(params model.StudentTableParams) (int64, []model.StudentRecord) {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return 0, nil
	}
	defer db.Close()

	var whereSql string
	if len(params.Filter) > 0 {
		whereSql = whereLastNameFilter(params.Filter)
	} else if len(params.Search) > 0 {
		whereSql = whereBasicSearch(params.Search)
	}
	return selectStudentsWithWhere(db, params, whereSql)
}

func SelectStudentsWithForm(params model.StudentTableParams, form model.AdvancedSearchForm) (int64, []model.StudentRecord) {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return 0, nil
	}
	defer db.Close()
	return selectStudentsWithWhere(db, params, whereAdvancedSearch(form))
}

func SelectStudent(id int) (model.StudentRecord, bool) {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return model.StudentRecord{}, false
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %v WHERE id = ? LIMIT 1", tableName)
	result, execErr := db.Execute(sql, id)
	if execErr != nil || len(result.Values) < 1 {
		return model.StudentRecord{}, false
	}

	return model.GetStudentRecordFromResult(result, 0), true
}

func InsertStudent(record model.StudentRecord) bool {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	paramsMap := record.ToMap()

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

	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", tableName, namesString, placeholdersString)
	_, execErr := db.Execute(sql, values...)
	return execErr == nil
}

func UpdateStudent(id int, record model.StudentRecord) bool {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	paramsMap := record.ToMap()

	setters := []string{}
	values := make([]interface{}, len(paramsMap)+1)

	i := 0
	for name, value := range paramsMap {
		setters = append(setters, fmt.Sprintf("%v = ?", name))
		values[i] = value
		i++
	}

	values[i] = id
	settersString := strings.Join(setters, ", ")

	sql := fmt.Sprintf("UPDATE %v SET %v WHERE id = ?", tableName, settersString)
	_, execErr := db.Execute(sql, values...)
	return execErr == nil
}

func DeleteStudent(id int) bool {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	sql := fmt.Sprintf("DELETE FROM %v WHERE id = ?", tableName)
	result, execErr := db.Execute(sql, id)
	return execErr == nil && result.AffectedRows == 1
}

func whereLastNameFilter(filter string) string {
	if len(filter) != 1 {
		return ""
	}
	return fmt.Sprintf("WHERE last_name LIKE '%v%%'", filter)
}

func whereBasicSearch(search string) string {
	var conditions []string
	for _, column := range model.GetStudentColumns() {
		if !column.BasicSearch {
			continue
		}

		if column.BasicSearchExact {
			conditions = append(conditions, fmt.Sprintf("%v LIKE '%v'", column.Name, search))
		} else {
			conditions = append(conditions, fmt.Sprintf("%v LIKE '%%%v%%'", column.Name, search))
		}
	}

	var sql string
	if len(conditions) > 0 {
		sql = fmt.Sprintf("WHERE %v", strings.Join(conditions, " OR "))
	}
	return sql
}

func whereAdvancedSearch(form model.AdvancedSearchForm) string {
	var conditions []string
	for name, value := range form.ToMap() {
		if len(value) == 0 {
			continue
		}

		column := model.GetStudentColumn(name)
		if len(column.Name) == 0 {
			continue
		}

		if column.AdvancedSearchExact {
			conditions = append(conditions, fmt.Sprintf("%v = %v", name, value))
		} else {
			conditions = append(conditions, fmt.Sprintf("%v LIKE '%%%v%%'", name, value))
		}
	}

	var sql string
	if len(conditions) > 0 {
		sql = fmt.Sprintf("WHERE %v", strings.Join(conditions, " AND "))
	}
	return sql
}

func selectStudentsWithWhere(db *client.Conn, params model.StudentTableParams, whereSql string) (int64, []model.StudentRecord) {
	totalSql := fmt.Sprintf("SELECT COUNT(*) AS total from %v %v", tableName, whereSql)
	totalResult, totalExecErr := db.Execute(totalSql)
	if totalExecErr != nil {
		return 0, nil
	}
	total, totalResultErr := totalResult.GetIntByName(0, "total")
	if totalResultErr != nil {
		return 0, nil
	}

	orderSql := fmt.Sprintf("ORDER BY %v %v", params.Sortby, params.Order)

	offset := (params.Page - 1) * paginateLimit
	pageSql := fmt.Sprintf("LIMIT %v OFFSET %v", paginateLimit, offset)

	recordsSql := fmt.Sprintf("SELECT * FROM %v %v %v %v", tableName, whereSql, orderSql, pageSql)
	recordsResult, recordsExecErr := db.Execute(recordsSql)
	if recordsExecErr != nil {
		return 0, nil
	}

	records := []model.StudentRecord{}

	for i := range recordsResult.Values {
		records = append(records, model.GetStudentRecordFromResult(recordsResult, i))
	}
	return total, records
}
