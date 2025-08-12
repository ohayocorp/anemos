package createreferencednamespaces

type Options struct {
	Predicate func(string) bool
}

func NewOptions() *Options {
	return &Options{}
}

func NewOptionsWithPredicate(predicate func(string) bool) *Options {
	return &Options{
		Predicate: predicate,
	}
}
