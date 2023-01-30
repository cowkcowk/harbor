package tag

type Tag struct {
	Immutable bool `json:"immutable"`
	Signed    bool `json:"signed"`
}

type Option struct {
	WithImmutableStatus bool
	WithSignature bool
	SignatureChecker 
}
