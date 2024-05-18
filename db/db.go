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

func InsertRecord(params app.RecordFormParams) bool {
	db, connectErr := client.Connect(dbAddr, dbUser, dbPassword, dbName)
	if connectErr != nil {
		return false
	}
	defer db.Close()

	paramsMap := params.GetMap(false)

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

	namesString := strings.Join(names, ",")
	placeholdersString := strings.Join(placeholders, ", ")

	sql := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", dbTable, namesString, placeholdersString)
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
