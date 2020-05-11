# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Reservation(pulumi.CustomResource):
    ignore_idle_slots: pulumi.Output[bool]
    """
    If false, any query using this reservation will use idle slots from other reservations within
    the same admin project. If true, a query using this reservation will execute with the slot
    capacity specified above at most.
    """
    location: pulumi.Output[str]
    """
    The geographic location where the transfer config should reside.
    Examples: US, EU, asia-northeast1. The default value is US.
    """
    name: pulumi.Output[str]
    """
    The name of the reservation. This field must only contain alphanumeric characters or dash.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    slot_capacity: pulumi.Output[float]
    """
    Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
    unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
    """
    def __init__(__self__, resource_name, opts=None, ignore_idle_slots=None, location=None, name=None, project=None, slot_capacity=None, __props__=None, __name__=None, __opts__=None):
        """
        A reservation is a mechanism used to guarantee BigQuery slots to users.

        To get more information about Reservation, see:

        * [API documentation](https://cloud.google.com/bigquery/docs/reference/reservations/rest/v1beta1/projects.locations.reservations/create)
        * How-to Guides
            * [Introduction to Reservations](https://cloud.google.com/bigquery/docs/reservations-intro)

        ## Example Usage - Bigquery Reservation Basic


        ```python
        import pulumi
        import pulumi_gcp as gcp

        reservation = gcp.bigquery.Reservation("reservation",
            location="asia-northeast1",
            slot_capacity=0,
            ignore_idle_slots=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] ignore_idle_slots: If false, any query using this reservation will use idle slots from other reservations within
               the same admin project. If true, a query using this reservation will execute with the slot
               capacity specified above at most.
        :param pulumi.Input[str] location: The geographic location where the transfer config should reside.
               Examples: US, EU, asia-northeast1. The default value is US.
        :param pulumi.Input[str] name: The name of the reservation. This field must only contain alphanumeric characters or dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[float] slot_capacity: Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
               unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
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

            __props__['ignore_idle_slots'] = ignore_idle_slots
            __props__['location'] = location
            __props__['name'] = name
            __props__['project'] = project
            if slot_capacity is None:
                raise TypeError("Missing required property 'slot_capacity'")
            __props__['slot_capacity'] = slot_capacity
        super(Reservation, __self__).__init__(
            'gcp:bigquery/reservation:Reservation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ignore_idle_slots=None, location=None, name=None, project=None, slot_capacity=None):
        """
        Get an existing Reservation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] ignore_idle_slots: If false, any query using this reservation will use idle slots from other reservations within
               the same admin project. If true, a query using this reservation will execute with the slot
               capacity specified above at most.
        :param pulumi.Input[str] location: The geographic location where the transfer config should reside.
               Examples: US, EU, asia-northeast1. The default value is US.
        :param pulumi.Input[str] name: The name of the reservation. This field must only contain alphanumeric characters or dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[float] slot_capacity: Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
               unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ignore_idle_slots"] = ignore_idle_slots
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["slot_capacity"] = slot_capacity
        return Reservation(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

