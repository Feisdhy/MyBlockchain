package state

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Txt2CsvForSequentialRead() {
	for i := 1; i <= 6; i++ {
		var (
			path     string
			division int64
		)

		switch i {
		case 1:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1W/experiment"
				division = 10000
			}
		case 2:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10W/experiment"
				division = 100000
			}
		case 3:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_100W/experiment"
				division = 1000000
			}
		case 4:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886/experiment"
				division = 1000000
			}
		case 5:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W/experiment"
				division = 1000000
			}
		case 6:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W/experiment"
				division = 1000000
			}
		}

		for j := 1; j <= 10; j++ {
			file, _ := os.Open(path + "/sequential_read_result" + strconv.Itoa(j) + ".txt")
			csvfile, _ := os.OpenFile(path+"/sequential read result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			writer := csv.NewWriter(csvfile)

			count := int64(0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 定义正则表达式
				//str := "2023/07/11 04:50:11 0xdF2eeBaa127019dc84d0725EFD84b77616276D99 166116"
				re := regexp.MustCompile(`(\d+)$`)
				matches := re.FindStringSubmatch(line)

				number, _ := strconv.ParseInt(matches[1], 10, 64)
				count += number
			}
			result1 := count / division
			fmt.Println(strconv.FormatInt(result1, 10))

			file, _ = os.Open(path + "/sequential_read_result" + strconv.Itoa(j) + "_cache.txt")

			count = int64(0)
			scanner = bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 定义正则表达式
				//str := "2023/07/11 04:50:11 0xdF2eeBaa127019dc84d0725EFD84b77616276D99 166116"
				re := regexp.MustCompile(`(\d+)$`)
				matches := re.FindStringSubmatch(line)

				number, _ := strconv.ParseInt(matches[1], 10, 64)
				count += number
			}
			result2 := count / division
			fmt.Println(strconv.FormatInt(result2, 10))

			writer.Write([]string{strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10)})
			writer.Flush()

			file.Close()
			csvfile.Close()
		}
		fmt.Println()
	}
}

func Txt2CsvForSequentialWrite() {
	for i := 1; i <= 6; i++ {
		var (
			path     string
			division int64
		)

		switch i {
		case 1:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1W/experiment"
				division = 10000
			}
		case 2:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10W/experiment"
				division = 100000
			}
		case 3:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_100W/experiment"
				division = 1000000
			}
		case 4:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886/experiment"
				division = 1000000
			}
		case 5:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W/experiment"
				division = 1000000
			}
		case 6:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W/experiment"
				division = 1000000
			}
		}

		for j := 1; j <= 10; j++ {
			file, _ := os.Open(path + "/sequential_write_result" + strconv.Itoa(j) + ".txt")
			csvfile, _ := os.OpenFile(path+"/sequential write result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			writer := csv.NewWriter(csvfile)

			count1 := int64(0)
			count2 := int64(0)
			count3 := int64(0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 定义正则表达式
				//str := "2023/07/11 12:44:36 0x0000944fc866186055e47c4db9c2d2db32e9d3de29588b727adcc767bd3615f7 268712 40415 22580"
				re := regexp.MustCompile(`(\d+) (\d+) (\d+)$`)
				matches := re.FindStringSubmatch(line)

				number1, _ := strconv.ParseInt(matches[1], 10, 64)
				number2, _ := strconv.ParseInt(matches[2], 10, 64)
				number3, _ := strconv.ParseInt(matches[3], 10, 64)

				count1 += number1
				count2 += number2
				count3 += number3

			}
			result1 := count1 / division
			result2 := count2 / division
			result3 := count3 / division

			fmt.Println(strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10), strconv.FormatInt(result3, 10))

			file, _ = os.Open(path + "/sequential_write_result" + strconv.Itoa(j) + "_cache.txt")

			count1 = int64(0)
			count2 = int64(0)
			count3 = int64(0)
			scanner = bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 定义正则表达式
				//str := "2023/07/11 12:44:36 0x0000944fc866186055e47c4db9c2d2db32e9d3de29588b727adcc767bd3615f7 268712 40415 22580"
				re := regexp.MustCompile(`(\d+) (\d+) (\d+)$`)
				matches := re.FindStringSubmatch(line)

				number1, _ := strconv.ParseInt(matches[1], 10, 64)
				number2, _ := strconv.ParseInt(matches[2], 10, 64)
				number3, _ := strconv.ParseInt(matches[3], 10, 64)

				count1 += number1
				count2 += number2
				count3 += number3
			}
			result4 := count1 / division
			result5 := count2 / division
			result6 := count3 / division

			fmt.Println(strconv.FormatInt(result4, 10), strconv.FormatInt(result5, 10), strconv.FormatInt(result6, 10))

			writer.Write([]string{strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10), strconv.FormatInt(result3, 10), strconv.FormatInt(result4, 10), strconv.FormatInt(result5, 10), strconv.FormatInt(result6, 10)})
			writer.Flush()

			file.Close()
			csvfile.Close()
		}
		fmt.Println()
	}
}

