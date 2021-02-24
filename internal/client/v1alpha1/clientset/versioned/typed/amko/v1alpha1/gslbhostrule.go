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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/vmware/global-load-balancing-services-for-kubernetes/internal/apis/amko/v1alpha1"
	scheme "github.com/vmware/global-load-balancing-services-for-kubernetes/internal/client/v1alpha1/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GSLBHostRulesGetter has a method to return a GSLBHostRuleInterface.
// A group's client should implement this interface.
type GSLBHostRulesGetter interface {
	GSLBHostRules(namespace string) GSLBHostRuleInterface
}

// GSLBHostRuleInterface has methods to work with GSLBHostRule resources.
type GSLBHostRuleInterface interface {
	Create(*v1alpha1.GSLBHostRule) (*v1alpha1.GSLBHostRule, error)
	Update(*v1alpha1.GSLBHostRule) (*v1alpha1.GSLBHostRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GSLBHostRule, error)
	List(opts v1.ListOptions) (*v1alpha1.GSLBHostRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GSLBHostRule, err error)
	GSLBHostRuleExpansion
}

// gSLBHostRules implements GSLBHostRuleInterface
type gSLBHostRules struct {
	client rest.Interface
	ns     string
}

// newGSLBHostRules returns a GSLBHostRules
func newGSLBHostRules(c *AmkoV1alpha1Client, namespace string) *gSLBHostRules {
	return &gSLBHostRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gSLBHostRule, and returns the corresponding gSLBHostRule object, and an error if there is any.
func (c *gSLBHostRules) Get(name string, options v1.GetOptions) (result *v1alpha1.GSLBHostRule, err error) {
	result = &v1alpha1.GSLBHostRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gslbhostrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GSLBHostRules that match those selectors.
func (c *gSLBHostRules) List(opts v1.ListOptions) (result *v1alpha1.GSLBHostRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GSLBHostRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gslbhostrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gSLBHostRules.
func (c *gSLBHostRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gslbhostrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gSLBHostRule and creates it.  Returns the server's representation of the gSLBHostRule, and an error, if there is any.
func (c *gSLBHostRules) Create(gSLBHostRule *v1alpha1.GSLBHostRule) (result *v1alpha1.GSLBHostRule, err error) {
	result = &v1alpha1.GSLBHostRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gslbhostrules").
		Body(gSLBHostRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gSLBHostRule and updates it. Returns the server's representation of the gSLBHostRule, and an error, if there is any.
func (c *gSLBHostRules) Update(gSLBHostRule *v1alpha1.GSLBHostRule) (result *v1alpha1.GSLBHostRule, err error) {
	result = &v1alpha1.GSLBHostRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gslbhostrules").
		Name(gSLBHostRule.Name).
		Body(gSLBHostRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the gSLBHostRule and deletes it. Returns an error if one occurs.
func (c *gSLBHostRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gslbhostrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gSLBHostRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gslbhostrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gSLBHostRule.
func (c *gSLBHostRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GSLBHostRule, err error) {
	result = &v1alpha1.GSLBHostRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gslbhostrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
