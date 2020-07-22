package main

import (
	"fmt"
	"github.com/TwinProduction/go-color"
)

func LogInfo(reference, data string) {
	fmt.Println(color.Ize(color.Green, "["+reference+"] --INF-- "+data))
}

func LogError(reference, data string) {
	fmt.Println(color.Ize(color.Red, "["+reference+"] --INF-- "+data))
}
