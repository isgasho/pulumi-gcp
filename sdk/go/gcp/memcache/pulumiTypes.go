// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package memcache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type InstanceMemcacheNode struct {
	Host   *string `pulumi:"host"`
	NodeId *string `pulumi:"nodeId"`
	Port   *int    `pulumi:"port"`
	State  *string `pulumi:"state"`
	Zone   *string `pulumi:"zone"`
}

// InstanceMemcacheNodeInput is an input type that accepts InstanceMemcacheNodeArgs and InstanceMemcacheNodeOutput values.
// You can construct a concrete instance of `InstanceMemcacheNodeInput` via:
//
// 		 InstanceMemcacheNodeArgs{...}
//
type InstanceMemcacheNodeInput interface {
	pulumi.Input

	ToInstanceMemcacheNodeOutput() InstanceMemcacheNodeOutput
	ToInstanceMemcacheNodeOutputWithContext(context.Context) InstanceMemcacheNodeOutput
}

type InstanceMemcacheNodeArgs struct {
	Host   pulumi.StringPtrInput `pulumi:"host"`
	NodeId pulumi.StringPtrInput `pulumi:"nodeId"`
	Port   pulumi.IntPtrInput    `pulumi:"port"`
	State  pulumi.StringPtrInput `pulumi:"state"`
	Zone   pulumi.StringPtrInput `pulumi:"zone"`
}

func (InstanceMemcacheNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMemcacheNode)(nil)).Elem()
}

func (i InstanceMemcacheNodeArgs) ToInstanceMemcacheNodeOutput() InstanceMemcacheNodeOutput {
	return i.ToInstanceMemcacheNodeOutputWithContext(context.Background())
}

func (i InstanceMemcacheNodeArgs) ToInstanceMemcacheNodeOutputWithContext(ctx context.Context) InstanceMemcacheNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemcacheNodeOutput)
}

// InstanceMemcacheNodeArrayInput is an input type that accepts InstanceMemcacheNodeArray and InstanceMemcacheNodeArrayOutput values.
// You can construct a concrete instance of `InstanceMemcacheNodeArrayInput` via:
//
// 		 InstanceMemcacheNodeArray{ InstanceMemcacheNodeArgs{...} }
//
type InstanceMemcacheNodeArrayInput interface {
	pulumi.Input

	ToInstanceMemcacheNodeArrayOutput() InstanceMemcacheNodeArrayOutput
	ToInstanceMemcacheNodeArrayOutputWithContext(context.Context) InstanceMemcacheNodeArrayOutput
}

type InstanceMemcacheNodeArray []InstanceMemcacheNodeInput

func (InstanceMemcacheNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceMemcacheNode)(nil)).Elem()
}

func (i InstanceMemcacheNodeArray) ToInstanceMemcacheNodeArrayOutput() InstanceMemcacheNodeArrayOutput {
	return i.ToInstanceMemcacheNodeArrayOutputWithContext(context.Background())
}

func (i InstanceMemcacheNodeArray) ToInstanceMemcacheNodeArrayOutputWithContext(ctx context.Context) InstanceMemcacheNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemcacheNodeArrayOutput)
}

type InstanceMemcacheNodeOutput struct{ *pulumi.OutputState }

func (InstanceMemcacheNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMemcacheNode)(nil)).Elem()
}

func (o InstanceMemcacheNodeOutput) ToInstanceMemcacheNodeOutput() InstanceMemcacheNodeOutput {
	return o
}

func (o InstanceMemcacheNodeOutput) ToInstanceMemcacheNodeOutputWithContext(ctx context.Context) InstanceMemcacheNodeOutput {
	return o
}

func (o InstanceMemcacheNodeOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceMemcacheNode) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o InstanceMemcacheNodeOutput) NodeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceMemcacheNode) *string { return v.NodeId }).(pulumi.StringPtrOutput)
}

func (o InstanceMemcacheNodeOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceMemcacheNode) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o InstanceMemcacheNodeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceMemcacheNode) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o InstanceMemcacheNodeOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceMemcacheNode) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

type InstanceMemcacheNodeArrayOutput struct{ *pulumi.OutputState }

func (InstanceMemcacheNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceMemcacheNode)(nil)).Elem()
}

func (o InstanceMemcacheNodeArrayOutput) ToInstanceMemcacheNodeArrayOutput() InstanceMemcacheNodeArrayOutput {
	return o
}

