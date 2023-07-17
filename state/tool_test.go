package state

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTxt2Csv(t *testing.T) {
	mapping := make(map[int]output, 1001)

	file1, _ := os.Open("file/output1.txt")
	defer file1.Close()

	// 创建一个带缓冲的读取器
	scanner := bufio.NewScanner(file1)

	for scanner.Scan() {
		line := scanner.Text()

		// 定义正则表达式
		//str := "2023-07-06 02:28:01.126553187 The number of accounts has achieved 0 The time of the commitment is 0"
		re := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{9}).*The number of accounts has achieved (\d+).*The time of the commitment is (\d+)`)

		matches := re.FindStringSubmatch(line)

		if len(matches) >= 4 {
			// 提取时间和剩下的两个数据
			data1 := matches[1]
			data2 := matches[2]
			data3 := matches[3]

			// 将时间字符串解析为Time类型
			tmp, _ := time.Parse("2006-01-02 15:04:05.999999999", data1)
			time1 := tmp.UnixNano()
			time2, _ := strconv.ParseInt(data3, 10, 64)

			fmt.Println(time1, data2, time2)
			index, _ := strconv.Atoi(data2)
			mapping[index] = output{time1, time2}
		}
	}

	basetime := mapping[0].Time1

	for i := 0; i <= 100000000; i += 100000 {
		tmp := mapping[i].Time1
		time1 := mapping[i].Time1
		time2 := mapping[i].Time2
		time1 -= basetime
		basetime = tmp
		mapping[i] = output{time1, time2}
	}

	file2, _ := os.Create("file/output1.csv")
	defer file2.Close()

	file3, _ := os.Create("file/output2.csv")
	defer file3.Close()

	// 创建CSV写入器
	writer1 := csv.NewWriter(file2)
	writer2 := csv.NewWriter(file3)

	for i := 0; i <= 100000000; i += 100000 {
		//fmt.Println(i, mapping[i].Time1, mapping[i].Time2)
		writer1.Write([]string{strconv.Itoa(i), strconv.FormatInt(mapping[i].Time1, 10)})
		writer2.Write([]string{strconv.Itoa(i), strconv.FormatInt(mapping[i].Time2, 10)})
	}

	// 刷新缓冲区
	writer1.Flush()
	writer2.Flush()
}

func TestTxt2CsvForSequentialRead(t *testing.T) {
	for i := 1; i <= 6; i++ {
		var (
			path     string
			division int64
		)
		switch i {
		case 1:
			{
				path = "1W accounts"
				division = 10000
			}
		case 2:
			{
				path = "10W accounts"
				division = 100000
			}
		case 3:
			{
				path = "100W accounts"
				division = 100000
			}
		case 4:
			{
				path = "2834886 accounts"
				division = 100000
			}
		case 5:
			{
				path = "1000W accounts"
				division = 100000
			}
		case 6:
			{
				path = "10000W accounts"
				division = 100000
			}

		}

		for j := 1; j <= 5; j++ {
			file, _ := os.Open("file/tmp/" + path + "/sequential_read_result" + strconv.Itoa(j) + ".txt")
			csvfile, _ := os.OpenFile("file/tmp/"+path+"/sequential read result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

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

			file, _ = os.Open("file/tmp/" + path + "/sequential_read_result" + strconv.Itoa(j) + "_cache.txt")

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

func TestTxt2CsvForSequentialWrite(t *testing.T) {
	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "1W accounts"
			}
		case 2:
			{
				path = "10W accounts"
			}
		case 3:
			{
				path = "100W accounts"
			}
		case 4:
			{
				path = "2834886 accounts"
			}
		case 5:
			{
				path = "1000W accounts"
			}
		case 6:
			{
				path = "10000W accounts"
			}

		}

		for j := 1; j <= 5; j++ {
			file, _ := os.Open("file/" + path + "/sequential_write_result" + strconv.Itoa(j) + ".txt")
			csvfile, _ := os.OpenFile("file/"+path+"/sequential write result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

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
			result1 := count1 / 100000
			result2 := count2 / 100000
			result3 := count3 / 100000

			fmt.Println(strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10), strconv.FormatInt(result3, 10))

			file, _ = os.Open("file/" + path + "/sequential_write_result" + strconv.Itoa(j) + "_cache.txt")

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
			result4 := count1 / 100000
			result5 := count2 / 100000
			result6 := count3 / 100000

			fmt.Println(strconv.FormatInt(result4, 10), strconv.FormatInt(result5, 10), strconv.FormatInt(result6, 10))

			writer.Write([]string{strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10), strconv.FormatInt(result3, 10), strconv.FormatInt(result4, 10), strconv.FormatInt(result5, 10), strconv.FormatInt(result6, 10)})
			writer.Flush()

			file.Close()
			csvfile.Close()
		}
		fmt.Println()
	}
}

func TestTxt2CsvForRandomRead(t *testing.T) {
	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "1W accounts"
			}
		case 2:
			{
				path = "10W accounts"
			}
		case 3:
			{
				path = "100W accounts"
			}
		case 4:
			{
				path = "2834886 accounts"
			}
		case 5:
			{
				path = "1000W accounts"
			}
		case 6:
			{
				path = "10000W accounts"
			}

		}

		for j := 1; j <= 5; j++ {
			file, _ := os.Open("file/" + path + "/random_read_result" + strconv.Itoa(j) + ".txt")
			csvfile, _ := os.OpenFile("file/"+path+"/random read result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

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
			result1 := count / 100000
			fmt.Println(strconv.FormatInt(result1, 10))

			file, _ = os.Open("file/" + path + "/random_read_result" + strconv.Itoa(j) + "_cache.txt")

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
			result2 := count / 100000
			fmt.Println(strconv.FormatInt(result2, 10))

			writer.Write([]string{strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10)})
			writer.Flush()

			file.Close()
			csvfile.Close()
		}
		fmt.Println()
	}
}

func TestTxt2CsvForRandomWrite(t *testing.T) {
	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "1W accounts"
			}
		case 2:
			{
				path = "10W accounts"
			}
		case 3:
			{
				path = "100W accounts"
			}
		case 4:
			{
				path = "2834886 accounts"
			}
		case 5:
			{
				path = "1000W accounts"
			}
		case 6:
			{
				path = "10000W accounts"
			}

		}

		for j := 1; j <= 5; j++ {
			file, _ := os.Open("file/" + path + "/random_write_result" + strconv.Itoa(j) + ".txt")
			csvfile, _ := os.OpenFile("file/"+path+"/random write result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

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
			result1 := count1 / 100000
			result2 := count2 / 100000
			result3 := count3 / 100000

			fmt.Println(strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10), strconv.FormatInt(result3, 10))

			file, _ = os.Open("file/" + path + "/random_write_result" + strconv.Itoa(j) + "_cache.txt")

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
			result4 := count1 / 100000
			result5 := count2 / 100000
			result6 := count3 / 100000

			fmt.Println(strconv.FormatInt(result4, 10), strconv.FormatInt(result5, 10), strconv.FormatInt(result6, 10))

			writer.Write([]string{strconv.FormatInt(result1, 10), strconv.FormatInt(result2, 10), strconv.FormatInt(result3, 10), strconv.FormatInt(result4, 10), strconv.FormatInt(result5, 10), strconv.FormatInt(result6, 10)})
			writer.Flush()

			file.Close()
			csvfile.Close()
		}
		fmt.Println()
	}
}

func TestCsv2CsvForSequentialRead(t *testing.T) {
	file, _ := os.OpenFile("file/tmp/sequential read result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "1W accounts"
			}
		case 2:
			{
				path = "10W accounts"
			}
		case 3:
			{
				path = "100W accounts"
			}
		case 4:
			{
				path = "2834886 accounts"
			}
		case 5:
			{
				path = "1000W accounts"
			}
		case 6:
			{
				path = "10000W accounts"
			}

		}

		csvfile, _ := os.Open("file/tmp/" + path + "/sequential read result.csv")
		defer csvfile.Close()

		scanner := bufio.NewScanner(csvfile)

		scanner.Scan()
		line := scanner.Text()
		results := strings.Split(line, ",")
		writer.Write([]string{results[0], results[1]})

		//number1, number2 := int64(0), int64(0)
		//for scanner.Scan() {
		//	line := scanner.Text()
		//	results := strings.Split(line, ",")
		//
		//	count, _ := strconv.ParseInt(results[0], 10, 64)
		//	number1 += count
		//
		//	count, _ = strconv.ParseInt(results[1], 10, 64)
		//	number2 += count
		//}
		//number1 /= 4
		//number2 /= 4
		//writer.Write([]string{strconv.FormatInt(number1, 10), strconv.FormatInt(number2, 10)})
	}
}

func TestCsv2CsvForSequentialWrite(t *testing.T) {
	file, _ := os.OpenFile("file/sequential write result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "1W accounts"
			}
		case 2:
			{
				path = "10W accounts"
			}
		case 3:
			{
				path = "100W accounts"
			}
		case 4:
			{
				path = "2834886 accounts"
			}
		case 5:
			{
				path = "1000W accounts"
			}
		case 6:
			{
				path = "10000W accounts"
			}

		}

		csvfile, _ := os.Open("file/" + path + "/sequential write result.csv")
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
		number1 /= 4
		number2 /= 4
		number3 /= 4
		number4 /= 4
		number5 /= 4
		number6 /= 4
		writer.Write([]string{strconv.FormatInt(number1, 10), strconv.FormatInt(number2, 10), strconv.FormatInt(number3, 10), strconv.FormatInt(number4, 10), strconv.FormatInt(number5, 10), strconv.FormatInt(number6, 10)})
	}
}

func TestCsv2CsvForRandomRead(t *testing.T) {
	file, _ := os.OpenFile("file/random read result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "1W accounts"
			}
		case 2:
			{
				path = "10W accounts"
			}
		case 3:
			{
				path = "100W accounts"
			}
		case 4:
			{
				path = "2834886 accounts"
			}
		case 5:
			{
				path = "1000W accounts"
			}
		case 6:
			{
				path = "10000W accounts"
			}

		}

		csvfile, _ := os.Open("file/" + path + "/random read result.csv")
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
		number1 /= 4
		number2 /= 4
		writer.Write([]string{strconv.FormatInt(number1, 10), strconv.FormatInt(number2, 10)})
	}
}

func TestCsv2CsvForRandomWrite(t *testing.T) {
	file, _ := os.OpenFile("file/random write result.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 6; i++ {
		var path string
		switch i {
		case 1:
			{
				path = "1W accounts"
			}
		case 2:
			{
				path = "10W accounts"
			}
		case 3:
			{
				path = "100W accounts"
			}
		case 4:
			{
				path = "2834886 accounts"
			}
		case 5:
			{
				path = "1000W accounts"
			}
		case 6:
			{
				path = "10000W accounts"
			}

		}

		csvfile, _ := os.Open("file/" + path + "/random write result.csv")
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
		number1 /= 4
		number2 /= 4
		number3 /= 4
		number4 /= 4
		number5 /= 4
		number6 /= 4
		writer.Write([]string{strconv.FormatInt(number1, 10), strconv.FormatInt(number2, 10), strconv.FormatInt(number3, 10), strconv.FormatInt(number4, 10), strconv.FormatInt(number5, 10), strconv.FormatInt(number6, 10)})
	}
}
