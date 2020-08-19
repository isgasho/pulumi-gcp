# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DicomStore(pulumi.CustomResource):
    dataset: pulumi.Output[str]
    """
    Identifies the dataset addressed by this request. Must be in the format
    'projects/{project}/locations/{location}/datasets/{dataset}'
    """
    labels: pulumi.Output[dict]
    """
    User-supplied key-value pairs used to organize DICOM stores.
    Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
    conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
    Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
    bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
    No more than 64 labels can be associated with a given store.
    An object containing a list of "key": value pairs.
    Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
    """
    name: pulumi.Output[str]
    """
    The resource name for the DicomStore.
    ** Changing this property may recreate the Dicom store (removing all data) **
    """
    notification_config: pulumi.Output[dict]
    """
    A nested object resource
    Structure is documented below.

      * `pubsubTopic` (`str`) - The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
        PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
        It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
        was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
        project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given
        Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
    """
    self_link: pulumi.Output[str]
    """
    The fully qualified name of this dataset
    """
    def __init__(__self__, resource_name, opts=None, dataset=None, labels=None, name=None, notification_config=None, __props__=None, __name__=None, __opts__=None):
        """
        A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
        (https://www.dicomstandard.org/about/) standard for Healthcare information exchange

        To get more information about DicomStore, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.dicomStores)
        * How-to Guides
            * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)

        ## Example Usage

        ### Healthcare Dicom Store Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        topic = gcp.pubsub.Topic("topic")
        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        default = gcp.healthcare.DicomStore("default",
            dataset=dataset.id,
            notification_config={
                "pubsubTopic": topic.id,
            },
            labels={
                "label1": "labelvalue1",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[dict] labels: User-supplied key-value pairs used to organize DICOM stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the DicomStore.
               ** Changing this property may recreate the Dicom store (removing all data) **
        :param pulumi.Input[dict] notification_config: A nested object resource
               Structure is documented below.

        The **notification_config** object supports the following:

          * `pubsubTopic` (`pulumi.Input[str]`) - The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
            PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
            It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
            was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
            project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given
            Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
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

            if dataset is None:
                raise TypeError("Missing required property 'dataset'")
            __props__['dataset'] = dataset
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['notification_config'] = notification_config
            __props__['self_link'] = None
        super(DicomStore, __self__).__init__(
            'gcp:healthcare/dicomStore:DicomStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dataset=None, labels=None, name=None, notification_config=None, self_link=None):
        """
        Get an existing DicomStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[dict] labels: User-supplied key-value pairs used to organize DICOM stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the DicomStore.
               ** Changing this property may recreate the Dicom store (removing all data) **
        :param pulumi.Input[dict] notification_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset

        The **notification_config** object supports the following:

          * `pubsubTopic` (`pulumi.Input[str]`) - The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
            PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
            It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
            was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
            project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given
            Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dataset"] = dataset
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["notification_config"] = notification_config
        __props__["self_link"] = self_link
        return DicomStore(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

