package main

import (
	"fmt"
	"os"
	"strings"
)

var latnIds = [...]string{"290", "910", "911", "912", "913", "914", "915", "916", "917", "919"}

// GenDropTable generator drop table
const GenDropTable string = "DROP TABLE "

// GenSelectTable generator select table
const GenSelectTable string = "SELECT * FROM "

// UnionAll union all
const UnionAll string = "UNION ALL"

// Space space
const Space string = " "

// Semicolon semicolon
const Semicolon string = ";"

// PlaceHolders placeHolders
const PlaceHolders string = "${latnId}"

func genSelectSQL(tableName string) (sql string, ok bool) {
	if strings.Index(tableName, PlaceHolders) != -1 {
		//sqlSlice := make([]string, 10)
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
			sql := strings.Join(sqlSlice, Space)
			return sql, true
		} else {
			return "", false
		}
	}
	return "", false
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("The input arguments is error: %v\n", os.Args)
		fmt.Printf("The Usage: gensqlbytable select/drop tablename")
		os.Exit(1)
	}

	operator := os.Args[1]
	tableName := os.Args[2]

	switch strings.ToUpper(operator) {
	case "SELECT":
		if sql, ok := genSelectSQL(tableName); len(sql) != 0 && ok {
			fmt.Println(sql)
		} else {
			fmt.Printf("tablename can not format :%s", tableName)
		}
	}
}
