// Copyright (c) 2016 LINE Corporation. All rights reserved.
// LINE Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.

package main

import (
	"flag"
	"fmt"
	"github.com/google/pprof/profile"
	"io/ioutil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	heaps := make([]string, 0)
	for i, arg := range args {
		if i > 0 {
			heaps = append(heaps, arg)
		}
	}
	profiles := make([]*profile.Profile, 0)

	for _, heap := range heaps {
		buf, err := ioutil.ReadFile(heap)
		if err != nil {
			panic(err)
		}
		pf, err := profile.ParseData(buf)
		if err != nil {
			panic(err)
		}
		profiles = append(profiles, pf)
	}

	//test
	pf := profiles[0]
	for _, sample := range pf.Sample {
		fmt.Println(sample.Value)
		fmt.Println(sample.Label)
		fmt.Println(sample.NumLabel)
		fmt.Println(sample.NumUnit)
		fmt.Println("--")
	}

	for _, fun := range pf.Function {
		fmt.Println(fun.Name)
		fmt.Println(fun.Filename)
		fmt.Println(fun.StartLine)
		fmt.Println(fun.SystemName)
		fmt.Println("---")
	}

	for _, vt := range pf.SampleType {
		fmt.Println(vt.Type)
		fmt.Println(vt.Unit)
	}

	for _, lc := range pf.Location {
		fmt.Println(lc.Address)
	}
}
