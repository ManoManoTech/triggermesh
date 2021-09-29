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

package googlecloudpubsubsource

import (
	"context"
	"time"

	"github.com/kelseyhightower/envconfig"

	"knative.dev/eventing/pkg/reconciler/source"
	k8sclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"

	"github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	informerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/informers/sources/v1alpha1/googlecloudpubsubsource"
	reconcilerv1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/injection/reconciler/sources/v1alpha1/googlecloudpubsubsource"
	"github.com/triggermesh/triggermesh/pkg/sources/client/gcloud/pubsub"
	"github.com/triggermesh/triggermesh/pkg/sources/reconciler/common"
)

// the resync period ensures we regularly re-check the state of Azure resources.
const informerResyncPeriod = time.Minute * 5

// NewController creates a Reconciler for the event source and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {

	typ := (*v1alpha1.GoogleCloudPubSubSource)(nil)
	app := common.ComponentName(typ)

	// Calling envconfig.Process() with a prefix appends that prefix
	// (uppercased) to the Go field name, e.g. MYSOURCE_IMAGE.
	adapterCfg := &adapterConfig{
		configs: source.WatchConfigurations(ctx, app, cmw, source.WithLogging, source.WithMetrics),
	}
	envconfig.MustProcess(app, adapterCfg)

	informer := informerv1alpha1.Get(ctx)

	r := &Reconciler{
		cg:         pubsub.NewClientGetter(k8sclient.Get(ctx).CoreV1().Secrets),
		srcLister:  informer.Lister().GoogleCloudPubSubSources,
		adapterCfg: adapterCfg,
	}
	impl := reconcilerv1alpha1.NewImpl(ctx, r)

	r.base = common.NewGenericDeploymentReconciler(
		ctx,
		typ.GetGroupVersionKind(),
		impl.Tracker,
		impl.EnqueueControllerOf,
	)

	informer.Informer().AddEventHandlerWithResyncPeriod(controller.HandleAll(impl.Enqueue), informerResyncPeriod)

	return impl
}
