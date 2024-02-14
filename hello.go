package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"encoding/json"

	"github.com/xuri/excelize/v2"
)



func meal_finder(day string, meal string) []string {
	var a []string
	var n int
	file, err := excelize.OpenFile("Sample-Menu.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	rows, _ := file.GetRows("Sheet1")
	row1 := rows[0]
	for i := 0; i < len(row1); i++ {
		if strings.EqualFold(row1[i], day) {
			n = i
			break
		}
	}
	cols, _ := file.GetCols("Sheet1")
	col_meal := cols[n]
	for j := 0; j < len(col_meal); j++ {
		if strings.EqualFold(col_meal[j], meal) {
			for k := j + 1; k <= len(col_meal)-1; k++ {
				if strings.EqualFold(col_meal[k], day) {
					break
				} else if col_meal[k] == "" {
					continue
				} else {
					a = append(a, col_meal[k])
				}
			}
			break
		}
	}
	return a
}
func item_counter(day string, meal string) int {
	//I have considered tea and coffee to be separate and bread and jam to be seperate so these are counted as 2
	//I have done the same for veg fried rice and egg fried rice
	var no_items int
	var n int
	file, err := excelize.OpenFile("Sample-Menu.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	rows, _ := file.GetRows("Sheet1")
	row1 := rows[0]
	for i := 0; i < len(row1); i++ {
		if strings.EqualFold(row1[i], day) {
			n = i
			break
		}
	}
	cols, _ := file.GetCols("Sheet1")
	col_meal := cols[n]
	for j := 0; j < len(col_meal); j++ {
		if strings.EqualFold(col_meal[j], meal) {
			for k := j + 1; k <= len(col_meal)-1; k++ {
				if strings.EqualFold(col_meal[k], day) {
					break
				} else if col_meal[k] == "" {
					continue
				} else if strings.Contains(col_meal[k], "+") || strings.Contains(col_meal[k], "/") {
					no_items += 2
				} else {
					no_items++
				}
			}
			break
		}
	}
	return no_items

}
func item_checker(day string, meal string, item string) bool {
	var n int
	file, err := excelize.OpenFile("Sample-Menu.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	rows, _ := file.GetRows("Sheet1")
	row1 := rows[0]
	for i := 0; i < len(row1); i++ {
		if strings.EqualFold(row1[i], day) {
			n = i
			break
		}
	}
	cols, _ := file.GetCols("Sheet1")
	col_meal := cols[n]
	for j := 0; j < len(col_meal); j++ {
		if strings.EqualFold(col_meal[j], meal) {
			for k := j + 1; k <= len(col_meal)-1; k++ {
				if strings.EqualFold(col_meal[k], day) {
					break
				} else if strings.EqualFold(col_meal[k], item) {
					return true
				} else if strings.EqualFold(item, "tea") || strings.EqualFold(item, "coffee") ||
					strings.EqualFold(item, "Veg fried rice") || strings.EqualFold(item, "egg fried rice") || strings.EqualFold(item, "bread") ||
					strings.EqualFold(item, "jam") || strings.EqualFold(item, "butter") || strings.EqualFold(item, "curd") || strings.EqualFold(item, "Tawa veg") {
					return true
				} else if col_meal[k] == "" {
					break
				} else {
					continue
				}
			}
			break
		}
	}
	return false

}

func myjson_converter() {
	type Rows struct {
		Monday    string `json:"Monday"`
		Tuesday   string `json:"Tuesday"`
		Wednesday string `json:"Wednesday"`
		Thursday  string `json:"Thursday"`
		Friday    string `json:"Friday"`
		Saturday  string `json:"Saturday"`
		Sunday    string `json:"Sunday"`
	}

	var mainrows []Rows
	File, err := excelize.OpenFile("Sample-Menu.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	rows, _ := File.GetRows("Sheet1")
	for i := 0; i < len(rows); i++ {
		var rows1 Rows
		for j := 0; j < len(rows[i]); j++ {
			header := rows[0][j]
			switch header {
			case "MONDAY":
				rows1.Monday = string(rows[i][j])
			case "TUESDAY":
				rows1.Tuesday = string(rows[i][j])
			case "WEDNESDAY":
				rows1.Wednesday = string(rows[i][j])
			case "THURSDAY":
				rows1.Thursday = string(rows[i][j])
			case "FRIDAY":
				rows1.Friday = string(rows[i][j])
			case "SATURDAY":
				rows1.Saturday = string(rows[i][j])
			case "SUNDAY":
				rows1.Sunday = string(rows[i][j])
			}
		}
		mainrows = append(mainrows, rows1)

	}
	jsonData, err := json.MarshalIndent(mainrows, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	jsonfile := os.WriteFile("output.json", jsonData, 0666)
	if jsonfile != nil {
		fmt.Println(jsonfile)
	}

}

type full_meal struct {
	day             string
	date            string
	Breakfast_items []string
	Lunch_items     []string
	Dinner_items    []string
}

func (m full_meal) details() {
	output := fmt.Sprintf("Day: %s \nDate: %s \nBreakfast items:%s\nLunch item:%s \nDinner_items%s",
		m.day, m.date, m.Breakfast_items, m.Lunch_items, m.Dinner_items)

	fmt.Println(output)

}

func main() {
	m := 1
	var day string
	var meal string
	var item string
	default_string := "Please choose one of the following options\n1. Want to know your meal\n2. Number of items\n3. Check if item in meal\n4. All day meal struct\n5. Exit"

	for m == 1 {
		var n int
		fmt.Println(default_string)
		fmt.Scan(&n)

		if n == 1 {
			fmt.Print("Please enter day")
			fmt.Scan(&day)
			fmt.Print("Please enter Meal")
			fmt.Scan(&meal)
			meals := meal_finder(day, meal)
			for i := 0; i < len(meals); i++ {
				fmt.Println(meals[i])
			}
		}
		if n == 2 {
			//I have considered tea and coffee to be separate and bread and jam to be seperate so these are counted as 2
			//I have done the same for veg fried rice and egg fried rice
			fmt.Print("Please enter day")
			fmt.Scan(&day)
			fmt.Print("Please enter Meal")
			fmt.Scan(&meal)
			meals := item_counter(day, meal)
			fmt.Printf("No of items: %v", meals)
		}
		if n == 3 {
			fmt.Print("Please enter day")
			fmt.Scan(&day)
			fmt.Print("Please enter Meal")
			fmt.Scan(&meal)
			fmt.Scanln()
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Please enter item you want to search for")
			item, _ = reader.ReadString('\n')
			cleaneditem := strings.ReplaceAll(item, "\n", "")
			cleaneditem = strings.ReplaceAll(cleaneditem, "\r", "")
			if item_checker(day, meal, cleaneditem) {
				fmt.Printf("The item is present in the meal")
			} else {
				fmt.Printf("The item is not present in the meal")

			}
		}
		if n == 4 {
			fmt.Print("Enter day")
			fmt.Scan(&day)
			File, err := excelize.OpenFile("Sample-Menu.xlsx")
			if err != nil {
				fmt.Println(err)
			}
			var t int
			rows, _ := File.GetRows("Sheet1")
			row1 := rows[0]
			for i := 0; i < len(row1); i++ {
				if strings.EqualFold(row1[i], day) {
					t = i
				}
			}
			date := rows[1][t]
			items_breakfast := meal_finder(day, "Breakfast")
			items_lunch := meal_finder(day, "Lunch")
			items_dinner := meal_finder(day, "Dinner")
			m := full_meal{day: day, date: date, Breakfast_items: items_breakfast, Lunch_items: items_lunch, Dinner_items: items_dinner}
			m.details()
		}
		if n == 5 {
			m = 0
		}
	}
	fmt.Println("Now we will convert this xslx file to json file and save it in the same directory")
	myjson_converter()

}