func (o InstanceMemcacheNodeArrayOutput) ToInstanceMemcacheNodeArrayOutputWithContext(ctx context.Context) InstanceMemcacheNodeArrayOutput {
	return o
}

func (o InstanceMemcacheNodeArrayOutput) Index(i pulumi.IntInput) InstanceMemcacheNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceMemcacheNode {
		return vs[0].([]InstanceMemcacheNode)[vs[1].(int)]
	}).(InstanceMemcacheNodeOutput)
}

type InstanceMemcacheParameters struct {
	// -
	// This is a unique ID associated with this set of parameters.
	Id *string `pulumi:"id"`
	// User-defined set of parameters to use in the memcache process.
	Params map[string]string `pulumi:"params"`
}

// InstanceMemcacheParametersInput is an input type that accepts InstanceMemcacheParametersArgs and InstanceMemcacheParametersOutput values.
// You can construct a concrete instance of `InstanceMemcacheParametersInput` via:
//
// 		 InstanceMemcacheParametersArgs{...}
//
type InstanceMemcacheParametersInput interface {
	pulumi.Input

	ToInstanceMemcacheParametersOutput() InstanceMemcacheParametersOutput
	ToInstanceMemcacheParametersOutputWithContext(context.Context) InstanceMemcacheParametersOutput
}

type InstanceMemcacheParametersArgs struct {
	// -
	// This is a unique ID associated with this set of parameters.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// User-defined set of parameters to use in the memcache process.
	Params pulumi.StringMapInput `pulumi:"params"`
}

func (InstanceMemcacheParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMemcacheParameters)(nil)).Elem()
}

func (i InstanceMemcacheParametersArgs) ToInstanceMemcacheParametersOutput() InstanceMemcacheParametersOutput {
	return i.ToInstanceMemcacheParametersOutputWithContext(context.Background())
}

func (i InstanceMemcacheParametersArgs) ToInstanceMemcacheParametersOutputWithContext(ctx context.Context) InstanceMemcacheParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemcacheParametersOutput)
}

func (i InstanceMemcacheParametersArgs) ToInstanceMemcacheParametersPtrOutput() InstanceMemcacheParametersPtrOutput {
	return i.ToInstanceMemcacheParametersPtrOutputWithContext(context.Background())
}

func (i InstanceMemcacheParametersArgs) ToInstanceMemcacheParametersPtrOutputWithContext(ctx context.Context) InstanceMemcacheParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemcacheParametersOutput).ToInstanceMemcacheParametersPtrOutputWithContext(ctx)
}

// InstanceMemcacheParametersPtrInput is an input type that accepts InstanceMemcacheParametersArgs, InstanceMemcacheParametersPtr and InstanceMemcacheParametersPtrOutput values.
// You can construct a concrete instance of `InstanceMemcacheParametersPtrInput` via:
//
// 		 InstanceMemcacheParametersArgs{...}
//
//  or:
//
// 		 nil
//
type InstanceMemcacheParametersPtrInput interface {
	pulumi.Input

	ToInstanceMemcacheParametersPtrOutput() InstanceMemcacheParametersPtrOutput
	ToInstanceMemcacheParametersPtrOutputWithContext(context.Context) InstanceMemcacheParametersPtrOutput
}

type instanceMemcacheParametersPtrType InstanceMemcacheParametersArgs

func InstanceMemcacheParametersPtr(v *InstanceMemcacheParametersArgs) InstanceMemcacheParametersPtrInput {
	return (*instanceMemcacheParametersPtrType)(v)
}

func (*instanceMemcacheParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceMemcacheParameters)(nil)).Elem()
}

func (i *instanceMemcacheParametersPtrType) ToInstanceMemcacheParametersPtrOutput() InstanceMemcacheParametersPtrOutput {
	return i.ToInstanceMemcacheParametersPtrOutputWithContext(context.Background())
}

func (i *instanceMemcacheParametersPtrType) ToInstanceMemcacheParametersPtrOutputWithContext(ctx context.Context) InstanceMemcacheParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMemcacheParametersPtrOutput)
}

type InstanceMemcacheParametersOutput struct{ *pulumi.OutputState }

func (InstanceMemcacheParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMemcacheParameters)(nil)).Elem()
}

func (o InstanceMemcacheParametersOutput) ToInstanceMemcacheParametersOutput() InstanceMemcacheParametersOutput {
	return o
}

