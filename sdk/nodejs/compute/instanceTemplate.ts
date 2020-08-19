// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a VM instance template resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instance-templates)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instanceTemplates).
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const foobar = new gcp.compute.Disk("foobar", {
 *     image: myImage.then(myImage => myImage.selfLink),
 *     size: 10,
 *     type: "pd-ssd",
 *     zone: "us-central1-a",
 * });
 * const _default = new gcp.compute.InstanceTemplate("default", {
 *     description: "This template is used to create app server instances.",
 *     tags: [
 *         "foo",
 *         "bar",
 *     ],
 *     labels: {
 *         environment: "dev",
 *     },
 *     instanceDescription: "description assigned to instances",
 *     machineType: "n1-standard-1",
 *     canIpForward: false,
 *     scheduling: {
 *         automaticRestart: true,
 *         onHostMaintenance: "MIGRATE",
 *     },
 *     disk: [
 *         {
 *             sourceImage: "debian-cloud/debian-9",
 *             autoDelete: true,
 *             boot: true,
 *         },
 *         {
 *             source: foobar.name,
 *             autoDelete: false,
 *             boot: false,
 *         },
 *     ],
 *     network_interface: [{
 *         network: "default",
 *     }],
 *     metadata: {
 *         foo: "bar",
 *     },
 *     service_account: {
 *         scopes: [
 *             "userinfo-email",
 *             "compute-ro",
 *             "storage-ro",
 *         ],
 *     },
 * });
 * ```
 *
 * ## Using with Instance Group Manager
 *
 * Instance Templates cannot be updated after creation with the Google
 * Cloud Platform API. In order to update an Instance Template, this provider will
 * create a replacement. In order to effectively
 * use an Instance Template resource with an [Instance Group Manager resource](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html).
 * Either omit the Instance Template `name` attribute, or specify a partial name
 * with `namePrefix`. Example:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instanceTemplate = new gcp.compute.InstanceTemplate("instanceTemplate", {
 *     namePrefix: "instance-template-",
 *     machineType: "n1-standard-1",
 *     region: "us-central1",
 *     disk: [{}],
 *     network_interface: [{}],
 * });
 * const instanceGroupManager = new gcp.compute.InstanceGroupManager("instanceGroupManager", {
 *     instanceTemplate: instanceTemplate.id,
 *     baseInstanceName: "instance-group-manager",
 *     zone: "us-central1-f",
 *     targetSize: "1",
 * });
 * ```
 *
 * With this setup, this provider generates a unique name for your Instance
 * Template and can then update the Instance Group manager without conflict before
 * destroying the previous Instance Template.
 *
 * ## Deploying the Latest Image
 *
 * A common way to use instance templates and managed instance groups is to deploy the
 * latest image in a family, usually the latest build of your application. There are two
 * ways to do this in the provider, and they have their pros and cons. The difference ends
 * up being in how "latest" is interpreted. You can either deploy the latest image available
 * when the provider runs, or you can have each instance check what the latest image is when
 * it's being created, either as part of a scaling event or being rebuilt by the instance
 * group manager.
 *
 * If you're not sure, we recommend deploying the latest image available when the provider runs,
 * because this means all the instances in your group will be based on the same image, always,
 * and means that no upgrades or changes to your instances happen outside of a `pulumi up`.
 * You can achieve this by using the `gcp.compute.Image`
 * data source, which will retrieve the latest image on every `pulumi apply`, and will update
 * the template to use that specific image:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const instanceTemplate = new gcp.compute.InstanceTemplate("instanceTemplate", {
 *     namePrefix: "instance-template-",
 *     machineType: "n1-standard-1",
 *     region: "us-central1",
 *     disk: [{
 *         sourceImage: google_compute_image.my_image.self_link,
 *     }],
 * });
 * ```
 *
 * To have instances update to the latest on every scaling event or instance re-creation,
 * use the family as the image for the disk, and it will use GCP's default behavior, setting
 * the image for the template to the family:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instanceTemplate = new gcp.compute.InstanceTemplate("instance_template", {
 *     // boot disk
 *     disks: [{
 *         sourceImage: "debian-cloud/debian-9",
 *     }],
 *     machineType: "n1-standard-1",
 *     namePrefix: "instance-template-",
 *     region: "us-central1",
 * });
 * ```
 */
