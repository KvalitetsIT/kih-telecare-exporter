package oioxds

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHeaders(t *testing.T) {

	log = logrus.New()
	cst := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Headers", r.Header.Get("Content-Type"))

		switch r.URL.Query().Get("k") {
		case "":
			w.Header().Set("Content-Disposition", `attachment; filename="file.txt"`)
		case "unq":
			w.Header().Set("Content-Disposition", `attachment; filename=file.txt`)
		case "lq":
			w.Header().Set("Content-Disposition", `attachment; filename="file.txt`)
		case "rq":
			w.Header().Set("Content-Disposition", `attachment; filename=file.txt"`)
		}
		w.Write([]byte(`Hello, World!`)) // nolint
	}))
	defer cst.Close()

	for _, sym := range []string{"", "unq", "lq", "rq"} {
		req, err := http.NewRequest(http.MethodPost, cst.URL+"?k="+sym, nil)
		if err != nil {
			t.Errorf("Error creaing reuqest %v", err)
		}
		req.Header.Add("Content-Type", "application/soap+xml;charset=\"utf-8\";ation=\"ihe:iti:2007:ProvideAndRegisterDocumentSet-b\"")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("Failed to perform plain get: %v", err)
			t.Errorf("Failed to perform plain get: %v", err)
		}
		blob, _ := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		log.Printf("sym=%-5q blob:%q\n", sym, blob)
	}
}
