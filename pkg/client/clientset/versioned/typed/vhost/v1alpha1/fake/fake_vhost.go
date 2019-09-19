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
	v1alpha1 "github.com/opslib/crd-nginx-vhost/pkg/apis/vhost/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVhosts implements VhostInterface
type FakeVhosts struct {
	Fake *FakeVhostV1alpha1
	ns   string
}

var vhostsResource = schema.GroupVersionResource{Group: "vhost.newops.cn", Version: "v1alpha1", Resource: "vhosts"}

var vhostsKind = schema.GroupVersionKind{Group: "vhost.newops.cn", Version: "v1alpha1", Kind: "Vhost"}

// Get takes name of the vhost, and returns the corresponding vhost object, and an error if there is any.
func (c *FakeVhosts) Get(name string, options v1.GetOptions) (result *v1alpha1.Vhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vhostsResource, c.ns, name), &v1alpha1.Vhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vhost), err
}

// List takes label and field selectors, and returns the list of Vhosts that match those selectors.
func (c *FakeVhosts) List(opts v1.ListOptions) (result *v1alpha1.VhostList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vhostsResource, vhostsKind, c.ns, opts), &v1alpha1.VhostList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VhostList{ListMeta: obj.(*v1alpha1.VhostList).ListMeta}
	for _, item := range obj.(*v1alpha1.VhostList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vhosts.
func (c *FakeVhosts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vhostsResource, c.ns, opts))

}

// Create takes the representation of a vhost and creates it.  Returns the server's representation of the vhost, and an error, if there is any.
func (c *FakeVhosts) Create(vhost *v1alpha1.Vhost) (result *v1alpha1.Vhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vhostsResource, c.ns, vhost), &v1alpha1.Vhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vhost), err
}

// Update takes the representation of a vhost and updates it. Returns the server's representation of the vhost, and an error, if there is any.
func (c *FakeVhosts) Update(vhost *v1alpha1.Vhost) (result *v1alpha1.Vhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vhostsResource, c.ns, vhost), &v1alpha1.Vhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vhost), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVhosts) UpdateStatus(vhost *v1alpha1.Vhost) (*v1alpha1.Vhost, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vhostsResource, "status", c.ns, vhost), &v1alpha1.Vhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vhost), err
}

// Delete takes name of the vhost and deletes it. Returns an error if one occurs.
func (c *FakeVhosts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vhostsResource, c.ns, name), &v1alpha1.Vhost{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVhosts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vhostsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VhostList{})
	return err
}

// Patch applies the patch and returns the patched vhost.
func (c *FakeVhosts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Vhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vhostsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Vhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vhost), err
}