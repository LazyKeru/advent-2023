package main

import (
	"log"
	"os"
	"regexp"
)

type node struct {
	name            string
	left_node_name  string
	right_node_name string
}
type instructions struct {
	LR string
}

func serializer(path string) ([]node, instructions) {
	var _instructions instructions
	var _nodes []node
	_body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	body := string(_body)
	instructions_regex, err := regexp.Compile("[LR]+")
	if err != nil {
		log.Fatal(err)
	}
	map_regex, err := regexp.Compile("[A-Z1-9]{3}[ ]*=[ ]*[(][A-Z1-9]{3}[ ]*,[ ]*[A-Z1-9]{3}[)]")
	if err != nil {
		log.Fatal(err)
	}
	map_reader_regex, err := regexp.Compile("[A-Z1-9]{3}")
	if err != nil {
		log.Fatal(err)
	}
	_instructions.LR = instructions_regex.FindString(body)
	_raw_maps := map_regex.FindAllString(body, -1)

	for _, _raw_map := range _raw_maps {
		tmp := map_reader_regex.FindAllString(_raw_map, -1)
		if len(tmp) != 3 {
			log.Fatal("Not enough elements in your map. It ain't the marauder that will autocomplete.")
		}
		_nodes = append(_nodes, node{tmp[0], tmp[1], tmp[2]})
	}

	return _nodes, _instructions
}

func apply(steps instructions, nodes []node) int {
	step := "AAA"
	target := "ZZZ"
	var i int
	for i = 0; step != target; i++ {
		for j := 0; j < len(nodes); j++ {
			if nodes[j].name == step {
				switch string(steps.LR[i%len(steps.LR)]) {
				case "L":
					step = nodes[j].left_node_name
				case "R":
					step = nodes[j].right_node_name
				}
				break
			}
		}
	}
	return i
}

func apply_as_ghost(steps instructions, nodes []node) int {
	// Find all starting nodes
	var starting_steps []string
	for _, node := range nodes {
		if string(node.name[2]) == "A" {
			starting_steps = append(starting_steps, node.name)
		}
	}
	var first_iterations_steps []int
	for _, starting_step := range starting_steps {
		first_iterations_steps = append(first_iterations_steps, step_ghost(starting_step, steps, nodes, 0))
	}
	return LCM(first_iterations_steps[0], first_iterations_steps[1], first_iterations_steps[2:]...)
}

func step_ghost(current_step string, steps instructions, nodes []node, index int) int {
	var next_step string
	for i := 0; i < len(nodes); i++ {
		if nodes[i].name == current_step {
			switch string(steps.LR[index%len(steps.LR)]) {
			case "L":
				next_step = nodes[i].left_node_name
			case "R":
				next_step = nodes[i].right_node_name
			}
			break
		}
	}
	if string(next_step[2]) != "Z" {
		return step_ghost(next_step, steps, nodes, index+1)
	}
	return index + 1
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	nodes_list, instructions_for_problem := serializer("input.txt")
	log.Print(apply(instructions_for_problem, nodes_list))
	nodes_list, instructions_for_problem = serializer("input.txt")
	log.Print(apply_as_ghost(instructions_for_problem, nodes_list))
}
