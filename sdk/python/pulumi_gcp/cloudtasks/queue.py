# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Queue(pulumi.CustomResource):
    app_engine_routing_override: pulumi.Output[dict]
    """
    Overrides for task-level appEngineRouting. These settings apply only to App Engine tasks in this queue

      * `host` (`str`)
      * `instance` (`str`)
      * `service` (`str`)
      * `version` (`str`)
    """
    location: pulumi.Output[str]
    """
    The location of the queue
    """
    name: pulumi.Output[str]
    """
    The queue name.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    rate_limits: pulumi.Output[dict]
    """
    Rate limits for task dispatches. The queue's actual dispatch rate is the result of: * Number of tasks in the queue *
    User-specified throttling: rateLimits, retryConfig, and the queue's state. * System throttling due to 429 (Too Many
    Requests) or 503 (Service Unavailable) responses from the worker, high error rates, or to smooth sudden large traffic
    spikes.

      * `maxBurstSize` (`float`)
      * `maxConcurrentDispatches` (`float`)
      * `maxDispatchesPerSecond` (`float`)
    """
    retry_config: pulumi.Output[dict]
    """
    Settings that determine the retry behavior.

      * `maxAttempts` (`float`)
      * `maxBackoff` (`str`)
      * `maxDoublings` (`float`)
      * `maxRetryDuration` (`str`)
      * `minBackoff` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, app_engine_routing_override=None, location=None, name=None, project=None, rate_limits=None, retry_config=None, __props__=None, __name__=None, __opts__=None):
        """
        A named resource to which messages are sent by publishers.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_tasks_queue.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only to App Engine tasks in this queue
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] rate_limits: Rate limits for task dispatches. The queue's actual dispatch rate is the result of: * Number of tasks in the queue *
               User-specified throttling: rateLimits, retryConfig, and the queue's state. * System throttling due to 429 (Too Many
               Requests) or 503 (Service Unavailable) responses from the worker, high error rates, or to smooth sudden large traffic
               spikes.
        :param pulumi.Input[dict] retry_config: Settings that determine the retry behavior.

        The **app_engine_routing_override** object supports the following:

          * `host` (`pulumi.Input[str]`)
          * `instance` (`pulumi.Input[str]`)
          * `service` (`pulumi.Input[str]`)
          * `version` (`pulumi.Input[str]`)

        The **rate_limits** object supports the following:

          * `maxBurstSize` (`pulumi.Input[float]`)
          * `maxConcurrentDispatches` (`pulumi.Input[float]`)
          * `maxDispatchesPerSecond` (`pulumi.Input[float]`)

        The **retry_config** object supports the following:

          * `maxAttempts` (`pulumi.Input[float]`)
          * `maxBackoff` (`pulumi.Input[str]`)
          * `maxDoublings` (`pulumi.Input[float]`)
          * `maxRetryDuration` (`pulumi.Input[str]`)
          * `minBackoff` (`pulumi.Input[str]`)
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

            __props__['app_engine_routing_override'] = app_engine_routing_override
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['name'] = name
            __props__['project'] = project
            __props__['rate_limits'] = rate_limits
            __props__['retry_config'] = retry_config
        super(Queue, __self__).__init__(
            'gcp:cloudtasks/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_engine_routing_override=None, location=None, name=None, project=None, rate_limits=None, retry_config=None):
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] app_engine_routing_override: Overrides for task-level appEngineRouting. These settings apply only to App Engine tasks in this queue
        :param pulumi.Input[str] location: The location of the queue
        :param pulumi.Input[str] name: The queue name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] rate_limits: Rate limits for task dispatches. The queue's actual dispatch rate is the result of: * Number of tasks in the queue *
               User-specified throttling: rateLimits, retryConfig, and the queue's state. * System throttling due to 429 (Too Many
               Requests) or 503 (Service Unavailable) responses from the worker, high error rates, or to smooth sudden large traffic
               spikes.
        :param pulumi.Input[dict] retry_config: Settings that determine the retry behavior.

        The **app_engine_routing_override** object supports the following:

          * `host` (`pulumi.Input[str]`)
          * `instance` (`pulumi.Input[str]`)
          * `service` (`pulumi.Input[str]`)
          * `version` (`pulumi.Input[str]`)

        The **rate_limits** object supports the following:

          * `maxBurstSize` (`pulumi.Input[float]`)
          * `maxConcurrentDispatches` (`pulumi.Input[float]`)
          * `maxDispatchesPerSecond` (`pulumi.Input[float]`)

        The **retry_config** object supports the following:

          * `maxAttempts` (`pulumi.Input[float]`)
          * `maxBackoff` (`pulumi.Input[str]`)
          * `maxDoublings` (`pulumi.Input[float]`)
          * `maxRetryDuration` (`pulumi.Input[str]`)
          * `minBackoff` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_engine_routing_override"] = app_engine_routing_override
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["rate_limits"] = rate_limits
        __props__["retry_config"] = retry_config
        return Queue(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

