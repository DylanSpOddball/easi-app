package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/cmsgov/easi-app/pkg/graph/generated"
	"github.com/cmsgov/easi-app/pkg/graph/model"
	"github.com/cmsgov/easi-app/pkg/models"
)

func (r *accessibilityRequestResolver) Documents(ctx context.Context, obj *models.AccessibilityRequest) ([]*model.AccessibilityRequestDocument, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *accessibilityRequestResolver) System(ctx context.Context, obj *models.AccessibilityRequest) (*models.System, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAccessibilityRequest(ctx context.Context, input *model.CreateAccessibilityRequestInput) (*model.CreateAccessibilityRequestPayload, error) {
	request, err := r.store.CreateAccessibilityRequest(ctx, &models.AccessibilityRequest{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}

	return &model.CreateAccessibilityRequestPayload{
		AccessibilityRequest: request,
		UserErrors:           nil,
	}, nil
}

func (r *mutationResolver) CreateTestDate(ctx context.Context, input *model.CreateTestDateInput) (*model.CreateTestDatePayload, error) {
	testDate, err := r.createTestDate(ctx, &models.TestDate{
		TestType:  input.TestType,
		Date:      input.Date,
		Score:     input.Score,
		RequestID: input.RequestID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CreateTestDatePayload{TestDate: testDate, UserErrors: nil}, nil
}

func (r *queryResolver) AccessibilityRequest(ctx context.Context, id uuid.UUID) (*models.AccessibilityRequest, error) {
	return r.store.FetchAccessibilityRequestByID(ctx, id)
}

func (r *queryResolver) AccessibilityRequests(ctx context.Context, after *string, first int) (*model.AccessibilityRequestsConnection, error) {
	requests, queryErr := r.store.FetchAccessibilityRequests(ctx)
	if queryErr != nil {
		return nil, gqlerror.Errorf("query error: %s", queryErr)
	}

	edges := []*model.AccessibilityRequestEdge{}

	for _, request := range requests {
		node := request
		edges = append(edges, &model.AccessibilityRequestEdge{
			Node: &node,
		})
	}

	return &model.AccessibilityRequestsConnection{Edges: edges}, nil
}

func (r *queryResolver) Systems(ctx context.Context, after *string, first int) (*model.SystemConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// AccessibilityRequest returns generated.AccessibilityRequestResolver implementation.
func (r *Resolver) AccessibilityRequest() generated.AccessibilityRequestResolver {
	return &accessibilityRequestResolver{r}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type accessibilityRequestResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
