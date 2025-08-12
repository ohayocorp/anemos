package collectnamespaces

type Options struct {
	DocumentGroupPath string
}

func NewOptions() *Options {
	options := Options{}

	return &options
}

func NewOptionsWithDocumentGroupPath(documentGroupPath string) *Options {
	return &Options{
		DocumentGroupPath: documentGroupPath,
	}
}
