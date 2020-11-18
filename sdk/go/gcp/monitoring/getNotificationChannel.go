// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A NotificationChannel is a medium through which an alert is delivered
// when a policy violation is detected. Examples of channels include email, SMS,
// and third-party messaging applications. Fields containing sensitive information
// like authentication tokens or contact info are only partially populated on retrieval.
//
// To get more information about NotificationChannel, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
// * How-to Guides
//     * [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
//     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
//
// ## Example Usage
// ### Notification Channel Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/monitoring"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Test Notification Channel"
// 		basic, err := monitoring.LookupNotificationChannel(ctx, &monitoring.LookupNotificationChannelArgs{
// 			DisplayName: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = monitoring.NewAlertPolicy(ctx, "alertPolicy", &monitoring.AlertPolicyArgs{
// 			DisplayName: pulumi.String("My Alert Policy"),
// 			NotificationChannels: pulumi.StringArray{
// 				pulumi.String(basic.Name),
// 			},
// 			Combiner: pulumi.String("OR"),
// 			Conditions: monitoring.AlertPolicyConditionArray{
// 				&monitoring.AlertPolicyConditionArgs{
// 					DisplayName: pulumi.String("test condition"),
// 					ConditionThreshold: &monitoring.AlertPolicyConditionConditionThresholdArgs{
// 						Filter:     pulumi.String("metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\""),
// 						Duration:   pulumi.String("60s"),
// 						Comparison: pulumi.String("COMPARISON_GT"),
// 						Aggregations: monitoring.AlertPolicyConditionConditionThresholdAggregationArray{
// 							&monitoring.AlertPolicyConditionConditionThresholdAggregationArgs{
// 								AlignmentPeriod:  pulumi.String("60s"),
// 								PerSeriesAligner: pulumi.String("ALIGN_RATE"),
// 							},
// 						},
// 					},
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
func LookupNotificationChannel(ctx *pulumi.Context, args *LookupNotificationChannelArgs, opts ...pulumi.InvokeOption) (*LookupNotificationChannelResult, error) {
	var rv LookupNotificationChannelResult
	err := ctx.Invoke("gcp:monitoring/getNotificationChannel:getNotificationChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNotificationChannel.
type LookupNotificationChannelArgs struct {
	// The display name for this notification channel.
	DisplayName *string `pulumi:"displayName"`
	// Labels (corresponding to the
	// NotificationChannelDescriptor schema) to filter the notification channels by.
	Labels map[string]string `pulumi:"labels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The type of the notification channel.
	Type *string `pulumi:"type"`
	// User-provided key-value labels to filter by.
	UserLabels map[string]string `pulumi:"userLabels"`
}

// A collection of values returned by getNotificationChannel.
type LookupNotificationChannelResult struct {
	Description string  `pulumi:"description"`
	DisplayName *string `pulumi:"displayName"`
	Enabled     bool    `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string                                 `pulumi:"id"`
	Labels             map[string]string                      `pulumi:"labels"`
	Name               string                                 `pulumi:"name"`
	Project            *string                                `pulumi:"project"`
	SensitiveLabels    []GetNotificationChannelSensitiveLabel `pulumi:"sensitiveLabels"`
	Type               *string                                `pulumi:"type"`
	UserLabels         map[string]string                      `pulumi:"userLabels"`
	VerificationStatus string                                 `pulumi:"verificationStatus"`
}