func Txt2CsvForRandomRead() {
	for i := 1; i <= 6; i++ {
		var (
			path     string
			division int64
		)

		switch i {
		case 1:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1W/experiment"
				division = 10000
			}
		case 2:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10W/experiment"
				division = 100000
			}
		case 3:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_100W/experiment"
				division = 1000000
			}
		case 4:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886/experiment"
				division = 1000000
			}
		case 5:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W/experiment"
				division = 1000000
			}
		case 6:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W/experiment"
				division = 1000000
			}
		}

		for j := 1; j <= 10; j++ {
			file, _ := os.Open(path + "/random_read_result" + strconv.Itoa(j) + ".txt")
			csvfile, _ := os.OpenFile(path+"/random read result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			writer := csv.NewWriter(csvfile)

			count := int64(0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 定义正则表达式
				//str := "2023/07/11 04:50:11 0xdF2eeBaa127019dc84d0725EFD84b77616276D99 166116"
				re := regexp.MustCompile(`(\d+)$`)
				matches := re.FindStringSubmatch(line)

				number, _ := strconv.ParseInt(matches[1], 10, 64)
				count += number
			}
			result1 := count / division
			fmt.Println(strconv.FormatInt(result1, 10))

			file, _ = os.Open(path + "/random_read_result" + strconv.Itoa(j) + "_cache.txt")

			count = int64(0)
			scanner = bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 定义正则表达式
				//str := "2023/07/11 04:50:11 0xdF2eeBaa127019dc84d0725EFD84b77616276D99 166116"
				re := regexp.MustCompile(`(\d+)$`)
				matches := re.FindStringSubmatch(line)

				number, _ := strconv.ParseInt(matches[1], 10, 64)
				count += number
			}
			result2 := count / division
			fmt.Println(strconv.FormatInt(result2, 10))

			writer.Write([]string{strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10)})
			writer.Flush()

			file.Close()
			csvfile.Close()
		}
		fmt.Println()
	}
}

