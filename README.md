# HOW TO MAKE A STANDALONE PACKAGE FOR OTHER APP TO USE?
- this project should be under `$GOPATH/src/github.com/onelazyugy/golang-common`
- create a project like this then run below cmd
- ```$ go mod init github.com/onelazyugy/golang-common```
- ```$ go install```
- add some code to this package 
- push to github

# TAG 
- make code change
- $ git add .
- $ git commit -m "some commit message"
- $ git push
- $ git tag v1.0.3
- $ git push --tags
- check github for latest release 

# HOW TO USE THIS FROM INSIDE ANOTHER APP?
- `$ go get -u github.com/onelazyugy/golang-common`
- after that cmd, you would get something like below

- get specific release version
- $ go get github.com/onelazyugy/golang-common@v1.0.3
```
go: downloading github.com/onelazyugy/golang-common v0.0.0-20210204213736-b0bd23d4e4aa
go: github.com/onelazyugy/golang-common upgrade => v0.0.0-20210204213736-b0bd23d4e4aa
```
- sample `main.go` from the app that is using this package
```
package main

import (
	"fmt"
	"log"

	common "github.com/onelazyugy/golang-common"
	"github.com/onelazyugy/golang-common/types"
	
)

func main() {
	log.Println("starting...")
	var c types.Common
	c.Name = "viet"

	fmt.Println(c)
	result := common.Add(1, 3)
	fmt.Println(result)	
}
```
-reference: https://medium.com/mindorks/how-to-create-a-package-in-go-ae4e79b95241