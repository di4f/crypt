package main

import (
	"os"
	"io"
	"fmt"
	"github.com/omnipunk/cli/mtool"
	"encoding/base64"
)

var root = mtool.T("crypt").Subs(
	mtool.T("caesar").Func(func(flags *mtool.Flags){
		var (
			decrypt bool
			key int
		)
		flags.BoolVar(&decrypt, "d", false, "decrypt instead of crypting")
		flags.IntVar(&key, "k", 1, "set the key")
		flags.Parse()
		bts, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		fn := func(n int32) int32 {
			return n + int32(key)
		}

		if decrypt {
			fn = func(n int32) int32 {
				return n - int32(key)
			}
		}

		str := string(bts)
		for _, r := range str {
			r = fn(int32(r))
			_, err = os.Stdout.Write([]byte(string(r)))
			if err != nil {
				panic(err)
			}
		}
	}),
	mtool.T("b64").Func(func(flags *mtool.Flags){
		var (
			decrypt bool
		)
		flags.BoolVar(&decrypt, "d", false, "decrypt")
		flags.Parse()
		msg, _ := io.ReadAll(os.Stdin)
		if decrypt {
			decoded, _ := base64.StdEncoding.DecodeString(string(msg))
			fmt.Print(string(decoded))
		} else {
			encoded := base64.StdEncoding.EncodeToString(msg)
			fmt.Print(encoded)
		}
	}),
)

func main() {
	root.Run(os.Args[1:])
}
