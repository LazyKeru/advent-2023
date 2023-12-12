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
	map_regex, err := regexp.Compile("[A-Z]{3}[ ]*=[ ]*[(][A-Z]{3}[ ]*,[ ]*[A-Z]{3}[)]")
	if err != nil {
		log.Fatal(err)
	}
	map_reader_regex, err := regexp.Compile("[A-Z]{3}")
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

func main() {
	nodes_list, instructions_for_problem := serializer("input.txt")
	log.Print(apply(instructions_for_problem, nodes_list))
}
