package methods

import (
	"ant/types"
	"fmt"
)

func RunAntsThroughTunnels(best_path [][]string, rules types.Rules) {
	ants := []types.AntStatus{}
	occupied_rooms_state := map[string]bool{}

	special_state := false
	total_ants := rules.Ants

	// Find the shortest path by length ->
	shortest_path := best_path[0]
	for _, path := range best_path {
		if len(path) < len(shortest_path) {
			shortest_path = path
		}
	}

	path_lengths := make(map[string]int)
	for _, path := range best_path {
		path_lengths[path[0]] = len(path)
	}

	for _, tunnel := range best_path {
		for _, path := range tunnel {
			occupied_rooms_state[path] = false
		}
	}

	for i := 0; i < total_ants; i++ {
		ants = append(ants, types.AntStatus{Ant: i, Tunnel: 0, CurrentRoom: rules.Start.Room, CurrentRoomIndex: 0, Finished: false})
	}

	finished_ants := 0
	running := true
	for running {
		for _, ant := range ants {
			if ant.Finished {
				continue
			}

			if ant.CurrentRoom == rules.Start.Room {
				for tunnel, path := range best_path {
					if !occupied_rooms_state[path[0]] && len(path) > 1 {

						if ant.Ant+1 == rules.Ants {
							if path_lengths[path[0]] > path_lengths[shortest_path[0]] {
								continue
							}
						}

						ant.Tunnel = tunnel
						ant.CurrentRoom = path[0]
						ant.CurrentRoomIndex = 0

						ants[ant.Ant] = ant
						occupied_rooms_state[path[0]] = true
						fmt.Print("L", ant.Ant+1, "-", path[0], " ")
						break
					}

					if len(path) == 1 && !special_state {
						// If the path is only 1 room long and the room is not in a special state
						// Special state is when the room is occupied, without it all ants would go this path

						ant.Finished = true
						special_state = true

						finished_ants += 1
						ants[ant.Ant] = ant
						fmt.Print("L", ant.Ant+1, "-", path[0], " ")
						break
					}

				}
			} else {
				// If the ant is not in the start room
				// Remove the room from the occupied_rooms_state map
				// Move the ant to the next room in the path
				// Add the room to the occupied_rooms_state map

				occupied_rooms_state[ant.CurrentRoom] = false
				ant.CurrentRoomIndex += 1
				ant.CurrentRoom = best_path[ant.Tunnel][ant.CurrentRoomIndex]
				occupied_rooms_state[ant.CurrentRoom] = true
				ants[ant.Ant] = ant
				fmt.Print("L", ant.Ant+1, "-", ant.CurrentRoom, " ")

			}

			if ant.CurrentRoom == rules.End.Room {
				// If the ant is in the end room
				finished_ants += 1
				ant.Finished = true
				ants[ant.Ant] = ant

			}

			if finished_ants == total_ants {
				// If all ants are finished
				running = false
			}
		}

		special_state = false // Reset special state
		fmt.Println()
	}
}

func ChooseAndConvertShortestPath(shortestPaths map[string][][]string, rules types.Rules) [][]string {
	// Chooses the best path from the shortest paths
	// The best path is the one that has the most paths to end

	// Eliminates where the shortest path is the same for all rooms
	for key, paths := range shortestPaths {
		identical_state := true
		last_path := []string{}

		if len(paths) == 1 {
			continue
		}

		for value, path := range paths {

			if value == 0 {
				last_path = path
				continue
			}

			if !AreArraysIdentical(last_path, path) {
				identical_state = false
				break
			}

		}

		if identical_state {
			delete(shortestPaths, key)
		}

	}

	// Doesn't allow path to go through the same room twice, check's if the path goes through Start room!
	for key, paths := range shortestPaths {

		skip_path := false
		start_room := rules.Start.Room

		for value, path := range paths {
			if value == 0 {
				continue
			}

			if path[1] != start_room && len(path) > 3 {
				skip_path = true
				break
			}

		}

		if skip_path {
			delete(shortestPaths, key)
		}
	}

	// Chooses the shortest path from remaining paths
	min_length := 0
	best_path := [][]string{}
	for _, paths := range shortestPaths {
		if len(paths) > min_length {
			min_length = len(paths)
			best_path = paths
		}
	}

	// Removes the start room from the path
	for index, path := range best_path {
		if path[1] == rules.Start.Room {
			best_path[index] = append(best_path[index][:0], best_path[index][2:]...)

		}
	}

	return best_path
}

func ConvertShortestPathToVisited(finish string, start string, visited map[string]bool, shortestPath []string) map[string]bool {
	// Convert shortestPath to map where the rooms are marked as visited
	for _, room := range shortestPath {
		if room == finish || room == start && len(shortestPath) > 3 {
			continue
		}
		visited[room] = true
	}
	return visited
}

func AreArraysIdentical(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func RemoveElement(slice []string, element string) []string {
	for i, item := range slice {
		if item == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func CheckIsFinished(ant_status types.AntStatus, finishRoom string) bool {
	if ant_status.CurrentRoom == finishRoom {
		return true
	}
	return false
}
