/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/h3poteto/aws-global-accelerator-controller/pkg/apis/endpointgroupbinding/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EndpointGroupBindingLister helps list EndpointGroupBindings.
// All objects returned here must be treated as read-only.
type EndpointGroupBindingLister interface {
	// List lists all EndpointGroupBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointGroupBinding, err error)
	// EndpointGroupBindings returns an object that can list and get EndpointGroupBindings.
	EndpointGroupBindings(namespace string) EndpointGroupBindingNamespaceLister
	EndpointGroupBindingListerExpansion
}

// endpointGroupBindingLister implements the EndpointGroupBindingLister interface.
type endpointGroupBindingLister struct {
	indexer cache.Indexer
}

// NewEndpointGroupBindingLister returns a new EndpointGroupBindingLister.
func NewEndpointGroupBindingLister(indexer cache.Indexer) EndpointGroupBindingLister {
	return &endpointGroupBindingLister{indexer: indexer}
}

// List lists all EndpointGroupBindings in the indexer.
func (s *endpointGroupBindingLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointGroupBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointGroupBinding))
	})
	return ret, err
}

// EndpointGroupBindings returns an object that can list and get EndpointGroupBindings.
func (s *endpointGroupBindingLister) EndpointGroupBindings(namespace string) EndpointGroupBindingNamespaceLister {
	return endpointGroupBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointGroupBindingNamespaceLister helps list and get EndpointGroupBindings.
// All objects returned here must be treated as read-only.
type EndpointGroupBindingNamespaceLister interface {
	// List lists all EndpointGroupBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointGroupBinding, err error)
	// Get retrieves the EndpointGroupBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EndpointGroupBinding, error)
	EndpointGroupBindingNamespaceListerExpansion
}

// endpointGroupBindingNamespaceLister implements the EndpointGroupBindingNamespaceLister
// interface.
type endpointGroupBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EndpointGroupBindings in the indexer for a given namespace.
func (s endpointGroupBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointGroupBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointGroupBinding))
	})
	return ret, err
}

// Get retrieves the EndpointGroupBinding from the indexer for a given namespace and name.
func (s endpointGroupBindingNamespaceLister) Get(name string) (*v1alpha1.EndpointGroupBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("endpointgroupbinding"), name)
	}
	return obj.(*v1alpha1.EndpointGroupBinding), nil
}
