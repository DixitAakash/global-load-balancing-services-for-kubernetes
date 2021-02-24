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

// FakeGSLBHostRules implements GSLBHostRuleInterface
type FakeGSLBHostRules struct {
	Fake *FakeAmkoV1alpha1
	ns   string
}

var gslbhostrulesResource = schema.GroupVersionResource{Group: "amko.vmware.com", Version: "v1alpha1", Resource: "gslbhostrules"}

var gslbhostrulesKind = schema.GroupVersionKind{Group: "amko.vmware.com", Version: "v1alpha1", Kind: "GSLBHostRule"}

// Get takes name of the gSLBHostRule, and returns the corresponding gSLBHostRule object, and an error if there is any.
func (c *FakeGSLBHostRules) Get(name string, options v1.GetOptions) (result *v1alpha1.GSLBHostRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gslbhostrulesResource, c.ns, name), &v1alpha1.GSLBHostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GSLBHostRule), err
}

// List takes label and field selectors, and returns the list of GSLBHostRules that match those selectors.
func (c *FakeGSLBHostRules) List(opts v1.ListOptions) (result *v1alpha1.GSLBHostRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gslbhostrulesResource, gslbhostrulesKind, c.ns, opts), &v1alpha1.GSLBHostRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GSLBHostRuleList{ListMeta: obj.(*v1alpha1.GSLBHostRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.GSLBHostRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gSLBHostRules.
func (c *FakeGSLBHostRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gslbhostrulesResource, c.ns, opts))

}

// Create takes the representation of a gSLBHostRule and creates it.  Returns the server's representation of the gSLBHostRule, and an error, if there is any.
func (c *FakeGSLBHostRules) Create(gSLBHostRule *v1alpha1.GSLBHostRule) (result *v1alpha1.GSLBHostRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gslbhostrulesResource, c.ns, gSLBHostRule), &v1alpha1.GSLBHostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GSLBHostRule), err
}

// Update takes the representation of a gSLBHostRule and updates it. Returns the server's representation of the gSLBHostRule, and an error, if there is any.
func (c *FakeGSLBHostRules) Update(gSLBHostRule *v1alpha1.GSLBHostRule) (result *v1alpha1.GSLBHostRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gslbhostrulesResource, c.ns, gSLBHostRule), &v1alpha1.GSLBHostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GSLBHostRule), err
}

// Delete takes name of the gSLBHostRule and deletes it. Returns an error if one occurs.
func (c *FakeGSLBHostRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gslbhostrulesResource, c.ns, name), &v1alpha1.GSLBHostRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGSLBHostRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gslbhostrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GSLBHostRuleList{})
	return err
}

// Patch applies the patch and returns the patched gSLBHostRule.
func (c *FakeGSLBHostRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GSLBHostRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gslbhostrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GSLBHostRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GSLBHostRule), err
}
