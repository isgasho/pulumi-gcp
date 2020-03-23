# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ProjectSink(pulumi.CustomResource):
    bigquery_options: pulumi.Output[dict]
    """
    Options that affect sinks exporting data to BigQuery. Structure documented below.

      * `usePartitionedTables` (`bool`)
    """
    destination: pulumi.Output[str]
    """
    The destination of the sink (or, in other words, where logs are written to). Can be a
    Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
    The writer associated with the sink must have access to write to the above resource.
    """
    filter: pulumi.Output[str]
    """
    The filter to apply when exporting logs. Only log entries that match the filter are exported.
    See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
    write a filter.
    """
    name: pulumi.Output[str]
    """
    The name of the logging sink.
    """
    project: pulumi.Output[str]
    """
    The ID of the project to create the sink in. If omitted, the project associated with the provider is
    used.
    """
    unique_writer_identity: pulumi.Output[bool]
    """
    Whether or not to create a unique identity associated with this sink. If `false`
    (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
    then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
    must set `unique_writer_identity` to true.
    """
    writer_identity: pulumi.Output[str]
    """
    The identity associated with this sink. This identity must be granted write access to the
    configured `destination`.
    """
    def __init__(__self__, resource_name, opts=None, bigquery_options=None, destination=None, filter=None, name=None, project=None, unique_writer_identity=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a project-level logging sink. For more information see
        [the official documentation](https://cloud.google.com/logging/docs/),
        [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs)
        and
        [API](https://cloud.google.com/logging/docs/reference/v2/rest/).

        > **Note:** You must have [granted the "Logs Configuration Writer"](https://cloud.google.com/logging/docs/access-control) IAM role (`roles/logging.configWriter`) to the credentials used with this provider.

        > **Note** You must [enable the Cloud Resource Manager API](https://console.cloud.google.com/apis/library/cloudresourcemanager.googleapis.com)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_project_sink.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[str] filter: The filter to apply when exporting logs. Only log entries that match the filter are exported.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging sink.
        :param pulumi.Input[str] project: The ID of the project to create the sink in. If omitted, the project associated with the provider is
               used.
        :param pulumi.Input[bool] unique_writer_identity: Whether or not to create a unique identity associated with this sink. If `false`
               (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
               then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
               must set `unique_writer_identity` to true.

        The **bigquery_options** object supports the following:

          * `usePartitionedTables` (`pulumi.Input[bool]`)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['bigquery_options'] = bigquery_options
            if destination is None:
                raise TypeError("Missing required property 'destination'")
            __props__['destination'] = destination
            __props__['filter'] = filter
            __props__['name'] = name
            __props__['project'] = project
            __props__['unique_writer_identity'] = unique_writer_identity
            __props__['writer_identity'] = None
        super(ProjectSink, __self__).__init__(
            'gcp:logging/projectSink:ProjectSink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bigquery_options=None, destination=None, filter=None, name=None, project=None, unique_writer_identity=None, writer_identity=None):
        """
        Get an existing ProjectSink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] bigquery_options: Options that affect sinks exporting data to BigQuery. Structure documented below.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[str] filter: The filter to apply when exporting logs. Only log entries that match the filter are exported.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging sink.
        :param pulumi.Input[str] project: The ID of the project to create the sink in. If omitted, the project associated with the provider is
               used.
        :param pulumi.Input[bool] unique_writer_identity: Whether or not to create a unique identity associated with this sink. If `false`
               (the default), then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
               then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
               must set `unique_writer_identity` to true.
        :param pulumi.Input[str] writer_identity: The identity associated with this sink. This identity must be granted write access to the
               configured `destination`.

        The **bigquery_options** object supports the following:

          * `usePartitionedTables` (`pulumi.Input[bool]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bigquery_options"] = bigquery_options
        __props__["destination"] = destination
        __props__["filter"] = filter
        __props__["name"] = name
        __props__["project"] = project
        __props__["unique_writer_identity"] = unique_writer_identity
        __props__["writer_identity"] = writer_identity
        return ProjectSink(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

