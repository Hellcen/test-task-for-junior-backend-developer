package metrics

import (
	"context"
	"example.com/taskservice/internal/domain/recurringrule"
	"example.com/taskservice/internal/usecase/recurring"
	"time"
)

func CollectMetrics(recurringUsecase recurring.Usecase) {
	for {
		rules, err := recurringUsecase.List(context.Background())
		if err == nil {
			activeCount := 0
			for _, rule := range rules {
				if rule.Status == recurringrule.StatusActive {
					activeCount++
				}
			}
			activeRecurringRules.Set(float64(activeCount))
		}
		time.Sleep(30 * time.Second)
	}
}
