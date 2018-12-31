package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const cloudCrmPath string = "D:/work/tydic/云化项目"
const mobilePath string = "D:/work/tydic/移动化受理/MOB/develop"

func main() {
	var pathName string
	var dateStr string
	var pathPrefix string
	var pathAll string

	var year, day int
	var month time.Month

	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&pathName, "path", "cloud", "cloud project")
	mySet.StringVar(&dateStr, "date", "sysdate", "system date")
	mySet.Parse(os.Args[1:])

	if len(os.Args) > 0 {
		switch pathName {
		case "cloud":
			pathPrefix = cloudCrmPath
		case "mobile":
			pathPrefix = mobilePath
		default:
			pathPrefix = cloudCrmPath
		}

		if dateStr != "sysdate" {
			if len(dateStr) != 8 {
				fmt.Printf("input error: The length of %s is not equals 8", dateStr)
				os.Exit(1)
			} else {
				year, month, day = getFormatTime(dateStr)
			}
		} else {
			year, month, day = getCurrentTime()
		}
	}

	yearString := strconv.Itoa(year)
	monthString := strconv.Itoa(int(month))
	if month < 10 {
		monthString = "0" + strconv.Itoa(int(month))
	}
	dayString := strconv.Itoa(day)
	if day < 10 {
		dayString = "0" + strconv.Itoa(day)
	}
	pathAll = pathPrefix + "/" + yearString + "/" + monthString + "/" + monthString + dayString + "/"
	fmt.Println("pathAll", pathAll)
	windowsFilePath := strings.Replace(pathAll, "/", "\\", -1)

	if _, err := os.Stat(pathAll); os.IsNotExist(err) {
		fmt.Printf("文件不存在:%s,%v", pathAll, err)
		os.MkdirAll(pathAll, os.ModeDir)
	}

	openDirByExplorer(windowsFilePath)
}

func getCurrentTime() (year int, month time.Month, day int) {
	return time.Now().Date()
}

func getFormatTime(dateStr string) (year int, month time.Month, day int) {
	year, err := strconv.Atoi(string(dateStr[0:4]))
	if err != nil {
		fmt.Printf("The input date: %s string can't format", dateStr)
		os.Exit(1)
	}

	monthInt, err := strconv.Atoi(string(dateStr[4:6]))
	if err != nil {
		fmt.Printf("The input date: %s string can't format", dateStr)
		os.Exit(1)
	}
	month = time.Month(monthInt)

	day, err = strconv.Atoi(string(dateStr[6:8]))
	if err != nil {
		fmt.Printf("The input date: %s string can't format", dateStr)
		os.Exit(1)
	}

	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	return year, month, day
}

func openDirByExplorer(path string) {
	cmd := exec.Command("explorer", "select,"+path)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("文件打开失败:%s,%v", path, err)
	}
}
