// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A named resource representing the stream of messages from a single,
// specific topic, to be delivered to the subscribing application.
//
// To get more information about Subscription, see:
//
// * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions)
// * How-to Guides
//     * [Managing Subscriptions](https://cloud.google.com/pubsub/docs/admin#managing_subscriptions)
//
// ## Example Usage
// ### Pubsub Subscription Push
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleTopic, err := pubsub.NewTopic(ctx, "exampleTopic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewSubscription(ctx, "exampleSubscription", &pubsub.SubscriptionArgs{
// 			Topic:              exampleTopic.Name,
// 			AckDeadlineSeconds: pulumi.Int(20),
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			PushConfig: &pubsub.SubscriptionPushConfigArgs{
// 				PushEndpoint: pulumi.String("https://example.com/push"),
// 				Attributes: pulumi.StringMap{
// 					"x-goog-version": pulumi.String("v1"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Pubsub Subscription Pull
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleTopic, err := pubsub.NewTopic(ctx, "exampleTopic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewSubscription(ctx, "exampleSubscription", &pubsub.SubscriptionArgs{
// 			Topic: exampleTopic.Name,
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 			MessageRetentionDuration: pulumi.String("1200s"),
// 			RetainAckedMessages:      pulumi.Bool(true),
// 			AckDeadlineSeconds:       pulumi.Int(20),
// 			ExpirationPolicy: &pubsub.SubscriptionExpirationPolicyArgs{
// 				Ttl: pulumi.String("300000.5s"),
// 			},
// 			RetryPolicy: &pubsub.SubscriptionRetryPolicyArgs{
// 				MinimumBackoff: pulumi.String("10s"),
// 			},
// 			EnableMessageOrdering: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Pubsub Subscription Different Project
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleTopic, err := pubsub.NewTopic(ctx, "exampleTopic", &pubsub.TopicArgs{
// 			Project: pulumi.String("topic-project"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewSubscription(ctx, "exampleSubscription", &pubsub.SubscriptionArgs{
// 			Project: pulumi.String("subscription-project"),
// 			Topic:   exampleTopic.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Pubsub Subscription Dead Letter
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleTopic, err := pubsub.NewTopic(ctx, "exampleTopic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleDeadLetter, err := pubsub.NewTopic(ctx, "exampleDeadLetter", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewSubscription(ctx, "exampleSubscription", &pubsub.SubscriptionArgs{
// 			Topic: exampleTopic.Name,
// 			DeadLetterPolicy: &pubsub.SubscriptionDeadLetterPolicyArgs{
// 				DeadLetterTopic:     exampleDeadLetter.ID(),
// 				MaxDeliveryAttempts: pulumi.Int(10),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Subscription can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:pubsub/subscription:Subscription default projects/{{project}}/subscriptions/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/subscription:Subscription default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/subscription:Subscription default {{name}}
// ```
type Subscription struct {
	pulumi.CustomResourceState

	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	// For pull subscriptions, this value is used as the initial value for
	// the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using
	// pull. The minimum custom deadline you can specify is 10 seconds. The
	// maximum custom deadline you can specify is 600 seconds (10 minutes).
	// If this parameter is 0, a default value of 10 seconds is used.
	// For push delivery, this value is also used to set the request timeout
	// for the call to the push endpoint.
	// If the subscriber never acknowledges the message, the Pub/Sub system
	// will eventually redeliver the message.
	AckDeadlineSeconds pulumi.IntOutput `pulumi:"ackDeadlineSeconds"`
	// A policy that specifies the conditions for dead lettering messages in
	// this subscription. If deadLetterPolicy is not set, dead lettering
	// is disabled.
	// The Cloud Pub/Sub service account associated with this subscriptions's
	// parent project (i.e.,
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	// permission to Acknowledge() messages on this subscription.
	// Structure is documented below.
	DeadLetterPolicy SubscriptionDeadLetterPolicyPtrOutput `pulumi:"deadLetterPolicy"`
	// If `true`, messages published with the same orderingKey in PubsubMessage will be delivered to
	// the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
	// may be delivered in any order.
	EnableMessageOrdering pulumi.BoolPtrOutput `pulumi:"enableMessageOrdering"`
	// A policy that specifies the conditions for this subscription's expiration.
	// A subscription is considered active as long as any connected subscriber
	// is successfully consuming messages from the subscription or is issuing
	// operations on the subscription. If expirationPolicy is not set, a default
	// policy with ttl of 31 days will be used.  If it is set but ttl is "", the
	// resource never expires.  The minimum allowed value for expirationPolicy.ttl
	// is 1 day.
	// Structure is documented below.
	ExpirationPolicy SubscriptionExpirationPolicyOutput `pulumi:"expirationPolicy"`
	// The subscription only delivers the messages that match the filter.
	// Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
	// by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
	// you can't modify the filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// A set of key/value label pairs to assign to this Subscription.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// How long to retain unacknowledged messages in the subscription's
	// backlog, from the moment a message is published. If
	// retainAckedMessages is true, then this also configures the retention
	// of acknowledged messages, and thus configures how far back in time a
	// subscriptions.seek can be done. Defaults to 7 days. Cannot be more
	// than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
	// A duration in seconds with up to nine fractional digits, terminated
	// by 's'. Example: `"600.5s"`.
	MessageRetentionDuration pulumi.StringPtrOutput `pulumi:"messageRetentionDuration"`
	// Name of the subscription.
	Name pulumi.StringOutput `pulumi:"name"`
	Path pulumi.StringOutput `pulumi:"path"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// If push delivery is used with this subscription, this field is used to
	// configure it. An empty pushConfig signifies that the subscriber will
	// pull and ack messages using API methods.
	// Structure is documented below.
	PushConfig SubscriptionPushConfigPtrOutput `pulumi:"pushConfig"`
	// Indicates whether to retain acknowledged messages. If `true`, then
	// messages are not expunged from the subscription's backlog, even if
	// they are acknowledged, until they fall out of the
	// messageRetentionDuration window.
	RetainAckedMessages pulumi.BoolPtrOutput `pulumi:"retainAckedMessages"`
	// A policy that specifies how Pub/Sub retries message delivery for this subscription.
	// If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
	// RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
	// Structure is documented below.
	RetryPolicy SubscriptionRetryPolicyPtrOutput `pulumi:"retryPolicy"`
	// A reference to a Topic resource.
	Topic pulumi.StringOutput `pulumi:"topic"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil || args.Topic == nil {
		return nil, errors.New("missing required argument 'Topic'")
	}
	if args == nil {
		args = &SubscriptionArgs{}
	}
	var resource Subscription
	err := ctx.RegisterResource("gcp:pubsub/subscription:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("gcp:pubsub/subscription:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	// For pull subscriptions, this value is used as the initial value for
	// the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using
	// pull. The minimum custom deadline you can specify is 10 seconds. The
	// maximum custom deadline you can specify is 600 seconds (10 minutes).
	// If this parameter is 0, a default value of 10 seconds is used.
	// For push delivery, this value is also used to set the request timeout
	// for the call to the push endpoint.
	// If the subscriber never acknowledges the message, the Pub/Sub system
	// will eventually redeliver the message.
	AckDeadlineSeconds *int `pulumi:"ackDeadlineSeconds"`
	// A policy that specifies the conditions for dead lettering messages in
	// this subscription. If deadLetterPolicy is not set, dead lettering
	// is disabled.
	// The Cloud Pub/Sub service account associated with this subscriptions's
	// parent project (i.e.,
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	// permission to Acknowledge() messages on this subscription.
	// Structure is documented below.
	DeadLetterPolicy *SubscriptionDeadLetterPolicy `pulumi:"deadLetterPolicy"`
	// If `true`, messages published with the same orderingKey in PubsubMessage will be delivered to
	// the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
	// may be delivered in any order.
	EnableMessageOrdering *bool `pulumi:"enableMessageOrdering"`
	// A policy that specifies the conditions for this subscription's expiration.
	// A subscription is considered active as long as any connected subscriber
	// is successfully consuming messages from the subscription or is issuing
	// operations on the subscription. If expirationPolicy is not set, a default
	// policy with ttl of 31 days will be used.  If it is set but ttl is "", the
	// resource never expires.  The minimum allowed value for expirationPolicy.ttl
	// is 1 day.
	// Structure is documented below.
	ExpirationPolicy *SubscriptionExpirationPolicy `pulumi:"expirationPolicy"`
	// The subscription only delivers the messages that match the filter.
	// Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
	// by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
	// you can't modify the filter.
	Filter *string `pulumi:"filter"`
	// A set of key/value label pairs to assign to this Subscription.
	Labels map[string]string `pulumi:"labels"`
	// How long to retain unacknowledged messages in the subscription's
	// backlog, from the moment a message is published. If
	// retainAckedMessages is true, then this also configures the retention
	// of acknowledged messages, and thus configures how far back in time a
	// subscriptions.seek can be done. Defaults to 7 days. Cannot be more
	// than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
	// A duration in seconds with up to nine fractional digits, terminated
	// by 's'. Example: `"600.5s"`.
	MessageRetentionDuration *string `pulumi:"messageRetentionDuration"`
	// Name of the subscription.
	Name *string `pulumi:"name"`
	Path *string `pulumi:"path"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// If push delivery is used with this subscription, this field is used to
	// configure it. An empty pushConfig signifies that the subscriber will
	// pull and ack messages using API methods.
	// Structure is documented below.
	PushConfig *SubscriptionPushConfig `pulumi:"pushConfig"`
	// Indicates whether to retain acknowledged messages. If `true`, then
	// messages are not expunged from the subscription's backlog, even if
	// they are acknowledged, until they fall out of the
	// messageRetentionDuration window.
	RetainAckedMessages *bool `pulumi:"retainAckedMessages"`
	// A policy that specifies how Pub/Sub retries message delivery for this subscription.
	// If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
	// RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
	// Structure is documented below.
	RetryPolicy *SubscriptionRetryPolicy `pulumi:"retryPolicy"`
	// A reference to a Topic resource.
	Topic *string `pulumi:"topic"`
}

type SubscriptionState struct {
	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	// For pull subscriptions, this value is used as the initial value for
	// the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using
	// pull. The minimum custom deadline you can specify is 10 seconds. The
	// maximum custom deadline you can specify is 600 seconds (10 minutes).
	// If this parameter is 0, a default value of 10 seconds is used.
	// For push delivery, this value is also used to set the request timeout
	// for the call to the push endpoint.
	// If the subscriber never acknowledges the message, the Pub/Sub system
	// will eventually redeliver the message.
	AckDeadlineSeconds pulumi.IntPtrInput
	// A policy that specifies the conditions for dead lettering messages in
	// this subscription. If deadLetterPolicy is not set, dead lettering
	// is disabled.
	// The Cloud Pub/Sub service account associated with this subscriptions's
	// parent project (i.e.,
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	// permission to Acknowledge() messages on this subscription.
	// Structure is documented below.
	DeadLetterPolicy SubscriptionDeadLetterPolicyPtrInput
	// If `true`, messages published with the same orderingKey in PubsubMessage will be delivered to
	// the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
	// may be delivered in any order.
	EnableMessageOrdering pulumi.BoolPtrInput
	// A policy that specifies the conditions for this subscription's expiration.
	// A subscription is considered active as long as any connected subscriber
	// is successfully consuming messages from the subscription or is issuing
	// operations on the subscription. If expirationPolicy is not set, a default
	// policy with ttl of 31 days will be used.  If it is set but ttl is "", the
	// resource never expires.  The minimum allowed value for expirationPolicy.ttl
	// is 1 day.
	// Structure is documented below.
	ExpirationPolicy SubscriptionExpirationPolicyPtrInput
	// The subscription only delivers the messages that match the filter.
	// Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
	// by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
	// you can't modify the filter.
	Filter pulumi.StringPtrInput
	// A set of key/value label pairs to assign to this Subscription.
	Labels pulumi.StringMapInput
	// How long to retain unacknowledged messages in the subscription's
	// backlog, from the moment a message is published. If
	// retainAckedMessages is true, then this also configures the retention
	// of acknowledged messages, and thus configures how far back in time a
	// subscriptions.seek can be done. Defaults to 7 days. Cannot be more
	// than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
	// A duration in seconds with up to nine fractional digits, terminated
	// by 's'. Example: `"600.5s"`.
	MessageRetentionDuration pulumi.StringPtrInput
	// Name of the subscription.
	Name pulumi.StringPtrInput
	Path pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// If push delivery is used with this subscription, this field is used to
	// configure it. An empty pushConfig signifies that the subscriber will
	// pull and ack messages using API methods.
	// Structure is documented below.
	PushConfig SubscriptionPushConfigPtrInput
	// Indicates whether to retain acknowledged messages. If `true`, then
	// messages are not expunged from the subscription's backlog, even if
	// they are acknowledged, until they fall out of the
	// messageRetentionDuration window.
	RetainAckedMessages pulumi.BoolPtrInput
	// A policy that specifies how Pub/Sub retries message delivery for this subscription.
	// If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
	// RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
	// Structure is documented below.
	RetryPolicy SubscriptionRetryPolicyPtrInput
	// A reference to a Topic resource.
	Topic pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	// For pull subscriptions, this value is used as the initial value for
	// the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using
	// pull. The minimum custom deadline you can specify is 10 seconds. The
	// maximum custom deadline you can specify is 600 seconds (10 minutes).
	// If this parameter is 0, a default value of 10 seconds is used.
	// For push delivery, this value is also used to set the request timeout
	// for the call to the push endpoint.
	// If the subscriber never acknowledges the message, the Pub/Sub system
	// will eventually redeliver the message.
	AckDeadlineSeconds *int `pulumi:"ackDeadlineSeconds"`
	// A policy that specifies the conditions for dead lettering messages in
	// this subscription. If deadLetterPolicy is not set, dead lettering
	// is disabled.
	// The Cloud Pub/Sub service account associated with this subscriptions's
	// parent project (i.e.,
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	// permission to Acknowledge() messages on this subscription.
	// Structure is documented below.
	DeadLetterPolicy *SubscriptionDeadLetterPolicy `pulumi:"deadLetterPolicy"`
	// If `true`, messages published with the same orderingKey in PubsubMessage will be delivered to
	// the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
	// may be delivered in any order.
	EnableMessageOrdering *bool `pulumi:"enableMessageOrdering"`
	// A policy that specifies the conditions for this subscription's expiration.
	// A subscription is considered active as long as any connected subscriber
	// is successfully consuming messages from the subscription or is issuing
	// operations on the subscription. If expirationPolicy is not set, a default
	// policy with ttl of 31 days will be used.  If it is set but ttl is "", the
	// resource never expires.  The minimum allowed value for expirationPolicy.ttl
	// is 1 day.
	// Structure is documented below.
	ExpirationPolicy *SubscriptionExpirationPolicy `pulumi:"expirationPolicy"`
	// The subscription only delivers the messages that match the filter.
	// Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
	// by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
	// you can't modify the filter.
	Filter *string `pulumi:"filter"`
	// A set of key/value label pairs to assign to this Subscription.
	Labels map[string]string `pulumi:"labels"`
	// How long to retain unacknowledged messages in the subscription's
	// backlog, from the moment a message is published. If
	// retainAckedMessages is true, then this also configures the retention
	// of acknowledged messages, and thus configures how far back in time a
	// subscriptions.seek can be done. Defaults to 7 days. Cannot be more
	// than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
	// A duration in seconds with up to nine fractional digits, terminated
	// by 's'. Example: `"600.5s"`.
	MessageRetentionDuration *string `pulumi:"messageRetentionDuration"`
	// Name of the subscription.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// If push delivery is used with this subscription, this field is used to
	// configure it. An empty pushConfig signifies that the subscriber will
	// pull and ack messages using API methods.
	// Structure is documented below.
	PushConfig *SubscriptionPushConfig `pulumi:"pushConfig"`
	// Indicates whether to retain acknowledged messages. If `true`, then
	// messages are not expunged from the subscription's backlog, even if
	// they are acknowledged, until they fall out of the
	// messageRetentionDuration window.
	RetainAckedMessages *bool `pulumi:"retainAckedMessages"`
	// A policy that specifies how Pub/Sub retries message delivery for this subscription.
	// If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
	// RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
	// Structure is documented below.
	RetryPolicy *SubscriptionRetryPolicy `pulumi:"retryPolicy"`
	// A reference to a Topic resource.
	Topic string `pulumi:"topic"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	// For pull subscriptions, this value is used as the initial value for
	// the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using
	// pull. The minimum custom deadline you can specify is 10 seconds. The
	// maximum custom deadline you can specify is 600 seconds (10 minutes).
	// If this parameter is 0, a default value of 10 seconds is used.
	// For push delivery, this value is also used to set the request timeout
	// for the call to the push endpoint.
	// If the subscriber never acknowledges the message, the Pub/Sub system
	// will eventually redeliver the message.
	AckDeadlineSeconds pulumi.IntPtrInput
	// A policy that specifies the conditions for dead lettering messages in
	// this subscription. If deadLetterPolicy is not set, dead lettering
	// is disabled.
	// The Cloud Pub/Sub service account associated with this subscriptions's
	// parent project (i.e.,
	// service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
	// permission to Acknowledge() messages on this subscription.
	// Structure is documented below.
	DeadLetterPolicy SubscriptionDeadLetterPolicyPtrInput
	// If `true`, messages published with the same orderingKey in PubsubMessage will be delivered to
	// the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
	// may be delivered in any order.
	EnableMessageOrdering pulumi.BoolPtrInput
	// A policy that specifies the conditions for this subscription's expiration.
	// A subscription is considered active as long as any connected subscriber
	// is successfully consuming messages from the subscription or is issuing
	// operations on the subscription. If expirationPolicy is not set, a default
	// policy with ttl of 31 days will be used.  If it is set but ttl is "", the
	// resource never expires.  The minimum allowed value for expirationPolicy.ttl
	// is 1 day.
	// Structure is documented below.
	ExpirationPolicy SubscriptionExpirationPolicyPtrInput
	// The subscription only delivers the messages that match the filter.
	// Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
	// by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
	// you can't modify the filter.
	Filter pulumi.StringPtrInput
	// A set of key/value label pairs to assign to this Subscription.
	Labels pulumi.StringMapInput
	// How long to retain unacknowledged messages in the subscription's
	// backlog, from the moment a message is published. If
	// retainAckedMessages is true, then this also configures the retention
	// of acknowledged messages, and thus configures how far back in time a
	// subscriptions.seek can be done. Defaults to 7 days. Cannot be more
	// than 7 days (`"604800s"`) or less than 10 minutes (`"600s"`).
	// A duration in seconds with up to nine fractional digits, terminated
	// by 's'. Example: `"600.5s"`.
	MessageRetentionDuration pulumi.StringPtrInput
	// Name of the subscription.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// If push delivery is used with this subscription, this field is used to
	// configure it. An empty pushConfig signifies that the subscriber will
	// pull and ack messages using API methods.
	// Structure is documented below.
	PushConfig SubscriptionPushConfigPtrInput
	// Indicates whether to retain acknowledged messages. If `true`, then
	// messages are not expunged from the subscription's backlog, even if
	// they are acknowledged, until they fall out of the
	// messageRetentionDuration window.
	RetainAckedMessages pulumi.BoolPtrInput
	// A policy that specifies how Pub/Sub retries message delivery for this subscription.
	// If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
	// RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
	// Structure is documented below.
	RetryPolicy SubscriptionRetryPolicyPtrInput
	// A reference to a Topic resource.
	Topic pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((*Subscription)(nil)).Elem()
}

func (i Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

type SubscriptionOutput struct {
	*pulumi.OutputState
}

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionOutput)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
}
