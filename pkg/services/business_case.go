package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/apperrors"
	"github.com/cmsgov/easi-app/pkg/appvalidation"
	"github.com/cmsgov/easi-app/pkg/models"
)

// NewFetchBusinessCaseByID is a service to fetch the business case by id
func NewFetchBusinessCaseByID(
	config Config,
	fetch func(id uuid.UUID) (*models.BusinessCase, error),
) func(id uuid.UUID) (*models.BusinessCase, error) {
	return func(id uuid.UUID) (*models.BusinessCase, error) {
		businessCase, err := fetch(id)
		if err != nil {
			config.logger.Error("failed to fetch business case")
			return &models.BusinessCase{}, &apperrors.QueryError{
				Err:       err,
				Model:     businessCase,
				Operation: apperrors.QueryFetch,
			}
		}
		return businessCase, nil
	}
}

// NewAuthorizeCreateBusinessCase returns a function
// that authorizes a user for creating a business case
func NewAuthorizeCreateBusinessCase(logger *zap.Logger) func(
	context context.Context,
	intake *models.SystemIntake,
) (bool, error) {
	return func(context context.Context, intake *models.SystemIntake) (bool, error) {
		if intake == nil {
			logger.With(zap.Bool("Authorized", false)).
				With(zap.String("Operation", "CreateBusinessCase")).
				Info("intake does not exist")
			return false, nil
		}
		euaID, ok := appcontext.EuaID(context)
		if !ok {
			// Default to failure to authorize and create a quick audit log
			logger.With(zap.Bool("Authorized", false)).
				With(zap.String("Operation", "CreateBusinessCase")).
				Info("something went wrong fetching the eua id from the context")
			return false, nil
		}
		// If intake is owned by user, authorize
		if euaID == intake.EUAUserID {
			logger.With(zap.Bool("Authorized", true)).
				With(zap.String("Operation", "CreateBusinessCase")).
				Info("user authorized to create business case")
			return true, nil
		}
		// Default to failure to authorize and create a quick audit log
		logger.With(zap.Bool("Authorized", false)).
			With(zap.String("Operation", "CreateBusinessCase")).
			Info("unauthorized attempt to create business case")
		return false, nil
	}
}

// NewCreateBusinessCase is a service to create a business case
func NewCreateBusinessCase(
	config Config,
	fetchIntake func(id uuid.UUID) (*models.SystemIntake, error),
	authorize func(context context.Context, intake *models.SystemIntake) (bool, error),
	create func(businessCase *models.BusinessCase) (*models.BusinessCase, error),
) func(context context.Context, businessCase *models.BusinessCase) (*models.BusinessCase, error) {
	return func(context context.Context, businessCase *models.BusinessCase) (*models.BusinessCase, error) {
		intake, err := fetchIntake(businessCase.SystemIntakeID)
		if err != nil {
			// We return an empty id in this error because the business case hasn't been created
			return &models.BusinessCase{}, &apperrors.ResourceConflictError{
				Err:        errors.New("system intake is required to create a business case"),
				Resource:   models.BusinessCase{},
				ResourceID: "",
			}
		}
		ok, err := authorize(context, intake)
		if err != nil {
			return &models.BusinessCase{}, err
		}
		if !ok {
			return &models.BusinessCase{}, &apperrors.UnauthorizedError{Err: err}
		}
		err = appvalidation.BusinessCaseForCreation(businessCase, intake)
		if err != nil {
			return &models.BusinessCase{}, err
		}
		businessCase, err = create(businessCase)
		if err != nil {
			config.logger.Error("failed to create a business case")
			return &models.BusinessCase{}, &apperrors.QueryError{
				Err:       err,
				Model:     businessCase,
				Operation: apperrors.QueryPost,
			}
		}
		return businessCase, nil
	}
}

// NewFetchBusinessCasesByEuaID is a service to fetch a list of business cases by EUA ID
func NewFetchBusinessCasesByEuaID(
	config Config,
	fetch func(euaID string) (models.BusinessCases, error),
) func(euaID string) (models.BusinessCases, error) {
	return func(euaID string) (models.BusinessCases, error) {
		businessCases, err := fetch(euaID)
		if err != nil {
			config.logger.Error("failed to fetch business cases")
			return models.BusinessCases{}, &apperrors.QueryError{
				Err:       err,
				Model:     "business cases",
				Operation: "fetch",
			}
		}
		return businessCases, nil
	}
}
