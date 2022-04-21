package queryperf

type Measure struct {
	Query          string  `json:"query"`
	Calls          int     `json:"calls"`
	MeanExecTime   float64 `json:"mean_exec_time"`
	TotalExecTime  float64 `json:"total_exec_time"`
	StdDevExecTime float64 `json:"stddev_exec_time"`
	NrOfRows       float64 `json:"rows"`
}
type Result []Measure
