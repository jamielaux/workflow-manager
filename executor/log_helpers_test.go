package executor

import (
	"testing"
	"time"

	counter "github.com/Clever/aws-sdk-go-counter"
	"github.com/Clever/workflow-manager/gen-go/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

func init() {
	err := logger.SetGlobalRouting("../kvconfig.yml")
	if err != nil {
		panic(err)
	}
}

func TestRoutingRules(t *testing.T) {
	t.Run("update-loop-lag-alert", func(t *testing.T) {
		mocklog := logger.NewMockCountLogger("workflow-manager")
		log = mocklog
		logPendingWorkflowUpdateLag(models.Workflow{
			WorkflowSummary: models.WorkflowSummary{
				ID:                 "id",
				LastUpdated:        strfmt.DateTime(time.Now()),
				WorkflowDefinition: &models.WorkflowDefinition{},
			},
		})
		counts := mocklog.RuleCounts()
		assert.Equal(t, 1, len(counts))
		assert.Contains(t, counts, "update-loop-lag-alert")
		assert.Equal(t, 1, counts["update-loop-lag-alert"])
	})

	t.Run("aws-sdk-go-counter", func(t *testing.T) {
		mocklog := logger.NewMockCountLogger("workflow-manager")
		log = mocklog
		LogSFNCounts([]counter.ServiceCount{
			{
				Service:   "sfn",
				Operation: "GetExecutionHistory",
				Count:     100,
			},
			{
				Service:   "sfn",
				Operation: "DescribeExecution",
				Count:     200,
			},
		})
		counts := mocklog.RuleCounts()
		assert.Equal(t, 1, len(counts))
		assert.Contains(t, counts, "aws-sdk-go-counter")
		assert.Equal(t, 2, counts["aws-sdk-go-counter"])
	})
}
