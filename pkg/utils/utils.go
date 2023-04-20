/*The utils package provides a function called ParseBody that reads the body of an HTTP request and parses it into a struct that is passed as a parameter.  */
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/* ParseBody function reads the request body and decodes it into the provided interface.
This function is used to extract the JSON data from a POST or PUT request.*/
func ParseBody(r *http.Request, x interface{}){
	// Read the request body
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		// If the read is successful, unmarshal the body into the provided interface
		if err := json.Unmarshal([]byte(body), x); err != nil{
			// If there is an error during unmarshaling, return
			return
		}
	}
}