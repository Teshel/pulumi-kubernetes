// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Secret holds secret data of a certain type. The total bytes of the values in the Data field must be less than MaxSecretSize bytes.
//
// Note: While Pulumi automatically encrypts the 'data' and 'stringData'
// fields, this encryption only applies to Pulumi's context, including the state file,
// the Service, the CLI, etc. Kubernetes does not encrypt Secret resources by default,
// and the contents are visible to users with access to the Secret in Kubernetes using
// tools like 'kubectl'.
//
// For more information on securing Kubernetes Secrets, see the following links:
// https://kubernetes.io/docs/concepts/configuration/secret/#security-properties
// https://kubernetes.io/docs/concepts/configuration/secret/#risks
type Secret struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	Data pulumi.StringMapOutput `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable pulumi.BoolPtrOutput `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
	StringData pulumi.StringMapOutput `pulumi:"stringData"`
	// Used to facilitate programmatic handling of secret data.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		args = &SecretArgs{}
	}

	args.ApiVersion = pulumi.StringPtr("v1")
	args.Kind = pulumi.StringPtr("Secret")
	if args.Data != nil {
		args.Data = pulumi.ToSecret(args.Data).(pulumi.StringMapOutput)
	}
	if args.StringData != nil {
		args.StringData = pulumi.ToSecret(args.StringData).(pulumi.StringMapOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"data",
		"stringData",
	})
	opts = append(opts, secrets)
	var resource Secret
	err := ctx.RegisterResource("kubernetes:core/v1:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("kubernetes:core/v1:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	Data map[string]string `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable *bool `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
	StringData map[string]string `pulumi:"stringData"`
	// Used to facilitate programmatic handling of secret data.
	Type *string `pulumi:"type"`
}

type SecretState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	Data pulumi.StringMapInput
	// Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable pulumi.BoolPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
	StringData pulumi.StringMapInput
	// Used to facilitate programmatic handling of secret data.
	Type pulumi.StringPtrInput
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	Data map[string]string `pulumi:"data"`
	// Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable *bool `pulumi:"immutable"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
	StringData map[string]string `pulumi:"stringData"`
	// Used to facilitate programmatic handling of secret data.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or '.'. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4
	Data pulumi.StringMapInput
	// Immutable, if set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
	Immutable pulumi.BoolPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// stringData allows specifying non-binary secret data in string form. It is provided as a write-only input field for convenience. All keys and values are merged into the data field on write, overwriting any existing values. The stringData field is never output when reading from the API.
	StringData pulumi.StringMapInput
	// Used to facilitate programmatic handling of secret data.
	Type pulumi.StringPtrInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}

type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(ctx context.Context) SecretOutput
}

func (*Secret) ElementType() reflect.Type {
	return reflect.TypeOf((*Secret)(nil))
}

func (i *Secret) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i *Secret) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}

func (i *Secret) ToSecretPtrOutput() SecretPtrOutput {
	return i.ToSecretPtrOutputWithContext(context.Background())
}

func (i *Secret) ToSecretPtrOutputWithContext(ctx context.Context) SecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPtrOutput)
}

type SecretPtrInput interface {
	pulumi.Input

	ToSecretPtrOutput() SecretPtrOutput
	ToSecretPtrOutputWithContext(ctx context.Context) SecretPtrOutput
}

type secretPtrType SecretArgs

func (*secretPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil))
}

func (i *secretPtrType) ToSecretPtrOutput() SecretPtrOutput {
	return i.ToSecretPtrOutputWithContext(context.Background())
}

func (i *secretPtrType) ToSecretPtrOutputWithContext(ctx context.Context) SecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPtrOutput)
}

// SecretArrayInput is an input type that accepts SecretArray and SecretArrayOutput values.
// You can construct a concrete instance of `SecretArrayInput` via:
//
//          SecretArray{ SecretArgs{...} }
type SecretArrayInput interface {
	pulumi.Input

	ToSecretArrayOutput() SecretArrayOutput
	ToSecretArrayOutputWithContext(context.Context) SecretArrayOutput
}

type SecretArray []SecretInput

func (SecretArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Secret)(nil))
}

func (i SecretArray) ToSecretArrayOutput() SecretArrayOutput {
	return i.ToSecretArrayOutputWithContext(context.Background())
}

func (i SecretArray) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretArrayOutput)
}

// SecretMapInput is an input type that accepts SecretMap and SecretMapOutput values.
// You can construct a concrete instance of `SecretMapInput` via:
//
//          SecretMap{ "key": SecretArgs{...} }
type SecretMapInput interface {
	pulumi.Input

	ToSecretMapOutput() SecretMapOutput
	ToSecretMapOutputWithContext(context.Context) SecretMapOutput
}

type SecretMap map[string]SecretInput

func (SecretMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Secret)(nil))
}

func (i SecretMap) ToSecretMapOutput() SecretMapOutput {
	return i.ToSecretMapOutputWithContext(context.Background())
}

func (i SecretMap) ToSecretMapOutputWithContext(ctx context.Context) SecretMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretMapOutput)
}

type SecretOutput struct {
	*pulumi.OutputState
}

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Secret)(nil))
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

func (o SecretOutput) ToSecretPtrOutput() SecretPtrOutput {
	return o.ToSecretPtrOutputWithContext(context.Background())
}

func (o SecretOutput) ToSecretPtrOutputWithContext(ctx context.Context) SecretPtrOutput {
	return o.ApplyT(func(v Secret) *Secret {
		return &v
	}).(SecretPtrOutput)
}

type SecretPtrOutput struct {
	*pulumi.OutputState
}

func (SecretPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Secret)(nil))
}

func (o SecretPtrOutput) ToSecretPtrOutput() SecretPtrOutput {
	return o
}

func (o SecretPtrOutput) ToSecretPtrOutputWithContext(ctx context.Context) SecretPtrOutput {
	return o
}

type SecretArrayOutput struct{ *pulumi.OutputState }

func (SecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Secret)(nil))
}

func (o SecretArrayOutput) ToSecretArrayOutput() SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) Index(i pulumi.IntInput) SecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Secret {
		return vs[0].([]Secret)[vs[1].(int)]
	}).(SecretOutput)
}

type SecretMapOutput struct{ *pulumi.OutputState }

func (SecretMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Secret)(nil))
}

func (o SecretMapOutput) ToSecretMapOutput() SecretMapOutput {
	return o
}

func (o SecretMapOutput) ToSecretMapOutputWithContext(ctx context.Context) SecretMapOutput {
	return o
}

func (o SecretMapOutput) MapIndex(k pulumi.StringInput) SecretOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Secret {
		return vs[0].(map[string]Secret)[vs[1].(string)]
	}).(SecretOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretOutput{})
	pulumi.RegisterOutputType(SecretPtrOutput{})
	pulumi.RegisterOutputType(SecretArrayOutput{})
	pulumi.RegisterOutputType(SecretMapOutput{})
}
