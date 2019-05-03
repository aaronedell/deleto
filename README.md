![deleto logo](https://github.com/aaronedell/deleto/blob/master/deleto.png)
[![Go Report Card](https://goreportcard.com/badge/github.com/aaronedell/deleto)](https://goreportcard.com/report/github.com/aaronedell/deleto)
# deleto

Randomly delete files from a directory.

If you're into machine learning and classification, then you know your classes need to be balanced. Sometimes, when I create large datasets to use with [imgclass](https://github.com/machinebox/toys/tree/master/imgclass) or [txtclass](https://github.com/machinebox/toys/tree/master/textclass), each folder can have a different number of files. In order to make them the same, I calculate the difference and use this tool to randomly delete X number of files from the target directory. 

## Usage

$ go run main.go -dir=path/to/target -num=[number of files to delete]



