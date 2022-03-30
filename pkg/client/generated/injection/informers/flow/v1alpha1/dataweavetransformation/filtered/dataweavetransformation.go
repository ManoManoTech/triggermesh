/*
Copyright 2022 TriggerMesh Inc.

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

// Code generated by injection-gen. DO NOT EDIT.

package filtered

import (
	context "context"

	apisflowv1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/flow/v1alpha1"
	internalclientset "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset"
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/informers/externalversions/flow/v1alpha1"
	client "github.com/triggermesh/triggermesh/pkg/client/generated/injection/client"
	filtered "github.com/triggermesh/triggermesh/pkg/client/generated/injection/informers/factory/filtered"
	flowv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/listers/flow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	cache "k8s.io/client-go/tools/cache"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterFilteredInformers(withInformer)
	injection.Dynamic.RegisterDynamicInformer(withDynamicInformer)
}

// Key is used for associating the Informer inside the context.Context.
type Key struct {
	Selector string
}

func withInformer(ctx context.Context) (context.Context, []controller.Informer) {
	untyped := ctx.Value(filtered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	infs := []controller.Informer{}
	for _, selector := range labelSelectors {
		f := filtered.Get(ctx, selector)
		inf := f.Flow().V1alpha1().DataWeaveTransformations()
		ctx = context.WithValue(ctx, Key{Selector: selector}, inf)
		infs = append(infs, inf.Informer())
	}
	return ctx, infs
}

func withDynamicInformer(ctx context.Context) context.Context {
	untyped := ctx.Value(filtered.LabelKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch labelkey from context.")
	}
	labelSelectors := untyped.([]string)
	for _, selector := range labelSelectors {
		inf := &wrapper{client: client.Get(ctx), selector: selector}
		ctx = context.WithValue(ctx, Key{Selector: selector}, inf)
	}
	return ctx
}

// Get extracts the typed informer from the context.
func Get(ctx context.Context, selector string) v1alpha1.DataWeaveTransformationInformer {
	untyped := ctx.Value(Key{Selector: selector})
	if untyped == nil {
		logging.FromContext(ctx).Panicf(
			"Unable to fetch github.com/triggermesh/triggermesh/pkg/client/generated/informers/externalversions/flow/v1alpha1.DataWeaveTransformationInformer with selector %s from context.", selector)
	}
	return untyped.(v1alpha1.DataWeaveTransformationInformer)
}

type wrapper struct {
	client internalclientset.Interface

	namespace string

	selector string
}

var _ v1alpha1.DataWeaveTransformationInformer = (*wrapper)(nil)
var _ flowv1alpha1.DataWeaveTransformationLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apisflowv1alpha1.DataWeaveTransformation{}, 0, nil)
}

func (w *wrapper) Lister() flowv1alpha1.DataWeaveTransformationLister {
	return w
}

func (w *wrapper) DataWeaveTransformations(namespace string) flowv1alpha1.DataWeaveTransformationNamespaceLister {
	return &wrapper{client: w.client, namespace: namespace, selector: w.selector}
}

func (w *wrapper) List(selector labels.Selector) (ret []*apisflowv1alpha1.DataWeaveTransformation, err error) {
	reqs, err := labels.ParseToRequirements(w.selector)
	if err != nil {
		return nil, err
	}
	selector = selector.Add(reqs...)
	lo, err := w.client.FlowV1alpha1().DataWeaveTransformations(w.namespace).List(context.TODO(), v1.ListOptions{
		LabelSelector: selector.String(),
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
	if err != nil {
		return nil, err
	}
	for idx := range lo.Items {
		ret = append(ret, &lo.Items[idx])
	}
	return ret, nil
}

func (w *wrapper) Get(name string) (*apisflowv1alpha1.DataWeaveTransformation, error) {
	// TODO(mattmoor): Check that the fetched object matches the selector.
	return w.client.FlowV1alpha1().DataWeaveTransformations(w.namespace).Get(context.TODO(), name, v1.GetOptions{
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
}
