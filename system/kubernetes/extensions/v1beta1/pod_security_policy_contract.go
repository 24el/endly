package v1beta1

import (
	"errors"
	"fmt"
	vvc "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

/*autogenerated contract adapter*/

// PodSecurityPolicyCreateRequest represents request
type PodSecurityPolicyCreateRequest struct {
	service_ v1beta1.PodSecurityPolicyInterface
	*vvc.PodSecurityPolicy
}

// PodSecurityPolicyUpdateRequest represents request
type PodSecurityPolicyUpdateRequest struct {
	service_ v1beta1.PodSecurityPolicyInterface
	*vvc.PodSecurityPolicy
}

// PodSecurityPolicyDeleteRequest represents request
type PodSecurityPolicyDeleteRequest struct {
	service_ v1beta1.PodSecurityPolicyInterface
	Name     string
	*v1.DeleteOptions
}

// PodSecurityPolicyDeleteCollectionRequest represents request
type PodSecurityPolicyDeleteCollectionRequest struct {
	service_ v1beta1.PodSecurityPolicyInterface
	*v1.DeleteOptions
	ListOptions v1.ListOptions
}

// PodSecurityPolicyGetRequest represents request
type PodSecurityPolicyGetRequest struct {
	service_ v1beta1.PodSecurityPolicyInterface
	Name     string
	v1.GetOptions
}

// PodSecurityPolicyListRequest represents request
type PodSecurityPolicyListRequest struct {
	service_ v1beta1.PodSecurityPolicyInterface
	v1.ListOptions
}

// PodSecurityPolicyWatchRequest represents request
type PodSecurityPolicyWatchRequest struct {
	service_ v1beta1.PodSecurityPolicyInterface
	v1.ListOptions
}

// PodSecurityPolicyPatchRequest represents request
type PodSecurityPolicyPatchRequest struct {
	service_     v1beta1.PodSecurityPolicyInterface
	Name         string
	Pt           types.PatchType
	Data         []byte
	Subresources []string
}

func init() {
	register(&PodSecurityPolicyCreateRequest{})
	register(&PodSecurityPolicyUpdateRequest{})
	register(&PodSecurityPolicyDeleteRequest{})
	register(&PodSecurityPolicyDeleteCollectionRequest{})
	register(&PodSecurityPolicyGetRequest{})
	register(&PodSecurityPolicyListRequest{})
	register(&PodSecurityPolicyWatchRequest{})
	register(&PodSecurityPolicyPatchRequest{})
}

func (r *PodSecurityPolicyCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodSecurityPolicyInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodSecurityPolicyInterface", service)
	}
	return nil
}

func (r *PodSecurityPolicyCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.PodSecurityPolicy)
	return result, err
}

func (r *PodSecurityPolicyCreateRequest) GetId() string {
	return "extensions/v1beta1.PodSecurityPolicy.Create"
}

func (r *PodSecurityPolicyUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodSecurityPolicyInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodSecurityPolicyInterface", service)
	}
	return nil
}

func (r *PodSecurityPolicyUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.PodSecurityPolicy)
	return result, err
}

func (r *PodSecurityPolicyUpdateRequest) GetId() string {
	return "extensions/v1beta1.PodSecurityPolicy.Update"
}

func (r *PodSecurityPolicyDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodSecurityPolicyInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodSecurityPolicyInterface", service)
	}
	return nil
}

func (r *PodSecurityPolicyDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name, r.DeleteOptions)
	return result, err
}

func (r *PodSecurityPolicyDeleteRequest) GetId() string {
	return "extensions/v1beta1.PodSecurityPolicy.Delete"
}

func (r *PodSecurityPolicyDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodSecurityPolicyInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodSecurityPolicyInterface", service)
	}
	return nil
}

func (r *PodSecurityPolicyDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions, r.ListOptions)
	return result, err
}

func (r *PodSecurityPolicyDeleteCollectionRequest) GetId() string {
	return "extensions/v1beta1.PodSecurityPolicy.DeleteCollection"
}

func (r *PodSecurityPolicyGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodSecurityPolicyInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodSecurityPolicyInterface", service)
	}
	return nil
}

func (r *PodSecurityPolicyGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name, r.GetOptions)
	return result, err
}

func (r *PodSecurityPolicyGetRequest) GetId() string {
	return "extensions/v1beta1.PodSecurityPolicy.Get"
}

func (r *PodSecurityPolicyListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodSecurityPolicyInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodSecurityPolicyInterface", service)
	}
	return nil
}

func (r *PodSecurityPolicyListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err
}

func (r *PodSecurityPolicyListRequest) GetId() string {
	return "extensions/v1beta1.PodSecurityPolicy.List"
}

func (r *PodSecurityPolicyWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodSecurityPolicyInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodSecurityPolicyInterface", service)
	}
	return nil
}

func (r *PodSecurityPolicyWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err
}

func (r *PodSecurityPolicyWatchRequest) GetId() string {
	return "extensions/v1beta1.PodSecurityPolicy.Watch"
}

func (r *PodSecurityPolicyPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.PodSecurityPolicyInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.PodSecurityPolicyInterface", service)
	}
	return nil
}

func (r *PodSecurityPolicyPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name, r.Pt, r.Data, r.Subresources...)
	return result, err
}

func (r *PodSecurityPolicyPatchRequest) GetId() string {
	return "extensions/v1beta1.PodSecurityPolicy.Patch"
}
