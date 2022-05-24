package apimpl

import (
	"context"
	"github.com/djcass44/go-utils/utilities/sliceutils"
	"github.com/go-logr/logr"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/api"
)

func NewAuthority(conf *Configuration, subjectHasGlobalRole SubjectHasGlobalRole, subjectCanDoAction SubjectCanDoAction) *Authority {
	return &Authority{
		conf:                 conf,
		subjectHasGlobalRole: subjectHasGlobalRole,
		subjectCanDoAction:   subjectCanDoAction,
	}
}

func (a *Authority) Can(ctx context.Context, request *api.AccessRequest) (*api.AccessResponse, error) {
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", request.GetSubject(), "Resource", request.GetResource(), "Action", request.GetAction().String())
	log.Info("checking if subject can perform action on resource")
	// check global roles
	ok, err := a.hasGlobalRole(ctx, request)
	if err != nil {
		return nil, err
	}
	if ok {
		return &api.AccessResponse{Message: "", Ok: true}, nil
	}
	return &api.AccessResponse{Message: "", Ok: false}, nil
}

func canPerformAction(actions []string, verb api.Verb) bool {
	// if the role has no actions, skip it
	if len(actions) == 0 {
		return false
	}
	// if the only action is SUDO don't skip it
	if len(actions) == 1 {
		if actions[0] == api.Verb_SUDO.String() {
			return true
		}
	}
	// otherwise check that the requested action
	// is in the list of supported actions
	return sliceutils.Includes(actions, verb.String())
}

func (a *Authority) hasGlobalRole(ctx context.Context, request *api.AccessRequest) (bool, error) {
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", request.GetSubject(), "Resource", request.GetResource(), "Action", request.GetAction().String())
	log.V(1).Info("checking global roles")
	// iterate through each of the global
	// roles
	for k, v := range a.conf.Globals {
		log = log.WithValues("Role", k)
		log.V(2).Info("checking global role")
		// check if the role matches the
		// verb that we're requesting
		if !canPerformAction(v, request.GetAction()) {
			log.V(2).Info("skipping global role as it does not include our Verb")
			continue
		}
		log.V(2).Info("verb matched, checking if subject has the role")
		// check if the subject actually has
		// the role
		ok, err := a.subjectHasGlobalRole(ctx, request.GetSubject(), k)
		if err != nil {
			log.Error(err, "failed to check if subject has role")
			return false, err
		}
		log.V(2).Info("successfully checked role membership", "Member", ok)
		if ok {
			return true, nil
		}
	}
	log.V(1).Info("failed to locate matching global role")
	return false, nil
}
