# go-with-tests
This is the repo I'm using to go through the [Go With Tests](https://quii.gitbook.io/learn-go-with-tests/) resource by Chris James. 

Left off [here](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#hello-world...-again)- implemented but not read the walkthrough or the code.

# Notes

## Hello, World
- in your return value, if you name your return, that named variable now exists in your func which allows you to just say `return`, ie: this is valid and will return an empty string as is. I could also just reference `prefix` without declaring it in the func and still just say `return` at the end
```
func greetingPrefix(language string) (prefix string){
    return
}
```

## Integers
- also contains an example for how to put example funcs in with your tests
- run `go test -v` to run your tests and your examples

## Iteration
- also contains an example for how to write a benchmark test to see how quick your functions can run locally

## Arrays and Slices
- arrays are fixed size, defined like `[5]int{1, 2, 3, 4, 5}`
- slices are not fixed size, defined like `[]int{1,2,3,4,5}`
- also has example of using `reflect.DeepEqual` to compare two slices in the test

## Stucts, Methods and Interfaces
- interfaces in go make it so that you can have one object at a high level that can be representative of many objects as long as all the objects implement the methods of the interface. In this case, a "shape" is anything that implements the Area() method
- You can use interfaces to make your tests more easily extendable

## 
