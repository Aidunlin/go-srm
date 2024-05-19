package db

import (
	"fmt"
	"strings"

	"github.com/Aidunlin/go-srm/app"
	"github.com/go-mysql-org/go-mysql/client"
)

const dbAddr = "localhost:3306"
const dbUser = "root"
const dbPassword = ""
const dbName = "ctec"
const dbTable = "student_v2"

func SelectRecords(params app.RecordTableParams) (int64, []app.StudentRecord) {
	db, connectErr := client.Connect(dbAddr, dbUser, dbPassword, dbName)
	if connectErr != nil {
		return 0, nil
	}
	defer db.Close()

	whereSql := ""

	if len(params.Filter) > 0 {
		whereSql = fmt.Sprintf("WHERE last_name LIKE '%v%%'", params.Filter)
	}

	totalSql := fmt.Sprintf("SELECT COUNT(*) AS total from %v %v", dbTable, whereSql)
	totalResult, totalExecErr := db.Execute(totalSql)
	if totalExecErr != nil {
		return 0, nil
	}
	total, totalResultErr := totalResult.GetIntByName(0, "total")
	if totalResultErr != nil {
		return 0, nil
	}

	orderSql := fmt.Sprintf("ORDER BY %v %v", params.Sortby, params.Order)

	offset := (params.Page - 1) * app.PaginateLimit
	pageSql := fmt.Sprintf("LIMIT %v OFFSET %v", app.PaginateLimit, offset)

	recordsSql := fmt.Sprintf("SELECT * FROM %v %v %v %v", dbTable, whereSql, orderSql, pageSql)
	recordsResult, recordsExecErr := db.Execute(recordsSql)
	if recordsExecErr != nil {
		return 0, nil
	}

	records := []app.StudentRecord{}

	for i := range recordsResult.Values {
		// Included to distinguish between records.
		id, _ := recordsResult.GetIntByName(i, "id")

		studentId, _ := recordsResult.GetIntByName(i, "student_id")
		firstName, _ := recordsResult.GetStringByName(i, "first_name")
		lastName, _ := recordsResult.GetStringByName(i, "last_name")
		email, _ := recordsResult.GetStringByName(i, "email")
		phone, _ := recordsResult.GetStringByName(i, "phone")
		degreeProgram, _ := recordsResult.GetStringByName(i, "degree_program")
		gpa, _ := recordsResult.GetFloatByName(i, "gpa")
		financialAid, _ := recordsResult.GetIntByName(i, "financial_aid")
		graduationDate, _ := recordsResult.GetStringByName(i, "graduation_date")

		records = append(records, app.StudentRecord{
			Id:             id,
			StudentId:      studentId,
			FirstName:      firstName,
			LastName:       lastName,
			Email:          email,
			Phone:          phone,
			DegreeProgram:  degreeProgram,
			Gpa:            gpa,
			FinancialAid:   financialAid,
			GraduationDate: graduationDate,
		})
	}
	return total, records
}

func SelectRecord(id int) (bool, app.StudentRecord) {
	db, connectErr := client.Connect(dbAddr, dbUser, dbPassword, dbName)
	if connectErr != nil {
		return false, app.StudentRecord{}
	}
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %v WHERE id = ? LIMIT 1", dbTable)
	stmt, prepareErr := db.Prepare(sql)
	if prepareErr != nil {
		return false, app.StudentRecord{}
	}

	result, execErr := stmt.Execute(id)
	if execErr != nil || len(result.Values) < 1 {
		return false, app.StudentRecord{}
	}

	studentId, _ := result.GetIntByName(0, "student_id")
	firstName, _ := result.GetStringByName(0, "first_name")
	lastName, _ := result.GetStringByName(0, "last_name")
	email, _ := result.GetStringByName(0, "email")
	phone, _ := result.GetStringByName(0, "phone")
	degreeProgram, _ := result.GetStringByName(0, "degree_program")
	gpa, _ := result.GetFloatByName(0, "gpa")
	financialAid, _ := result.GetIntByName(0, "financial_aid")
	graduationDate, _ := result.GetStringByName(0, "graduation_date")

	return true, app.StudentRecord{
		StudentId:      studentId,
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

func InsertRecord(record app.StudentRecord) bool {
	db, connectErr := client.Connect(dbAddr, dbUser, dbPassword, dbName)
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

	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", dbTable, namesString, placeholdersString)
	stmt, prepareErr := db.Prepare(sql)
	if prepareErr != nil {
		return false
	}

	_, execErr := stmt.Execute(values...)
	return execErr == nil
}

func UpdateRecord(id int, record app.StudentRecord) bool {
	db, connectErr := client.Connect(dbAddr, dbUser, dbPassword, dbName)
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

	sql := fmt.Sprintf("UPDATE %v SET %v WHERE id = ?", dbTable, settersString)
	stmt, prepareErr := db.Prepare(sql)
	if prepareErr != nil {
		return false
	}

	_, execErr := stmt.Execute(values...)
	return execErr == nil
}

func DeleteRecord(id int) bool {
	db, connectErr := client.Connect(dbAddr, dbUser, dbPassword, dbName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	sql := fmt.Sprintf("DELETE FROM %v WHERE id = ?", dbTable)
	stmt, prepareErr := db.Prepare(sql)
	if prepareErr != nil {
		return false
	}

	result, execErr := stmt.Execute(id)
	return execErr == nil && result.AffectedRows == 1
}
