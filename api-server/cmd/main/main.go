package main

import (
	"fmt"
	"github.com/Vivirinter/sdr-analysis-system/api-server/internal"
	"github.com/Vivirinter/sdr-analysis-system/api-server/pkg"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "SDR Data")
	})

	internal.InternalLogic()
	pkg.PkgLogic()

	http.ListenAndServe(":8080", nil)
}
