package methods

import "ant/types"

func BFS(start string, rules types.Rules, links map[string][]string, usedrooms map[string]bool) []string {
	// BFS algorithm that returns all possible shortest paths from start to end
	visited := make(map[string]bool)
	for k, v := range usedrooms {
		visited[k] = v
	}

	finished := false
	queue := []string{start}
	parents := map[string]string{}

	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]

		if room == rules.End.Room {
			finished = true
			break
		}
		for _, neighbour := range links[room] {
			if !visited[neighbour] {
				queue = append(queue, neighbour)
				visited[neighbour] = true
				parents[neighbour] = room
			}
		}
	}

	if !finished {
		return nil
	}

	// Backtrack to find the correct path
	shortestPath := []string{rules.End.Room}
	currentVertex := rules.End.Room
	for currentVertex != start {
		parent := parents[currentVertex]
		shortestPath = append([]string{parent}, shortestPath...)
		currentVertex = parent
	}
	return shortestPath
}
