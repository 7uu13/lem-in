
---

# Lem-in: Ant Farm Simulation in Go

**Lem-in** is a Go program that simulates an ant farm. It reads input from a file, describing ants and a colony, and finds the quickest path for ants to traverse the colony from the start room to the end room.

## Table of Contents

- [Introduction](#introduction)
- [How It Works](#how-it-works)
- [Usage](#usage)
- [Input Format](#input-format)
- [Output Format](#output-format)
- [Examples](#examples)
- [Installation](#installation)
- [Running The Program](#running-the-program)

## Introduction

In this project, you create an ant farm with tunnels and rooms. You place ants in one room and observe how they find the exit. The goal is to find the quickest way to get "n" ants across the colony, composed of rooms and tunnels. The shortest path may not always be the simplest, and the program must handle various colony configurations.

## How It Works

Here's an overview of how the program works:

- At the start of the game, all ants are in the room `##start`, and the objective is to bring them to the room `##end` with as few moves as possible.

- The program finds the quickest path for ants to traverse the colony. Some colonies may have many rooms, links, and specific constraints, which the program must handle.

- It displays the results in the standard output in the following format:
    ```
    number_of_ants
    the_rooms
    the_links
    
    Lx-y Lz-w Lr-o ...
    ```
    - `x`, `z`, `r` represent the ant numbers (from 1 to `number_of_ants`).
    - `y`, `w`, `o` represent the room names.
    - Rooms are defined as "name coord_x coord_y," and links are defined as "name1-name2."

## Usage

To run the program, use the following command:

```bash
go run . input_file.txt
```

- Replace `input_file.txt` with the path to your input file.

## Input Format

The input file should follow these guidelines:

- Define rooms, tunnels, and ants.
- Rooms should not start with "L" or "#" and should have no spaces.
- Use tunnels to connect rooms, with each tunnel connecting only two rooms.
- Ensure that two rooms have at most one tunnel connecting them.
- Specify the number of ants at the beginning of the file.
- Include a `##start` room and a `##end` room.

## Output Format

The program outputs results in the following format:

```
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

- `number_of_ants` is the number of ants in the colony.
- `the_rooms` lists the rooms in the colony.
- `the_links` lists the links (tunnels) between rooms.
- `Lx-y Lz-w Lr-o ...` shows the moves of each ant from room to room.

## Examples

Here are some example usages of the program:

### Example 1

```bash
go run . test0.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
$
```

### Example 2

```bash
go run . test1.txt
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
$
```

### Example 3

```bash
go run . test2.txt
3
2 5 0
##start
0 1 2
##end
1 9 2
3 5 4
0-2
0-3
2-1
3-1
2-3

L1-2 L2-3
L1-1 L2-1 L3-2
L3-1
$
```

## Installation

You can install the program by cloning this repository:

```bash
git clone https://01.kood.tech/git/ktuule/lem-in
```

## Running The Program
Please save your time and run the automated test-script. \
By any means, you can type each file individually, but i just wanted to save your time:
```bash
bash test_script.sh
```

