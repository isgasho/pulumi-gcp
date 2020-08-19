# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class CryptoKeyIAMPolicy(pulumi.CustomResource):
    crypto_key_id: pulumi.Output[str]
    """
    The crypto key ID, in the form
    `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
    `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
    the provider's project setting will be used as a fallback.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the project's IAM policy.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by
    a `organizations.getIAMPolicy` data source.
    """
    def __init__(__self__, resource_name, opts=None, crypto_key_id=None, policy_data=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:

        * `kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
        * `kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
        * `kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.

        > **Note:** `kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `kms.CryptoKeyIAMBinding` and `kms.CryptoKeyIAMMember` or they will fight over what your policy should be.

        > **Note:** `kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        key = gcp.kms.CryptoKey("key",
            key_ring=keyring.id,
            rotation_period="100000s")
        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/cloudkms.cryptoKeyEncrypter",
            "members": ["user:jane@example.com"],
        }])
        crypto_key = gcp.kms.CryptoKeyIAMPolicy("cryptoKey",
            crypto_key_id=key.id,
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "condition": {
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            "members": ["user:jane@example.com"],
            "role": "roles/cloudkms.cryptoKeyEncrypter",
        }])
        ```

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMBinding("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMBinding("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            })
        ```

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMMember("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMMember("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_key_id: The crypto key ID, in the form
               `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
               `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
               the provider's project setting will be used as a fallback.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
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

            if crypto_key_id is None:
                raise TypeError("Missing required property 'crypto_key_id'")
            __props__['crypto_key_id'] = crypto_key_id
            if policy_data is None:
                raise TypeError("Missing required property 'policy_data'")
            __props__['policy_data'] = policy_data
            __props__['etag'] = None
        super(CryptoKeyIAMPolicy, __self__).__init__(
            'gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, crypto_key_id=None, etag=None, policy_data=None):
        """
        Get an existing CryptoKeyIAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_key_id: The crypto key ID, in the form
               `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
               `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
               the provider's project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["crypto_key_id"] = crypto_key_id
        __props__["etag"] = etag
        __props__["policy_data"] = policy_data
        return CryptoKeyIAMPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

