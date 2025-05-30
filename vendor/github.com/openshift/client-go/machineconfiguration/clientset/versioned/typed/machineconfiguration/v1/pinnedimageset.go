// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	machineconfigurationv1 "github.com/openshift/api/machineconfiguration/v1"
	applyconfigurationsmachineconfigurationv1 "github.com/openshift/client-go/machineconfiguration/applyconfigurations/machineconfiguration/v1"
	scheme "github.com/openshift/client-go/machineconfiguration/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// PinnedImageSetsGetter has a method to return a PinnedImageSetInterface.
// A group's client should implement this interface.
type PinnedImageSetsGetter interface {
	PinnedImageSets() PinnedImageSetInterface
}

// PinnedImageSetInterface has methods to work with PinnedImageSet resources.
type PinnedImageSetInterface interface {
	Create(ctx context.Context, pinnedImageSet *machineconfigurationv1.PinnedImageSet, opts metav1.CreateOptions) (*machineconfigurationv1.PinnedImageSet, error)
	Update(ctx context.Context, pinnedImageSet *machineconfigurationv1.PinnedImageSet, opts metav1.UpdateOptions) (*machineconfigurationv1.PinnedImageSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*machineconfigurationv1.PinnedImageSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*machineconfigurationv1.PinnedImageSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *machineconfigurationv1.PinnedImageSet, err error)
	Apply(ctx context.Context, pinnedImageSet *applyconfigurationsmachineconfigurationv1.PinnedImageSetApplyConfiguration, opts metav1.ApplyOptions) (result *machineconfigurationv1.PinnedImageSet, err error)
	PinnedImageSetExpansion
}

// pinnedImageSets implements PinnedImageSetInterface
type pinnedImageSets struct {
	*gentype.ClientWithListAndApply[*machineconfigurationv1.PinnedImageSet, *machineconfigurationv1.PinnedImageSetList, *applyconfigurationsmachineconfigurationv1.PinnedImageSetApplyConfiguration]
}

// newPinnedImageSets returns a PinnedImageSets
func newPinnedImageSets(c *MachineconfigurationV1Client) *pinnedImageSets {
	return &pinnedImageSets{
		gentype.NewClientWithListAndApply[*machineconfigurationv1.PinnedImageSet, *machineconfigurationv1.PinnedImageSetList, *applyconfigurationsmachineconfigurationv1.PinnedImageSetApplyConfiguration](
			"pinnedimagesets",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *machineconfigurationv1.PinnedImageSet { return &machineconfigurationv1.PinnedImageSet{} },
			func() *machineconfigurationv1.PinnedImageSetList { return &machineconfigurationv1.PinnedImageSetList{} },
		),
	}
}
