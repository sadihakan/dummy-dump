package main

import (
	"context"
	"fmt"
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New()

	if err != nil {
		fmt.Println("DummyDump new error: ", err)
	}

	dd.SetBinaryConfig(config.MySQL, false, true)

	ctx := context.Background()

	binary, version := dd.GetBinary(ctx)
	fmt.Println("Bin: ", binary)
	fmt.Println("Version: ", version)

}
