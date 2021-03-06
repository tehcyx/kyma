// +build !ignore_autogenerated

// Code generated by main. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Asset) DeepCopyInto(out *Asset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Asset.
func (in *Asset) DeepCopy() *Asset {
	if in == nil {
		return nil
	}
	out := new(Asset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Asset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetBucketRef) DeepCopyInto(out *AssetBucketRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetBucketRef.
func (in *AssetBucketRef) DeepCopy() *AssetBucketRef {
	if in == nil {
		return nil
	}
	out := new(AssetBucketRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetFile) DeepCopyInto(out *AssetFile) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetFile.
func (in *AssetFile) DeepCopy() *AssetFile {
	if in == nil {
		return nil
	}
	out := new(AssetFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetList) DeepCopyInto(out *AssetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Asset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetList.
func (in *AssetList) DeepCopy() *AssetList {
	if in == nil {
		return nil
	}
	out := new(AssetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetSource) DeepCopyInto(out *AssetSource) {
	*out = *in
	if in.ValidationWebhookService != nil {
		in, out := &in.ValidationWebhookService, &out.ValidationWebhookService
		*out = make([]AssetWebhookService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MutationWebhookService != nil {
		in, out := &in.MutationWebhookService, &out.MutationWebhookService
		*out = make([]AssetWebhookService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetadataWebhookService != nil {
		in, out := &in.MetadataWebhookService, &out.MetadataWebhookService
		*out = make([]WebhookService, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetSource.
func (in *AssetSource) DeepCopy() *AssetSource {
	if in == nil {
		return nil
	}
	out := new(AssetSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetSpec) DeepCopyInto(out *AssetSpec) {
	*out = *in
	in.CommonAssetSpec.DeepCopyInto(&out.CommonAssetSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetSpec.
func (in *AssetSpec) DeepCopy() *AssetSpec {
	if in == nil {
		return nil
	}
	out := new(AssetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetStatus) DeepCopyInto(out *AssetStatus) {
	*out = *in
	in.CommonAssetStatus.DeepCopyInto(&out.CommonAssetStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetStatus.
func (in *AssetStatus) DeepCopy() *AssetStatus {
	if in == nil {
		return nil
	}
	out := new(AssetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetStatusRef) DeepCopyInto(out *AssetStatusRef) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]AssetFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetStatusRef.
func (in *AssetStatusRef) DeepCopy() *AssetStatusRef {
	if in == nil {
		return nil
	}
	out := new(AssetStatusRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssetWebhookService) DeepCopyInto(out *AssetWebhookService) {
	*out = *in
	out.WebhookService = in.WebhookService
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssetWebhookService.
func (in *AssetWebhookService) DeepCopy() *AssetWebhookService {
	if in == nil {
		return nil
	}
	out := new(AssetWebhookService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	out.CommonBucketSpec = in.CommonBucketSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	in.CommonBucketStatus.DeepCopyInto(&out.CommonBucketStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAsset) DeepCopyInto(out *ClusterAsset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAsset.
func (in *ClusterAsset) DeepCopy() *ClusterAsset {
	if in == nil {
		return nil
	}
	out := new(ClusterAsset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAsset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAssetList) DeepCopyInto(out *ClusterAssetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAsset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAssetList.
func (in *ClusterAssetList) DeepCopy() *ClusterAssetList {
	if in == nil {
		return nil
	}
	out := new(ClusterAssetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAssetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAssetSpec) DeepCopyInto(out *ClusterAssetSpec) {
	*out = *in
	in.CommonAssetSpec.DeepCopyInto(&out.CommonAssetSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAssetSpec.
func (in *ClusterAssetSpec) DeepCopy() *ClusterAssetSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAssetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAssetStatus) DeepCopyInto(out *ClusterAssetStatus) {
	*out = *in
	in.CommonAssetStatus.DeepCopyInto(&out.CommonAssetStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAssetStatus.
func (in *ClusterAssetStatus) DeepCopy() *ClusterAssetStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAssetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBucket) DeepCopyInto(out *ClusterBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBucket.
func (in *ClusterBucket) DeepCopy() *ClusterBucket {
	if in == nil {
		return nil
	}
	out := new(ClusterBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBucketList) DeepCopyInto(out *ClusterBucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterBucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBucketList.
func (in *ClusterBucketList) DeepCopy() *ClusterBucketList {
	if in == nil {
		return nil
	}
	out := new(ClusterBucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterBucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBucketSpec) DeepCopyInto(out *ClusterBucketSpec) {
	*out = *in
	out.CommonBucketSpec = in.CommonBucketSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBucketSpec.
func (in *ClusterBucketSpec) DeepCopy() *ClusterBucketSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterBucketStatus) DeepCopyInto(out *ClusterBucketStatus) {
	*out = *in
	in.CommonBucketStatus.DeepCopyInto(&out.CommonBucketStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterBucketStatus.
func (in *ClusterBucketStatus) DeepCopy() *ClusterBucketStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterBucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonAssetSpec) DeepCopyInto(out *CommonAssetSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	out.BucketRef = in.BucketRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonAssetSpec.
func (in *CommonAssetSpec) DeepCopy() *CommonAssetSpec {
	if in == nil {
		return nil
	}
	out := new(CommonAssetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonAssetStatus) DeepCopyInto(out *CommonAssetStatus) {
	*out = *in
	in.AssetRef.DeepCopyInto(&out.AssetRef)
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonAssetStatus.
func (in *CommonAssetStatus) DeepCopy() *CommonAssetStatus {
	if in == nil {
		return nil
	}
	out := new(CommonAssetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonBucketSpec) DeepCopyInto(out *CommonBucketSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonBucketSpec.
func (in *CommonBucketSpec) DeepCopy() *CommonBucketSpec {
	if in == nil {
		return nil
	}
	out := new(CommonBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonBucketStatus) DeepCopyInto(out *CommonBucketStatus) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonBucketStatus.
func (in *CommonBucketStatus) DeepCopy() *CommonBucketStatus {
	if in == nil {
		return nil
	}
	out := new(CommonBucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookService) DeepCopyInto(out *WebhookService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookService.
func (in *WebhookService) DeepCopy() *WebhookService {
	if in == nil {
		return nil
	}
	out := new(WebhookService)
	in.DeepCopyInto(out)
	return out
}
