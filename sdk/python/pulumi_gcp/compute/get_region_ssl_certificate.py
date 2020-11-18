# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetRegionSslCertificateResult',
    'AwaitableGetRegionSslCertificateResult',
    'get_region_ssl_certificate',
]

@pulumi.output_type
class GetRegionSslCertificateResult:
    """
    A collection of values returned by getRegionSslCertificate.
    """
    def __init__(__self__, certificate=None, certificate_id=None, creation_timestamp=None, description=None, id=None, name=None, name_prefix=None, private_key=None, project=None, region=None, self_link=None):
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        pulumi.set(__self__, "certificate", certificate)
        if certificate_id and not isinstance(certificate_id, int):
            raise TypeError("Expected argument 'certificate_id' to be a int")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        pulumi.set(__self__, "name_prefix", name_prefix)
        if private_key and not isinstance(private_key, str):
            raise TypeError("Expected argument 'private_key' to be a str")
        pulumi.set(__self__, "private_key", private_key)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def certificate(self) -> str:
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> int:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> str:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> str:
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        return pulumi.get(self, "self_link")


class AwaitableGetRegionSslCertificateResult(GetRegionSslCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionSslCertificateResult(
            certificate=self.certificate,
            certificate_id=self.certificate_id,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            id=self.id,
            name=self.name,
            name_prefix=self.name_prefix,
            private_key=self.private_key,
            project=self.project,
            region=self.region,
            self_link=self.self_link)


def get_region_ssl_certificate(name: Optional[str] = None,
                               project: Optional[str] = None,
                               region: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionSslCertificateResult:
    """
    Get info about a Region Google Compute SSL Certificate from its name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_cert = gcp.compute.get_region_ssl_certificate(name="my-cert")
    pulumi.export("certificate", my_cert.certificate)
    pulumi.export("certificateId", my_cert.certificate_id)
    pulumi.export("selfLink", my_cert.self_link)
    ```


    :param str name: The name of the certificate.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str region: The region in which the resource belongs. If it
           is not provided, the provider region is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getRegionSslCertificate:getRegionSslCertificate', __args__, opts=opts, typ=GetRegionSslCertificateResult).value

    return AwaitableGetRegionSslCertificateResult(
        certificate=__ret__.certificate,
        certificate_id=__ret__.certificate_id,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        name_prefix=__ret__.name_prefix,
        private_key=__ret__.private_key,
        project=__ret__.project,
        region=__ret__.region,
        self_link=__ret__.self_link)
