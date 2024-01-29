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

// IngressCreateRequest represents request
type IngressCreateRequest struct {
	service_ v1beta1.IngressInterface
	*vvc.Ingress
}

// IngressUpdateRequest represents request
type IngressUpdateRequest struct {
	service_ v1beta1.IngressInterface
	*vvc.Ingress
}

// IngressUpdateStatusRequest represents request
type IngressUpdateStatusRequest struct {
	service_ v1beta1.IngressInterface
	*vvc.Ingress
}

// IngressDeleteRequest represents request
type IngressDeleteRequest struct {
	service_ v1beta1.IngressInterface
	Name     string
	*v1.DeleteOptions
}

// IngressDeleteCollectionRequest represents request
type IngressDeleteCollectionRequest struct {
	service_ v1beta1.IngressInterface
	*v1.DeleteOptions
	ListOptions v1.ListOptions
}

// IngressGetRequest represents request
type IngressGetRequest struct {
	service_ v1beta1.IngressInterface
	Name     string
	v1.GetOptions
}

// IngressListRequest represents request
type IngressListRequest struct {
	service_ v1beta1.IngressInterface
	v1.ListOptions
}

// IngressWatchRequest represents request
type IngressWatchRequest struct {
	service_ v1beta1.IngressInterface
	v1.ListOptions
}

// IngressPatchRequest represents request
type IngressPatchRequest struct {
	service_     v1beta1.IngressInterface
	Name         string
	Pt           types.PatchType
	Data         []byte
	Subresources []string
}

func init() {
	register(&IngressCreateRequest{})
	register(&IngressUpdateRequest{})
	register(&IngressUpdateStatusRequest{})
	register(&IngressDeleteRequest{})
	register(&IngressDeleteCollectionRequest{})
	register(&IngressGetRequest{})
	register(&IngressListRequest{})
	register(&IngressWatchRequest{})
	register(&IngressPatchRequest{})
}

func (r *IngressCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.Ingress)
	return result, err
}

func (r *IngressCreateRequest) GetId() string {
	return "extensions/v1beta1.Ingress.Create"
}

func (r *IngressUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.Ingress)
	return result, err
}

func (r *IngressUpdateRequest) GetId() string {
	return "extensions/v1beta1.Ingress.Update"
}

func (r *IngressUpdateStatusRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressUpdateStatusRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.UpdateStatus(r.Ingress)
	return result, err
}

func (r *IngressUpdateStatusRequest) GetId() string {
	return "extensions/v1beta1.Ingress.UpdateStatus"
}

func (r *IngressDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name, r.DeleteOptions)
	return result, err
}

func (r *IngressDeleteRequest) GetId() string {
	return "extensions/v1beta1.Ingress.Delete"
}

func (r *IngressDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions, r.ListOptions)
	return result, err
}

func (r *IngressDeleteCollectionRequest) GetId() string {
	return "extensions/v1beta1.Ingress.DeleteCollection"
}

func (r *IngressGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name, r.GetOptions)
	return result, err
}

func (r *IngressGetRequest) GetId() string {
	return "extensions/v1beta1.Ingress.Get"
}

func (r *IngressListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err
}

func (r *IngressListRequest) GetId() string {
	return "extensions/v1beta1.Ingress.List"
}

func (r *IngressWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err
}

func (r *IngressWatchRequest) GetId() string {
	return "extensions/v1beta1.Ingress.Watch"
}

func (r *IngressPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1beta1.IngressInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1beta1.IngressInterface", service)
	}
	return nil
}

func (r *IngressPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name, r.Pt, r.Data, r.Subresources...)
	return result, err
}

func (r *IngressPatchRequest) GetId() string {
	return "extensions/v1beta1.Ingress.Patch"
}
