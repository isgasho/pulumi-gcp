# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Autoscalar']

warnings.warn("""gcp.compute.Autoscalar has been deprecated in favor of gcp.compute.Autoscaler""", DeprecationWarning)


class Autoscalar(pulumi.CustomResource):
    warnings.warn("""gcp.compute.Autoscalar has been deprecated in favor of gcp.compute.Autoscaler""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['AutoscalarAutoscalingPolicyArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents an Autoscaler resource.

        Autoscalers allow you to automatically scale virtual machine instances in
        managed instance groups according to an autoscaling policy that you
        define.

        To get more information about Autoscaler, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/autoscalers)
        * How-to Guides
            * [Autoscaling Groups of Instances](https://cloud.google.com/compute/docs/autoscaler/)

        ## Example Usage
        ### Autoscaler Single Instance

        ```python
        import pulumi
        import pulumi_gcp as gcp

        debian9 = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        default_instance_template = gcp.compute.InstanceTemplate("defaultInstanceTemplate",
            machine_type="e2-medium",
            can_ip_forward=False,
            tags=[
                "foo",
                "bar",
            ],
            disks=[gcp.compute.InstanceTemplateDiskArgs(
                source_image=debian9.id,
            )],
            network_interfaces=[gcp.compute.InstanceTemplateNetworkInterfaceArgs(
                network="default",
            )],
            metadata={
                "foo": "bar",
            },
            service_account=gcp.compute.InstanceTemplateServiceAccountArgs(
                scopes=[
                    "userinfo-email",
                    "compute-ro",
                    "storage-ro",
                ],
            ),
            opts=ResourceOptions(provider=google_beta))
        default_target_pool = gcp.compute.TargetPool("defaultTargetPool", opts=ResourceOptions(provider=google_beta))
        default_instance_group_manager = gcp.compute.InstanceGroupManager("defaultInstanceGroupManager",
            zone="us-central1-f",
            versions=[gcp.compute.InstanceGroupManagerVersionArgs(
                instance_template=default_instance_template.id,
                name="primary",
            )],
            target_pools=[default_target_pool.id],
            base_instance_name="autoscaler-sample",
            opts=ResourceOptions(provider=google_beta))
        default_autoscaler = gcp.compute.Autoscaler("defaultAutoscaler",
            zone="us-central1-f",
            target=default_instance_group_manager.id,
            autoscaling_policy=gcp.compute.AutoscalerAutoscalingPolicyArgs(
                max_replicas=5,
                min_replicas=1,
                cooldown_period=60,
                metrics=[gcp.compute.AutoscalerAutoscalingPolicyMetricArgs(
                    name="pubsub.googleapis.com/subscription/num_undelivered_messages",
                    filter="resource.type = pubsub_subscription AND resource.label.subscription_id = our-subscription",
                    single_instance_assignment=65535,
                )],
            ),
            opts=ResourceOptions(provider=google_beta))
        ```
        ### Autoscaler Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        debian9 = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        foobar_instance_template = gcp.compute.InstanceTemplate("foobarInstanceTemplate",
            machine_type="e2-medium",
            can_ip_forward=False,
            tags=[
                "foo",
                "bar",
            ],
            disks=[gcp.compute.InstanceTemplateDiskArgs(
                source_image=debian9.id,
            )],
            network_interfaces=[gcp.compute.InstanceTemplateNetworkInterfaceArgs(
                network="default",
            )],
            metadata={
                "foo": "bar",
            },
            service_account=gcp.compute.InstanceTemplateServiceAccountArgs(
                scopes=[
                    "userinfo-email",
                    "compute-ro",
                    "storage-ro",
                ],
            ))
        foobar_target_pool = gcp.compute.TargetPool("foobarTargetPool")
        foobar_instance_group_manager = gcp.compute.InstanceGroupManager("foobarInstanceGroupManager",
            zone="us-central1-f",
            versions=[gcp.compute.InstanceGroupManagerVersionArgs(
                instance_template=foobar_instance_template.id,
                name="primary",
            )],
            target_pools=[foobar_target_pool.id],
            base_instance_name="foobar")
        foobar_autoscaler = gcp.compute.Autoscaler("foobarAutoscaler",
            zone="us-central1-f",
            target=foobar_instance_group_manager.id,
            autoscaling_policy=gcp.compute.AutoscalerAutoscalingPolicyArgs(
                max_replicas=5,
                min_replicas=1,
                cooldown_period=60,
                cpu_utilization=gcp.compute.AutoscalerAutoscalingPolicyCpuUtilizationArgs(
                    target=0.5,
                ),
            ))
        ```

        ## Import

        Autoscaler can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/autoscalar:Autoscalar default projects/{{project}}/zones/{{zone}}/autoscalers/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/autoscalar:Autoscalar default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/autoscalar:Autoscalar default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/autoscalar:Autoscalar default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AutoscalarAutoscalingPolicyArgs']] autoscaling_policy: The configuration parameters for the autoscaling algorithm. You can
               define one or more of the policies for an autoscaler: cpuUtilization,
               customMetricUtilizations, and loadBalancingUtilization.
               If none of these are specified, the default will be to autoscale based
               on cpuUtilization to 0.6 or 60%.
               Structure is documented below.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: The identifier (type) of the Stackdriver Monitoring metric.
               The metric cannot have negative values.
               The metric must have a value type of INT64 or DOUBLE.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] target: Fraction of backend capacity utilization (set in HTTP(s) load
               balancing configuration) that autoscaler should maintain. Must
               be a positive float value. If not defined, the default is 0.8.
        :param pulumi.Input[str] zone: URL of the zone where the instance group resides.
        """
        pulumi.log.warn("Autoscalar is deprecated: gcp.compute.Autoscalar has been deprecated in favor of gcp.compute.Autoscaler")
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

            if autoscaling_policy is None:
                raise TypeError("Missing required property 'autoscaling_policy'")
            __props__['autoscaling_policy'] = autoscaling_policy
            __props__['description'] = description
            __props__['name'] = name
            __props__['project'] = project
            if target is None:
                raise TypeError("Missing required property 'target'")
            __props__['target'] = target
            __props__['zone'] = zone
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(Autoscalar, __self__).__init__(
            'gcp:compute/autoscalar:Autoscalar',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['AutoscalarAutoscalingPolicyArgs']]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Autoscalar':
        """
        Get an existing Autoscalar resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AutoscalarAutoscalingPolicyArgs']] autoscaling_policy: The configuration parameters for the autoscaling algorithm. You can
               define one or more of the policies for an autoscaler: cpuUtilization,
               customMetricUtilizations, and loadBalancingUtilization.
               If none of these are specified, the default will be to autoscale based
               on cpuUtilization to 0.6 or 60%.
               Structure is documented below.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: The identifier (type) of the Stackdriver Monitoring metric.
               The metric cannot have negative values.
               The metric must have a value type of INT64 or DOUBLE.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] target: Fraction of backend capacity utilization (set in HTTP(s) load
               balancing configuration) that autoscaler should maintain. Must
               be a positive float value. If not defined, the default is 0.8.
        :param pulumi.Input[str] zone: URL of the zone where the instance group resides.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["autoscaling_policy"] = autoscaling_policy
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["target"] = target
        __props__["zone"] = zone
        return Autoscalar(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> pulumi.Output['outputs.AutoscalarAutoscalingPolicy']:
        """
        The configuration parameters for the autoscaling algorithm. You can
        define one or more of the policies for an autoscaler: cpuUtilization,
        customMetricUtilizations, and loadBalancingUtilization.
        If none of these are specified, the default will be to autoscale based
        on cpuUtilization to 0.6 or 60%.
        Structure is documented below.
        """
        return pulumi.get(self, "autoscaling_policy")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The identifier (type) of the Stackdriver Monitoring metric.
        The metric cannot have negative values.
        The metric must have a value type of INT64 or DOUBLE.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[str]:
        """
        Fraction of backend capacity utilization (set in HTTP(s) load
        balancing configuration) that autoscaler should maintain. Must
        be a positive float value. If not defined, the default is 0.8.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        URL of the zone where the instance group resides.
        """
        return pulumi.get(self, "zone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

