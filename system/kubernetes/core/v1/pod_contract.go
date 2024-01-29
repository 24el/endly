package v1

import (
	"errors"
	"fmt"
	vvc "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

/*autogenerated contract adapter*/

// PodCreateRequest represents request
type PodCreateRequest struct {
	service_ v1.PodInterface
	*vvc.Pod
}

// PodUpdateRequest represents request
type PodUpdateRequest struct {
	service_ v1.PodInterface
	*vvc.Pod
}

// PodUpdateStatusRequest represents request
type PodUpdateStatusRequest struct {
	service_ v1.PodInterface
	*vvc.Pod
}

// PodDeleteRequest represents request
type PodDeleteRequest struct {
	service_ v1.PodInterface
	Name     string
	*metav1.DeleteOptions
}

// PodDeleteCollectionRequest represents request
type PodDeleteCollectionRequest struct {
	service_ v1.PodInterface
	*metav1.DeleteOptions
	ListOptions metav1.ListOptions
}

// PodGetRequest represents request
type PodGetRequest struct {
	service_ v1.PodInterface
	Name     string
	metav1.GetOptions
}

// PodListRequest represents request
type PodListRequest struct {
	service_ v1.PodInterface
	metav1.ListOptions
}

// PodWatchRequest represents request
type PodWatchRequest struct {
	service_ v1.PodInterface
	metav1.ListOptions
}

// PodPatchRequest represents request
type PodPatchRequest struct {
	service_     v1.PodInterface
	Name         string
	Pt           types.PatchType
	Data         []byte
	Subresources []string
}

func init() {
	register(&PodCreateRequest{})
	register(&PodUpdateRequest{})
	register(&PodUpdateStatusRequest{})
	register(&PodDeleteRequest{})
	register(&PodDeleteCollectionRequest{})
	register(&PodGetRequest{})
	register(&PodListRequest{})
	register(&PodWatchRequest{})
	register(&PodPatchRequest{})
}

func (r *PodCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.Pod)
	return result, err
}

func (r *PodCreateRequest) GetId() string {
	return "v1.Pod.Create"
}

func (r *PodUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.Pod)
	return result, err
}

func (r *PodUpdateRequest) GetId() string {
	return "v1.Pod.Update"
}

func (r *PodUpdateStatusRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodUpdateStatusRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.UpdateStatus(r.Pod)
	return result, err
}

func (r *PodUpdateStatusRequest) GetId() string {
	return "v1.Pod.UpdateStatus"
}

func (r *PodDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name, r.DeleteOptions)
	return result, err
}

func (r *PodDeleteRequest) GetId() string {
	return "v1.Pod.Delete"
}

func (r *PodDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions, r.ListOptions)
	return result, err
}

func (r *PodDeleteCollectionRequest) GetId() string {
	return "v1.Pod.DeleteCollection"
}

func (r *PodGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name, r.GetOptions)
	return result, err
}

func (r *PodGetRequest) GetId() string {
	return "v1.Pod.Get"
}

func (r *PodListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err
}

func (r *PodListRequest) GetId() string {
	return "v1.Pod.List"
}

func (r *PodWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err
}

func (r *PodWatchRequest) GetId() string {
	return "v1.Pod.Watch"
}

func (r *PodPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodInterface", service)
	}
	return nil
}

func (r *PodPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name, r.Pt, r.Data, r.Subresources...)
	return result, err
}

func (r *PodPatchRequest) GetId() string {
	return "v1.Pod.Patch"
}
