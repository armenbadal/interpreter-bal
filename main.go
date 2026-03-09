package main

import (
	"bal/interpreter"
	"bal/parser"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func processOneFile(filename string, printAST bool) int {
	// բացել ֆայլը
	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("ՍԽԱԼ։ %s ֆայլը չգտնվեց:\n", filename)
		} else {
			fmt.Printf("ՍԽԱԼ։ %s ֆայլը բացելը ձախողվեց:\n", filename)
		}
		return 1
	}
	defer file.Close()

	parser := parser.New(bufio.NewReader(file))

	tree, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		return 3
	}

	if printAST {
		fmt.Printf("%v", tree)
		return 0
	}

	err = interpreter.Execute(tree)
	if err != nil {
		fmt.Println(err)
		return 4
	}

	return 0
}

const version = "0.0.1"

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Ինտերպրետատորի օգտագործումը.\n")
		fmt.Fprintf(w, "  bal --help\n")
		fmt.Fprintf(w, "  bal --version\n")
		fmt.Fprintf(w, "  bal --ast example.{bal|բալ|bas}\n")
		fmt.Fprintf(w, "  bal example.{bal|բալ|bas}\n\n")
		flag.PrintDefaults()
	}

	claVersion := flag.Bool("version", false, "Ծրագրի տարբերակը")

	var printAST bool
	flag.BoolVar(&printAST, "ast", false, "Տպել վերացական շարահյուսական ծառը")

	flag.Parse()

	if *claVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		fmt.Printf("Բալ ալգորիթմական լեզու v%s\n", version)
		os.Exit(0)
	}

	filename := flag.Arg(0)
	extension := filepath.Ext(filename)
	if extension != ".bal" && extension != ".բալ" && extension != ".bas" {
		fmt.Println("ՍԽԱԼ։ ծրագրի ֆայլը պետք է ունենա '.bal', '.բալ' կամ '.bas' վերջավորություն։")
		os.Exit(1)
	}

	status := processOneFile(flag.Arg(0), printAST)
	os.Exit(status)
}
