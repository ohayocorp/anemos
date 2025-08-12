package collectcrds

type Options struct {
	DocumentGroupPath string
}

func NewOptions() *Options {
	return &Options{}
}

func NewOptionsWithDocumentGroupPath(documentGroupPath string) *Options {
	return &Options{
		DocumentGroupPath: documentGroupPath,
	}
}
