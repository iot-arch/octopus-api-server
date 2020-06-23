/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/cnrancher/edge-api-server/pkg/apis/edgeapi.cattle.io/v1alpha1"
	"github.com/cnrancher/edge-api-server/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type EdgeapiV1alpha1Interface interface {
	RESTClient() rest.Interface
	CatalogsGetter
	DeviceTemplatesGetter
	DeviceTemplateRevisionsGetter
	SettingsGetter
}

// EdgeapiV1alpha1Client is used to interact with features provided by the edgeapi.cattle.io group.
type EdgeapiV1alpha1Client struct {
	restClient rest.Interface
}

func (c *EdgeapiV1alpha1Client) Catalogs(namespace string) CatalogInterface {
	return newCatalogs(c, namespace)
}

func (c *EdgeapiV1alpha1Client) DeviceTemplates(namespace string) DeviceTemplateInterface {
	return newDeviceTemplates(c, namespace)
}

func (c *EdgeapiV1alpha1Client) DeviceTemplateRevisions(namespace string) DeviceTemplateRevisionInterface {
	return newDeviceTemplateRevisions(c, namespace)
}

func (c *EdgeapiV1alpha1Client) Settings() SettingInterface {
	return newSettings(c)
}

// NewForConfig creates a new EdgeapiV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*EdgeapiV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &EdgeapiV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new EdgeapiV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *EdgeapiV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new EdgeapiV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *EdgeapiV1alpha1Client {
	return &EdgeapiV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *EdgeapiV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
