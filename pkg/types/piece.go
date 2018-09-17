package types

type Piece int
type Side int

const (
	None Side = iota
	A
	B
)

func (s Side) Other() Side {
	if s == A {
		return B
	}
	return A
}

const (
	Empty Piece = iota
	AMouse
	ACat
	AWolf
	ADog
	AHyena
	ATiger
	ALion
	AElephant
	BMouse
	BCat
	BWolf
	BDog
	BHyena
	BTiger
	BLion
	BElephant
)

func (p Piece) Side() Side {
	switch p {
	case
		AMouse,
		ACat,
		AWolf,
		ADog,
		AHyena,
		ATiger,
		ALion:
		return A
	case
		BMouse,
		BCat,
		BWolf,
		BDog,
		BHyena,
		BTiger,
		BLion:
		return B
	default:
		return None
	}
}

func (p Piece) CanJump() bool {
	switch p {
	case
		ATiger,
		ALion,
		BTiger,
		BLion:
		return true
	default:
		return false
	}
}

func (p Piece) CanSwim() bool {
	switch p {
	case
		AMouse,
		BMouse:
		return true
	default:
		return false
	}
}

var pieceCanTake = map[Piece]map[Piece]bool{
	AMouse: {
		Empty:     true,
		BMouse:    true,
		BElephant: true,
	},
	ACat: {
		Empty:  true,
		BMouse: true,
		BCat:   true,
	},
	AWolf: {
		Empty:  true,
		BMouse: true,
		BCat:   true,
		BWolf:  true,
	},
	ADog: {
		Empty:  true,
		BMouse: true,
		BCat:   true,
		BWolf:  true,
		BDog:   true,
	},
	AHyena: {
		Empty:  true,
		BMouse: true,
		BCat:   true,
		BWolf:  true,
		BDog:   true,
		BHyena: true,
	},
	ATiger: {
		Empty:  true,
		BMouse: true,
		BCat:   true,
		BWolf:  true,
		BDog:   true,
		BHyena: true,
		BTiger: true,
	},
	ALion: {
		Empty:  true,
		BMouse: true,
		BCat:   true,
		BWolf:  true,
		BDog:   true,
		BHyena: true,
		BTiger: true,
		BLion:  true,
	},
	AElephant: {
		Empty:     true,
		BCat:      true,
		BWolf:     true,
		BDog:      true,
		BHyena:    true,
		BTiger:    true,
		BLion:     true,
		BElephant: true,
	},
	BMouse: {
		Empty:     true,
		AMouse:    true,
		AElephant: true,
	},
	BCat: {
		Empty:  true,
		AMouse: true,
		ACat:   true,
	},
	BWolf: {
		Empty:  true,
		AMouse: true,
		ACat:   true,
		AWolf:  true,
	},
	BDog: {
		Empty:  true,
		AMouse: true,
		ACat:   true,
		AWolf:  true,
		ADog:   true,
	},
	BHyena: {
		Empty:  true,
		AMouse: true,
		ACat:   true,
		AWolf:  true,
		ADog:   true,
		AHyena: true,
	},
	BTiger: {
		Empty:  true,
		AMouse: true,
		ACat:   true,
		AWolf:  true,
		ADog:   true,
		AHyena: true,
		ATiger: true,
	},
	BLion: {
		Empty:  true,
		AMouse: true,
		ACat:   true,
		AWolf:  true,
		ADog:   true,
		AHyena: true,
		ATiger: true,
		ALion:  true,
	},
	BElephant: {
		Empty:     true,
		ACat:      true,
		AWolf:     true,
		ADog:      true,
		AHyena:    true,
		ATiger:    true,
		ALion:     true,
		AElephant: true,
	},
}

func (p Piece) CanTake(opponent Piece) bool {
	if _, ok := pieceCanTake[p][opponent]; ok {
		return true
	}
	return false
}

func (p Piece) String() string {
	switch p {
	case AMouse:
		return "a1"
	case ACat:
		return "a2"
	case AWolf:
		return "a3"
	case ADog:
		return "a4"
	case AHyena:
		return "a5"
	case ATiger:
		return "a6"
	case ALion:
		return "a7"
	case AElephant:
		return "a8"
	case BMouse:
		return "b1"
	case BCat:
		return "b2"
	case BWolf:
		return "b3"
	case BDog:
		return "b4"
	case BHyena:
		return "b5"
	case BTiger:
		return "b6"
	case BLion:
		return "b7"
	case BElephant:
		return "b8"
	default:
		return "\u00A0\u00A0" // non-breaking space
	}
}
