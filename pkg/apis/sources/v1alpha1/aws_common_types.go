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

package v1alpha1

import "knative.dev/pkg/apis"

// AWSSecurityCredentials represents a set of AWS security credentials.
// See https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html
type AWSSecurityCredentials struct {
	AccessKeyID     ValueFromField `json:"accessKeyID"`
	SecretAccessKey ValueFromField `json:"secretAccessKey"`
}

// AWSEndpoint contains parameters which are used to override the destination
// of REST API calls to AWS services.
// It allows, for example, to target API-compatible alternatives to the public
// AWS cloud (Localstack, Minio, ElasticMQ, ...).
type AWSEndpoint struct {
	// URL of the endpoint.
	URL *apis.URL `json:"url,omitempty"`
}