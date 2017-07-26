package main

//http web server practice
/*import (
	"io"
	"net/http"
)*/

import (
	"fmt"

	"github.com/b10023037/stringutil"
)

/*func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello http~")
}

func main() {
	//fmt.Printf("hello,Go World\n")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
*/
func main() {
	fmt.Printf(stringutil.Reverse("!oG,olleH"))
}
