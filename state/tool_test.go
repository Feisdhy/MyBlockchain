package state

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func TestTxt2Csv(t *testing.T) {
	mapping := make(map[int]output, 1001)

	file1, _ := os.Open("file/output.txt")
	defer file1.Close()

	// 创建一个带缓冲的读取器
	scanner := bufio.NewScanner(file1)

	for scanner.Scan() {
		line := scanner.Text()

		// 定义正则表达式
		//str := "2023/06/29 11:02:20 The number of accounts has achieved 0 The time of the commitment is 0"
		re := regexp.MustCompile(`(\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}).*The number of accounts has achieved (\d+).*The time of the commitment is (\d+)`)

		matches := re.FindStringSubmatch(line)

		if len(matches) >= 4 {
			// 提取时间和剩下的两个数据
			time1 := matches[1]
			data := matches[2]
			time2 := matches[3]

			// 将时间字符串解析为Time类型
			time, _ := time.Parse("2006/01/02 15:04:05", time1)

			// 将时间转换为秒数
			seconds := time.Unix()

			//fmt.Printf("时间：%s\n", time1)
			//fmt.Printf("数据1：%s\n", data)
			//fmt.Printf("数据2：%s\n", time2)

			index, _ := strconv.Atoi(data)
			nanoseconds, _ := strconv.ParseInt(time2, 10, 64)
			mapping[index] = output{seconds, nanoseconds}
		}
	}

	basetime := mapping[0].Time1

	for i := 100000; i <= 100000000; i += 100000 {
		time1 := mapping[i].Time1
		time2 := mapping[i].Time2
		time1 -= basetime
		mapping[i] = output{time1, time2}
	}

	mapping[0] = output{0, mapping[0].Time2}

	file2, _ := os.Create("file/output_1.csv")
	defer file2.Close()

	file3, _ := os.Create("file/output_2.csv")
	defer file3.Close()

	// 创建CSV写入器
	writer1 := csv.NewWriter(file2)
	writer2 := csv.NewWriter(file3)

	writer1.Write([]string{"accounts", "time"})
	writer2.Write([]string{"accounts", "time"})

	for i := 0; i <= 100000000; i += 100000 {
		fmt.Println(i, mapping[i].Time1, mapping[i].Time2)
		writer1.Write([]string{strconv.Itoa(i), strconv.FormatInt(mapping[i].Time1, 10)})
		writer2.Write([]string{strconv.Itoa(i), strconv.FormatInt(mapping[i].Time2, 10)})
	}

	// 刷新缓冲区
	writer1.Flush()
	writer2.Flush()
}
