package db

import (
	"fmt"
	"math"
	"strings"

	"github.com/Aidunlin/go-srm/app"
	"github.com/go-mysql-org/go-mysql/client"
	"github.com/go-mysql-org/go-mysql/mysql"
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

func getRecordFromResult(row int, result *mysql.Result) app.StudentRecord {
	id, _ := result.GetIntByName(row, "id")
	firstName, _ := result.GetStringByName(row, "first_name")
	lastName, _ := result.GetStringByName(row, "last_name")
	email, _ := result.GetStringByName(row, "email")
	phone, _ := result.GetStringByName(row, "phone")
	degreeProgram, _ := result.GetStringByName(row, "degree_program")
	gpa, _ := result.GetFloatByName(row, "gpa")
	financialAid, _ := result.GetIntByName(row, "financial_aid")
	graduationDate, _ := result.GetStringByName(row, "graduation_date")

	return app.StudentRecord{
		Id:             id,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		Phone:          phone,
		DegreeProgram:  degreeProgram,
		Gpa:            gpa,
		FinancialAid:   financialAid,
		GraduationDate: graduationDate,
	}
}

func whereLastNameFilter(filter string) string {
	if len(filter) != 1 {
		return ""
	}
	return fmt.Sprintf("WHERE last_name LIKE '%v%%'", filter)
}

func whereBasicSearch(search string) string {
	var conditions []string
	for _, column := range app.GetStudentColumns() {
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

func whereAdvancedSearch(form app.StudentRecord) string {
	var conditions []string
	for name, value := range form.GetRawMap() {
		if len(value) == 0 {
			continue
		}

		column := app.GetStudentColumn(name)
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

func selectRecordsWithWhere(db *client.Conn, params app.RecordTableParams, whereSql string) (int64, []app.StudentRecord) {
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

	records := []app.StudentRecord{}

	for i := range recordsResult.Values {
		records = append(records, getRecordFromResult(i, recordsResult))
	}
	return total, records
}

func SelectRecords(params app.RecordTableParams) (int64, []app.StudentRecord) {
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
	return selectRecordsWithWhere(db, params, whereSql)
}

func SelectRecordsWithForm(params app.RecordTableParams, form app.StudentRecord) (int64, []app.StudentRecord) {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return 0, nil
	}
	defer db.Close()
	return selectRecordsWithWhere(db, params, whereAdvancedSearch(form))
}

func SelectRecord(id int) (bool, app.StudentRecord) {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return false, app.StudentRecord{}
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %v WHERE id = ? LIMIT 1", tableName)
	result, execErr := db.Execute(sql, id)
	if execErr != nil || len(result.Values) < 1 {
		return false, app.StudentRecord{}
	}

	return true, getRecordFromResult(0, result)
}

func InsertRecord(record app.StudentRecord) bool {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	paramsMap := record.GetMap()

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

func UpdateRecord(id int, record app.StudentRecord) bool {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	paramsMap := record.GetMap()

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

func DeleteRecord(id int) bool {
	db, connectErr := client.Connect(address, user, password, databaseName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	sql := fmt.Sprintf("DELETE FROM %v WHERE id = ?", tableName)
	result, execErr := db.Execute(sql, id)
	return execErr == nil && result.AffectedRows == 1
}
