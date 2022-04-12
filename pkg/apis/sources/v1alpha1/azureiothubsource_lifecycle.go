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

package v1alpha1

import (
	"github.com/triggermesh/triggermesh/pkg/apis/sources"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/triggermesh/triggermesh/pkg/apis/common/v1alpha1"
)

// Supported event types
const (
	AzureIOTHubGenericEventType = "message"
)

// GetEventTypes returns the event types generated by the source.
func (s *AzureIOTHubSource) GetEventTypes() []string {
	return []string{
		AzureEventType(sources.AzureIOTHub, AzureIOTHubGenericEventType),
	}
}

// GetGroupVersionKind implements kmeta.OwnerRefable.
func (s *AzureIOTHubSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("AzureIOTHubSource")
}

// GetConditionSet implements duckv1.KRShaped.
func (s *AzureIOTHubSource) GetConditionSet() apis.ConditionSet {
	return v1alpha1.EventSenderConditionSet
}

// GetStatus implements duckv1.KRShaped.
func (s *AzureIOTHubSource) GetStatus() *duckv1.Status {
	return &s.Status.Status
}

// GetSink implements EventSender.
func (s *AzureIOTHubSource) GetSink() *duckv1.Destination {
	return &s.Spec.Sink
}

// GetStatusManager implements Reconcilable.
func (s *AzureIOTHubSource) GetStatusManager() *v1alpha1.StatusManager {
	return &v1alpha1.StatusManager{
		ConditionSet: s.GetConditionSet(),
		Status:       &s.Status,
	}
}

// AsEventSource implements EventSource.
func (s *AzureIOTHubSource) AsEventSource() string {
	return AzureIOTHubSourceName(s.Namespace, s.Name)
}

// AzureIOTHubSourceName returns a unique reference to the source suitable for use
// as as a CloudEvent source.
func AzureIOTHubSourceName(namespace, name string) string {
	return "io.triggermesh.azureiothubsource/" + namespace + "/" + name
}