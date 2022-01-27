package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// check ip address valid or not
func ip4or6(ip string) string {

	if len(ip) == 0 {
		return "invalid IP"
	}

	var countdot int = strings.Count(ip, ".")

	var countcolon int = strings.Count(ip, ":")

	if countdot == 3 {

		var sub string = ""
		ip += "."

		for i := 0; i < len(ip); i++ {

			if ip[i] == '.' {

				num, err := strconv.ParseInt(sub, 10, 64)

				if len(sub) == 0 {
					return "invalid IP"

				} else if num < 0 || num > 255 {
					return "invalid IP"
				} else if err != nil {
					return "invalid IP"
				}
				sub = ""

			} else {
				sub += string(ip[i])

			}
		}

		return "valid IPv4"

	} else if countcolon == 7 {

		var sub string = ""
		ip += ":"

		for i := 0; i < len(ip); i++ {

			if ip[i] == ':' {

				if len(sub) > 4 {
					return "invalid IP"
				}

				sub = ""

			} else {

				if (ip[i] < 'A' || ip[i] > 'F') && (ip[i] < 'a' && ip[i] > 'f') && (ip[i] < '0' || ip[i] > '9') {
					return "invalid IP"
				}

				sub += string(ip[i])
			}
		}
		return "valid IPv6"

	} else {
		return "Invalid IP"
	}

}

func process(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "form.html")

	case "POST":

		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)

		}

		ipaddress := r.FormValue("ip")

		fmt.Fprintf(w, " %s is  %s \n", ipaddress, ip4or6(ipaddress))

		// name := r.FormValue("name")
		// occupation := r.FormValue("occupation")

		// fmt.Fprintf(w, "%s is a %s\n", name, occupation)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func main() {

	http.HandleFunc("/", process)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
