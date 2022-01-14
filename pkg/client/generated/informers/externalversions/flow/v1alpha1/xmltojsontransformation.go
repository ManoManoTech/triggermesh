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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	flowv1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/flow/v1alpha1"
	internalclientset "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset"
	internalinterfaces "github.com/triggermesh/triggermesh/pkg/client/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/listers/flow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// XMLToJSONTransformationInformer provides access to a shared informer and lister for
// XMLToJSONTransformations.
type XMLToJSONTransformationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.XMLToJSONTransformationLister
}

type xMLToJSONTransformationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewXMLToJSONTransformationInformer constructs a new informer for XMLToJSONTransformation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewXMLToJSONTransformationInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredXMLToJSONTransformationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredXMLToJSONTransformationInformer constructs a new informer for XMLToJSONTransformation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredXMLToJSONTransformationInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowV1alpha1().XMLToJSONTransformations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowV1alpha1().XMLToJSONTransformations(namespace).Watch(context.TODO(), options)
			},
		},
		&flowv1alpha1.XMLToJSONTransformation{},
		resyncPeriod,
		indexers,
	)
}

func (f *xMLToJSONTransformationInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredXMLToJSONTransformationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *xMLToJSONTransformationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flowv1alpha1.XMLToJSONTransformation{}, f.defaultInformer)
}

func (f *xMLToJSONTransformationInformer) Lister() v1alpha1.XMLToJSONTransformationLister {
	return v1alpha1.NewXMLToJSONTransformationLister(f.Informer().GetIndexer())
}
