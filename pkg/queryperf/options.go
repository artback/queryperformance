package queryperf

type Options struct {
	SortBy     string   `query:"sort_by"`
	ASC        bool     `query:"asc"`
	Offset     uint64   `query:"offset"`
	Limit      uint64   `query:"limit"`
	Statements []string `query:"statement"`
	MinCalls   int      `query:"mincalls"`
}
