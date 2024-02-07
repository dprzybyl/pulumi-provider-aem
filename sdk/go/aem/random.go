// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aem

import (
	"context"
	"reflect"

	"errors"
	"github.com/dprzybyl/pulumi-provider-aem/sdk/go/aem/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Random struct {
	pulumi.CustomResourceState

	Length pulumi.IntOutput    `pulumi:"length"`
	Result pulumi.StringOutput `pulumi:"result"`
}

// NewRandom registers a new resource with the given unique name, arguments, and options.
func NewRandom(ctx *pulumi.Context,
	name string, args *RandomArgs, opts ...pulumi.ResourceOption) (*Random, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Random
	err := ctx.RegisterResource("aem:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandom gets an existing Random resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomState, opts ...pulumi.ResourceOption) (*Random, error) {
	var resource Random
	err := ctx.ReadResource("aem:index:Random", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Random resources.
type randomState struct {
}

type RandomState struct {
}

func (RandomState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomState)(nil)).Elem()
}

type randomArgs struct {
	Length int `pulumi:"length"`
}

// The set of arguments for constructing a Random resource.
type RandomArgs struct {
	Length pulumi.IntInput
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}

type RandomInput interface {
	pulumi.Input

	ToRandomOutput() RandomOutput
	ToRandomOutputWithContext(ctx context.Context) RandomOutput
}

func (*Random) ElementType() reflect.Type {
	return reflect.TypeOf((**Random)(nil)).Elem()
}

func (i *Random) ToRandomOutput() RandomOutput {
	return i.ToRandomOutputWithContext(context.Background())
}

func (i *Random) ToRandomOutputWithContext(ctx context.Context) RandomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomOutput)
}

type RandomOutput struct{ *pulumi.OutputState }

func (RandomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Random)(nil)).Elem()
}

func (o RandomOutput) ToRandomOutput() RandomOutput {
	return o
}

func (o RandomOutput) ToRandomOutputWithContext(ctx context.Context) RandomOutput {
	return o
}

func (o RandomOutput) Length() pulumi.IntOutput {
	return o.ApplyT(func(v *Random) pulumi.IntOutput { return v.Length }).(pulumi.IntOutput)
}

func (o RandomOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v *Random) pulumi.StringOutput { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomInput)(nil)).Elem(), &Random{})
	pulumi.RegisterOutputType(RandomOutput{})
}
