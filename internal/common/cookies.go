package common

import (
	"net/http"
	"strings"
)

/*
JSESSIONID=290624BC4A930E64208F00786A2D2650; Path=/oesc; HttpOnly
538D2EDB2C736C7FCF04E3A6FE2E70CD=0%2b4%2bp57aINYoUfk3oQK6kg%3d%3d; Path=/; HttpOnly
538D2EDB2C736C7FCF04E3A6FE2E70CD=jWSsvkme8YfjawEBMfyz2Mov%2f8PNQxJFpJCdQ4P5%2bm8lbloerJTVB44ISCmuIItP; Path=/; HttpOnly
*/

// Parse parses cookies from a string
func Parse(args []string) []*http.Cookie {
	cookies := []*http.Cookie{}
	for _, arg := range args {
		if strings.Contains(arg, "=") {
			cookieKeyVals := strings.Split(arg, ";")
			cookie := new(http.Cookie)
			for _, cookieKeyVal := range cookieKeyVals {
				if strings.Contains(cookieKeyVal, "=") {
					cookieVal := strings.Split(cookieKeyVal, "=")
					switch cookieVal[0] {
					case "Path":
						cookie.Path = cookieVal[1]
					default:
						cookie.Name = cookieVal[0]
						cookie.Value = cookieVal[1]
					}
				}
				if cookieKeyVal == "HttpOnly" {
					cookie.HttpOnly = true
				}
			}
			cookies = append(cookies, cookie)
		}
	}
	return cookies
}
