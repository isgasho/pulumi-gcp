// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
 * (https://www.dicomstandard.org/about/) standard for Healthcare information exchange
 *
 * To get more information about DicomStore, see:
 *
 * * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.dicomStores)
 * * How-to Guides
 *     * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)
 *
 * ## Example Usage
 *
 * ### Healthcare Dicom Store Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const topic = new gcp.pubsub.Topic("topic", {});
 * const dataset = new gcp.healthcare.Dataset("dataset", {location: "us-central1"});
 * const _default = new gcp.healthcare.DicomStore("default", {
 *     dataset: dataset.id,
 *     notification_config: {
 *         pubsubTopic: topic.id,
 *     },
 *     labels: {
 *         label1: "labelvalue1",
 *     },
 * });
 * ```
 */
export class DicomStore extends pulumi.CustomResource {
    /**
     * Get an existing DicomStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DicomStoreState, opts?: pulumi.CustomResourceOptions): DicomStore {
        return new DicomStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/dicomStore:DicomStore';

    /**
     * Returns true if the given object is an instance of DicomStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DicomStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DicomStore.__pulumiType;
    }

    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    public readonly dataset!: pulumi.Output<string>;
    /**
     * User-supplied key-value pairs used to organize DICOM stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource name for the DicomStore.
     * ** Changing this property may recreate the Dicom store (removing all data) **
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A nested object resource
     * Structure is documented below.
     */
    public readonly notificationConfig!: pulumi.Output<outputs.healthcare.DicomStoreNotificationConfig | undefined>;
    /**
     * The fully qualified name of this dataset
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

    /**
     * Create a DicomStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DicomStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DicomStoreArgs | DicomStoreState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DicomStoreState | undefined;
            inputs["dataset"] = state ? state.dataset : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationConfig"] = state ? state.notificationConfig : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as DicomStoreArgs | undefined;
            if (!args || args.dataset === undefined) {
                throw new Error("Missing required property 'dataset'");
            }
            inputs["dataset"] = args ? args.dataset : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DicomStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DicomStore resources.
 */
export interface DicomStoreState {
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    readonly dataset?: pulumi.Input<string>;
    /**
     * User-supplied key-value pairs used to organize DICOM stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the DicomStore.
     * ** Changing this property may recreate the Dicom store (removing all data) **
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A nested object resource
     * Structure is documented below.
     */
    readonly notificationConfig?: pulumi.Input<inputs.healthcare.DicomStoreNotificationConfig>;
    /**
     * The fully qualified name of this dataset
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DicomStore resource.
 */
export interface DicomStoreArgs {
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    readonly dataset: pulumi.Input<string>;
    /**
     * User-supplied key-value pairs used to organize DICOM stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the DicomStore.
     * ** Changing this property may recreate the Dicom store (removing all data) **
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A nested object resource
     * Structure is documented below.
     */
    readonly notificationConfig?: pulumi.Input<inputs.healthcare.DicomStoreNotificationConfig>;
}
