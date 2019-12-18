// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_project_sink.html.markdown.
type ProjectSink struct {
	s *pulumi.ResourceState
}

// NewProjectSink registers a new resource with the given unique name, arguments, and options.
func NewProjectSink(ctx *pulumi.Context,
	name string, args *ProjectSinkArgs, opts ...pulumi.ResourceOpt) (*ProjectSink, error) {
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["bigqueryOptions"] = nil
		inputs["destination"] = nil
		inputs["filter"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["uniqueWriterIdentity"] = nil
	} else {
		inputs["bigqueryOptions"] = args.BigqueryOptions
		inputs["destination"] = args.Destination
		inputs["filter"] = args.Filter
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["uniqueWriterIdentity"] = args.UniqueWriterIdentity
	}
	inputs["writerIdentity"] = nil
	s, err := ctx.RegisterResource("gcp:logging/projectSink:ProjectSink", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProjectSink{s: s}, nil
}

// GetProjectSink gets an existing ProjectSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectSink(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ProjectSinkState, opts ...pulumi.ResourceOpt) (*ProjectSink, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["bigqueryOptions"] = state.BigqueryOptions
		inputs["destination"] = state.Destination
		inputs["filter"] = state.Filter
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["uniqueWriterIdentity"] = state.UniqueWriterIdentity
		inputs["writerIdentity"] = state.WriterIdentity
	}
	s, err := ctx.ReadResource("gcp:logging/projectSink:ProjectSink", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ProjectSink{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ProjectSink) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ProjectSink) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Options that affect sinks exporting data to BigQuery. Structure documented below.
func (r *ProjectSink) BigqueryOptions() pulumi.Output {
	return r.s.State["bigqueryOptions"]
}

// The destination of the sink (or, in other words, where logs are written to). Can be a
// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
// The writer associated with the sink must have access to write to the above resource.
func (r *ProjectSink) Destination() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["destination"])
}

// The filter to apply when exporting logs. Only log entries that match the filter are exported.
// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
// write a filter.
func (r *ProjectSink) Filter() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["filter"])
}

// The name of the logging sink.
func (r *ProjectSink) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project to create the sink in. If omitted, the project associated with the provider is
// used.
func (r *ProjectSink) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Whether or not to create a unique identity associated with this sink. If `false`
// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
// must set `uniqueWriterIdentity` to true.
func (r *ProjectSink) UniqueWriterIdentity() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["uniqueWriterIdentity"])
}

// The identity associated with this sink. This identity must be granted write access to the
// configured `destination`.
func (r *ProjectSink) WriterIdentity() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["writerIdentity"])
}

// Input properties used for looking up and filtering ProjectSink resources.
type ProjectSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions interface{}
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
	// The writer associated with the sink must have access to write to the above resource.
	Destination interface{}
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter interface{}
	// The name of the logging sink.
	Name interface{}
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project interface{}
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity interface{}
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity interface{}
}

// The set of arguments for constructing a ProjectSink resource.
type ProjectSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions interface{}
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
	// The writer associated with the sink must have access to write to the above resource.
	Destination interface{}
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter interface{}
	// The name of the logging sink.
	Name interface{}
	// The ID of the project to create the sink in. If omitted, the project associated with the provider is
	// used.
	Project interface{}
	// Whether or not to create a unique identity associated with this sink. If `false`
	// (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
	// then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
	// must set `uniqueWriterIdentity` to true.
	UniqueWriterIdentity interface{}
}
