# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Instance']


class Instance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 num_nodes: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An isolated set of Cloud Spanner resources on which databases can be
        hosted.

        To get more information about Instance, see:

        * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/spanner/)

        ## Example Usage
        ### Spanner Instance Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.spanner.Instance("example",
            config="regional-us-central1",
            display_name="Test Spanner Instance",
            labels={
                "foo": "bar",
            },
            num_nodes=2)
        ```
        ### Spanner Instance Multi Regional

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.spanner.Instance("example",
            config="nam-eur-asia1",
            display_name="Multi Regional Instance",
            labels={
                "foo": "bar",
            },
            num_nodes=2)
        ```

        ## Import

        Instance can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:spanner/instance:Instance default projects/{{project}}/instances/{{name}}
        ```

        ```sh
         $ pulumi import gcp:spanner/instance:Instance default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:spanner/instance:Instance default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The name of the instance's configuration (similar but not
               quite the same as a region) which defines defines the geographic placement and
               replication of your databases in this instance. It determines where your data
               is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
               In order to obtain a valid list please consult the
               [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
        :param pulumi.Input[str] display_name: The descriptive name for this instance as it appears in UIs. Must be
               unique per project and between 4 and 30 characters in length.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: A unique identifier for the instance, which cannot be changed after
               the instance is created. The name must be between 6 and 30 characters
               in length.
        :param pulumi.Input[int] num_nodes: The number of nodes allocated to this instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if config is None:
                raise TypeError("Missing required property 'config'")
            __props__['config'] = config
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['num_nodes'] = num_nodes
            __props__['project'] = project
            __props__['state'] = None
        super(Instance, __self__).__init__(
            'gcp:spanner/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            num_nodes: Optional[pulumi.Input[int]] = None,
            project: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config: The name of the instance's configuration (similar but not
               quite the same as a region) which defines defines the geographic placement and
               replication of your databases in this instance. It determines where your data
               is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
               In order to obtain a valid list please consult the
               [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
        :param pulumi.Input[str] display_name: The descriptive name for this instance as it appears in UIs. Must be
               unique per project and between 4 and 30 characters in length.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: A unique identifier for the instance, which cannot be changed after
               the instance is created. The name must be between 6 and 30 characters
               in length.
        :param pulumi.Input[int] num_nodes: The number of nodes allocated to this instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: Instance status: 'CREATING' or 'READY'.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["config"] = config
        __props__["display_name"] = display_name
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["num_nodes"] = num_nodes
        __props__["project"] = project
        __props__["state"] = state
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        The name of the instance's configuration (similar but not
        quite the same as a region) which defines defines the geographic placement and
        replication of your databases in this instance. It determines where your data
        is stored. Values are typically of the form `regional-europe-west1` , `us-central` etc.
        In order to obtain a valid list please consult the
        [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The descriptive name for this instance as it appears in UIs. Must be
        unique per project and between 4 and 30 characters in length.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique identifier for the instance, which cannot be changed after
        the instance is created. The name must be between 6 and 30 characters
        in length.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numNodes")
    def num_nodes(self) -> pulumi.Output[Optional[int]]:
        """
        The number of nodes allocated to this instance.
        """
        return pulumi.get(self, "num_nodes")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Instance status: 'CREATING' or 'READY'.
        """
        return pulumi.get(self, "state")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

