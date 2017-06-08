# Go Performances Plotter

```go

    //max size of the dataset
    max := 10
    //increment dataset by 
    increment := 1
    //how many tries for each dataset size
    //we keep the median
	retries := 10
    //The name to be displayed on the graph
    label := "nothing"

    gripper.PerfPlotter().
    Analyze(
        myFunctionToAnalyze,
        max,
        increment,
        retries,
        label,
    ).Plot("x", "y", "x/y", "testing.png")
```