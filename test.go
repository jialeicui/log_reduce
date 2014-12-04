package main

import (
	"fmt"
	"net"
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

func deal_ip_port(ip_port string) (int, int) {
	ip_port = strings.Trim(ip_port, ":")
	ip_port = strings.Replace(ip_port, "#", ":", 1)
	host, port, _ := net.SplitHostPort(ip_port)
	ip := net.ParseIP(host)
	fmt.Print(ip)
	_ = port
	return 1, 2
}

func main() {
	line := "[2014-12-04 13:27:54.758 LEVEL = QUERY] client 61.220.4.123#63013: view chunghwa-telecom: query:int.dpool.sina.com.cn. IN 1"
	depart := strings.Split(line, " ")
	t := deal_time(depart[0], depart[1])
	a, _ := deal_ip_port(depart[6])
	_ = t
	_ = a
}
