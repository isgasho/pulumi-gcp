# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['BillingAccountExclusion']


class BillingAccountExclusion(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a billing account logging exclusion. For more information see
        [the official documentation](https://cloud.google.com/logging/docs/) and
        [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).

        Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
        granted to the credentials used with the provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_exclusion = gcp.logging.BillingAccountExclusion("my-exclusion",
            billing_account="ABCDEF-012345-GHIJKL",
            description="Exclude GCE instance debug logs",
            filter="resource.type = gce_instance AND severity <= DEBUG")
        ```

        ## Import

        Billing account logging exclusions can be imported using their URI, e.g.

        ```sh
         $ pulumi import gcp:logging/billingAccountExclusion:BillingAccountExclusion my_exclusion billingAccounts/my-billing_account/exclusions/my-exclusion
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account: The billing account to create the exclusion for.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging exclusion.
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

            if billing_account is None:
                raise TypeError("Missing required property 'billing_account'")
            __props__['billing_account'] = billing_account
            __props__['description'] = description
            __props__['disabled'] = disabled
            if filter is None:
                raise TypeError("Missing required property 'filter'")
            __props__['filter'] = filter
            __props__['name'] = name
        super(BillingAccountExclusion, __self__).__init__(
            'gcp:logging/billingAccountExclusion:BillingAccountExclusion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            billing_account: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            filter: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'BillingAccountExclusion':
        """
        Get an existing BillingAccountExclusion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account: The billing account to create the exclusion for.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging exclusion.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["billing_account"] = billing_account
        __props__["description"] = description
        __props__["disabled"] = disabled
        __props__["filter"] = filter
        __props__["name"] = name
        return BillingAccountExclusion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> pulumi.Output[str]:
        """
        The billing account to create the exclusion for.
        """
        return pulumi.get(self, "billing_account")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A human-readable description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether this exclusion rule should be disabled or not. This defaults to
        false.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[str]:
        """
        The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        write a filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the logging exclusion.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

