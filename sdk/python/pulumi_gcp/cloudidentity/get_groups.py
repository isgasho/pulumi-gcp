# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetGroupsResult:
    """
    A collection of values returned by getGroups.
    """
    def __init__(__self__, groups=None, id=None, parent=None):
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        __self__.groups = groups
        """
        The list of groups under the provided customer or namespace. Structure is documented below.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if parent and not isinstance(parent, str):
            raise TypeError("Expected argument 'parent' to be a str")
        __self__.parent = parent
class AwaitableGetGroupsResult(GetGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupsResult(
            groups=self.groups,
            id=self.id,
            parent=self.parent)

def get_groups(parent=None,opts=None):
    """
    Use this data source to get list of the Cloud Identity Groups under a customer or namespace.

    https://cloud.google.com/identity/docs/concepts/overview#groups

    ## Example Usage



    ```python
    import pulumi
    import pulumi_gcp as gcp

    groups = gcp.cloudidentity.get_groups(parent="customers/A01b123xz")
    ```


    :param str parent: The parent resource under which to list all Groups. Must be of the form identitysources/{identity_source_id} for external- identity-mapped groups or customers/{customer_id} for Google Groups.
    """
    __args__ = dict()


    __args__['parent'] = parent
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:cloudidentity/getGroups:getGroups', __args__, opts=opts).value

    return AwaitableGetGroupsResult(
        groups=__ret__.get('groups'),
        id=__ret__.get('id'),
        parent=__ret__.get('parent'))
