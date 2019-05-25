/*
Eatcpu is a command that just uses CPU time in a hopefully controlled way.

usage: eatcpu [<N>]

If given an N, eat N CPUs worth. Otherwise, eat all of the CPUs.

*/

package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func spinner() {
	var i int
	for {
		i++
	}
}

func main() {
	flag.Parse()
	maxcpus := uint64(runtime.GOMAXPROCS(0))
	var howmany uint64
	var e error

	switch flag.NArg() {
	case 0:
		howmany = maxcpus
	case 1:
		howmany, e = strconv.ParseUint(flag.Arg(0), 0, 64)
		if e != nil {
			log.Fatalf("could not understand '%s' as how many CPUs to eat\n", flag.Arg(0))
		}
		if howmany > maxcpus {
			howmany = maxcpus
		}
	default:
		log.Fatalf("usage: eatcpu [<N>]\n")
	}
	if howmany == 0 {
		log.Fatalf("asked to eat 0 CPUs, doing nothing\n")
	}

	var i uint64
	fmt.Printf("eatcpu: eating %d CPUs\n", howmany)
	for i = 0; i < howmany-1; i++ {
		go spinner()
	}
	spinner()
}
