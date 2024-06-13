package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&dst)
	
	if err != nil {

		var synctaxError *json.SyntaxError
		var invalidUnMarshalError *json.InvalidUnmarshalError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch  {
		case errors.As(err, &synctaxError):
			return fmt.Errorf("body contains badly-formed JSON syntax at character %d", synctaxError.Offset)
		case errors.As(err, &unmarshalTypeError):
			return fmt.Errorf("invalid type a t %d", unmarshalTypeError.Type)
		case errors.As(err, &invalidUnMarshalError):
			return fmt.Errorf("invalid argument")
		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("body contains badly-formed JSON")
		default:
			return err
		}
	}

	return nil
}
