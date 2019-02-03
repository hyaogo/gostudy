package go01

import (
	"fmt"
	"net/http"
)

func JiujiuBiao(w http.ResponseWriter, r *http.Request) {
	jiujiu(1, w)
}

func jiujiu(n int, w http.ResponseWriter) {
	for i := 1; i <= n; i++ {
		fmt.Fprint(w, i, "*", n, "=", i*n, "\t")
	}
	if n == 9 {
		return
	}
	fmt.Fprintln(w)
	//递归
	jiujiu(n+1, w)
}