func (o InstanceMemcacheParametersOutput) ToInstanceMemcacheParametersOutputWithContext(ctx context.Context) InstanceMemcacheParametersOutput {
	return o
}

func (o InstanceMemcacheParametersOutput) ToInstanceMemcacheParametersPtrOutput() InstanceMemcacheParametersPtrOutput {
	return o.ToInstanceMemcacheParametersPtrOutputWithContext(context.Background())
}

func (o InstanceMemcacheParametersOutput) ToInstanceMemcacheParametersPtrOutputWithContext(ctx context.Context) InstanceMemcacheParametersPtrOutput {
	return o.ApplyT(func(v InstanceMemcacheParameters) *InstanceMemcacheParameters {
		return &v
	}).(InstanceMemcacheParametersPtrOutput)
}

// -
// This is a unique ID associated with this set of parameters.
func (o InstanceMemcacheParametersOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceMemcacheParameters) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// User-defined set of parameters to use in the memcache process.
func (o InstanceMemcacheParametersOutput) Params() pulumi.StringMapOutput {
	return o.ApplyT(func(v InstanceMemcacheParameters) map[string]string { return v.Params }).(pulumi.StringMapOutput)
}

type InstanceMemcacheParametersPtrOutput struct{ *pulumi.OutputState }

func (InstanceMemcacheParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceMemcacheParameters)(nil)).Elem()
}

func (o InstanceMemcacheParametersPtrOutput) ToInstanceMemcacheParametersPtrOutput() InstanceMemcacheParametersPtrOutput {
	return o
}

func (o InstanceMemcacheParametersPtrOutput) ToInstanceMemcacheParametersPtrOutputWithContext(ctx context.Context) InstanceMemcacheParametersPtrOutput {
	return o
}

func (o InstanceMemcacheParametersPtrOutput) Elem() InstanceMemcacheParametersOutput {
	return o.ApplyT(func(v *InstanceMemcacheParameters) InstanceMemcacheParameters { return *v }).(InstanceMemcacheParametersOutput)
}

// -
// This is a unique ID associated with this set of parameters.
func (o InstanceMemcacheParametersPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceMemcacheParameters) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// User-defined set of parameters to use in the memcache process.
func (o InstanceMemcacheParametersPtrOutput) Params() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstanceMemcacheParameters) map[string]string {
		if v == nil {
			return nil
		}
		return v.Params
	}).(pulumi.StringMapOutput)
}

type InstanceNodeConfig struct {
	// Number of CPUs per node.
	CpuCount int `pulumi:"cpuCount"`
	// Memory size in Mebibytes for each memcache node.
	MemorySizeMb int `pulumi:"memorySizeMb"`
}

// InstanceNodeConfigInput is an input type that accepts InstanceNodeConfigArgs and InstanceNodeConfigOutput values.
// You can construct a concrete instance of `InstanceNodeConfigInput` via:
//
// 		 InstanceNodeConfigArgs{...}
//
type InstanceNodeConfigInput interface {
	pulumi.Input

	ToInstanceNodeConfigOutput() InstanceNodeConfigOutput
	ToInstanceNodeConfigOutputWithContext(context.Context) InstanceNodeConfigOutput
}

type InstanceNodeConfigArgs struct {
	// Number of CPUs per node.
	CpuCount pulumi.IntInput `pulumi:"cpuCount"`
	// Memory size in Mebibytes for each memcache node.
	MemorySizeMb pulumi.IntInput `pulumi:"memorySizeMb"`
}

func (InstanceNodeConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNodeConfig)(nil)).Elem()
}

func (i InstanceNodeConfigArgs) ToInstanceNodeConfigOutput() InstanceNodeConfigOutput {
	return i.ToInstanceNodeConfigOutputWithContext(context.Background())
}

func (i InstanceNodeConfigArgs) ToInstanceNodeConfigOutputWithContext(ctx context.Context) InstanceNodeConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNodeConfigOutput)
}

func (i InstanceNodeConfigArgs) ToInstanceNodeConfigPtrOutput() InstanceNodeConfigPtrOutput {
	return i.ToInstanceNodeConfigPtrOutputWithContext(context.Background())
}

func (i InstanceNodeConfigArgs) ToInstanceNodeConfigPtrOutputWithContext(ctx context.Context) InstanceNodeConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNodeConfigOutput).ToInstanceNodeConfigPtrOutputWithContext(ctx)
}

