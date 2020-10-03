package linker

import (
	"github.com/gdower/bhlinker/domain/entity"
	"github.com/gdower/bhlinker/domain/usecase"
)

// Linker is an implementation of usecase/Plugger interface.
type Linker struct {
	usecase.Referencer
}

func NewLinker(r usecase.Referencer) Linker {
	return Linker{Referencer: r}
}

// GetLink takes name-string with its reference data. The reference data is
// expected to be a paper with original nomenclatural description of the
// the name-string. The method tries to find the best BHL match to that
// reference and sends back a BHL reference metadata as well as URL link to
// the reference.
func (l Linker) GetLink(input entity.Input) (entity.Output, error) {
	name := input.Name.Name
	bhlRefs, err := l.Refs(name)
	if err != nil {
		return entity.Output{}, err
	}
	return bestMatchBHL(input, bhlRefs.References), nil
}

// GetLinks takes a stream of name-strings with their reference data. The
// reference data for each name-string is expected to be a paper with original
// nomenclatural description of the the name-string.  The method tries to find
// the best BHL match to that reference and sends back a BHL reference metadata
// as well as URL link to the reference.
//
// The streams are implemented as channels. This approach allows to work with
// inputs of any size.
func (l Linker) GetLinks(chIn <-chan entity.Input, chOut chan<- entity.Output) {}
