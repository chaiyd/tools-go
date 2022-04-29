package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
	"github.com/xuri/excelize/v2"
)

func main() {

	var fineNum = "1"

	f := excelize.NewFile()
	if err := f.SaveAs("gitlab-1.xlsx"); err != nil {
		fmt.Println(err)
	}

	jsonfile, err := os.Open("/gitlab/json/gitlab_" + fineNum + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonfile.Close()
	byteValue, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("%v%v", i, ".name")
		path_with_namespace := fmt.Sprintf("%v%v", i, ".path_with_namespace")
		http_url_to_repo := fmt.Sprintf("%v%v", i, ".http_url_to_repo")
		description := fmt.Sprintf("%v%v", i, ".description")

		git_name := gjson.Get(string(byteValue), name)
		git_des := gjson.GetBytes(byteValue, description)
		git_namespace := gjson.Get(string(byteValue), path_with_namespace)
		git_url := gjson.Get(string(byteValue), http_url_to_repo)

		f, err := excelize.OpenFile("gitlab-1.xlsx")
		if err != nil {
			fmt.Println(err)
			excelize.NewFile()
		}

		//lineStr := fmt.Sprintf(strconv.Itoa(i) + "00")
		//line, err := strconv.Atoi(lineStr)
		//fmt.Println(line)
		////line = line + 1

		//fmt.Println(fineline)
		f.SetCellValue("Sheet1", fmt.Sprint("A", 100+i+1), git_name.String())
		f.SetCellValue("Sheet1", fmt.Sprint("B", 100+i+1), git_des.String())
		f.SetCellValue("Sheet1", fmt.Sprint("C", 100+i+1), git_namespace.String())
		f.SetCellValue("Sheet1", fmt.Sprint("D", 100+i+1), git_url.String())

		if err := f.SaveAs("gitlab-1.xlsx"); err != nil {
			fmt.Println(err)
		}
		//fmt.Println(git_name.String(),git_namespace.String(), git_des, git_url.String())
	}
}
