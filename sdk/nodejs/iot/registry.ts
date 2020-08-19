// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Google Cloud IoT Core device registry.
 *
 *
 * To get more information about DeviceRegistry, see:
 *
 * * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/iot/docs/)
 *
 * ## Example Usage
 *
 * ### Cloudiot Device Registry Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const test_registry = new gcp.iot.Registry("test-registry", {});
 * ```
 *
 * ### Cloudiot Device Registry Single Event Notification Configs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default_telemetry = new gcp.pubsub.Topic("default-telemetry", {});
 * const test_registry = new gcp.iot.Registry("test-registry", {event_notification_configs: [{
 *     pubsubTopicName: default_telemetry.id,
 *     subfolderMatches: "",
 * }]});
 * ```
 *
 * ### Cloudiot Device Registry Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * from "fs";
 *
 * const default_devicestatus = new gcp.pubsub.Topic("default-devicestatus", {});
 * const default_telemetry = new gcp.pubsub.Topic("default-telemetry", {});
 * const additional_telemetry = new gcp.pubsub.Topic("additional-telemetry", {});
 * const test_registry = new gcp.iot.Registry("test-registry", {
 *     event_notification_configs: [
 *         {
 *             pubsubTopicName: additional_telemetry.id,
 *             subfolderMatches: "test/path",
 *         },
 *         {
 *             pubsubTopicName: default_telemetry.id,
 *             subfolderMatches: "",
 *         },
 *     ],
 *     stateNotificationConfig: {
 *         pubsub_topic_name: default_devicestatus.id,
 *     },
 *     mqttConfig: {
 *         mqtt_enabled_state: "MQTT_ENABLED",
 *     },
 *     httpConfig: {
 *         http_enabled_state: "HTTP_ENABLED",
 *     },
 *     logLevel: "INFO",
 *     credentials: [{
 *         publicKeyCertificate: {
 *             format: "X509_CERTIFICATE_PEM",
 *             certificate: fs.readFileSync("test-fixtures/rsa_cert.pem"),
 *         },
 *     }],
 * });
 * ```
 */
export class Registry extends pulumi.CustomResource {
    /**
     * Get an existing Registry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryState, opts?: pulumi.CustomResourceOptions): Registry {
        return new Registry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iot/registry:Registry';

    /**
     * Returns true if the given object is an instance of Registry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registry.__pulumiType;
    }

    /**
     * List of public key certificates to authenticate devices.
     * The structure is documented below.
     */
    public readonly credentials!: pulumi.Output<outputs.iot.RegistryCredential[] | undefined>;
    /**
     * List of configurations for event notifications, such as PubSub topics
     * to publish device events to.
     * Structure is documented below.
     */
    public readonly eventNotificationConfigs!: pulumi.Output<outputs.iot.RegistryEventNotificationConfigItem[]>;
    /**
     * Activate or deactivate HTTP.
     * The structure is documented below.
     */
    public readonly httpConfig!: pulumi.Output<outputs.iot.RegistryHttpConfig>;
    /**
     * The default logging verbosity for activity from devices in this
     * registry. Specifies which events should be written to logs. For
     * example, if the LogLevel is ERROR, only events that terminate in
     * errors will be logged. LogLevel is inclusive; enabling INFO logging
     * will also enable ERROR logging.
     * Default value is `NONE`.
     * Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
     */
    public readonly logLevel!: pulumi.Output<string | undefined>;
    /**
     * Activate or deactivate MQTT.
     * The structure is documented below.
     */
    public readonly mqttConfig!: pulumi.Output<outputs.iot.RegistryMqttConfig>;
    /**
     * A unique name for the resource, required by device registry.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region in which the created registry should reside.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A PubSub topic to publish device state updates.
     * The structure is documented below.
     */
    public readonly stateNotificationConfig!: pulumi.Output<outputs.iot.RegistryStateNotificationConfig | undefined>;

    /**
     * Create a Registry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryArgs | RegistryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegistryState | undefined;
            inputs["credentials"] = state ? state.credentials : undefined;
            inputs["eventNotificationConfigs"] = state ? state.eventNotificationConfigs : undefined;
            inputs["httpConfig"] = state ? state.httpConfig : undefined;
            inputs["logLevel"] = state ? state.logLevel : undefined;
            inputs["mqttConfig"] = state ? state.mqttConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["stateNotificationConfig"] = state ? state.stateNotificationConfig : undefined;
        } else {
            const args = argsOrState as RegistryArgs | undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["eventNotificationConfigs"] = args ? args.eventNotificationConfigs : undefined;
            inputs["httpConfig"] = args ? args.httpConfig : undefined;
            inputs["logLevel"] = args ? args.logLevel : undefined;
            inputs["mqttConfig"] = args ? args.mqttConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["stateNotificationConfig"] = args ? args.stateNotificationConfig : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "gcp:kms/registry:Registry" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Registry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Registry resources.
 */
export interface RegistryState {
    /**
     * List of public key certificates to authenticate devices.
     * The structure is documented below.
     */
    readonly credentials?: pulumi.Input<pulumi.Input<inputs.iot.RegistryCredential>[]>;
    /**
     * List of configurations for event notifications, such as PubSub topics
     * to publish device events to.
     * Structure is documented below.
     */
    readonly eventNotificationConfigs?: pulumi.Input<pulumi.Input<inputs.iot.RegistryEventNotificationConfigItem>[]>;
    /**
     * Activate or deactivate HTTP.
     * The structure is documented below.
     */
    readonly httpConfig?: pulumi.Input<inputs.iot.RegistryHttpConfig>;
    /**
     * The default logging verbosity for activity from devices in this
     * registry. Specifies which events should be written to logs. For
     * example, if the LogLevel is ERROR, only events that terminate in
     * errors will be logged. LogLevel is inclusive; enabling INFO logging
     * will also enable ERROR logging.
     * Default value is `NONE`.
     * Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
     */
    readonly logLevel?: pulumi.Input<string>;
    /**
     * Activate or deactivate MQTT.
     * The structure is documented below.
     */
    readonly mqttConfig?: pulumi.Input<inputs.iot.RegistryMqttConfig>;
    /**
     * A unique name for the resource, required by device registry.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the created registry should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A PubSub topic to publish device state updates.
     * The structure is documented below.
     */
    readonly stateNotificationConfig?: pulumi.Input<inputs.iot.RegistryStateNotificationConfig>;
}

/**
 * The set of arguments for constructing a Registry resource.
 */
export interface RegistryArgs {
    /**
     * List of public key certificates to authenticate devices.
     * The structure is documented below.
     */
    readonly credentials?: pulumi.Input<pulumi.Input<inputs.iot.RegistryCredential>[]>;
    /**
     * List of configurations for event notifications, such as PubSub topics
     * to publish device events to.
     * Structure is documented below.
     */
    readonly eventNotificationConfigs?: pulumi.Input<pulumi.Input<inputs.iot.RegistryEventNotificationConfigItem>[]>;
    /**
     * Activate or deactivate HTTP.
     * The structure is documented below.
     */
    readonly httpConfig?: pulumi.Input<inputs.iot.RegistryHttpConfig>;
    /**
     * The default logging verbosity for activity from devices in this
     * registry. Specifies which events should be written to logs. For
     * example, if the LogLevel is ERROR, only events that terminate in
     * errors will be logged. LogLevel is inclusive; enabling INFO logging
     * will also enable ERROR logging.
     * Default value is `NONE`.
     * Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
     */
    readonly logLevel?: pulumi.Input<string>;
    /**
     * Activate or deactivate MQTT.
     * The structure is documented below.
     */
    readonly mqttConfig?: pulumi.Input<inputs.iot.RegistryMqttConfig>;
    /**
     * A unique name for the resource, required by device registry.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the created registry should reside.
     * If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A PubSub topic to publish device state updates.
     * The structure is documented below.
     */
    readonly stateNotificationConfig?: pulumi.Input<inputs.iot.RegistryStateNotificationConfig>;
}