export class InstanceTemplate extends pulumi.CustomResource {
    /**
     * Get an existing InstanceTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceTemplateState, opts?: pulumi.CustomResourceOptions): InstanceTemplate {
        return new InstanceTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/instanceTemplate:InstanceTemplate';

    /**
     * Returns true if the given object is an instance of InstanceTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceTemplate.__pulumiType;
    }

    /**
     * Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs. This defaults to false.
     */
    public readonly canIpForward!: pulumi.Output<boolean | undefined>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    public readonly confidentialInstanceConfig!: pulumi.Output<outputs.compute.InstanceTemplateConfidentialInstanceConfig>;
    /**
     * A brief description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Disks to attach to instances created from this template.
     * This can be specified multiple times for multiple disks. Structure is
     * documented below.
     */
    public readonly disks!: pulumi.Output<outputs.compute.InstanceTemplateDisk[]>;
    /**
     * Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
     */
    public readonly enableDisplay!: pulumi.Output<boolean | undefined>;
    /**
     * List of the type and count of accelerator cards attached to the instance. Structure documented below.
     */
    public readonly guestAccelerators!: pulumi.Output<outputs.compute.InstanceTemplateGuestAccelerator[] | undefined>;
    /**
     * A brief description to use for instances
     * created from this template.
     */
    public readonly instanceDescription!: pulumi.Output<string | undefined>;
    /**
     * A set of key/value label pairs to assign to instances
     * created from this template,
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The machine type to create.
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * Metadata key/value pairs to make available from
     * within instances created from this template.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The unique fingerprint of the metadata.
     */
    public /*out*/ readonly metadataFingerprint!: pulumi.Output<string>;
    /**
     * An alternative to using the
     * startup-script metadata key, mostly to match the computeInstance resource.
     * This replaces the startup-script metadata key on the created instance and
     * thus the two mechanisms are not allowed to be used simultaneously.
     */
    public readonly metadataStartupScript!: pulumi.Output<string | undefined>;
    /**
     * Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     */
    public readonly minCpuPlatform!: pulumi.Output<string | undefined>;
    /**
     * The name of the instance template. If you leave
     * this blank, the provider will auto-generate a unique name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * Networks to attach to instances created from
     * this template. This can be specified multiple times for multiple networks.
     * Structure is documented below.
     */
    public readonly networkInterfaces!: pulumi.Output<outputs.compute.InstanceTemplateNetworkInterface[] | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * An instance template is a global resource that is not
     * bound to a zone or a region. However, you can still specify some regional
     * resources in an instance template, which restricts the template to the
     * region where that resource resides. For example, a custom `subnetwork`
     * resource is tied to a specific region. Defaults to the region of the
     * Provider if no value is given.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     */
    public readonly scheduling!: pulumi.Output<outputs.compute.InstanceTemplateScheduling>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Service account to attach to the instance. Structure is documented below.
     */
    public readonly serviceAccount!: pulumi.Output<outputs.compute.InstanceTemplateServiceAccount | undefined>;
    /**
     * Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     */
    public readonly shieldedInstanceConfig!: pulumi.Output<outputs.compute.InstanceTemplateShieldedInstanceConfig>;
    /**
     * Tags to attach to the instance.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The unique fingerprint of the tags.
     */
    public /*out*/ readonly tagsFingerprint!: pulumi.Output<string>;

