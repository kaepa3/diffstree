package main

import (
	"os"

	"github.com/kaepa3/diffstree/sub"
)

//引数は実行ディレクトリ、実行パス、コミット番号１、コミット番号２
func main() {
	sourcetree.CheckArg(os.Args)
	sourcetree.MakeBatAndDo(os.Args)
	sourcetree.GoExec()
}
