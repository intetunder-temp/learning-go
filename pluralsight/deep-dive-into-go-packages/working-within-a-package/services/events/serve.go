package events

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/intetunder-temp/learning-go/pluralsight/deep-dive-into-go-packages/working-within-a-package/services/internal/ports"
)

var port = 42

func StartServer() error {
	sm := http.NewServeMux()
	sm.Handle("/", new(eventHandler))
	return http.ListenAndServe(":"+strconv.Itoa(port), sm)
}

func init() {
	fmt.Println("serve.go 1", port)
	port = ports.EventService
	fmt.Println("serve.go 2", port)
}
