gen-graph-struct-int:
	genny -in /Users/nick/go/src/bitbucket.org/MykolaKuropatkin/golang_algorithms/algorithms/graph/graph.go -out /Users/nick/go/src/bitbucket.org/MykolaKuropatkin/golang_algorithms/algorithms/graph/gen/graph-int.go -pkg gen gen "Item=int"

gen-graph-struct-str:
	genny -in /Users/nick/go/src/bitbucket.org/MykolaKuropatkin/golang_algorithms/algorithms/graph/graph.go -out /Users/nick/go/src/bitbucket.org/MykolaKuropatkin/golang_algorithms/algorithms/graph/gen/graph-string.go -pkg gen gen "Item=string"