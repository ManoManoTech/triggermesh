/*
Copyright 2021 TriggerMesh Inc.

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
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/flow/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// XMLToJSONTransformationLister helps list XMLToJSONTransformations.
// All objects returned here must be treated as read-only.
type XMLToJSONTransformationLister interface {
	// List lists all XMLToJSONTransformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.XMLToJSONTransformation, err error)
	// XMLToJSONTransformations returns an object that can list and get XMLToJSONTransformations.
	XMLToJSONTransformations(namespace string) XMLToJSONTransformationNamespaceLister
	XMLToJSONTransformationListerExpansion
}

// xMLToJSONTransformationLister implements the XMLToJSONTransformationLister interface.
type xMLToJSONTransformationLister struct {
	indexer cache.Indexer
}

// NewXMLToJSONTransformationLister returns a new XMLToJSONTransformationLister.
func NewXMLToJSONTransformationLister(indexer cache.Indexer) XMLToJSONTransformationLister {
	return &xMLToJSONTransformationLister{indexer: indexer}
}

// List lists all XMLToJSONTransformations in the indexer.
func (s *xMLToJSONTransformationLister) List(selector labels.Selector) (ret []*v1alpha1.XMLToJSONTransformation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XMLToJSONTransformation))
	})
	return ret, err
}

// XMLToJSONTransformations returns an object that can list and get XMLToJSONTransformations.
func (s *xMLToJSONTransformationLister) XMLToJSONTransformations(namespace string) XMLToJSONTransformationNamespaceLister {
	return xMLToJSONTransformationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// XMLToJSONTransformationNamespaceLister helps list and get XMLToJSONTransformations.
// All objects returned here must be treated as read-only.
type XMLToJSONTransformationNamespaceLister interface {
	// List lists all XMLToJSONTransformations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.XMLToJSONTransformation, err error)
	// Get retrieves the XMLToJSONTransformation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.XMLToJSONTransformation, error)
	XMLToJSONTransformationNamespaceListerExpansion
}

// xMLToJSONTransformationNamespaceLister implements the XMLToJSONTransformationNamespaceLister
// interface.
type xMLToJSONTransformationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all XMLToJSONTransformations in the indexer for a given namespace.
func (s xMLToJSONTransformationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.XMLToJSONTransformation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XMLToJSONTransformation))
	})
	return ret, err
}

// Get retrieves the XMLToJSONTransformation from the indexer for a given namespace and name.
func (s xMLToJSONTransformationNamespaceLister) Get(name string) (*v1alpha1.XMLToJSONTransformation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("xmltojsontransformation"), name)
	}
	return obj.(*v1alpha1.XMLToJSONTransformation), nil
}
