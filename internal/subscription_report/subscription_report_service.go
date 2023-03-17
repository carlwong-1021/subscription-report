package subscription_report

type _reportService struct {
	options *ReportOption
}

type ReportService interface {
	Exec()
}

type ReportOption struct {
	From string
	To   string
}

func NewReportService(options *ReportOption) ReportService {
	return &_reportService{
		options: options,
	}
}

func (r *_reportService) Exec() {

}
