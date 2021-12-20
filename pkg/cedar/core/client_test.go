package cedarcore

import (
	"context"
	"net/http"
	"testing"

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

/*
type ClientService interface {
	SystemDetailAdd(params *SystemDetailAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemDetailAddOK, error)

	SystemDetailFindByID(params *SystemDetailFindByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemDetailFindByIDOK, error)

	SystemDetailUpdate(params *SystemDetailUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemDetailUpdateOK, error)

	SystemSummaryFindByID(params *SystemSummaryFindByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemSummaryFindByIDOK, error)

	SystemSummaryFindList(params *SystemSummaryFindListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SystemSummaryFindListOK, error)

	SetTransport(transport runtime.ClientTransport)
}
*/

type StubClientService struct {
}

func (stub *StubClientService) SystemDetailAdd(params *system.SystemDetailAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...system.ClientOption) (*system.SystemDetailAddOK, error) {
	panic("unimplemented")
}

func (stub *StubClientService) SystemDetailFindByID(params *system.SystemDetailFindByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...system.ClientOption) (*system.SystemDetailFindByIDOK, error) {
	panic("unimplemented")
}

func (stub *StubClientService) SystemDetailUpdate(params *system.SystemDetailUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...system.ClientOption) (*system.SystemDetailUpdateOK, error) {
	panic("unimplemented")
}

func (stub *StubClientService) SystemSummaryFindByID(params *system.SystemSummaryFindByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...system.ClientOption) (*system.SystemSummaryFindByIDOK, error) {
	return &system.SystemSummaryFindByIDOK{
		Payload: &genmodels.SystemSummaryResponse{
			SystemSummary: []*genmodels.SystemSummary{},
		},
	}, nil
}

func (stub *StubClientService) SystemSummaryFindList(params *system.SystemSummaryFindListParams, authInfo runtime.ClientAuthInfoWriter, opts ...system.ClientOption) (*system.SystemSummaryFindListOK, error) {
	panic("unimplemented")
}

func (stub *StubClientService) SetTransport(transport runtime.ClientTransport) {
	panic("unimplemented")
}

func CreateTestClient() Client {
	client := Client{
		cedarCoreEnabled: func(c context.Context) bool {
			return true
		},
		hc:   http.DefaultClient, // maybe mock this?
		auth: httptransport.BasicAuth("testUser", "testPassword"),
		sdk: &client.CEDARCoreAPI{
			System: &StubClientService{},
		},
	}

	return client
}

func (s ClientTestSuite) TestClientWithStub() {
	ctx := appcontext.WithLogger(context.Background(), s.logger)

	s.Run("GetSystem returns ResourceNotFoundError for nonexistent system", func() {
		c := CreateTestClient()
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
