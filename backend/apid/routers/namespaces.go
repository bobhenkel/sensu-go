package routers

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/sensu/sensu-go/types"
)

// NamespacesController represents the controller needs of the NamespacesRouter.
type NamespacesController interface {
	Query(ctx context.Context) ([]*types.Namespace, error)
	Find(ctx context.Context, name string) (*types.Namespace, error)
	Create(ctx context.Context, newOrg types.Namespace) error
	CreateOrReplace(ctx context.Context, newOrg types.Namespace) error
	Update(ctx context.Context, given types.Namespace) error
	Destroy(ctx context.Context, name string) error
}

// NamespacesRouter handles requests for /namespaces
type NamespacesRouter struct {
	controller NamespacesController
}

// NewNamespacesRouter instantiates new router for controlling check resources
func NewNamespacesRouter(ctrl NamespacesController) *NamespacesRouter {
	return &NamespacesRouter{
		controller: ctrl,
	}
}

// Mount the NamespacesRouter to a parent Router
func (r *NamespacesRouter) Mount(parent *mux.Router) {
	routes := ResourceRoute{Router: parent, PathPrefix: "/rbac/namespaces"}
	routes.GetAll(r.list)
	routes.Get(r.find)
	routes.Post(r.create)
	routes.Del(r.destroy)
	routes.Put(r.createOrReplace)
}

func (r *NamespacesRouter) list(req *http.Request) (interface{}, error) {
	records, err := r.controller.Query(req.Context())
	return records, err
}

func (r *NamespacesRouter) find(req *http.Request) (interface{}, error) {
	params := mux.Vars(req)
	id, err := url.PathUnescape(params["id"])
	if err != nil {
		return nil, err
	}
	record, err := r.controller.Find(req.Context(), id)
	return record, err
}

func (r *NamespacesRouter) create(req *http.Request) (interface{}, error) {
	org := types.Namespace{}
	if err := UnmarshalBody(req, &org); err != nil {
		return nil, err
	}

	err := r.controller.Create(req.Context(), org)
	return org, err
}

func (r *NamespacesRouter) createOrReplace(req *http.Request) (interface{}, error) {
	org := types.Namespace{}
	if err := UnmarshalBody(req, &org); err != nil {
		return nil, err
	}

	err := r.controller.CreateOrReplace(req.Context(), org)
	return org, err
}

func (r *NamespacesRouter) destroy(req *http.Request) (interface{}, error) {
	params := mux.Vars(req)
	id, err := url.PathUnescape(params["id"])
	if err != nil {
		return nil, err
	}
	err = r.controller.Destroy(req.Context(), id)
	return nil, err
}
