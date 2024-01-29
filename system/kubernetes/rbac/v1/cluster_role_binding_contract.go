package v1

import (
	"errors"
	"fmt"
	vvc "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/typed/rbac/v1"
)

/*autogenerated contract adapter*/

// ClusterRoleBindingCreateRequest represents request
type ClusterRoleBindingCreateRequest struct {
	service_ v1.ClusterRoleBindingInterface
	*vvc.ClusterRoleBinding
}

// ClusterRoleBindingUpdateRequest represents request
type ClusterRoleBindingUpdateRequest struct {
	service_ v1.ClusterRoleBindingInterface
	*vvc.ClusterRoleBinding
}

// ClusterRoleBindingDeleteRequest represents request
type ClusterRoleBindingDeleteRequest struct {
	service_ v1.ClusterRoleBindingInterface
	Name     string
	*metav1.DeleteOptions
}

// ClusterRoleBindingDeleteCollectionRequest represents request
type ClusterRoleBindingDeleteCollectionRequest struct {
	service_ v1.ClusterRoleBindingInterface
	*metav1.DeleteOptions
	ListOptions metav1.ListOptions
}

// ClusterRoleBindingGetRequest represents request
type ClusterRoleBindingGetRequest struct {
	service_ v1.ClusterRoleBindingInterface
	Name     string
	metav1.GetOptions
}

// ClusterRoleBindingListRequest represents request
type ClusterRoleBindingListRequest struct {
	service_ v1.ClusterRoleBindingInterface
	metav1.ListOptions
}

// ClusterRoleBindingWatchRequest represents request
type ClusterRoleBindingWatchRequest struct {
	service_ v1.ClusterRoleBindingInterface
	metav1.ListOptions
}

// ClusterRoleBindingPatchRequest represents request
type ClusterRoleBindingPatchRequest struct {
	service_     v1.ClusterRoleBindingInterface
	Name         string
	Pt           types.PatchType
	Data         []byte
	Subresources []string
}

func init() {
	register(&ClusterRoleBindingCreateRequest{})
	register(&ClusterRoleBindingUpdateRequest{})
	register(&ClusterRoleBindingDeleteRequest{})
	register(&ClusterRoleBindingDeleteCollectionRequest{})
	register(&ClusterRoleBindingGetRequest{})
	register(&ClusterRoleBindingListRequest{})
	register(&ClusterRoleBindingWatchRequest{})
	register(&ClusterRoleBindingPatchRequest{})
}

func (r *ClusterRoleBindingCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ClusterRoleBindingInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ClusterRoleBindingInterface", service)
	}
	return nil
}

func (r *ClusterRoleBindingCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.ClusterRoleBinding)
	return result, err
}

func (r *ClusterRoleBindingCreateRequest) GetId() string {
	return "rbac/v1.ClusterRoleBinding.Create"
}

func (r *ClusterRoleBindingUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ClusterRoleBindingInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ClusterRoleBindingInterface", service)
	}
	return nil
}

func (r *ClusterRoleBindingUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.ClusterRoleBinding)
	return result, err
}

func (r *ClusterRoleBindingUpdateRequest) GetId() string {
	return "rbac/v1.ClusterRoleBinding.Update"
}

func (r *ClusterRoleBindingDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ClusterRoleBindingInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ClusterRoleBindingInterface", service)
	}
	return nil
}

func (r *ClusterRoleBindingDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name, r.DeleteOptions)
	return result, err
}

func (r *ClusterRoleBindingDeleteRequest) GetId() string {
	return "rbac/v1.ClusterRoleBinding.Delete"
}

func (r *ClusterRoleBindingDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ClusterRoleBindingInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ClusterRoleBindingInterface", service)
	}
	return nil
}

func (r *ClusterRoleBindingDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions, r.ListOptions)
	return result, err
}

func (r *ClusterRoleBindingDeleteCollectionRequest) GetId() string {
	return "rbac/v1.ClusterRoleBinding.DeleteCollection"
}

func (r *ClusterRoleBindingGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ClusterRoleBindingInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ClusterRoleBindingInterface", service)
	}
	return nil
}

func (r *ClusterRoleBindingGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name, r.GetOptions)
	return result, err
}

func (r *ClusterRoleBindingGetRequest) GetId() string {
	return "rbac/v1.ClusterRoleBinding.Get"
}

func (r *ClusterRoleBindingListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ClusterRoleBindingInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ClusterRoleBindingInterface", service)
	}
	return nil
}

func (r *ClusterRoleBindingListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err
}

func (r *ClusterRoleBindingListRequest) GetId() string {
	return "rbac/v1.ClusterRoleBinding.List"
}

func (r *ClusterRoleBindingWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ClusterRoleBindingInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ClusterRoleBindingInterface", service)
	}
	return nil
}

func (r *ClusterRoleBindingWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err
}

func (r *ClusterRoleBindingWatchRequest) GetId() string {
	return "rbac/v1.ClusterRoleBinding.Watch"
}

func (r *ClusterRoleBindingPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ClusterRoleBindingInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ClusterRoleBindingInterface", service)
	}
	return nil
}

func (r *ClusterRoleBindingPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name, r.Pt, r.Data, r.Subresources...)
	return result, err
}

func (r *ClusterRoleBindingPatchRequest) GetId() string {
	return "rbac/v1.ClusterRoleBinding.Patch"
}
