package methods

import (
	"ant/types"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScanFile(fileName string) ([]string, error) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var textLines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}

	return textLines, nil
}

func ConvertFileToRules(textLines []string) (types.Rules, error) {
	// Function that converts the text file to a struct of rules
	Rules := types.Rules{}
	Rules.Ants, _ = strconv.Atoi(textLines[0])

	if Rules.Ants == 0 {
		return Rules, errors.New("error: No ants")
	}

	for i := 1; i < len(textLines); i++ {
		element := textLines[i]
		if element == "" {
			continue
		}
		fchar := string(element[0]) // first character of the line

		if fchar == "L" || fchar == " " {
			continue
		}

		if fchar == "#" {
			if element == "##start" {
				Rules.Start = ConvertTextToRoom(textLines[i+1])
				i++

			} else if element == "##end" {
				Rules.End = ConvertTextToRoom(textLines[i+1])
				i++
			} else {
				// if the line is a comment
				continue
			}
		}

		if len(strings.Split(element, " ")) == 3 {
			// Room
			Rules.Rooms = append(Rules.Rooms, ConvertTextToRoom(element))
		}

		if len(strings.Split(element, "-")) == 2 {
			// Link
			Rules.Links = append(Rules.Links, ConvertTextToLink(element))
		}

	}

	if Rules.Start.Room == "" || Rules.End.Room == "" {
		return Rules, errors.New("error: No start or end room")
	}

	return Rules, nil
}

func ConvertTextToRoom(text string) types.Room {
	text_list := strings.Split(text, " ")
	room := text_list[0]
	cord_x, err2 := strconv.Atoi(text_list[1])
	cord_y, err3 := strconv.Atoi(text_list[2])

	if err2 != nil || err3 != nil {
		fmt.Println("Error:", err2, err3)
	}

	return types.Room{Room: room, Location: [2]int{cord_x, cord_y}}
}

func ConvertTextToLink(text string) types.Link {
	text_list := strings.Split(text, "-")
	room_a := text_list[0]
	room_b := text_list[1]

	return types.Link{RoomA: room_a, RoomB: room_b}
}
