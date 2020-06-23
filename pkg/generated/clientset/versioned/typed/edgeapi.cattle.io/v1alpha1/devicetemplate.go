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
	"context"
	"time"

	v1alpha1 "github.com/cnrancher/edge-api-server/pkg/apis/edgeapi.cattle.io/v1alpha1"
	scheme "github.com/cnrancher/edge-api-server/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeviceTemplatesGetter has a method to return a DeviceTemplateInterface.
// A group's client should implement this interface.
type DeviceTemplatesGetter interface {
	DeviceTemplates(namespace string) DeviceTemplateInterface
}

// DeviceTemplateInterface has methods to work with DeviceTemplate resources.
type DeviceTemplateInterface interface {
	Create(ctx context.Context, deviceTemplate *v1alpha1.DeviceTemplate, opts v1.CreateOptions) (*v1alpha1.DeviceTemplate, error)
	Update(ctx context.Context, deviceTemplate *v1alpha1.DeviceTemplate, opts v1.UpdateOptions) (*v1alpha1.DeviceTemplate, error)
	UpdateStatus(ctx context.Context, deviceTemplate *v1alpha1.DeviceTemplate, opts v1.UpdateOptions) (*v1alpha1.DeviceTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DeviceTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DeviceTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeviceTemplate, err error)
	DeviceTemplateExpansion
}

// deviceTemplates implements DeviceTemplateInterface
type deviceTemplates struct {
	client rest.Interface
	ns     string
}

// newDeviceTemplates returns a DeviceTemplates
func newDeviceTemplates(c *EdgeapiV1alpha1Client, namespace string) *deviceTemplates {
	return &deviceTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deviceTemplate, and returns the corresponding deviceTemplate object, and an error if there is any.
func (c *deviceTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeviceTemplate, err error) {
	result = &v1alpha1.DeviceTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeviceTemplates that match those selectors.
func (c *deviceTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeviceTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DeviceTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deviceTemplates.
func (c *deviceTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devicetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deviceTemplate and creates it.  Returns the server's representation of the deviceTemplate, and an error, if there is any.
func (c *deviceTemplates) Create(ctx context.Context, deviceTemplate *v1alpha1.DeviceTemplate, opts v1.CreateOptions) (result *v1alpha1.DeviceTemplate, err error) {
	result = &v1alpha1.DeviceTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devicetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deviceTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deviceTemplate and updates it. Returns the server's representation of the deviceTemplate, and an error, if there is any.
func (c *deviceTemplates) Update(ctx context.Context, deviceTemplate *v1alpha1.DeviceTemplate, opts v1.UpdateOptions) (result *v1alpha1.DeviceTemplate, err error) {
	result = &v1alpha1.DeviceTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicetemplates").
		Name(deviceTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deviceTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deviceTemplates) UpdateStatus(ctx context.Context, deviceTemplate *v1alpha1.DeviceTemplate, opts v1.UpdateOptions) (result *v1alpha1.DeviceTemplate, err error) {
	result = &v1alpha1.DeviceTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicetemplates").
		Name(deviceTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deviceTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deviceTemplate and deletes it. Returns an error if one occurs.
func (c *deviceTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicetemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deviceTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicetemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deviceTemplate.
func (c *deviceTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeviceTemplate, err error) {
	result = &v1alpha1.DeviceTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devicetemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
