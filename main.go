package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Skill represents a skill with its name and progress
type Skill struct {
	Name     string
	Progress float64
}

// ParseInput reads the input file and returns the dependencies and progress maps
func ParseInput(filePath string) (map[string][]string, map[string]float64) {
	dependencies := make(map[string][]string)
	progress := make(map[string]float64)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "->") {
			parts := strings.Split(line, "->")
			parent, child := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
			dependencies[parent] = append(dependencies[parent], child)

			// Ensure all nodes are in the progress map
			if _, exists := progress[parent]; !exists {
				progress[parent] = 0.0
			}
			if _, exists := progress[child]; !exists {
				progress[child] = 0.0
			}
		} else if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			skill, prog := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
			var progressValue float64
			_, err := fmt.Sscanf(prog, "%f", &progressValue)
			if err != nil {
				fmt.Println("Error parsing progress value:", err)
				continue
			}
			progress[skill] = progressValue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return dependencies, progress
}

// TopologicalSort sorts the skills based on dependencies and progress
func TopologicalSort(dependencies map[string][]string, progress map[string]float64) []string {
	inDegree := make(map[string]int)
	for _, children := range dependencies {
		for _, child := range children {
			inDegree[child]++
		}
	}

	inDegreeZeroQueue := []Skill{}
	for skill := range progress {
		if inDegree[skill] == 0 {
			inDegreeZeroQueue = append(inDegreeZeroQueue, Skill{skill, progress[skill]})
		}
	}

	validOrder := []string{}
	for len(inDegreeZeroQueue) > 0 {
		sort.Slice(inDegreeZeroQueue, func(i, j int) bool {
			if inDegreeZeroQueue[i].Progress == inDegreeZeroQueue[j].Progress {
				return inDegreeZeroQueue[i].Name < inDegreeZeroQueue[j].Name
			}
			return inDegreeZeroQueue[i].Progress > inDegreeZeroQueue[j].Progress
		})
		
	// Dequeue the first skill
		skill := inDegreeZeroQueue[0]
		inDegreeZeroQueue = inDegreeZeroQueue[1:]
		validOrder = append(validOrder, skill.Name)

		for _, edgeNode := range dependencies[skill.Name] {
			inDegree[edgeNode]--
			if inDegree[edgeNode] == 0 {
				inDegreeZeroQueue = append(inDegreeZeroQueue, Skill{edgeNode, progress[edgeNode]})
			}
		}
	}

	return validOrder
}

func main() {
	filePath := "input.txt"
	dependencies, progress := ParseInput(filePath)

	sortedSkills := TopologicalSort(dependencies, progress)
	for _, skill := range sortedSkills {
		fmt.Println(skill)
	}
}
