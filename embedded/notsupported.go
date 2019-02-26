package embedded

import (
	"context"
	"errors"

	"github.com/Clever/workflow-manager/gen-go/client"
	"github.com/Clever/workflow-manager/gen-go/models"
)

// ErrNotSupported is returned when the method is not supported.
var ErrNotSupported = errors.New("not supported")

func (e embedded) PostStateResource(ctx context.Context, i *models.NewStateResource) (*models.StateResource, error) {
	return nil, ErrNotSupported
}

func (e embedded) DeleteStateResource(ctx context.Context, i *models.DeleteStateResourceInput) error {
	return ErrNotSupported
}

func (e embedded) GetStateResource(ctx context.Context, i *models.GetStateResourceInput) (*models.StateResource, error) {
	return nil, ErrNotSupported
}

func (e embedded) PutStateResource(ctx context.Context, i *models.PutStateResourceInput) (*models.StateResource, error) {
	return nil, ErrNotSupported
}

func (e embedded) GetWorkflowDefinitionVersionsByName(ctx context.Context, i *models.GetWorkflowDefinitionVersionsByNameInput) ([]models.WorkflowDefinition, error) {
	return nil, ErrNotSupported
}

func (e embedded) NewWorkflowDefinition(ctx context.Context, i *models.NewWorkflowDefinitionRequest) (*models.WorkflowDefinition, error) {
	return nil, ErrNotSupported
}

func (e embedded) UpdateWorkflowDefinition(ctx context.Context, i *models.UpdateWorkflowDefinitionInput) (*models.WorkflowDefinition, error) {
	return nil, ErrNotSupported
}

func (e embedded) NewGetWorkflowsIter(ctx context.Context, i *models.GetWorkflowsInput) (client.GetWorkflowsIter, error) {
	return nil, ErrNotSupported
}

func (e embedded) ResumeWorkflowByID(ctx context.Context, i *models.ResumeWorkflowByIDInput) (*models.Workflow, error) {
	return nil, ErrNotSupported
}

func (e embedded) ResolveWorkflowByID(ctx context.Context, workflowID string) error {
	return ErrNotSupported
}
