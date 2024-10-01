/**
	create a graph to represent files and directories, given the file path
	to the root folder
**/

package main

import (
	"fmt"
	"log"
	"os"
)

var graph = make(map[string][]string)
var visited = make(map[string]bool)
var fileCount = 0

// check if path is a file or directory
func getType(path string) (interface{}, error) {
	entry, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("isFile->os.Stat: %v", err)
	}

	if entry.IsDir() {
		return false, nil
	}

	return true, nil
}

// get the adjacent nodes of a node
func adjNode(path string) ([]string, error) {
	adjNodes := []string{}

	isFile, err := getType(path)
	if err != nil {
		return nil, fmt.Errorf("adjNode->getType: %v", err)
	}

	// path is a file, return
	if isFile == true {
		fileCount++
		return adjNodes, nil
	}

	// path is a folder. Visit folder
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("adjNodes->os.ReadDir: %v", err)
	}

	for _, entry := range entries {
		fullPath := fmt.Sprintf("%v/%v", path, entry.Name())
		adjNodes = append(adjNodes, fullPath)
	}

	return adjNodes, nil
}

// create a graph
func makeGraph(node string) {
	// check if node is visited
	if visited[node] {
		log.Printf("node {{%v}} already visited\n", node)
		return
	}

	// get adjacent nodes of the node
	adjNodes, err := adjNode(node)
	if err != nil {
		log.Printf("makeGraph->adjNode: *** Error getting adjNodes[] *** of {{%v}}\n", node)
		log.Printf("makeGraph->adjNode: %v\n", err)
		return
	}

	// add node to graph, mark it as visited
	graph[node] = adjNodes
	visited[node] = true
	log.Printf("*** {{%v}} added to graph and marked ***\n", node)

	length := len(adjNodes)
	if length == 0 {
		log.Printf("node {{%v}} does not have adjacent nodes\n", node)
		return
	}

	// loop through each adjacent node to discover new nodes
	for i := 0; i < length; i++ {
		log.Printf("*** visiting adjNodes[] of {{%v}} ***\n", node)
		newNode := adjNodes[i]
		makeGraph(newNode)
	}
}

func printGraph(graph map[string][]string, fileCount int) {
	fmt.Printf("\n\n*** GRAPH (total files: %v) ***\n\n", fileCount)
	i := 1
	for node, adjNodes := range graph {
		fmt.Printf("%v. {{%v}}: %v\n", i, node, adjNodes)
		i++
	}
}

func main() {
	var rootNode = "/path/to/your/file/or/directory"
	path := rootNode

	makeGraph(path)
	printGraph(graph, fileCount)

	// control: try to visit an already visited node
	fmt.Println()
	makeGraph(path)
}
