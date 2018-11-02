# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetProjectServicesResult(object):
    """
    A collection of values returned by getProjectServices.
    """
    def __init__(__self__, disable_on_destroy=None, services=None, id=None):
        if disable_on_destroy and not isinstance(disable_on_destroy, bool):
            raise TypeError('Expected argument disable_on_destroy to be a bool')
        __self__.disable_on_destroy = disable_on_destroy
        if services and not isinstance(services, list):
            raise TypeError('Expected argument services to be a list')
        __self__.services = services
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_project_services(project=None):
    """
    Use this data source to get details on the enabled project services.
    
    For a list of services available, visit the
    [API library page](https://console.cloud.google.com/apis/library) or run `gcloud services list`.
    """
    __args__ = dict()

    __args__['project'] = project
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getProjectServices:getProjectServices', __args__)

    return GetProjectServicesResult(
        disable_on_destroy=__ret__.get('disableOnDestroy'),
        services=__ret__.get('services'),
        id=__ret__.get('id'))