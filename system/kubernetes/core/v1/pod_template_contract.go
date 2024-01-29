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

// PodTemplateCreateRequest represents request
type PodTemplateCreateRequest struct {
	service_ v1.PodTemplateInterface
	*vvc.PodTemplate
}

// PodTemplateUpdateRequest represents request
type PodTemplateUpdateRequest struct {
	service_ v1.PodTemplateInterface
	*vvc.PodTemplate
}

// PodTemplateDeleteRequest represents request
type PodTemplateDeleteRequest struct {
	service_ v1.PodTemplateInterface
	Name     string
	*metav1.DeleteOptions
}

// PodTemplateDeleteCollectionRequest represents request
type PodTemplateDeleteCollectionRequest struct {
	service_ v1.PodTemplateInterface
	*metav1.DeleteOptions
	ListOptions metav1.ListOptions
}

// PodTemplateGetRequest represents request
type PodTemplateGetRequest struct {
	service_ v1.PodTemplateInterface
	Name     string
	metav1.GetOptions
}

// PodTemplateListRequest represents request
type PodTemplateListRequest struct {
	service_ v1.PodTemplateInterface
	metav1.ListOptions
}

// PodTemplateWatchRequest represents request
type PodTemplateWatchRequest struct {
	service_ v1.PodTemplateInterface
	metav1.ListOptions
}

// PodTemplatePatchRequest represents request
type PodTemplatePatchRequest struct {
	service_     v1.PodTemplateInterface
	Name         string
	Pt           types.PatchType
	Data         []byte
	Subresources []string
}

func init() {
	register(&PodTemplateCreateRequest{})
	register(&PodTemplateUpdateRequest{})
	register(&PodTemplateDeleteRequest{})
	register(&PodTemplateDeleteCollectionRequest{})
	register(&PodTemplateGetRequest{})
	register(&PodTemplateListRequest{})
	register(&PodTemplateWatchRequest{})
	register(&PodTemplatePatchRequest{})
}

func (r *PodTemplateCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodTemplateInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodTemplateInterface", service)
	}
	return nil
}

func (r *PodTemplateCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.PodTemplate)
	return result, err
}

func (r *PodTemplateCreateRequest) GetId() string {
	return "v1.PodTemplate.Create"
}

func (r *PodTemplateUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodTemplateInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodTemplateInterface", service)
	}
	return nil
}

func (r *PodTemplateUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.PodTemplate)
	return result, err
}

func (r *PodTemplateUpdateRequest) GetId() string {
	return "v1.PodTemplate.Update"
}

func (r *PodTemplateDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodTemplateInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodTemplateInterface", service)
	}
	return nil
}

func (r *PodTemplateDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name, r.DeleteOptions)
	return result, err
}

func (r *PodTemplateDeleteRequest) GetId() string {
	return "v1.PodTemplate.Delete"
}

func (r *PodTemplateDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodTemplateInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodTemplateInterface", service)
	}
	return nil
}

func (r *PodTemplateDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.DeleteOptions, r.ListOptions)
	return result, err
}

func (r *PodTemplateDeleteCollectionRequest) GetId() string {
	return "v1.PodTemplate.DeleteCollection"
}

func (r *PodTemplateGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodTemplateInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodTemplateInterface", service)
	}
	return nil
}

func (r *PodTemplateGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name, r.GetOptions)
	return result, err
}

func (r *PodTemplateGetRequest) GetId() string {
	return "v1.PodTemplate.Get"
}

func (r *PodTemplateListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodTemplateInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodTemplateInterface", service)
	}
	return nil
}

func (r *PodTemplateListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.ListOptions)
	return result, err
}

func (r *PodTemplateListRequest) GetId() string {
	return "v1.PodTemplate.List"
}

func (r *PodTemplateWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodTemplateInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodTemplateInterface", service)
	}
	return nil
}

func (r *PodTemplateWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.ListOptions)
	return result, err
}

func (r *PodTemplateWatchRequest) GetId() string {
	return "v1.PodTemplate.Watch"
}

func (r *PodTemplatePatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.PodTemplateInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.PodTemplateInterface", service)
	}
	return nil
}

func (r *PodTemplatePatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name, r.Pt, r.Data, r.Subresources...)
	return result, err
}

func (r *PodTemplatePatchRequest) GetId() string {
	return "v1.PodTemplate.Patch"
}
