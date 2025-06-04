package name_verification

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type CrimName struct {
	FirstName  string `json:"fname"`
	MiddleName string `json:"mname"`
	LastName   string `json:"lname"`
}

type ForManualReview struct {
	SubmittedName string      `json:"submitted_name"`
	Matches       fuzzy.Ranks `json:"matches"`
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

	var non_alphabetic_regex = regexp.MustCompile(`[^A-ZÑ ]+`)
	for _, row := range csv_data[1:] {
		first_name := non_alphabetic_regex.ReplaceAllString(strings.ToUpper(row[fname]), "")
		middle_name := non_alphabetic_regex.ReplaceAllString(strings.ToUpper(row[mname]), "")
		last_name := non_alphabetic_regex.ReplaceAllString(strings.ToUpper(row[lname]), "")
		name := CrimName{
			FirstName:  first_name,
			MiddleName: middle_name,
			LastName:   last_name,
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

func CompareNames(ver_names_fp string, gforms_names_fp string) ([]ForManualReview, error) {
	gforms_names, err := ReadGformsNames(gforms_names_fp, math.MaxInt)
	if err != nil {
		return nil, err
	}

	verified_names, err := ReadVerifiedNames(ver_names_fp, math.MaxInt)
	if err != nil {
		return nil, err
	}

	var for_man_review []ForManualReview
	fmt.Println("Number of names read from .csv: ", len(gforms_names))
	for _, crim := range gforms_names {
		var crim_arr []string

		if crim.MiddleName == "" {
			crim_arr = append(crim_arr, crim.LastName, crim.FirstName)
		} else {
			crim_arr = append(crim_arr, crim.LastName, crim.FirstName, crim.MiddleName)
		}

		crim_name := strings.TrimSpace(strings.Join(crim_arr, " "))

		// Check for exact matches
		var exact_match_found = false
		for _, ver_name := range verified_names {
			if ver_name == crim_name {
				exact_match_found = true
			}
		}

		if exact_match_found {
			continue
		}

		// Use fuzzy to look for partial matches
		match_res := fuzzy.RankFind(crim_name, verified_names)

		switch match_res.Len() {
		case 0:
			for_man_review = append(for_man_review, ForManualReview{
				SubmittedName: crim_name,
				Matches:       nil,
			})
			continue
		case 1:
			// HACK: Redundant perfect match check
			// but idk sometimes this doesn't catch
			// perfect matches
			if match_res[0].Distance == 0 {
				continue
			}
			for_man_review = append(for_man_review, ForManualReview{
				SubmittedName: crim_name,
				Matches:       match_res,
			})
		default:
			for_man_review = append(for_man_review, ForManualReview{
				SubmittedName: crim_name,
				Matches:       match_res,
			})
		}
	}

	fmt.Println("Number of names for manual review: ", len(for_man_review))
	return for_man_review, nil
}
