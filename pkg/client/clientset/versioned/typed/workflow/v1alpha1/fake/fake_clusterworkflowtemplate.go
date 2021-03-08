// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/argoproj/argo/v2/pkg/apis/workflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterWorkflowTemplates implements ClusterWorkflowTemplateInterface
type FakeClusterWorkflowTemplates struct {
	Fake *FakeArgoprojV1alpha1
}

var clusterworkflowtemplatesResource = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "clusterworkflowtemplates"}

var clusterworkflowtemplatesKind = schema.GroupVersionKind{Group: "argoproj.io", Version: "v1alpha1", Kind: "ClusterWorkflowTemplate"}

// Get takes name of the clusterWorkflowTemplate, and returns the corresponding clusterWorkflowTemplate object, and an error if there is any.
func (c *FakeClusterWorkflowTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterWorkflowTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterworkflowtemplatesResource, name), &v1alpha1.ClusterWorkflowTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkflowTemplate), err
}

// List takes label and field selectors, and returns the list of ClusterWorkflowTemplates that match those selectors.
func (c *FakeClusterWorkflowTemplates) List(opts v1.ListOptions) (result *v1alpha1.ClusterWorkflowTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterworkflowtemplatesResource, clusterworkflowtemplatesKind, opts), &v1alpha1.ClusterWorkflowTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterWorkflowTemplateList{ListMeta: obj.(*v1alpha1.ClusterWorkflowTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterWorkflowTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterWorkflowTemplates.
func (c *FakeClusterWorkflowTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterworkflowtemplatesResource, opts))
}

// Create takes the representation of a clusterWorkflowTemplate and creates it.  Returns the server's representation of the clusterWorkflowTemplate, and an error, if there is any.
func (c *FakeClusterWorkflowTemplates) Create(clusterWorkflowTemplate *v1alpha1.ClusterWorkflowTemplate) (result *v1alpha1.ClusterWorkflowTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterworkflowtemplatesResource, clusterWorkflowTemplate), &v1alpha1.ClusterWorkflowTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkflowTemplate), err
}

// Update takes the representation of a clusterWorkflowTemplate and updates it. Returns the server's representation of the clusterWorkflowTemplate, and an error, if there is any.
func (c *FakeClusterWorkflowTemplates) Update(clusterWorkflowTemplate *v1alpha1.ClusterWorkflowTemplate) (result *v1alpha1.ClusterWorkflowTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterworkflowtemplatesResource, clusterWorkflowTemplate), &v1alpha1.ClusterWorkflowTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkflowTemplate), err
}

// Delete takes name of the clusterWorkflowTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClusterWorkflowTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterworkflowtemplatesResource, name), &v1alpha1.ClusterWorkflowTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterWorkflowTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterworkflowtemplatesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterWorkflowTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clusterWorkflowTemplate.
func (c *FakeClusterWorkflowTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterWorkflowTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterworkflowtemplatesResource, name, pt, data, subresources...), &v1alpha1.ClusterWorkflowTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterWorkflowTemplate), err
}