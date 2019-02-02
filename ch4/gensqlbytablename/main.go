package main

import (
	"fmt"
	"os"
	"strings"
)

var latnIds = [...]string{"290", "910", "911", "912", "913", "914", "915", "916", "917", "919"}

// GenDropTable generator drop table
const GenDropTable string = "DROP TABLE "

// GenTruncateTable generator truncate table
const GenTruncateTable string = "TRUNCATE "

// GenSelectTable generator select table
const GenSelectTable string = "SELECT * FROM "

// UnionAll union all
const UnionAll string = "UNION ALL"

// Space space
const Space string = " "

// Wrap wrap
const Wrap string = "\r\n"

// Semicolon semicolon
const Semicolon string = ";"

// PlaceHolders placeHolders
const PlaceHolders string = "${latnId}"

func genSelectSQL(tableName string) (sql string, ok bool) {
	if strings.Index(tableName, PlaceHolders) != -1 {
		var sqlSlice []string
		for index, latnID := range latnIds {
			vSQL := ""
			if index != 9 {
				vSQL = GenSelectTable + strings.Replace(tableName, PlaceHolders, latnID, -1) + Space + UnionAll
			} else {
				vSQL = GenSelectTable + strings.Replace(tableName, PlaceHolders, latnID, -1)
			}
			sqlSlice = append(sqlSlice, vSQL)
		}
		if sqlSlice != nil {
			sql := strings.Join(sqlSlice, Wrap)
			return sql, true
		}
	}
	return "", false
}

func genDropOrTruncateSQL(arg string, tableName string) (sql string, ok bool) {
	var sqlSlice []string

	for _, latnID := range latnIds {
		vSQL := ""
		switch arg {
		case "DROP":
			vSQL = GenDropTable + strings.Replace(tableName, PlaceHolders, latnID, -1) + Semicolon
		case "TRUNCATE":
			vSQL = GenTruncateTable + strings.Replace(tableName, PlaceHolders, latnID, -1) + Semicolon
		default:

		}
		if len(vSQL) > 0 {
			sqlSlice = append(sqlSlice, vSQL)
		}

	}
	if sqlSlice != nil {
		sql := strings.Join(sqlSlice, Wrap)
		return sql, true
	}

	return "", false
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("The input arguments is error: %v\n", os.Args)
		fmt.Printf("The Usage: gensqlbytablename select/drop/truncate tablename")
		os.Exit(1)
	}

	operator := os.Args[1]
	upperOperator := strings.ToUpper(operator)
	tableName := os.Args[2]

	switch {
	case upperOperator == "SELECT":
		if sql, ok := genSelectSQL(tableName); len(sql) != 0 && ok {
			fmt.Println(sql)
		} else {
			fmt.Printf("tablename can not format :%s", tableName)
		}
	case (upperOperator == "DROP" || upperOperator == "TRUNCATE"):
		if sql, ok := genDropOrTruncateSQL(upperOperator, tableName); len(sql) != 0 && ok {
			fmt.Println(sql)
		} else {
			fmt.Printf("tablename can not format :%s", tableName)
		}
	default:
		fmt.Printf("The args[1]: %s is not supply!\n", operator)
		os.Exit(1)
	}

	os.Exit(0)
}
