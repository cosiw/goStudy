package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "vegas:!mgrsol123@tcp(int.trustnhope.com:6306)/h00258")
	if err != nil {
		log.Fatal(err)
	}
	startrow := 0
	ppl := 100
	statementList := make([]Statement, 0)
	for {
		statementList, err = getStatementData(db, startrow, ppl)
		insertPayment(db, statementList)

		startrow += ppl

		fmt.Printf("%d / %d\n", startrow, ppl)
		if len(statementList) < ppl {
			break
		}
	}
	//schList := make([]Schedule, 0)

	//schList, err = getScheduleData(db, startrow, ppl)

}

func GetJSONFromRows(rows *sql.Rows) {
	columns, _ := rows.Columns()

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	jsonData, _ := json.Marshal(tableData)

	fmt.Printf("%s", string(jsonData))
}

// func getScheduleData(db *sql.DB, startRow int, ppl int) ([]int, error) {

// 	var query = fmt.Sprintf(`SELECT SCHEDULEID, ORGID, SCHEDULEDATE, SCHEDULETIME, CUSTOMERID, SCHEDULESTATUS
// 	FROM TCUSTOMERSCHEDULE WHERE LEFT(SCHEDULEDATE,4) = '2022' AND SCHEDULESTATUS = 5 AND INSTYPE = 1 LIMIT %d, %d`, startRow, ppl)

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		fmt.Printf("Query error = %s\n", err)
// 	}
// 	defer rows.Close()
// 	schList := make([]Schedule, 0)

// 	for rows.Next() {
// 		var r Schedule
// 		err = rows.Scan(&r.ScheduleId, &r.OrgId, &r.ScheduleDate, &r.ScheduleTime, &r.CustomerId, &r.Status)
// 		if err != nil {
// 			fmt.Printf("Scan error = %s\n", err)
// 		}
// 		schList = append(schList, r)
// 	}
// 	//fmt.Println(schList)
// 	//SchMap := make([]map[string]Schedule, 0)

// 	return schList, err
// }

func getStatementData(db *sql.DB, startRow int, ppl int) ([]Statement, error) {

	var query = fmt.Sprintf(`SELECT STATEMENTID, ORGID, SCHEDULEID, CUSTOMERID, LIABILITYAMT, NONINSAMT, TAXABLEAMT
	FROM TSALESTATEMENT WHERE SCHEDULEID IN (SELECT SCHEDULEID FROM TCUSTOMERSCHEDULE WHERE LEFT(SCHEDULEDATE,4)= '2022' AND SCHEDULESTATUS = 5 AND INSTYPE = 1)
	ORDER BY SCHEDULEID LIMIT %d, %d`, startRow, ppl)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Query error = %s\n", err)
	}
	defer rows.Close()
	stateList := make([]Statement, 0)

	for rows.Next() {
		var r Statement
		err = rows.Scan(&r.StatementId, &r.OrgId, &r.ScheduleId, &r.CustomerId, &r.LiabilityAmt, &r.NoninsAmt, &r.TaxableAmt)
		if err != nil {
			fmt.Printf("Scan error = %s\n", err)
		}
		stateList = append(stateList, r)
	}

	return stateList, err
}

func insertPayment(db *sql.DB, stateList []Statement) error {
	size := len(stateList)
	for i := 0; i < size; i++ {
		query := fmt.Sprintf(`INSERT INTO TPAYMENT (ORGID, CUSTOMERID, SCHEDULEID, PAYMENTAMT, PAYLIABILITYAMT, PAYMENTCODE, PAYDATE, ORDERDATE, TAX, CRTIME) VALUES
		 (%d, %d, %d, %d, %d, 'E01',  '20230105', '20230105', %f, '20230105000000' )`, stateList[i].OrgId, stateList[i].CustomerId, stateList[i].ScheduleId, stateList[i].LiabilityAmt+stateList[i].NoninsAmt,
			stateList[i].LiabilityAmt, stateList[i].TaxableAmt)

		_, sqlErr := db.Exec(query)

		if sqlErr != nil {
			fmt.Printf("Failed to insert data (%s)", sqlErr)
			return sqlErr
		}
	}

	return nil
}
