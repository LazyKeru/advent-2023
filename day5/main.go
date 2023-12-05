package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type locations struct {
	location []int
}

type location_range struct {
	start int
	end   int
}

type converter struct {
	destination_range_start int
	source_range_start      int
	range_length            int
}

type converter_map struct {
	converters []converter
}

func string_list_to_int_list(list []string) []int {
	var res []int
	for i := 0; i < len(list); i++ {
		tmp, err := strconv.Atoi(list[i])
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, tmp)
	}
	return res
}

func input(path string) (locations, []converter_map) {
	_body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	body := string(_body)
	seeds_regex, err := regexp.Compile("seeds: [0-9 ]*")
	maps_regex, err := regexp.Compile("[a-z -]*:[\n\r]+([0-9 ]*[\n\r]+)*")
	number, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	var res_source locations
	var res_converter_maps []converter_map
	res_source.location = string_list_to_int_list(number.FindAllString((seeds_regex.FindString(body)), -1))
	tmp_raw_maps := maps_regex.FindAllString(body, -1)
	for _, raw_map := range tmp_raw_maps {
		lines := strings.Split(raw_map, "\n")
		var tmp_converter_map converter_map
		for _, line := range lines {
			tmp := string_list_to_int_list(number.FindAllString(line, -1))
			if len(tmp) == 3 {
				var tmp_converter converter
				tmp_converter.destination_range_start = tmp[0]
				tmp_converter.source_range_start = tmp[1]
				tmp_converter.range_length = tmp[2]
				tmp_converter_map.converters = append(tmp_converter_map.converters, tmp_converter)
			}
		}
		res_converter_maps = append(res_converter_maps, tmp_converter_map)
	}
	return res_source, res_converter_maps
}

func find_destination(source int, converter_maps []converter_map) int {
	destination := source
	for _, converter_map := range converter_maps {
		for _, converter := range converter_map.converters {
			if converter.source_range_start <= destination && destination <= converter.source_range_start+converter.range_length {
				destination = (destination - converter.source_range_start) + converter.destination_range_start
				break
			}
		}
	}
	return destination
}

// i for converter_maps
// k for map
func destination_range(source_range location_range, converter_maps []converter_map, i int, k int) []location_range {
	var destination_ranges []location_range
	if !(i == len(converter_maps) || k == len(converter_maps[i].converters)) {
		map_source_start := converter_maps[i].converters[k].source_range_start
		map_source_end := converter_maps[i].converters[k].source_range_start + converter_maps[i].converters[k].range_length - 1
		map_destination_start := converter_maps[i].converters[k].destination_range_start
		map_destination_end := converter_maps[i].converters[k].destination_range_start + converter_maps[i].converters[k].range_length - 1
		if map_source_start <= source_range.start && source_range.end <= map_source_end {
			tmp1 := destination_range(location_range{(source_range.start - map_source_start) + map_destination_start, (source_range.end - map_source_start) + map_destination_start}, converter_maps, i+1, 0)
			destination_ranges = append(destination_ranges, tmp1...)
		} else if source_range.start <= map_source_start && map_source_start <= source_range.end && source_range.end <= map_source_end {
			tmp1 := destination_range(location_range{map_destination_start, (source_range.end - map_source_start) + map_destination_start}, converter_maps, i+1, 0)
			tmp2 := destination_range(location_range{source_range.start, map_source_start - 1}, converter_maps, i, k+1)
			destination_ranges = append(destination_ranges, tmp1...)
			destination_ranges = append(destination_ranges, tmp2...)
		} else if map_source_start <= source_range.start && source_range.start <= map_source_end && map_source_end <= source_range.end {
			tmp1 := destination_range(location_range{(source_range.start - map_source_start) + map_destination_start, map_destination_end}, converter_maps, i+1, 0)
			tmp2 := destination_range(location_range{map_source_end + 1, source_range.end}, converter_maps, i, k+1)
			destination_ranges = append(destination_ranges, tmp1...)
			destination_ranges = append(destination_ranges, tmp2...)
		} else if source_range.start <= map_source_start && map_source_end <= source_range.end {
			tmp1 := destination_range(location_range{map_destination_start, map_destination_end}, converter_maps, i+1, 0)
			tmp2 := destination_range(location_range{source_range.start, map_source_start - 1}, converter_maps, i, k+1)
			tmp3 := destination_range(location_range{map_source_end + 1, source_range.end}, converter_maps, i, k+1)
			destination_ranges = append(destination_ranges, tmp1...)
			destination_ranges = append(destination_ranges, tmp2...)
			destination_ranges = append(destination_ranges, tmp3...)
		} else {
			tmp1 := destination_range(location_range{source_range.start, source_range.end}, converter_maps, i, k+1)
			destination_ranges = append(destination_ranges, tmp1...)
		}
	}
	if i == len(converter_maps) {
		destination_ranges = append(destination_ranges, source_range)
	} else if k == len(converter_maps[i].converters) {
		tmp1 := destination_range(location_range{source_range.start, source_range.end}, converter_maps, i+1, 0)
		destination_ranges = append(destination_ranges, tmp1...)
	}
	return destination_ranges
}

func locations_to_ranges(input locations) []location_range {
	if len(input.location)%2 != 0 {
		log.Fatal()
	}
	var output []location_range
	for i := 0; i < len(input.location); i += 2 {
		output = append(output, location_range{input.location[i], input.location[i] + input.location[i+1] - 1})
	}
	return output
}

func main() {
	seeds_locations, converter_maps := input("input.txt")
	var seed_destinations []int
	for _, seed_location := range seeds_locations.location {
		seed_destinations = append(seed_destinations, find_destination(seed_location, converter_maps))
	}
	smallest_destination := seed_destinations[0]
	for i := 1; i < len(seed_destinations); i++ {
		if seed_destinations[i] < smallest_destination {
			smallest_destination = seed_destinations[i]
		}
	}
	log.Print("Part 1, lowest location :")
	log.Print(smallest_destination)
	log.Print("Everyone will starve we need more seeds")
	locations_ranges := locations_to_ranges(seeds_locations)
	var destinations_ranges []location_range
	for _, seed_location_range := range locations_ranges {
		destinations_ranges = append(destinations_ranges, destination_range(seed_location_range, converter_maps, 0, 0)...)
	}
	smallest_destination = destinations_ranges[0].start
	for i := 1; i < len(destinations_ranges); i++ {
		if destinations_ranges[i].start < smallest_destination {
			smallest_destination = destinations_ranges[i].start
		}
	}
	log.Print("Part 2, lowest location :")
	log.Print(smallest_destination)
}
