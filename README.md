# graphie-fs
An algorithm used to organize a set of files/directories using a directed graph

# Usage
To use the graphie-fs algorithm, pass the root file/directory as the starting node. The algorithm will recursively traverse the directory structure and build a directed graph, organizing files and directories efficiently.

## Pseudocode
```go
// NOTE: this syntax is not language specific, but a pseudocode
START
declare map "graph" {key:string, value: string[]}
declare map "visited" {key:string, value: boolean}
init fileCount = 0  (keeps track of total no. of files)
rootNode = /path/to/the/root/file/or/directory
makeGraph(node)
END
```

## makeGraph(node) Function
```go
// NOTE: this syntax is not language specific, but a pseudocode
makeGraph(node){
    CurrentNode = node
    if isVisited(CurrentNode){
        return
    }

    adjacentNodes[] = getAdjacentNodesOfCurrentNode()
    // add CurrentNode to graph, with its adjacent nodes
    graph[key] = CurrentNode
    value = adjNodes[]
    
    CurrentNode = visited

    // CurrentNode has adjacent/child nodes if length(adjNodes[]) > 0
    if hasAdjacentNodes(CurrentNode){
        // CurrentNode is a folder, repeat makeGraph() for each adjacentNode
        n = length(adjacentNodes[]) - 1
        from i = 0 to i == n {
            makeGraph(adjacentNodes[i])
        }
    }else{
        // CurrentNode does not have adjacent nodes
        return
    }
}
```

## NOTE
1. A node can be a file or a directory
2. A node has adjacent nodes if the node is a directory and is not empty
3. The graph formed from the algorithm is a "directed graph". Therefore, if the edge is moving from
   node A to node B, then we say that node A is adjacent to node B, but node B is not adjacent to node A