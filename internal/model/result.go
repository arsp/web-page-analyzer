package model

// AnalysisResult represents the aggregated output
// of a web page analysis.
//
// This struct contains only data and no business logic,
type AnalysisResult struct {
	URL         string
	HTMLVersion string
	Title       string
}
