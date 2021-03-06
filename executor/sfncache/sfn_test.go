package sfncache

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Clever/workflow-manager/mocks"
)

func TestSFNCache_DescribeStateMachineWithContext(t *testing.T) {
	testARN := aws.String("arn:aws:states:us-east-1:111122223333:stateMachine:HelloWorld-StateMachine")
	testOutput := &sfn.DescribeStateMachineOutput{StateMachineArn: testARN}

	tests := []struct {
		name         string
		input        *sfn.DescribeStateMachineInput
		expectations func(m *mocks.MockSFNAPI)
		assertions   func(t *testing.T, s *SFNCache, got *sfn.DescribeStateMachineOutput, err error)
	}{
		{
			name:  "happy path + caches for an arn",
			input: &sfn.DescribeStateMachineInput{StateMachineArn: testARN},
			expectations: func(m *mocks.MockSFNAPI) {
				m.EXPECT().
					DescribeStateMachineWithContext(gomock.Any(), gomock.Any()).
					Return(testOutput, nil).
					Times(1)
			},
			assertions: func(t *testing.T, s *SFNCache, got *sfn.DescribeStateMachineOutput, err error) {
				require.NoError(t, err)
				require.Equal(t, testOutput, got)
				// ensure the object is cached
				require.Equal(t, 1, s.describeStateMachineWithContextCache.Len())
				for i := 0; i < 1000; i++ {
					// ensure we don't make any other API calls
					output, err := s.DescribeStateMachineWithContext(context.Background(), &sfn.DescribeStateMachineInput{StateMachineArn: testARN})
					require.NoError(t, err)
					require.Equal(t, testOutput, output)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()
			mockSFNAPI := mocks.NewMockSFNAPI(mockController)
			tt.expectations(mockSFNAPI)
			s, err := New(mockSFNAPI)
			require.NoError(t, err)

			got, err := s.DescribeStateMachineWithContext(context.Background(), tt.input)
			tt.assertions(t, s, got, err)
		})
	}
}
