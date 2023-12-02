package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	numMap = map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",

		"orez":  "0",
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}
)

func main() {
	sol1()
	sol2()
}

func sol2() {
	reStart := regexp.MustCompile(`[0-9]|zero|one|two|three|four|five|six|seven|eight|nine`)
	reEnd := regexp.MustCompile(`[0-9]|orez|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)

	lines := strings.Split(input1, "\n")
	var toSum []int64
	for _, v := range lines {
		st := reStart.FindString(v)
		end := reEnd.FindString(reverse(v))
		if st == "" || end == "" {
			fmt.Println("Error: found string without numbers:", v)
			os.Exit(1)
		}
		if v, ok := numMap[st]; ok {
			st = v
		}
		if v, ok := numMap[end]; ok {
			end = v
		}

		num, err := strconv.ParseInt(fmt.Sprintf("%s%s", st, end), 10, 64)
		if err != nil {
			fmt.Println("Error: could not parse number:", err)
			os.Exit(1)
		}
		toSum = append(toSum, num)
	}
	var sum int64
	for _, v := range toSum {
		sum += v
	}
	fmt.Println(sum)
}

func reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func sol1() {
	re := regexp.MustCompile(`[0-9]`)
	lines := strings.Split(input1, "\n")
	var toSum []int64
	for _, v := range lines {
		matches := re.FindAllString(v, -1)
		if len(matches) < 1 {
			fmt.Println("Error: found string without numbers:", v)
			os.Exit(1)
		}

		num, err := strconv.ParseInt(fmt.Sprintf("%s%s", matches[0], matches[len(matches)-1]), 10, 64)
		if err != nil {
			fmt.Println("Error: could not parse number:", err)
			os.Exit(1)
		}
		toSum = append(toSum, num)
	}
	var sum int64
	for _, v := range toSum {
		sum += v
	}
	fmt.Println(sum)
}