    /**
     * Create a InstanceTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceTemplateArgs | InstanceTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceTemplateState | undefined;
            inputs["canIpForward"] = state ? state.canIpForward : undefined;
            inputs["confidentialInstanceConfig"] = state ? state.confidentialInstanceConfig : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disks"] = state ? state.disks : undefined;
            inputs["enableDisplay"] = state ? state.enableDisplay : undefined;
            inputs["guestAccelerators"] = state ? state.guestAccelerators : undefined;
            inputs["instanceDescription"] = state ? state.instanceDescription : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["machineType"] = state ? state.machineType : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["metadataFingerprint"] = state ? state.metadataFingerprint : undefined;
            inputs["metadataStartupScript"] = state ? state.metadataStartupScript : undefined;
            inputs["minCpuPlatform"] = state ? state.minCpuPlatform : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["scheduling"] = state ? state.scheduling : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            inputs["shieldedInstanceConfig"] = state ? state.shieldedInstanceConfig : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsFingerprint"] = state ? state.tagsFingerprint : undefined;
        } else {
            const args = argsOrState as InstanceTemplateArgs | undefined;
            if (!args || args.disks === undefined) {
                throw new Error("Missing required property 'disks'");
            }
            if (!args || args.machineType === undefined) {
                throw new Error("Missing required property 'machineType'");
            }
            inputs["canIpForward"] = args ? args.canIpForward : undefined;
            inputs["confidentialInstanceConfig"] = args ? args.confidentialInstanceConfig : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disks"] = args ? args.disks : undefined;
            inputs["enableDisplay"] = args ? args.enableDisplay : undefined;
            inputs["guestAccelerators"] = args ? args.guestAccelerators : undefined;
            inputs["instanceDescription"] = args ? args.instanceDescription : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["machineType"] = args ? args.machineType : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["metadataStartupScript"] = args ? args.metadataStartupScript : undefined;
            inputs["minCpuPlatform"] = args ? args.minCpuPlatform : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["scheduling"] = args ? args.scheduling : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["shieldedInstanceConfig"] = args ? args.shieldedInstanceConfig : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["metadataFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["tagsFingerprint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceTemplate resources.
 */
export interface InstanceTemplateState {
    /**
     * Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs. This defaults to false.
     */
    readonly canIpForward?: pulumi.Input<boolean>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    readonly confidentialInstanceConfig?: pulumi.Input<inputs.compute.InstanceTemplateConfidentialInstanceConfig>;
    /**
     * A brief description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Disks to attach to instances created from this template.
     * This can be specified multiple times for multiple disks. Structure is
     * documented below.
     */
    readonly disks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceTemplateDisk>[]>;
    /**
     * Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
     */
    readonly enableDisplay?: pulumi.Input<boolean>;
    /**
     * List of the type and count of accelerator cards attached to the instance. Structure documented below.
     */
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceTemplateGuestAccelerator>[]>;
    /**
     * A brief description to use for instances
     * created from this template.
     */
    readonly instanceDescription?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to instances
     * created from this template,
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The machine type to create.
     */
    readonly machineType?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to make available from
     * within instances created from this template.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The unique fingerprint of the metadata.
     */
    readonly metadataFingerprint?: pulumi.Input<string>;
    /**
     * An alternative to using the
     * startup-script metadata key, mostly to match the computeInstance resource.
     * This replaces the startup-script metadata key on the created instance and
     * thus the two mechanisms are not allowed to be used simultaneously.
     */
    readonly metadataStartupScript?: pulumi.Input<string>;
    /**
     * Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     */
    readonly minCpuPlatform?: pulumi.Input<string>;
    /**
     * The name of the instance template. If you leave
     * this blank, the provider will auto-generate a unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * Networks to attach to instances created from
     * this template. This can be specified multiple times for multiple networks.
     * Structure is documented below.
     */
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceTemplateNetworkInterface>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * An instance template is a global resource that is not
     * bound to a zone or a region. However, you can still specify some regional
     * resources in an instance template, which restricts the template to the
     * region where that resource resides. For example, a custom `subnetwork`
     * resource is tied to a specific region. Defaults to the region of the
     * Provider if no value is given.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     */
    readonly scheduling?: pulumi.Input<inputs.compute.InstanceTemplateScheduling>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Service account to attach to the instance. Structure is documented below.
     */
    readonly serviceAccount?: pulumi.Input<inputs.compute.InstanceTemplateServiceAccount>;
    /**
     * Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     */
    readonly shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceTemplateShieldedInstanceConfig>;
    /**
     * Tags to attach to the instance.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique fingerprint of the tags.
     */
    readonly tagsFingerprint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceTemplate resource.
 */
export interface InstanceTemplateArgs {
    /**
     * Whether to allow sending and receiving of
     * packets with non-matching source or destination IPs. This defaults to false.
     */
    readonly canIpForward?: pulumi.Input<boolean>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    readonly confidentialInstanceConfig?: pulumi.Input<inputs.compute.InstanceTemplateConfidentialInstanceConfig>;
    /**
     * A brief description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Disks to attach to instances created from this template.
     * This can be specified multiple times for multiple disks. Structure is
     * documented below.
     */
    readonly disks: pulumi.Input<pulumi.Input<inputs.compute.InstanceTemplateDisk>[]>;
    /**
     * Enable [Virtual Displays](https://cloud.google.com/compute/docs/instances/enable-instance-virtual-display#verify_display_driver) on this instance.
     * **Note**: `allowStoppingForUpdate` must be set to true in order to update this field.
     */
    readonly enableDisplay?: pulumi.Input<boolean>;
    /**
     * List of the type and count of accelerator cards attached to the instance. Structure documented below.
     */
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceTemplateGuestAccelerator>[]>;
    /**
     * A brief description to use for instances
     * created from this template.
     */
    readonly instanceDescription?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to instances
     * created from this template,
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The machine type to create.
     */
    readonly machineType: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to make available from
     * within instances created from this template.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * An alternative to using the
     * startup-script metadata key, mostly to match the computeInstance resource.
     * This replaces the startup-script metadata key on the created instance and
     * thus the two mechanisms are not allowed to be used simultaneously.
     */
    readonly metadataStartupScript?: pulumi.Input<string>;
    /**
     * Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as
     * `Intel Haswell` or `Intel Skylake`. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
     */
    readonly minCpuPlatform?: pulumi.Input<string>;
    /**
     * The name of the instance template. If you leave
     * this blank, the provider will auto-generate a unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * Networks to attach to instances created from
     * this template. This can be specified multiple times for multiple networks.
     * Structure is documented below.
     */
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceTemplateNetworkInterface>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * An instance template is a global resource that is not
     * bound to a zone or a region. However, you can still specify some regional
     * resources in an instance template, which restricts the template to the
     * region where that resource resides. For example, a custom `subnetwork`
     * resource is tied to a specific region. Defaults to the region of the
     * Provider if no value is given.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The scheduling strategy to use. More details about
     * this configuration option are detailed below.
     */
    readonly scheduling?: pulumi.Input<inputs.compute.InstanceTemplateScheduling>;
    /**
     * Service account to attach to the instance. Structure is documented below.
     */
    readonly serviceAccount?: pulumi.Input<inputs.compute.InstanceTemplateServiceAccount>;
    /**
     * Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.
     * **Note**: `shieldedInstanceConfig` can only be used with boot images with shielded vm support. See the complete list [here](https://cloud.google.com/compute/docs/images#shielded-images).
     */
    readonly shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceTemplateShieldedInstanceConfig>;
    /**
     * Tags to attach to the instance.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
