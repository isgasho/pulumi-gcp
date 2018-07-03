# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class SubscriptionIAMMember(pulumi.CustomResource):
    """
    Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
    
    * `google_pubsub_subscription_iam_policy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
    * `google_pubsub_subscription_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
    * `google_pubsub_subscription_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
    
    ~> **Note:** `google_pubsub_subscription_iam_policy` **cannot** be used in conjunction with `google_pubsub_subscription_iam_binding` and `google_pubsub_subscription_iam_member` or they will fight over what your policy should be.
    
    ~> **Note:** `google_pubsub_subscription_iam_binding` resources **can be** used in conjunction with `google_pubsub_subscription_iam_member` resources **only if** they do not grant privilege to the same role.
    """
    def __init__(__self__, __name__, __opts__=None, member=None, project=None, role=None, subscription=None):
        """Create a SubscriptionIAMMember resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not member:
            raise TypeError('Missing required property member')
        elif not isinstance(member, basestring):
            raise TypeError('Expected property member to be a basestring')
        __self__.member = member
        __props__['member'] = member

        if project and not isinstance(project, basestring):
            raise TypeError('Expected property project to be a basestring')
        __self__.project = project
        """
        The project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        __props__['project'] = project

        if not role:
            raise TypeError('Missing required property role')
        elif not isinstance(role, basestring):
            raise TypeError('Expected property role to be a basestring')
        __self__.role = role
        """
        The role that should be applied. Only one
        `google_pubsub_subscription_iam_binding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        __props__['role'] = role

        if not subscription:
            raise TypeError('Missing required property subscription')
        elif not isinstance(subscription, basestring):
            raise TypeError('Expected property subscription to be a basestring')
        __self__.subscription = subscription
        """
        The subscription name or id to bind to attach IAM policy to.
        """
        __props__['subscription'] = subscription

        __self__.etag = pulumi.runtime.UNKNOWN
        """
        (Computed) The etag of the subscription's IAM policy.
        """

        super(SubscriptionIAMMember, __self__).__init__(
            'gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'etag' in outs:
            self.etag = outs['etag']
        if 'member' in outs:
            self.member = outs['member']
        if 'project' in outs:
            self.project = outs['project']
        if 'role' in outs:
            self.role = outs['role']
        if 'subscription' in outs:
            self.subscription = outs['subscription']