func Txt2CsvForRandomWrite() {
	for i := 1; i <= 6; i++ {
		var (
			path     string
			division int64
		)

		switch i {
		case 1:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1W/experiment"
				division = 10000
			}
		case 2:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10W/experiment"
				division = 100000
			}
		case 3:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_100W/experiment"
				division = 1000000
			}
		case 4:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886/experiment"
				division = 1000000
			}
		case 5:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W/experiment"
				division = 1000000
			}
		case 6:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W/experiment"
				division = 1000000
			}
		}

		for j := 1; j <= 10; j++ {
			file, _ := os.Open(path + "/random_write_result" + strconv.Itoa(j) + ".txt")
			csvfile, _ := os.OpenFile(path+"/random write result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			writer := csv.NewWriter(csvfile)

			count1 := int64(0)
			count2 := int64(0)
			count3 := int64(0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 定义正则表达式
				//str := "2023/07/11 12:44:36 0x0000944fc866186055e47c4db9c2d2db32e9d3de29588b727adcc767bd3615f7 268712 40415 22580"
				re := regexp.MustCompile(`(\d+) (\d+) (\d+)$`)
				matches := re.FindStringSubmatch(line)

				number1, _ := strconv.ParseInt(matches[1], 10, 64)
				number2, _ := strconv.ParseInt(matches[2], 10, 64)
				number3, _ := strconv.ParseInt(matches[3], 10, 64)

				count1 += number1
				count2 += number2
				count3 += number3

			}
			result1 := count1 / division
			result2 := count2 / division
			result3 := count3 / division

			fmt.Println(strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10), strconv.FormatInt(result3, 10))

			file, _ = os.Open(path + "/random_write_result" + strconv.Itoa(j) + "_cache.txt")

			count1 = int64(0)
			count2 = int64(0)
			count3 = int64(0)
			scanner = bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 定义正则表达式
				//str := "2023/07/11 12:44:36 0x0000944fc866186055e47c4db9c2d2db32e9d3de29588b727adcc767bd3615f7 268712 40415 22580"
				re := regexp.MustCompile(`(\d+) (\d+) (\d+)$`)
				matches := re.FindStringSubmatch(line)

				number1, _ := strconv.ParseInt(matches[1], 10, 64)
				number2, _ := strconv.ParseInt(matches[2], 10, 64)
				number3, _ := strconv.ParseInt(matches[3], 10, 64)

				count1 += number1
				count2 += number2
				count3 += number3
			}
			result4 := count1 / division
			result5 := count2 / division
			result6 := count3 / division

			fmt.Println(strconv.FormatInt(result4, 10), strconv.FormatInt(result5, 10), strconv.FormatInt(result6, 10))

			writer.Write([]string{strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10), strconv.FormatInt(result3, 10), strconv.FormatInt(result4, 10), strconv.FormatInt(result5, 10), strconv.FormatInt(result6, 10)})
			writer.Flush()

			file.Close()
			csvfile.Close()
		}
		fmt.Println()
	}
}

func Csv2CsvForSequentialRead() {
	file, _ := os.OpenFile(filepath+"/sequential read result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1W/experiment"
			}
		case 2:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10W/experiment"
			}
		case 3:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_100W/experiment"
			}
		case 4:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886/experiment"
			}
		case 5:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W/experiment"
			}
		case 6:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W/experiment"
			}
		}

		csvfile, _ := os.Open(path + "/sequential read result.csv")
		defer csvfile.Close()

		scanner := bufio.NewScanner(csvfile)

		//scanner.Scan()
		//line := scanner.Text()
		//results := strings.Split(line, ",")
		//writer.Write([]string{results[0], results[1]})

		number1, number2 := int64(0), int64(0)
		for scanner.Scan() {
			line := scanner.Text()
			results := strings.Split(line, ",")

			count, _ := strconv.ParseInt(results[0], 10, 64)
			number1 += count

			count, _ = strconv.ParseInt(results[1], 10, 64)
			number2 += count
		}
		number1 /= 10
		number2 /= 10
		writer.Write([]string{strconv.FormatInt(number1, 10), strconv.FormatInt(number2, 10)})
	}
}

