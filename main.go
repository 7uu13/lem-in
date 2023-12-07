package main

import (
	"ant/methods"
	"fmt"
	"os"
)

func main() {
	// Check if the correct number of command-line arguments is provided.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	scannedFile, err := methods.ScanFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	rules, err := methods.ConvertFileToRules(scannedFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	links := map[string][]string{}
	for _, link := range rules.Links {
		links[link.RoomA] = append(links[link.RoomA], link.RoomB)
		links[link.RoomB] = append(links[link.RoomB], link.RoomA)
	}

	shortestPaths := map[string][][]string{}
	for _, room := range links[rules.Start.Room] {
		visited := map[string]bool{}
		for i := 0; i < len(links[rules.Start.Room]); i++ {
			shortestPath := methods.BFS(room, rules, links, visited)
			if shortestPath != nil {
				shortestPaths[room] = append(shortestPaths[room], shortestPath)
			}
			visited = methods.ConvertShortestPathToVisited(rules.End.Room, rules.Start.Room, visited, shortestPath)
		}
	}

	bestPath := methods.ChooseAndConvertShortestPath(shortestPaths, rules)
	if len(bestPath) == 0 {
		fmt.Println("ERROR: invalid data format, no paths found.")
		return
	}

	methods.RunAntsThroughTunnels(bestPath, rules)

}