// InstanceNodeConfigPtrInput is an input type that accepts InstanceNodeConfigArgs, InstanceNodeConfigPtr and InstanceNodeConfigPtrOutput values.
// You can construct a concrete instance of `InstanceNodeConfigPtrInput` via:
//
// 		 InstanceNodeConfigArgs{...}
//
//  or:
//
// 		 nil
//
type InstanceNodeConfigPtrInput interface {
	pulumi.Input

	ToInstanceNodeConfigPtrOutput() InstanceNodeConfigPtrOutput
	ToInstanceNodeConfigPtrOutputWithContext(context.Context) InstanceNodeConfigPtrOutput
}

type instanceNodeConfigPtrType InstanceNodeConfigArgs

func InstanceNodeConfigPtr(v *InstanceNodeConfigArgs) InstanceNodeConfigPtrInput {
	return (*instanceNodeConfigPtrType)(v)
}

func (*instanceNodeConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceNodeConfig)(nil)).Elem()
}

func (i *instanceNodeConfigPtrType) ToInstanceNodeConfigPtrOutput() InstanceNodeConfigPtrOutput {
	return i.ToInstanceNodeConfigPtrOutputWithContext(context.Background())
}

func (i *instanceNodeConfigPtrType) ToInstanceNodeConfigPtrOutputWithContext(ctx context.Context) InstanceNodeConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNodeConfigPtrOutput)
}

type InstanceNodeConfigOutput struct{ *pulumi.OutputState }

func (InstanceNodeConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNodeConfig)(nil)).Elem()
}

func (o InstanceNodeConfigOutput) ToInstanceNodeConfigOutput() InstanceNodeConfigOutput {
	return o
}

func (o InstanceNodeConfigOutput) ToInstanceNodeConfigOutputWithContext(ctx context.Context) InstanceNodeConfigOutput {
	return o
}

func (o InstanceNodeConfigOutput) ToInstanceNodeConfigPtrOutput() InstanceNodeConfigPtrOutput {
	return o.ToInstanceNodeConfigPtrOutputWithContext(context.Background())
}

func (o InstanceNodeConfigOutput) ToInstanceNodeConfigPtrOutputWithContext(ctx context.Context) InstanceNodeConfigPtrOutput {
	return o.ApplyT(func(v InstanceNodeConfig) *InstanceNodeConfig {
		return &v
	}).(InstanceNodeConfigPtrOutput)
}

// Number of CPUs per node.
func (o InstanceNodeConfigOutput) CpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceNodeConfig) int { return v.CpuCount }).(pulumi.IntOutput)
}

// Memory size in Mebibytes for each memcache node.
func (o InstanceNodeConfigOutput) MemorySizeMb() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceNodeConfig) int { return v.MemorySizeMb }).(pulumi.IntOutput)
}

type InstanceNodeConfigPtrOutput struct{ *pulumi.OutputState }

func (InstanceNodeConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceNodeConfig)(nil)).Elem()
}

func (o InstanceNodeConfigPtrOutput) ToInstanceNodeConfigPtrOutput() InstanceNodeConfigPtrOutput {
	return o
}

func (o InstanceNodeConfigPtrOutput) ToInstanceNodeConfigPtrOutputWithContext(ctx context.Context) InstanceNodeConfigPtrOutput {
	return o
}

func (o InstanceNodeConfigPtrOutput) Elem() InstanceNodeConfigOutput {
	return o.ApplyT(func(v *InstanceNodeConfig) InstanceNodeConfig { return *v }).(InstanceNodeConfigOutput)
}

// Number of CPUs per node.
func (o InstanceNodeConfigPtrOutput) CpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InstanceNodeConfig) *int {
		if v == nil {
			return nil
		}
		return &v.CpuCount
	}).(pulumi.IntPtrOutput)
}

// Memory size in Mebibytes for each memcache node.
func (o InstanceNodeConfigPtrOutput) MemorySizeMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InstanceNodeConfig) *int {
		if v == nil {
			return nil
		}
		return &v.MemorySizeMb
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceMemcacheNodeOutput{})
	pulumi.RegisterOutputType(InstanceMemcacheNodeArrayOutput{})
	pulumi.RegisterOutputType(InstanceMemcacheParametersOutput{})
	pulumi.RegisterOutputType(InstanceMemcacheParametersPtrOutput{})
	pulumi.RegisterOutputType(InstanceNodeConfigOutput{})
	pulumi.RegisterOutputType(InstanceNodeConfigPtrOutput{})
}
