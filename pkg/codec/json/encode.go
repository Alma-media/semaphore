package json

import (
	"log"

	"github.com/francoispqt/gojay"
	"github.com/jexia/semaphore/pkg/references"
	"github.com/jexia/semaphore/pkg/specs"
)

// TODO: use the path as a resource
func encodeElement(encoder *gojay.Encoder, template *specs.Template, store references.Store) {
	log.Printf("%#v", template)

	switch {
	case template.Reference != nil:
		log.Println("Reference")

	case template.Message != nil:
		log.Println("Object")

		encoder.Object(
			NewObject(template.Message, store),
		)
	case template.Repeated != nil:
		log.Println("Array")

		encoder.Array(
			NewArray(template.Repeated, template.Reference, store),
		)
	case template.Enum != nil:
		log.Println("Enum")

		NewEnum("", template.Enum, template.Reference, store).MarshalJSONEnum(encoder)
	case template.Scalar != nil:
		log.Println("Scalar")

		NewScalar("", template.Scalar, template.Reference, store).MarshalJSONScalar(encoder)
	}
}

func encodeElementKey(encoder *gojay.Encoder, key string, template *specs.Template, store references.Store) {
	switch {
	case template.Message != nil:
		encoder.ObjectKey(
			key,
			NewObject(template.Message, store),
		)
	case template.Repeated != nil:
		encoder.ArrayKey(
			key,
			NewArray(template.Repeated, template.Reference, store),
		)
	case template.Enum != nil:
		NewEnum(key, template.Enum, template.Reference, store).MarshalJSONEnumKey(encoder)
	case template.Scalar != nil:
		NewScalar(key, template.Scalar, template.Reference, store).MarshalJSONScalarKey(encoder)
	}
}
