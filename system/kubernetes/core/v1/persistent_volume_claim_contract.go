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

// PersistentVolumeClaimCreateRequest represents request
type PersistentVolumeClaimCreateRequest struct {
	service_ v1.PersistentVolumeClaimInterface
	*vvc.PersistentVolumeClaim
}

// PersistentVolumeClaimUpdateRequest represents request
type PersistentVolumeClaimUpdateRequest struct {
	service_ v1.PersistentVolumeClaimInterface
	*vvc.PersistentVolumeClaim
}

// PersistentVolumeClaimUpdateStatusRequest represents request
type PersistentVolumeClaimUpdateStatusRequest struct {
	service_ v1.PersistentVolumeClaimInterface
	*vvc.PersistentVolumeClaim
}

// PersistentVolumeClaimDeleteRequest represents request
type PersistentVolumeClaimDeleteRequest struct {
	service_ v1.PersistentVolumeClaimInterface
	Name     string
	*metav1.DeleteOptions
}

// PersistentVolumeClaimDeleteCollectionRequest represents request
type PersistentVolumeClaimDeleteCollectionRequest struct {
	service_ v1.PersistentVolumeClaimInterface
	*metav1.DeleteOptions
	ListOptions metav1.ListOptions
}

// PersistentVolumeClaimGetRequest represents request
type PersistentVolumeClaimGetRequest struct {
	service_ v1.PersistentVolumeClaimInterface
	Name     string
	metav1.GetOptions
}

// PersistentVolumeClaimListRequest represents request
type PersistentVolumeClaimListRequest struct {
	service_ v1.PersistentVolumeClaimInterface
	metav1.ListOptions
}

// PersistentVolumeClaimWatchRequest represents request
type PersistentVolumeClaimWatchRequest struct {
	service_ v1.PersistentVolumeClaimInterface
	metav1.ListOptions
}

// PersistentVolumeClaimPatchRequest represents request
type PersistentVolumeClaimPatchRequest struct {
	service_     v1.PersistentVolumeClaimInterface
	Name         string
	Pt           types.PatchType
	Data         []byte
	Subresources []string
}

func init() {
	register(&PersistentVolumeClaimCreateRequest{})
	register(&PersistentVolumeClaimUpdateRequest{})
	register(&PersistentVolumeClaimUpdateStatusRequest{})
	register(&PersistentVolumeClaimDeleteRequest{})
	register(&PersistentVolumeClaimDeleteCollectionRequest{})
	register(&PersistentVolumeClaimGetRequest{})
	register(&PersistentVolumeClaimListRequest{})
	register(&PersistentVolumeClaimWatchRequest{})
	register(&PersistentVolumeClaimPatchRequest{})
}

func (r *PersistentVolumeClaimCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.PersistentVolumeClaim)
	return result, err
}

func (r *PersistentVolumeClaimCreateRequest) GetId() string {
	return "v1.PersistentVolumeClaim.Create"
}

func (r *PersistentVolumeClaimUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.PersistentVolumeClaim)
	return result, err
}

func (r *PersistentVolumeClaimUpdateRequest) GetId() string {
	return "v1.PersistentVolumeClaim.Update"
}

func (r *PersistentVolumeClaimUpdateStatusRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimUpdateStatusRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.UpdateStatus(r.PersistentVolumeClaim)
	return result, err
}

func (r *PersistentVolumeClaimUpdateStatusRequest) GetId() string {
	return "v1.PersistentVolumeClaim.UpdateStatus"
}

func (r *PersistentVolumeClaimDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name, r.DeleteOptions)
	return result, err
}

func (r *PersistentVolumeClaimDeleteRequest) GetId() string {
	return "v1.PersistentVolumeClaim.Delete"
}

func (r *PersistentVolumeClaimDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions, r.ListOptions)
	return result, err
}

func (r *PersistentVolumeClaimDeleteCollectionRequest) GetId() string {
	return "v1.PersistentVolumeClaim.DeleteCollection"
}

func (r *PersistentVolumeClaimGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name, r.GetOptions)
	return result, err
}

func (r *PersistentVolumeClaimGetRequest) GetId() string {
	return "v1.PersistentVolumeClaim.Get"
}

func (r *PersistentVolumeClaimListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err
}

func (r *PersistentVolumeClaimListRequest) GetId() string {
	return "v1.PersistentVolumeClaim.List"
}

func (r *PersistentVolumeClaimWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err
}

func (r *PersistentVolumeClaimWatchRequest) GetId() string {
	return "v1.PersistentVolumeClaim.Watch"
}

func (r *PersistentVolumeClaimPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PersistentVolumeClaimInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PersistentVolumeClaimInterface", service)
	}
	return nil
}

func (r *PersistentVolumeClaimPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name, r.Pt, r.Data, r.Subresources...)
	return result, err
}

func (r *PersistentVolumeClaimPatchRequest) GetId() string {
	return "v1.PersistentVolumeClaim.Patch"
}
