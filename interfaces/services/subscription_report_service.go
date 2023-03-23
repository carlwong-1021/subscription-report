package services_interfaces

import (
	steps_interfaces "subscription-report/interfaces/steps"
)

type ReportService interface {
	Exec(steps_interfaces.FetchAppSubscriptionStep, steps_interfaces.GenerateReportStep)
}
