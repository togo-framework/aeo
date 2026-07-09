// Package aeo is the AEO togo plugin: Answer-Engine Optimization: per-page /md markdown twins, llms.txt index, sitemap sections and jsonb+i18n unwrapping — the clean content a model sees.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package aeo

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("aeo", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/aeo/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"aeo","status":"ok"}`))
		})
		return nil
	})
}
