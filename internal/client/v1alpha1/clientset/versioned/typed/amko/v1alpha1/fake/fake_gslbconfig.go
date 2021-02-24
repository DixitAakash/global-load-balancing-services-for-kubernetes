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

package fake

import (
	v1alpha1 "github.com/vmware/global-load-balancing-services-for-kubernetes/internal/apis/amko/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGSLBConfigs implements GSLBConfigInterface
type FakeGSLBConfigs struct {
	Fake *FakeAmkoV1alpha1
	ns   string
}

var gslbconfigsResource = schema.GroupVersionResource{Group: "amko.vmware.com", Version: "v1alpha1", Resource: "gslbconfigs"}

var gslbconfigsKind = schema.GroupVersionKind{Group: "amko.vmware.com", Version: "v1alpha1", Kind: "GSLBConfig"}

// Get takes name of the gSLBConfig, and returns the corresponding gSLBConfig object, and an error if there is any.
func (c *FakeGSLBConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.GSLBConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gslbconfigsResource, c.ns, name), &v1alpha1.GSLBConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GSLBConfig), err
}

// List takes label and field selectors, and returns the list of GSLBConfigs that match those selectors.
func (c *FakeGSLBConfigs) List(opts v1.ListOptions) (result *v1alpha1.GSLBConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gslbconfigsResource, gslbconfigsKind, c.ns, opts), &v1alpha1.GSLBConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GSLBConfigList{ListMeta: obj.(*v1alpha1.GSLBConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.GSLBConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gSLBConfigs.
func (c *FakeGSLBConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gslbconfigsResource, c.ns, opts))

}

// Create takes the representation of a gSLBConfig and creates it.  Returns the server's representation of the gSLBConfig, and an error, if there is any.
func (c *FakeGSLBConfigs) Create(gSLBConfig *v1alpha1.GSLBConfig) (result *v1alpha1.GSLBConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gslbconfigsResource, c.ns, gSLBConfig), &v1alpha1.GSLBConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GSLBConfig), err
}

// Update takes the representation of a gSLBConfig and updates it. Returns the server's representation of the gSLBConfig, and an error, if there is any.
func (c *FakeGSLBConfigs) Update(gSLBConfig *v1alpha1.GSLBConfig) (result *v1alpha1.GSLBConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gslbconfigsResource, c.ns, gSLBConfig), &v1alpha1.GSLBConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GSLBConfig), err
}

// Delete takes name of the gSLBConfig and deletes it. Returns an error if one occurs.
func (c *FakeGSLBConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gslbconfigsResource, c.ns, name), &v1alpha1.GSLBConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGSLBConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gslbconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GSLBConfigList{})
	return err
}

// Patch applies the patch and returns the patched gSLBConfig.
func (c *FakeGSLBConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GSLBConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gslbconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GSLBConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GSLBConfig), err
}