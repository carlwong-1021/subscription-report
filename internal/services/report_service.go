package services

type _reportService struct {
	steps []Step
}

type ReportService interface {
	Exec(from string, to string) any
}

type Step interface {
	Exec(input any, option *ReportOption) (any, *ReportOption)
}

type ReportOption struct {
	From      string
	To        string
	BatchSize int
}

func NewReportService(
	steps []Step) ReportService {
	return &_reportService{
		steps: steps,
	}
}

func (r *_reportService) Exec(from string, to string) any {
	var report any
	saveOption := &ReportOption{From: from, To: to, BatchSize: 5}
	for _, step := range r.steps {
		report, saveOption = step.Exec(report, saveOption)
	}
	return report
}