func Csv2CsvForSequentialWrite() {
	file, _ := os.OpenFile(filepath+"/sequential write result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1W/experiment"
			}
		case 2:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10W/experiment"
			}
		case 3:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_100W/experiment"
			}
		case 4:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886/experiment"
			}
		case 5:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W/experiment"
			}
		case 6:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W/experiment"
			}
		}

		csvfile, _ := os.Open(path + "/sequential write result.csv")
		defer csvfile.Close()

		scanner := bufio.NewScanner(csvfile)

		//scanner.Scan()
		//line := scanner.Text()
		//results := strings.Split(line, ",")
		//writer.Write([]string{results[0], results[1], results[2], results[3], results[4], results[5]})

		number1, number2, number3, number4, number5, number6 := int64(0), int64(0), int64(0), int64(0), int64(0), int64(0)
		for scanner.Scan() {
			line := scanner.Text()
			results := strings.Split(line, ",")

			count, _ := strconv.ParseInt(results[0], 10, 64)
			number1 += count

			count, _ = strconv.ParseInt(results[1], 10, 64)
			number2 += count

			count, _ = strconv.ParseInt(results[2], 10, 64)
			number3 += count

			count, _ = strconv.ParseInt(results[3], 10, 64)
			number4 += count

			count, _ = strconv.ParseInt(results[4], 10, 64)
			number5 += count

			count, _ = strconv.ParseInt(results[5], 10, 64)
			number6 += count
		}
		number1 /= 10
		number2 /= 10
		number3 /= 10
		number4 /= 10
		number5 /= 10
		number6 /= 10
		writer.Write([]string{strconv.FormatInt(number1, 10), strconv.FormatInt(number2, 10), strconv.FormatInt(number3, 10), strconv.FormatInt(number4, 10), strconv.FormatInt(number5, 10), strconv.FormatInt(number6, 10)})
	}
}

func Csv2CsvForRandomRead() {
	file, _ := os.OpenFile(filepath+"/random read result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1W/experiment"
			}
		case 2:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10W/experiment"
			}
		case 3:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_100W/experiment"
			}
		case 4:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886/experiment"
			}
		case 5:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W/experiment"
			}
		case 6:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W/experiment"
			}
		}

		csvfile, _ := os.Open(path + "/random read result.csv")
		defer csvfile.Close()

		scanner := bufio.NewScanner(csvfile)

		//scanner.Scan()
		//line := scanner.Text()
		//results := strings.Split(line, ",")
		//writer.Write([]string{results[0], results[1]})

		number1, number2 := int64(0), int64(0)
		for scanner.Scan() {
			line := scanner.Text()
			results := strings.Split(line, ",")

			count, _ := strconv.ParseInt(results[0], 10, 64)
			number1 += count

			count, _ = strconv.ParseInt(results[1], 10, 64)
			number2 += count
		}
		number1 /= 10
		number2 /= 10
		writer.Write([]string{strconv.FormatInt(number1, 10), strconv.FormatInt(number2, 10)})
	}
}

func Csv2CsvForRandomWrite() {
	file, _ := os.OpenFile(filepath+"/random write result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1W/experiment"
			}
		case 2:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10W/experiment"
			}
		case 3:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_100W/experiment"
			}
		case 4:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_2834886/experiment"
			}
		case 5:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_1000W/experiment"
			}
		case 6:
			{
				path = "/home/fuzh/leveldb/state/trie_leveldb_in_10000W/experiment"
			}
		}

		csvfile, _ := os.Open(path + "/random write result.csv")
		defer csvfile.Close()

		scanner := bufio.NewScanner(csvfile)

		//scanner.Scan()
		//line := scanner.Text()
		//results := strings.Split(line, ",")
		//writer.Write([]string{results[0], results[1], results[2], results[3], results[4], results[5]})

		number1, number2, number3, number4, number5, number6 := int64(0), int64(0), int64(0), int64(0), int64(0), int64(0)
		for scanner.Scan() {
			line := scanner.Text()
			results := strings.Split(line, ",")

			count, _ := strconv.ParseInt(results[0], 10, 64)
			number1 += count

			count, _ = strconv.ParseInt(results[1], 10, 64)
			number2 += count

			count, _ = strconv.ParseInt(results[2], 10, 64)
			number3 += count

			count, _ = strconv.ParseInt(results[3], 10, 64)
			number4 += count

			count, _ = strconv.ParseInt(results[4], 10, 64)
			number5 += count

			count, _ = strconv.ParseInt(results[5], 10, 64)
			number6 += count
		}
		number1 /= 10
		number2 /= 10
		number3 /= 10
		number4 /= 10
		number5 /= 10
		number6 /= 10
		writer.Write([]string{strconv.FormatInt(number1, 10), strconv.FormatInt(number2, 10), strconv.FormatInt(number3, 10), strconv.FormatInt(number4, 10), strconv.FormatInt(number5, 10), strconv.FormatInt(number6, 10)})
	}
}
