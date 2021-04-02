package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
)

func main() {
	dd, err := dummydump.New()




	if err != nil {
		panic(err)
	}


	if _, err = dd.Check().Export().Run(); err != nil {
		panic(err)
	}

}


