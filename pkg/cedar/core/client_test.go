package cedarcore

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	ld "gopkg.in/launchdarkly/go-server-sdk.v5"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/apperrors"
	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/client"
	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/client/system"
	genmodels "github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
	"github.com/cmsgov/easi-app/pkg/models"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	"github.com/cmsgov/easi-app/pkg/cedar/core/mocks"
)

type ClientTestSuite struct {
	suite.Suite
	logger *zap.Logger
}

func TestClientTestSuite(t *testing.T) {
	tests := &ClientTestSuite{
		Suite:  suite.Suite{},
		logger: zap.NewExample(),
	}
	suite.Run(t, tests)
}

func CreateTestClient(clientService *mocks.ClientService) Client {
	client := Client{
		cedarCoreEnabled: func(c context.Context) bool {
			return true
		},
		hc:   http.DefaultClient, // maybe mock this?
		auth: httptransport.BasicAuth("testUser", "testPassword"),
		sdk: &client.CEDARCoreAPI{
			System: clientService,
		},
	}

	return client
}

func (s ClientTestSuite) TestClientWithStub() {
	ctx := appcontext.WithLogger(context.Background(), s.logger)

	s.Run("GetSystem returns ResourceNotFoundError for nonexistent system", func() {
		mockSystemClient := new(mocks.ClientService)
		mockSystemClient.On("SystemSummaryFindByID", mock.MatchedBy(func(params *system.SystemSummaryFindByIDParams) bool {
			return true
		}), mock.MatchedBy(func(authInfo runtime.ClientAuthInfoWriter) bool {
			return true
		})).Return(&system.SystemSummaryFindByIDOK{
			Payload: &genmodels.SystemSummaryResponse{
				SystemSummary: []*genmodels.SystemSummary{},
			},
		}, nil)

		c := CreateTestClient(mockSystemClient)
		resp, err := c.GetSystem(ctx, "doesntexist")
		s.Nil(resp)
		s.Assertions.IsType(&apperrors.ResourceNotFoundError{}, err)
	})
}

func (s ClientTestSuite) TestClient() {
	ctx := appcontext.WithLogger(context.Background(), s.logger)

	ldClient, err := ld.MakeCustomClient("fake", ld.Config{Offline: true}, 0)
	s.NoError(err)

	s.Run("Instantiation successful", func() {
		c := NewClient("fake", "fake", ldClient)
		s.NotNil(c)
	})

	s.Run("LD defaults protects invocation of GetSystemSummary", func() {
		c := NewClient("fake", "fake", ldClient)
		resp, err := c.GetSystemSummary(ctx)
		s.NoError(err)

		blankSummary := []*models.CedarSystem{}
		s.Equal(resp, blankSummary)
	})
	s.Run("LD defaults protects invocation of GetSystem", func() {
		c := NewClient("fake", "fake", ldClient)
		resp, err := c.GetSystem(ctx, "fake")
		s.NoError(err)

		blankSummary := models.CedarSystem{}
		s.Equal(*resp, blankSummary)
	})
}
