package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func actionSelect() {
	fmt.Printf("action(swap, deposit, withdraw): ")
	inputReaderUser := bufio.NewReader(os.Stdin)
	action, _ = inputReaderUser.ReadString('\r')
	action = strings.Replace(action, "\r", "", -1)
}

func swapSelect() (int, float64) {
	fmt.Printf("需要花费币的序号(0 or 1): ")
	inputReaderIndex := bufio.NewReader(os.Stdin)
	indexStr, _ := inputReaderIndex.ReadString('\r')
	indexStr = strings.Replace(indexStr, "\r", "", -1)
	index, _ := strconv.Atoi(indexStr)

	fmt.Printf("需要花费的量: ")
	inputReaderValue := bufio.NewReader(os.Stdin)
	valueStr, _ := inputReaderValue.ReadString('\r')
	valueStr = strings.Replace(valueStr, "\r", "", -1)
	value, _ := strconv.ParseFloat(valueStr, 64)

	return index, value
}

func depositSelect() (int, float64) {
	fmt.Printf("需要抵押的币的序号(0 or 1): ")
	inputReaderIndex := bufio.NewReader(os.Stdin)
	indexStr, _ := inputReaderIndex.ReadString('\r')
	indexStr = strings.Replace(indexStr, "\r", "", -1)
	index, _ := strconv.Atoi(indexStr)

	fmt.Printf("需要抵押的量: ")
	inputReaderValue := bufio.NewReader(os.Stdin)
	valueStr, _ := inputReaderValue.ReadString('\r')
	valueStr = strings.Replace(valueStr, "\r", "", -1)
	value, _ := strconv.ParseFloat(valueStr, 64)
	return index, value
}
