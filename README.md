# Group Project: GO LANG

General Assembly WDI NYC
September, 2107

#### The Master Minds
- Ryan Edwards
- Asher Shaheen

Table of Contents
=================
+  [Description](#description)
+  [Go Docs](#link-to-go-docs)
+  [Who Uses Go](#who-uses-go)
+  [Getting Started With Go](#getting-started-with-go)
+  [Other Resources](#other-helpful-resources)

## Description 
The Go programming language is an open source project to make programmers more productive.
![Gopher image](https://github.com/golang/go/blob/master/doc/gopher/fiveyears.jpg)
*Gopher image by _Renee French_, licensed under _Creative Commons 3.0 Attributions license_.*

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

## Link to GO DOCS
[GO Docs](https://golang.org/doc/)

## Who Uses GO
-  [x] Google™
-  [x] YouTube™
-  [x] Apple™

## Getting Started with Go
![](https://github.com/Asher978/SWEETLIBS/blob/master/assets/instructionA.png)

![](https://github.com/Asher978/SWEETLIBS/blob/master/assets/instructionB.png)

<details>
<summary>Detailed Instructions</summary>

  *  Download the package file, open it, and follow the prompts to install the Go tools. The package installs the Go distribution to /usr/local/go.
The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.

  *  Check that Go is installed correctly by setting up a workspace and building a simple program, as follows. Create your workspace directory, $HOME/go (in the same directory as your 'wdi' folder). (If you'd like to use a different directory, you will need to set the GOPATH environment variable.)
Next, make the directory 'src/hello' inside your workspace, and in that directory create a file named hello.go that looks like:

```
            package main
            import "fmt"
            func main() {
                fmt.Printf("hello, world\n")
}
```
  *  Then build it with the go tool:
```
            $ cd $HOME/go/src/hello
            $ go build
```
  *  The command above will build an executable named hello in the directory alongside your source code. Execute it to see the greeting:
```
            $ ./hello
            hello, world
```
</details>

#### Other Helpful Resources 
  *  [GOLANG WEB APP TUTORIAL 1](https://www.youtube.com/watch?v=iIztjjNTSjs)
  *  [GOLANG WEB APP TUTORIAL 2](https://www.youtube.com/watch?v=Vlie-srOU8c)

## Algorithms With Go
[Algorithms](https://github.com/Asher978/SWEETLIBS/blob/master/algorithms.go)

