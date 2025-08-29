package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fish171204/gin-framework/lesson"
)

func main() {
	lessonPtr := flag.String("lesson", "lesson01", "Select lesson to run")
	flag.Parse()

	switch *lessonPtr {
	case "lesson01":
		fmt.Println("Running Lesson 01: Package HTTP")
		lesson.Lesson01PackageHTTP()
	case "lesson02":
		fmt.Println("Running Lesson 02: Gin Starter")
		lesson.Lesson02GinStarter()
	case "lesson03":
		fmt.Println("Running Lesson 03: Route Group")
		lesson.Lesson03RouteGroup()
	case "":
		fmt.Println("Please specify a lesson to run:")
		fmt.Println("  go run . -lesson=lesson01  (for HTTP package lesson)")
		fmt.Println("  go run . -lesson=lesson02  (for Gin starter lesson)")
		fmt.Println("  go run . -lesson=lesson03  (for Route Group lesson)")
	default:
		log.Fatalf("Unknown lesson: %s. Available lessons: lesson01, lesson02, lesson03", *lessonPtr)
	}
}
