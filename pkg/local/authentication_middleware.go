package local

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-openapi/swag"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/authentication"
)

// DevUserConfig is the set of values that can be passed in a request header
type DevUserConfig struct {
	EUA      string   `json:"euaId"`
	JobCodes []string `json:"jobCodes"`
}

func authenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := appcontext.ZLogger(r.Context())
		logger.Info("Using local authorization middleware")

		if len(r.Header["Authorization"]) == 0 {
			logger.Info("No local auth header present")
			next.ServeHTTP(w, r)
			return
		}

		// don't attempt to handle local auth if the Authorization Header doesn't start with "Local"
		if !strings.HasPrefix(r.Header["Authorization"][0], "Local") {
			next.ServeHTTP(w, r)
			return
		}

		tokenParts := strings.Split(r.Header["Authorization"][0], "Local ")
		if len(tokenParts) < 2 {
			logger.Error("Invalid local auth header")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		devUserConfigJSON := tokenParts[1]
		if devUserConfigJSON == "" {
			logger.Error("Empty dev user config JSON")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		config := DevUserConfig{}

		if parseErr := json.Unmarshal([]byte(devUserConfigJSON), &config); parseErr != nil {
			logger.Error("Could not parse local auth JSON")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		logger.Info("Using local authorization middleware and populating EUA ID and job codes")
		ctx := appcontext.WithPrincipal(r.Context(), &authentication.EUAPrincipal{
			EUAID:            strings.ToUpper(config.EUA),
			JobCodeEASi:      true,
			JobCodeGRT:       swag.ContainsStrings(config.JobCodes, "EASI_D_GOVTEAM"),
			JobCode508User:   swag.ContainsStrings(config.JobCodes, "EASI_D_508_USER"),
			JobCode508Tester: swag.ContainsStrings(config.JobCodes, "EASI_D_508_TESTER"),
			JobCodeTRBAdmin:  swag.ContainsStrings(config.JobCodes, "EASI_TRB_ADMIN_D"),
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewLocalAuthenticationMiddleware stubs out context info for local (non-Okta) authentication
func NewLocalAuthenticationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return authenticateMiddleware(next)
	}
}
