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

// LimitRangeCreateRequest represents request
type LimitRangeCreateRequest struct {
	service_ v1.LimitRangeInterface
	*vvc.LimitRange
}

// LimitRangeUpdateRequest represents request
type LimitRangeUpdateRequest struct {
	service_ v1.LimitRangeInterface
	*vvc.LimitRange
}

// LimitRangeDeleteRequest represents request
type LimitRangeDeleteRequest struct {
	service_ v1.LimitRangeInterface
	Name     string
	*metav1.DeleteOptions
}

// LimitRangeDeleteCollectionRequest represents request
type LimitRangeDeleteCollectionRequest struct {
	service_ v1.LimitRangeInterface
	*metav1.DeleteOptions
	ListOptions metav1.ListOptions
}

// LimitRangeGetRequest represents request
type LimitRangeGetRequest struct {
	service_ v1.LimitRangeInterface
	Name     string
	metav1.GetOptions
}

// LimitRangeListRequest represents request
type LimitRangeListRequest struct {
	service_ v1.LimitRangeInterface
	metav1.ListOptions
}

// LimitRangeWatchRequest represents request
type LimitRangeWatchRequest struct {
	service_ v1.LimitRangeInterface
	metav1.ListOptions
}

// LimitRangePatchRequest represents request
type LimitRangePatchRequest struct {
	service_     v1.LimitRangeInterface
	Name         string
	Pt           types.PatchType
	Data         []byte
	Subresources []string
}

func init() {
	register(&LimitRangeCreateRequest{})
	register(&LimitRangeUpdateRequest{})
	register(&LimitRangeDeleteRequest{})
	register(&LimitRangeDeleteCollectionRequest{})
	register(&LimitRangeGetRequest{})
	register(&LimitRangeListRequest{})
	register(&LimitRangeWatchRequest{})
	register(&LimitRangePatchRequest{})
}

func (r *LimitRangeCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.LimitRangeInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.LimitRangeInterface", service)
	}
	return nil
}

func (r *LimitRangeCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.LimitRange)
	return result, err
}

func (r *LimitRangeCreateRequest) GetId() string {
	return "v1.LimitRange.Create"
}

func (r *LimitRangeUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.LimitRangeInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.LimitRangeInterface", service)
	}
	return nil
}

func (r *LimitRangeUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.LimitRange)
	return result, err
}

func (r *LimitRangeUpdateRequest) GetId() string {
	return "v1.LimitRange.Update"
}

func (r *LimitRangeDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.LimitRangeInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.LimitRangeInterface", service)
	}
	return nil
}

func (r *LimitRangeDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name, r.DeleteOptions)
	return result, err
}

func (r *LimitRangeDeleteRequest) GetId() string {
	return "v1.LimitRange.Delete"
}

func (r *LimitRangeDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.LimitRangeInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.LimitRangeInterface", service)
	}
	return nil
}

func (r *LimitRangeDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions, r.ListOptions)
	return result, err
}

func (r *LimitRangeDeleteCollectionRequest) GetId() string {
	return "v1.LimitRange.DeleteCollection"
}

func (r *LimitRangeGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.LimitRangeInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.LimitRangeInterface", service)
	}
	return nil
}

func (r *LimitRangeGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name, r.GetOptions)
	return result, err
}

func (r *LimitRangeGetRequest) GetId() string {
	return "v1.LimitRange.Get"
}

func (r *LimitRangeListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.LimitRangeInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.LimitRangeInterface", service)
	}
	return nil
}

func (r *LimitRangeListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err
}

func (r *LimitRangeListRequest) GetId() string {
	return "v1.LimitRange.List"
}

func (r *LimitRangeWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.LimitRangeInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.LimitRangeInterface", service)
	}
	return nil
}

func (r *LimitRangeWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err
}

func (r *LimitRangeWatchRequest) GetId() string {
	return "v1.LimitRange.Watch"
}

func (r *LimitRangePatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.LimitRangeInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.LimitRangeInterface", service)
	}
	return nil
}

func (r *LimitRangePatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name, r.Pt, r.Data, r.Subresources...)
	return result, err
}

func (r *LimitRangePatchRequest) GetId() string {
	return "v1.LimitRange.Patch"
}
