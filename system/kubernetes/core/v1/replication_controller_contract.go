package v1

import (
	"errors"
	"fmt"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	vvc "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

/*autogenerated contract adapter*/

// ReplicationControllerCreateRequest represents request
type ReplicationControllerCreateRequest struct {
	service_ v1.ReplicationControllerInterface
	*vvc.ReplicationController
}

// ReplicationControllerUpdateRequest represents request
type ReplicationControllerUpdateRequest struct {
	service_ v1.ReplicationControllerInterface
	*vvc.ReplicationController
}

// ReplicationControllerUpdateStatusRequest represents request
type ReplicationControllerUpdateStatusRequest struct {
	service_ v1.ReplicationControllerInterface
	*vvc.ReplicationController
}

// ReplicationControllerDeleteRequest represents request
type ReplicationControllerDeleteRequest struct {
	service_ v1.ReplicationControllerInterface
	Name     string
	*metav1.DeleteOptions
}

// ReplicationControllerDeleteCollectionRequest represents request
type ReplicationControllerDeleteCollectionRequest struct {
	service_ v1.ReplicationControllerInterface
	*metav1.DeleteOptions
	ListOptions metav1.ListOptions
}

// ReplicationControllerGetRequest represents request
type ReplicationControllerGetRequest struct {
	service_ v1.ReplicationControllerInterface
	Name     string
	metav1.GetOptions
}

// ReplicationControllerListRequest represents request
type ReplicationControllerListRequest struct {
	service_ v1.ReplicationControllerInterface
	metav1.ListOptions
}

// ReplicationControllerWatchRequest represents request
type ReplicationControllerWatchRequest struct {
	service_ v1.ReplicationControllerInterface
	metav1.ListOptions
}

// ReplicationControllerPatchRequest represents request
type ReplicationControllerPatchRequest struct {
	service_     v1.ReplicationControllerInterface
	Name         string
	Pt           types.PatchType
	Data         []byte
	Subresources []string
}

// ReplicationControllerGetScaleRequest represents request
type ReplicationControllerGetScaleRequest struct {
	service_                  v1.ReplicationControllerInterface
	ReplicationControllerName string
	metav1.GetOptions
}

// ReplicationControllerUpdateScaleRequest represents request
type ReplicationControllerUpdateScaleRequest struct {
	service_                  v1.ReplicationControllerInterface
	ReplicationControllerName string
	*autoscalingv1.Scale
}

func init() {
	register(&ReplicationControllerCreateRequest{})
	register(&ReplicationControllerUpdateRequest{})
	register(&ReplicationControllerUpdateStatusRequest{})
	register(&ReplicationControllerDeleteRequest{})
	register(&ReplicationControllerDeleteCollectionRequest{})
	register(&ReplicationControllerGetRequest{})
	register(&ReplicationControllerListRequest{})
	register(&ReplicationControllerWatchRequest{})
	register(&ReplicationControllerPatchRequest{})
	register(&ReplicationControllerGetScaleRequest{})
	register(&ReplicationControllerUpdateScaleRequest{})
}

func (r *ReplicationControllerCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.ReplicationController)
	return result, err
}

func (r *ReplicationControllerCreateRequest) GetId() string {
	return "v1.ReplicationController.Create"
}

func (r *ReplicationControllerUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.ReplicationController)
	return result, err
}

func (r *ReplicationControllerUpdateRequest) GetId() string {
	return "v1.ReplicationController.Update"
}

func (r *ReplicationControllerUpdateStatusRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerUpdateStatusRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.UpdateStatus(r.ReplicationController)
	return result, err
}

func (r *ReplicationControllerUpdateStatusRequest) GetId() string {
	return "v1.ReplicationController.UpdateStatus"
}

func (r *ReplicationControllerDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name, r.DeleteOptions)
	return result, err
}

func (r *ReplicationControllerDeleteRequest) GetId() string {
	return "v1.ReplicationController.Delete"
}

func (r *ReplicationControllerDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions, r.ListOptions)
	return result, err
}

func (r *ReplicationControllerDeleteCollectionRequest) GetId() string {
	return "v1.ReplicationController.DeleteCollection"
}

func (r *ReplicationControllerGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name, r.GetOptions)
	return result, err
}

func (r *ReplicationControllerGetRequest) GetId() string {
	return "v1.ReplicationController.Get"
}

func (r *ReplicationControllerListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err
}

func (r *ReplicationControllerListRequest) GetId() string {
	return "v1.ReplicationController.List"
}

func (r *ReplicationControllerWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err
}

func (r *ReplicationControllerWatchRequest) GetId() string {
	return "v1.ReplicationController.Watch"
}

func (r *ReplicationControllerPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name, r.Pt, r.Data, r.Subresources...)
	return result, err
}

func (r *ReplicationControllerPatchRequest) GetId() string {
	return "v1.ReplicationController.Patch"
}

func (r *ReplicationControllerGetScaleRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerGetScaleRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.GetScale(r.ReplicationControllerName, r.GetOptions)
	return result, err
}

func (r *ReplicationControllerGetScaleRequest) GetId() string {
	return "v1.ReplicationController.GetScale"
}

func (r *ReplicationControllerUpdateScaleRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ReplicationControllerInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ReplicationControllerInterface", service)
	}
	return nil
}

func (r *ReplicationControllerUpdateScaleRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.UpdateScale(r.ReplicationControllerName, r.Scale)
	return result, err
}

func (r *ReplicationControllerUpdateScaleRequest) GetId() string {
	return "v1.ReplicationController.UpdateScale"
}
