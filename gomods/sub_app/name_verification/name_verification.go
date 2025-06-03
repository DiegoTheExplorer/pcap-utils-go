package name_verification

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CrimName struct {
	FirstName  string `json:"fname"`
	MiddleName string `json:"mname"`
	LastName   string `json:"lname"`
}

func ReadVerifiedNames(filepath string, preview_count int) ([]string, error) {
	fp, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	var verified_names []string
	var non_alphabetic_regex = regexp.MustCompile(`[^A-ZÑ ]+`)

	// TODO: Check how the file is formatted

	// HACK: Currently only works with the copy-pasted
	// list of passers from the website

	// Loop through text file
	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, " ")

		// Try to convert the first string to an integer
		_, err := strconv.Atoi(split_line[0])

		// Skip line if convertion is a fail
		if err != nil {
			continue
		}

		// Remove passer number
		split_line = split_line[1:]
		name := strings.ToUpper(strings.Join(split_line, " "))
		name = non_alphabetic_regex.ReplaceAllString(name, "")
		verified_names = append(verified_names, name)
	}

	// Check if the number of preview rows requested
	// is greater than the nubmer of rows read from
	// the file
	if len(verified_names) < preview_count {
		return verified_names, nil
	}

	return verified_names[0 : preview_count-1], nil
}

func ReadGformsNames(filepath string, preview_count int) ([]CrimName, error) {
	csv_file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer csv_file.Close()

	// initialize the csv reader
	csv_data, err := csv.NewReader(csv_file).ReadAll()

	if err != nil {
		return nil, err
	}

	var gforms_names []CrimName

	// Identify the row index for last name, first name, and middle name
	var fname, mname, lname = -1, -1, -1
	for ind, row_name := range csv_data[0] {
		row_name = strings.ToLower(row_name)
		switch row_name {
		case "first name":
			fname = ind
		case "middle name":
			mname = ind
		case "last name":
			lname = ind
		default:
		}
	}

	var err_string []string
	err_string = append(err_string, "The following rows are missing: ")
	row_missing := false
	if fname < 0 {
		err_string = append(err_string, "first name ")
		row_missing = true
	}
	if mname < 0 {
		err_string = append(err_string, "middle name ")
		row_missing = true
	}
	if lname < 0 {
		err_string = append(err_string, "last name ")
		row_missing = true
	}

	if row_missing {
		return nil, fmt.Errorf(strings.Join(err_string, ""))
	}

	for _, row := range csv_data[1:] {
		name := CrimName{
			FirstName:  row[fname],
			MiddleName: row[mname],
			LastName:   row[lname],
		}

		gforms_names = append(gforms_names, name)
	}

	// Check if the number of preview rows requested
	// is greater than the nubmer of rows read from
	// the file
	if len(gforms_names) < preview_count {
		return gforms_names, nil
	}

	return gforms_names[0:preview_count], nil
}
