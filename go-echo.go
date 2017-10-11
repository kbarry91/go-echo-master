// An echo web application.
// https://ianmcloughlin.github.io :: 2017-10-05
// kevin barry 11-10-17

package main

import (
	"fmt"
	"net/http"
	"bytes"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-Type","text/html");// allows browser to render html tags

	fmt.Fprintln(w, "<h4>r.Method:           </h4>",  r.Method           )
	fmt.Fprintln(w, "<h4>r.URL:              </h4>",  r.URL              )
	fmt.Fprintln(w, "<h4>r.Proto:            </h4>",  r.Proto            )
	fmt.Fprintln(w, "<h4>r.ContentLength:    </h4>",  r.ContentLength    )
	fmt.Fprintln(w, "<h4>r.TransferEncoding: </h4>",  r.TransferEncoding )
	fmt.Fprintln(w, "<h4>r.Close:            </h4>",  r.Close            )
	fmt.Fprintln(w, "<h4>r.Host:             </h4>",  r.Host             )
	fmt.Fprintln(w, "<h4>r.Form:             </h4>",  r.Form             )
	fmt.Fprintln(w, "<h4>r.PostForm:         </h4>",  r.PostForm         )
	fmt.Fprintln(w, "<h4>r.RemoteAddr:       </h4>",  r.RemoteAddr       )
	fmt.Fprintln(w, "<h4>r.RequestURI:       </h4>",  r.RequestURI       )

	fmt.Fprintln(w, "r.URL.Opaque:       ", r.URL.Opaque        )
	fmt.Fprintln(w, "r.URL.Scheme:       ", r.URL.Scheme        )
	fmt.Fprintln(w, "r.URL.Host:         ", r.URL.Host          )
	fmt.Fprintln(w, "r.URL.Path:         ", r.URL.Path          )
	fmt.Fprintln(w, "r.URL.RawPath:      ", r.URL.RawPath       )
	fmt.Fprintln(w, "r.URL.RawQuery:     ", r.URL.RawQuery      )
	fmt.Fprintln(w, "r.URL.Fragment:     ", r.URL.Fragment      )

	fmt.Fprintln(w, "Header:")
	for key, value := range r.Header {
		fmt.Fprintln(w, "\t" + key + ":", value)
	}

	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	fmt.Fprintln(w, "r.Body:             ",  body.String())
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
