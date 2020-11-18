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

__all__ = ['Reservation']


class Reservation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 specific_reservation: Optional[pulumi.Input[pulumi.InputType['ReservationSpecificReservationArgs']]] = None,
                 specific_reservation_required: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a reservation resource. A reservation ensures that capacity is
        held in a specific zone even if the reserved VMs are not running.

        Reservations apply only to Compute Engine, Cloud Dataproc, and Google
        Kubernetes Engine VM usage.Reservations do not apply to `f1-micro` or
        `g1-small` machine types, preemptible VMs, sole tenant nodes, or other
        services not listed above
        like Cloud SQL and Dataflow.

        To get more information about Reservation, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/reservations)
        * How-to Guides
            * [Reserving zonal resources](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources)

        ## Example Usage
        ### Reservation Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        gce_reservation = gcp.compute.Reservation("gceReservation",
            specific_reservation=gcp.compute.ReservationSpecificReservationArgs(
                count=1,
                instance_properties=gcp.compute.ReservationSpecificReservationInstancePropertiesArgs(
                    machine_type="n2-standard-2",
                    min_cpu_platform="Intel Cascade Lake",
                ),
            ),
            zone="us-central1-a")
        ```

        ## Import

        Reservation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/reservation:Reservation default projects/{{project}}/zones/{{zone}}/reservations/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/reservation:Reservation default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/reservation:Reservation default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/reservation:Reservation default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['ReservationSpecificReservationArgs']] specific_reservation: Reservation for instances with specific machine shapes.
               Structure is documented below.
        :param pulumi.Input[bool] specific_reservation_required: When set to true, only VMs that target this reservation by name can
               consume this reservation. Otherwise, it can be consumed by VMs with
               affinity for any reservation. Defaults to false.
        :param pulumi.Input[str] zone: The zone where the reservation is made.
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

            __props__['description'] = description
            __props__['name'] = name
            __props__['project'] = project
            if specific_reservation is None:
                raise TypeError("Missing required property 'specific_reservation'")
            __props__['specific_reservation'] = specific_reservation
            __props__['specific_reservation_required'] = specific_reservation_required
            if zone is None:
                raise TypeError("Missing required property 'zone'")
            __props__['zone'] = zone
            __props__['commitment'] = None
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
            __props__['status'] = None
        super(Reservation, __self__).__init__(
            'gcp:compute/reservation:Reservation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            commitment: Optional[pulumi.Input[str]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            specific_reservation: Optional[pulumi.Input[pulumi.InputType['ReservationSpecificReservationArgs']]] = None,
            specific_reservation_required: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Reservation':
        """
        Get an existing Reservation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] commitment: Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[pulumi.InputType['ReservationSpecificReservationArgs']] specific_reservation: Reservation for instances with specific machine shapes.
               Structure is documented below.
        :param pulumi.Input[bool] specific_reservation_required: When set to true, only VMs that target this reservation by name can
               consume this reservation. Otherwise, it can be consumed by VMs with
               affinity for any reservation. Defaults to false.
        :param pulumi.Input[str] status: The status of the reservation.
        :param pulumi.Input[str] zone: The zone where the reservation is made.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["commitment"] = commitment
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["specific_reservation"] = specific_reservation
        __props__["specific_reservation_required"] = specific_reservation_required
        __props__["status"] = status
        __props__["zone"] = zone
        return Reservation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def commitment(self) -> pulumi.Output[str]:
        """
        Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        """
        return pulumi.get(self, "commitment")

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
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
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
    @pulumi.getter(name="specificReservation")
    def specific_reservation(self) -> pulumi.Output['outputs.ReservationSpecificReservation']:
        """
        Reservation for instances with specific machine shapes.
        Structure is documented below.
        """
        return pulumi.get(self, "specific_reservation")

    @property
    @pulumi.getter(name="specificReservationRequired")
    def specific_reservation_required(self) -> pulumi.Output[Optional[bool]]:
        """
        When set to true, only VMs that target this reservation by name can
        consume this reservation. Otherwise, it can be consumed by VMs with
        affinity for any reservation. Defaults to false.
        """
        return pulumi.get(self, "specific_reservation_required")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the reservation.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone where the reservation is made.
        """
        return pulumi.get(self, "zone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

