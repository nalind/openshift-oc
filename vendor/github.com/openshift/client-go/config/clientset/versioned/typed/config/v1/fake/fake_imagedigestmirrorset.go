// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImageDigestMirrorSets implements ImageDigestMirrorSetInterface
type FakeImageDigestMirrorSets struct {
	Fake *FakeConfigV1
}

var imagedigestmirrorsetsResource = v1.SchemeGroupVersion.WithResource("imagedigestmirrorsets")

var imagedigestmirrorsetsKind = v1.SchemeGroupVersion.WithKind("ImageDigestMirrorSet")

// Get takes name of the imageDigestMirrorSet, and returns the corresponding imageDigestMirrorSet object, and an error if there is any.
func (c *FakeImageDigestMirrorSets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ImageDigestMirrorSet, err error) {
	emptyResult := &v1.ImageDigestMirrorSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(imagedigestmirrorsetsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ImageDigestMirrorSet), err
}

// List takes label and field selectors, and returns the list of ImageDigestMirrorSets that match those selectors.
func (c *FakeImageDigestMirrorSets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ImageDigestMirrorSetList, err error) {
	emptyResult := &v1.ImageDigestMirrorSetList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(imagedigestmirrorsetsResource, imagedigestmirrorsetsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ImageDigestMirrorSetList{ListMeta: obj.(*v1.ImageDigestMirrorSetList).ListMeta}
	for _, item := range obj.(*v1.ImageDigestMirrorSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageDigestMirrorSets.
func (c *FakeImageDigestMirrorSets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(imagedigestmirrorsetsResource, opts))
}

// Create takes the representation of a imageDigestMirrorSet and creates it.  Returns the server's representation of the imageDigestMirrorSet, and an error, if there is any.
func (c *FakeImageDigestMirrorSets) Create(ctx context.Context, imageDigestMirrorSet *v1.ImageDigestMirrorSet, opts metav1.CreateOptions) (result *v1.ImageDigestMirrorSet, err error) {
	emptyResult := &v1.ImageDigestMirrorSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(imagedigestmirrorsetsResource, imageDigestMirrorSet, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ImageDigestMirrorSet), err
}

// Update takes the representation of a imageDigestMirrorSet and updates it. Returns the server's representation of the imageDigestMirrorSet, and an error, if there is any.
func (c *FakeImageDigestMirrorSets) Update(ctx context.Context, imageDigestMirrorSet *v1.ImageDigestMirrorSet, opts metav1.UpdateOptions) (result *v1.ImageDigestMirrorSet, err error) {
	emptyResult := &v1.ImageDigestMirrorSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(imagedigestmirrorsetsResource, imageDigestMirrorSet, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ImageDigestMirrorSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageDigestMirrorSets) UpdateStatus(ctx context.Context, imageDigestMirrorSet *v1.ImageDigestMirrorSet, opts metav1.UpdateOptions) (result *v1.ImageDigestMirrorSet, err error) {
	emptyResult := &v1.ImageDigestMirrorSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(imagedigestmirrorsetsResource, "status", imageDigestMirrorSet, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ImageDigestMirrorSet), err
}

// Delete takes name of the imageDigestMirrorSet and deletes it. Returns an error if one occurs.
func (c *FakeImageDigestMirrorSets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(imagedigestmirrorsetsResource, name, opts), &v1.ImageDigestMirrorSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageDigestMirrorSets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(imagedigestmirrorsetsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ImageDigestMirrorSetList{})
	return err
}

// Patch applies the patch and returns the patched imageDigestMirrorSet.
func (c *FakeImageDigestMirrorSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ImageDigestMirrorSet, err error) {
	emptyResult := &v1.ImageDigestMirrorSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(imagedigestmirrorsetsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ImageDigestMirrorSet), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied imageDigestMirrorSet.
func (c *FakeImageDigestMirrorSets) Apply(ctx context.Context, imageDigestMirrorSet *configv1.ImageDigestMirrorSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ImageDigestMirrorSet, err error) {
	if imageDigestMirrorSet == nil {
		return nil, fmt.Errorf("imageDigestMirrorSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(imageDigestMirrorSet)
	if err != nil {
		return nil, err
	}
	name := imageDigestMirrorSet.Name
	if name == nil {
		return nil, fmt.Errorf("imageDigestMirrorSet.Name must be provided to Apply")
	}
	emptyResult := &v1.ImageDigestMirrorSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(imagedigestmirrorsetsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ImageDigestMirrorSet), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeImageDigestMirrorSets) ApplyStatus(ctx context.Context, imageDigestMirrorSet *configv1.ImageDigestMirrorSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ImageDigestMirrorSet, err error) {
	if imageDigestMirrorSet == nil {
		return nil, fmt.Errorf("imageDigestMirrorSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(imageDigestMirrorSet)
	if err != nil {
		return nil, err
	}
	name := imageDigestMirrorSet.Name
	if name == nil {
		return nil, fmt.Errorf("imageDigestMirrorSet.Name must be provided to Apply")
	}
	emptyResult := &v1.ImageDigestMirrorSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(imagedigestmirrorsetsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ImageDigestMirrorSet), err
}
