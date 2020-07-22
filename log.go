package main

import (
	"fmt"
	"github.com/TwinProduction/go-color"
	"time"
)

func LogInfo(reference, data string) {
	fmt.Println(color.Ize(color.Green, time.Now().Format("2006-01-02 15:04:05.000")+" ["+reference+"] --INF-- "+data))
}

func LogError(reference, data string) {
	fmt.Println(color.Ize(color.Red, time.Now().Format("2006-01-02 15:04:05.000")+" ["+reference+"] --INF-- "+data))
}
