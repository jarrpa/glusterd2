package snapshotcommands

import (
	"net/http"

	"github.com/gluster/glusterd2/glusterd2/servers/rest/route"
	restutils "github.com/gluster/glusterd2/glusterd2/servers/rest/utils"
	"github.com/prashanthpai/sunrpc"
)

// Plugin is a structure which implements GlusterdPlugin interface
type Plugin struct {
}

// Name returns name of plugin
func (p *Plugin) Name() string {
	return "snapshot"
}

// SunRPCProgram returns sunrpc program to register with Glusterd
func (p *Plugin) SunRPCProgram() sunrpc.Program {
	return nil
}

// RestRoutes returns list of REST API routes to register with Glusterd
func (p *Plugin) RestRoutes() route.Routes {
	return route.Routes{
		route.Route{
			Name:        "SnapshotCreate",
			Method:      "PUT",
			Pattern:     "/snapshot/",
			Version:     1,
			HandlerFunc: snapshotCreateHandler},
		route.Route{
			Name:        "SnapshotActivate",
			Method:      "POST",
			Pattern:     "/snapshot/{snapname}/activate",
			Version:     1,
			HandlerFunc: snapshotActivateHandler},
		route.Route{
			Name:        "SnapshotDeactivate",
			Method:      "POST",
			Pattern:     "/snapshot/{snapname}/deactivate",
			Version:     1,
			HandlerFunc: snapshotDeactivateHandler},
		route.Route{
			Name:        "SnapshotClone",
			Method:      "POST",
			Pattern:     "/snapshot/{snapname}/clone",
			Version:     1,
			HandlerFunc: snapshotCloneHandler},
		route.Route{
			Name:        "SnapshotRestore",
			Method:      "POST",
			Pattern:     "/snapshot/{snapname}/restore",
			Version:     1,
			HandlerFunc: snapshotRestoreHandler},
		route.Route{
			Name:        "SnapshotInfo",
			Method:      "GET",
			Pattern:     "/snapshots/{snapname}",
			Version:     1,
			HandlerFunc: snapshotInfoHandler},

		route.Route{
			Name:        "SnapshotList",
			Method:      "GET",
			Pattern:     "/snapshots",
			Version:     1,
			HandlerFunc: snapshotListHandler},
		/*
			route.Route{
				Name:        "SnapshotStatus",
				Method:      "GET",
				Pattern:     "/snapshot",
				Version:     1,
				HandlerFunc: snapshotStatusHandler},
		*/
		route.Route{
			Name:        "SnapshotDelete",
			Method:      "DELETE",
			Pattern:     "/snapshot",
			Version:     1,
			HandlerFunc: snapshotDeleteHandler},
		route.Route{
			Name:        "SnapshotConfigGet",
			Method:      "GET",
			Pattern:     "/snapshot/config",
			Version:     1,
			HandlerFunc: snapshotConfigGetHandler},
		route.Route{
			Name:        "SnapshotConfigSet",
			Method:      "POST",
			Pattern:     "/snapshot/config",
			Version:     1,
			HandlerFunc: snapshotConfigSetHandler},
		route.Route{
			Name:        "SnapshotConfigReset",
			Method:      "DELETE",
			Pattern:     "/snapshot/config",
			Version:     1,
			HandlerFunc: snapshotConfigResetHandler},
	}
}

func snapshotCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Create")
}
func snapshotActivateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Activate")
}
func snapshotDeactivateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Deactivate")
}
func snapshotInfoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Info")
}
func snapshotListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot List")
}

func snapshotCloneHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Clone")
}

func snapshotRestoreHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Restore")
}

func snapshotStatusHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Status/Info")
}

func snapshotDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Delete")
}

func snapshotConfigGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Config Get")
}

func snapshotConfigSetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Config Set")
}

func snapshotConfigResetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, "Snapshot Config Reset")
}

// RegisterStepFuncs registers transaction step functions with
// Glusterd Transaction framework
func (p *Plugin) RegisterStepFuncs() {
	return
}
