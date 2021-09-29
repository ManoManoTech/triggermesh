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

// Code generated by injection-gen. DO NOT EDIT.

package filtered

import (
	context "context"

	apistargetsv1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	internalclientset "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset"
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/informers/externalversions/targets/v1alpha1"
	client "github.com/triggermesh/triggermesh/pkg/client/generated/injection/client"
	filtered "github.com/triggermesh/triggermesh/pkg/client/generated/injection/informers/factory/filtered"
	targetsv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/listers/targets/v1alpha1"
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
		inf := f.Targets().V1alpha1().AlibabaOSSTargets()
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
func Get(ctx context.Context, selector string) v1alpha1.AlibabaOSSTargetInformer {
	untyped := ctx.Value(Key{Selector: selector})
	if untyped == nil {
		logging.FromContext(ctx).Panicf(
			"Unable to fetch github.com/triggermesh/triggermesh/pkg/client/generated/informers/externalversions/targets/v1alpha1.AlibabaOSSTargetInformer with selector %s from context.", selector)
	}
	return untyped.(v1alpha1.AlibabaOSSTargetInformer)
}

type wrapper struct {
	client internalclientset.Interface

	namespace string

	selector string
}

var _ v1alpha1.AlibabaOSSTargetInformer = (*wrapper)(nil)
var _ targetsv1alpha1.AlibabaOSSTargetLister = (*wrapper)(nil)

func (w *wrapper) Informer() cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(nil, &apistargetsv1alpha1.AlibabaOSSTarget{}, 0, nil)
}

func (w *wrapper) Lister() targetsv1alpha1.AlibabaOSSTargetLister {
	return w
}

func (w *wrapper) AlibabaOSSTargets(namespace string) targetsv1alpha1.AlibabaOSSTargetNamespaceLister {
	return &wrapper{client: w.client, namespace: namespace, selector: w.selector}
}

func (w *wrapper) List(selector labels.Selector) (ret []*apistargetsv1alpha1.AlibabaOSSTarget, err error) {
	reqs, err := labels.ParseToRequirements(w.selector)
	if err != nil {
		return nil, err
	}
	selector = selector.Add(reqs...)
	lo, err := w.client.TargetsV1alpha1().AlibabaOSSTargets(w.namespace).List(context.TODO(), v1.ListOptions{
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

func (w *wrapper) Get(name string) (*apistargetsv1alpha1.AlibabaOSSTarget, error) {
	// TODO(mattmoor): Check that the fetched object matches the selector.
	return w.client.TargetsV1alpha1().AlibabaOSSTargets(w.namespace).Get(context.TODO(), name, v1.GetOptions{
		// TODO(mattmoor): Incorporate resourceVersion bounds based on staleness criteria.
	})
}
