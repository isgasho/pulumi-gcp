// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A named resource to which messages are sent by publishers.
 *
 * ## Example Usage
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueState, opts?: pulumi.CustomResourceOptions): Queue {
        return new Queue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudtasks/queue:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * Overrides for task-level appEngineRouting. These settings apply only
     * to App Engine tasks in this queue
     * Structure is documented below.
     */
    public readonly appEngineRoutingOverride!: pulumi.Output<outputs.cloudtasks.QueueAppEngineRoutingOverride | undefined>;
    /**
     * The location of the queue
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The queue name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Rate limits for task dispatches.
     * The queue's actual dispatch rate is the result of:
     * * Number of tasks in the queue
     * * User-specified throttling: rateLimits, retryConfig, and the queue's state.
     * * System throttling due to 429 (Too Many Requests) or 503 (Service
     * Unavailable) responses from the worker, high error rates, or to
     * smooth sudden large traffic spikes.
     * Structure is documented below.
     */
    public readonly rateLimits!: pulumi.Output<outputs.cloudtasks.QueueRateLimits>;
    /**
     * Settings that determine the retry behavior.
     * Structure is documented below.
     */
    public readonly retryConfig!: pulumi.Output<outputs.cloudtasks.QueueRetryConfig>;
    /**
     * Configuration options for writing logs to Stackdriver Logging.
     * Structure is documented below.
     */
    public readonly stackdriverLoggingConfig!: pulumi.Output<outputs.cloudtasks.QueueStackdriverLoggingConfig | undefined>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueueArgs | QueueState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as QueueState | undefined;
            inputs["appEngineRoutingOverride"] = state ? state.appEngineRoutingOverride : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["rateLimits"] = state ? state.rateLimits : undefined;
            inputs["retryConfig"] = state ? state.retryConfig : undefined;
            inputs["stackdriverLoggingConfig"] = state ? state.stackdriverLoggingConfig : undefined;
        } else {
            const args = argsOrState as QueueArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            inputs["appEngineRoutingOverride"] = args ? args.appEngineRoutingOverride : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["rateLimits"] = args ? args.rateLimits : undefined;
            inputs["retryConfig"] = args ? args.retryConfig : undefined;
            inputs["stackdriverLoggingConfig"] = args ? args.stackdriverLoggingConfig : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Queue.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queue resources.
 */
export interface QueueState {
    /**
     * Overrides for task-level appEngineRouting. These settings apply only
     * to App Engine tasks in this queue
     * Structure is documented below.
     */
    readonly appEngineRoutingOverride?: pulumi.Input<inputs.cloudtasks.QueueAppEngineRoutingOverride>;
    /**
     * The location of the queue
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The queue name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Rate limits for task dispatches.
     * The queue's actual dispatch rate is the result of:
     * * Number of tasks in the queue
     * * User-specified throttling: rateLimits, retryConfig, and the queue's state.
     * * System throttling due to 429 (Too Many Requests) or 503 (Service
     * Unavailable) responses from the worker, high error rates, or to
     * smooth sudden large traffic spikes.
     * Structure is documented below.
     */
    readonly rateLimits?: pulumi.Input<inputs.cloudtasks.QueueRateLimits>;
    /**
     * Settings that determine the retry behavior.
     * Structure is documented below.
     */
    readonly retryConfig?: pulumi.Input<inputs.cloudtasks.QueueRetryConfig>;
    /**
     * Configuration options for writing logs to Stackdriver Logging.
     * Structure is documented below.
     */
    readonly stackdriverLoggingConfig?: pulumi.Input<inputs.cloudtasks.QueueStackdriverLoggingConfig>;
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * Overrides for task-level appEngineRouting. These settings apply only
     * to App Engine tasks in this queue
     * Structure is documented below.
     */
    readonly appEngineRoutingOverride?: pulumi.Input<inputs.cloudtasks.QueueAppEngineRoutingOverride>;
    /**
     * The location of the queue
     */
    readonly location: pulumi.Input<string>;
    /**
     * The queue name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Rate limits for task dispatches.
     * The queue's actual dispatch rate is the result of:
     * * Number of tasks in the queue
     * * User-specified throttling: rateLimits, retryConfig, and the queue's state.
     * * System throttling due to 429 (Too Many Requests) or 503 (Service
     * Unavailable) responses from the worker, high error rates, or to
     * smooth sudden large traffic spikes.
     * Structure is documented below.
     */
    readonly rateLimits?: pulumi.Input<inputs.cloudtasks.QueueRateLimits>;
    /**
     * Settings that determine the retry behavior.
     * Structure is documented below.
     */
    readonly retryConfig?: pulumi.Input<inputs.cloudtasks.QueueRetryConfig>;
    /**
     * Configuration options for writing logs to Stackdriver Logging.
     * Structure is documented below.
     */
    readonly stackdriverLoggingConfig?: pulumi.Input<inputs.cloudtasks.QueueStackdriverLoggingConfig>;
}
