package writereports

const (
	ReportOutputTypeHtml     ReportOutputType = "html"
	ReportOutputTypeMarkdown ReportOutputType = "markdown"
)

type ReportOutputType string

type Options struct {
	// OutputTypes determines the file format of the report.
	OutputTypes []ReportOutputType
}

func NewOptions() *Options {
	return &Options{}
}

func NewOptionsWithOutputTypes(outputTypes []ReportOutputType) *Options {
	return &Options{
		OutputTypes: outputTypes,
	}
}
