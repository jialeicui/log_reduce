package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type proc_log func(line string)

// 处理时间
func deal_time(lhs, rhs string) int64 {
	l, _ := time.LoadLocation("Local")
	time_str := lhs + " " + rhs
	time_str = strings.Trim(time_str, "[]")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, l)

	return t.Unix()
}

// ip to int
func inet_aton(ipnr net.IP) int {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int

	sum += int(b0) << 24
	sum += int(b1) << 16
	sum += int(b2) << 8
	sum += int(b3)

	return sum
}

// 处理ip和端口
func deal_ip_port(ip_port string) (int, int) {
	ip_port = strings.Trim(ip_port, ":")
	ip_port = strings.Replace(ip_port, "#", ":", 1)
	host, port_str, _ := net.SplitHostPort(ip_port)
	ip_o := net.ParseIP(host)

	port, _ := strconv.Atoi(port_str)
	return inet_aton(ip_o), port
}

func deal_view(view string) int {
	// 后期需要从数据库中查找
	view = strings.Trim(view, ":")
	return math.MaxInt32
}

func deal_query(query string) int32 {
	// 后期需要从数据库中查找
	s := strings.Split(query, ":")
	query = s[1]
	return math.MaxInt32
}

func foreach_line(filename string, proc_func proc_log) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proc_func(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
		return
	}
}

func main() {
	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	proc_one_line := func(line string) {
		depart := strings.Split(line, " ")
		t := deal_time(depart[0], depart[1])
		ip, port := deal_ip_port(depart[6])
		view := deal_view(depart[8])
		domain := deal_query(depart[9])
		qtype, _ := strconv.Atoi(depart[11])
		if false {
			res := fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", t, ip, port, view, domain, qtype)
			fo.WriteString(res)
		} else {
			// 由于binary.Write需要明确长度的值,所以这里需要进行int的转换
			var data = []interface{}{t, int32(ip), int32(port), int32(view), int32(domain), int32(qtype)}
			for _, v := range data {
				err := binary.Write(fo, binary.LittleEndian, v)
				if err != nil {
					fmt.Println("binary.Write failed:", err)
				}
			}
		}
	}

	foreach_line("log", proc_one_line)
}
