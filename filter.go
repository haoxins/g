package g

// I disagree with Rob Pike
//   I just use "for" loops.
//   You shouldn't use it either.
import "robpike.io/filter"

var (
	Apply         = filter.Apply
	ApplyInPlace  = filter.ApplyInPlace
	Choose        = filter.Choose
	ChooseInPlace = filter.ChooseInPlace
	Drop          = filter.Drop
	DropInPlace   = filter.DropInPlace
	Reduce        = filter.Reduce
)
