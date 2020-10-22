package references

import (
	"encoding/json"
	"fmt"
	"io"
)

// EncodeJSON writes JSON reference representation to the provided writer.
func (reference *Reference) EncodeJSON(writer io.Writer) error {
	switch {
	case reference.Scalar != nil:
		encoded, err := json.Marshal(reference.Scalar)
		if err != nil {
			return err
		}

		_, err = writer.Write(encoded)

		return err
	case reference.Enum != nil:
		_, err := writer.Write([]byte(fmt.Sprintf("%d", *reference.Enum)))

		return err // TODO: use enum key
	case reference.Repeated != nil:
		if _, err := fmt.Fprint(writer, "["); err != nil {
			return err
		}

		for index, storage := range reference.Repeated {
			if index > 0 {
				if _, err := fmt.Fprint(writer, ","); err != nil {
					return err
				}
			}

			encoder, ok := storage.(interface{ EncodeJSON(io.Writer) error })
			if ok {
				if err := encoder.EncodeJSON(writer); err != nil {
					return err
				}
			} else {
				if err := json.NewEncoder(writer).Encode(storage); err != nil {
					return err
				}
			}
		}

		_, err := fmt.Fprint(writer, "]")

		return err
	case reference.Message != nil:
		encoder, ok := reference.Message.(interface{ EncodeJSON(io.Writer) error })
		if ok {
			return encoder.EncodeJSON(writer)
		}

		return json.NewEncoder(writer).Encode(reference.Message)
	default:
		_, err := fmt.Fprint(writer, "null")

		return err
	}
}

func (store *store) EncodeJSON(writer io.Writer) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	var separate bool

	if _, err := fmt.Fprint(writer, "{"); err != nil {
		return err
	}

	for key, reference := range store.values {
		if separate {
			if _, err := fmt.Fprint(writer, ","); err != nil {
				return err
			}
		} else {
			separate = true
		}

		if _, err := fmt.Fprintf(writer, "%q:", key); err != nil {
			return err
		}

		if err := reference.EncodeJSON(writer); err != nil {
			return err
		}
	}

	_, err := fmt.Fprint(writer, "}")

	return err
}
