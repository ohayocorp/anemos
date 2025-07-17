package collectnamespaces

type Options struct {
	Directory string
}

func NewOptions() *Options {
	options := Options{}

	return &options
}

func NewOptionsWithDirectory(directory string) *Options {
	return &Options{
		Directory: directory,
	}
}
