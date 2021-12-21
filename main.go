package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func timeit(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s duration: %+v\n", name, elapsed)
}

func part1(fd *os.File) (result int, err error) {
	var pos, dep int
	defer timeit(time.Now(), "part1")
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		direction := strings.ToLower(fields[0])
		val, err := strconv.Atoi(fields[1])
		check(err)
		switch direction {
		case "forward":
			pos += val
		case "up":
			dep -= val
		case "down":
			dep += val
		}
	}
	err = scanner.Err()
	check(err)
	result = pos * dep
	return result, err
}

func part2(fd *os.File) (result int, err error) {
	var pos, aim, dep int
	defer timeit(time.Now(), "part2")
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		direction := strings.ToLower(fields[0])
		val, err := strconv.Atoi(fields[1])
		check(err)
		switch direction {
		case "forward":
			pos += val
			dep += (val * aim)
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}
	err = scanner.Err()
	check(err)
	result = pos * dep
	return result, err
}

func main() {
	defer timeit(time.Now(), "main")
	if len(os.Args) != 2 {
		fmt.Println("please provide a filename argument")
		os.Exit(1)
	}
	filename := os.Args[1]

	fd, err := os.Open(filename)
	defer fd.Close()
	check(err)

	result1, err := part1(fd)
	fmt.Printf("part1 result: %+v\n", result1)

	fd.Seek(0, io.SeekStart)

	result2, err := part2(fd)
	fmt.Printf("part2 result: %+v\n", result2)
}
