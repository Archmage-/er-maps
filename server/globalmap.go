package server

import (
	"github.com/Archmage-/er-maps/core"
	"github.com/Archmage-/er-maps/modules/minemap"
	"net/http"
	//"net/url"
)

type globalMap struct {
	provider minemapprovider.IProvider
}

func NewGlobalMap(provider minemapprovider.IProvider) *globalMap {
	return &globalMap{provider: provider}
}

func (this *globalMap) Handle(r *http.Request, rc *core.RequestContext) (data []byte, e error) {
	//check input params
	//params := r.URL
	//show map

	return nil, nil
}
