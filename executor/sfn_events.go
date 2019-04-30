package executor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sfn"
)

func reverseHistory(events []*sfn.HistoryEvent) []*sfn.HistoryEvent {
	numEvents := len(events)
	reversedEvents := make([]*sfn.HistoryEvent, numEvents)
	for i := 0; i < numEvents; i++ {
		reversedEvents[i] = events[numEvents-1-i]
	}
	return reversedEvents
}

func isSupportedStateEnteredEvent(evt *sfn.HistoryEvent) bool {
	switch aws.StringValue(evt.Type) {
	case sfn.HistoryEventTypeTaskStateEntered,
		sfn.HistoryEventTypeChoiceStateEntered,
		sfn.HistoryEventTypeSucceedStateEntered:
		return true
	default:
		return false
	}
}
