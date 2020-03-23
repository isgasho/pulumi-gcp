# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSecretVersionResult:
    """
    A collection of values returned by getSecretVersion.
    """
    def __init__(__self__, create_time=None, destroy_time=None, enabled=None, id=None, name=None, project=None, secret=None, secret_data=None, version=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        __self__.create_time = create_time
        """
        The time at which the Secret was created.
        """
        if destroy_time and not isinstance(destroy_time, str):
            raise TypeError("Expected argument 'destroy_time' to be a str")
        __self__.destroy_time = destroy_time
        """
        The time at which the Secret was destroyed. Only present if state is DESTROYED.
        """
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        """
        True if the current state of the SecretVersion is enabled.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The resource name of the SecretVersion. Format:
        `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if secret and not isinstance(secret, str):
            raise TypeError("Expected argument 'secret' to be a str")
        __self__.secret = secret
        if secret_data and not isinstance(secret_data, str):
            raise TypeError("Expected argument 'secret_data' to be a str")
        __self__.secret_data = secret_data
        """
        The secret data. No larger than 64KiB.
        """
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        __self__.version = version
class AwaitableGetSecretVersionResult(GetSecretVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretVersionResult(
            create_time=self.create_time,
            destroy_time=self.destroy_time,
            enabled=self.enabled,
            id=self.id,
            name=self.name,
            project=self.project,
            secret=self.secret,
            secret_data=self.secret_data,
            version=self.version)

def get_secret_version(project=None,secret=None,version=None,opts=None):
    """
    Get a Secret Manager secret's version. For more information see the [official documentation](https://cloud.google.com/secret-manager/docs/) and [API](https://cloud.google.com/secret-manager/docs/reference/rest/v1beta1/projects.secrets.versions).

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_secret_manager_secret_version.html.markdown.


    :param str project: The project to get the secret version for. If it
           is not provided, the provider project is used.
    :param str secret: The secret to get the secret version for.
    :param str version: The version of the secret to get. If it
           is not provided, the latest version is retrieved.
    """
    __args__ = dict()


    __args__['project'] = project
    __args__['secret'] = secret
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:secretmanager/getSecretVersion:getSecretVersion', __args__, opts=opts).value

    return AwaitableGetSecretVersionResult(
        create_time=__ret__.get('createTime'),
        destroy_time=__ret__.get('destroyTime'),
        enabled=__ret__.get('enabled'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        secret=__ret__.get('secret'),
        secret_data=__ret__.get('secretData'),
        version=__ret__.get('version'))
