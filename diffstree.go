package main

import (
	"os"

	"./sub/sourcetree"
)

//引数は実行ディレクトリ、実行パス、コミット番号１、コミット番号２
func main() {
	sourcetree.CheckArg(os.Args)

}
