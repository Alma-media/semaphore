package json

import (
	"github.com/francoispqt/gojay"
	"github.com/jexia/semaphore/pkg/references"
	"github.com/jexia/semaphore/pkg/specs"
)

func decodeElement(decoder *gojay.Decoder, resource, path string, template *specs.Template, store references.Store) error {
	var reference = specs.PropertyReference{
		Resource: resource,
		Path:     path,
	}

	switch {
	case template.Message != nil:
		reference := &references.Reference{
			Path:    path,
			Message: references.NewReferenceStore(0),
		}

		store.StoreReference(resource, reference)

		object := NewObject(resource, path, template.Message, reference.Message)
		if object == nil {
			break
		}

		return decoder.Object(object)
	case template.Repeated != nil:
		store.StoreReference(resource, &references.Reference{Path: path})

		return decoder.Array(
			NewArray(resource, path, template.Repeated, &reference, store),
		)
	case template.Enum != nil:
		return NewEnum("", template.Enum, &reference, store).UnmarshalJSONEnum(decoder)
	case template.Scalar != nil:
		return NewScalar("", template.Scalar, &reference, store).UnmarshalJSONScalar(decoder)
	}

	return nil
}
