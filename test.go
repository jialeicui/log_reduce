package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func deal_time(lhs, rhs string) int64 {
	l, _ := time.LoadLocation("Local")
	time_str := lhs + " " + rhs
	time_str = strings.Trim(time_str, "[]")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, l)

	return t.Unix()
}

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

func deal_ip_port(ip_port string) (int, int) {
	ip_port = strings.Trim(ip_port, ":")
	ip_port = strings.Replace(ip_port, "#", ":", 1)
	host, port_str, _ := net.SplitHostPort(ip_port)
	ip_o := net.ParseIP(host)

	port, _ := strconv.Atoi(port_str)
	return inet_aton(ip_o), port
}

func main() {
	line := "[2014-12-04 13:27:54.758 LEVEL = QUERY] client 61.220.4.123#63013: view chunghwa-telecom: query:int.dpool.sina.com.cn. IN 1"
	depart := strings.Split(line, " ")
	t := deal_time(depart[0], depart[1])
	a, p := deal_ip_port(depart[6])
	_ = t
	fmt.Print(a, p)
}
