// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown.
 */
export class FhirStore extends pulumi.CustomResource {
    /**
     * Get an existing FhirStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FhirStoreState, opts?: pulumi.CustomResourceOptions): FhirStore {
        return new FhirStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/fhirStore:FhirStore';

    /**
     * Returns true if the given object is an instance of FhirStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FhirStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FhirStore.__pulumiType;
    }

    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    public readonly dataset!: pulumi.Output<string>;
    /**
     * Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The
     * default value is false, meaning that the API will enforce referential integrity and fail the requests that will
     * result in inconsistent state in the FHIR store. When this field is set to true, the API will skip referential
     * integrity check. Consequently, operations that rely on references, such as Patient.get$everything, will not return
     * all the results if broken references exist. ** Changing this property may recreate the FHIR store (removing all
     * data) **
     */
    public readonly disableReferentialIntegrity!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR
     * store. If set to false, which is the default behavior, all write operations will cause historical versions to be
     * recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If
     * set to true, no historical versions will be kept. The server will send back errors for attempts to read the
     * historical versions. ** Changing this property may recreate the FHIR store (removing all data) **
     */
    public readonly disableResourceVersioning!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to allow the bulk import API to accept history bundles and directly insert historical resource versions into
     * the FHIR store. Importing resource histories creates resource interactions that appear to have occurred in the past,
     * which clients may not want to allow. If set to false, history bundles within an import will fail with an error. **
     * Changing this property may recreate the FHIR store (removing all data) ** ** This property can be changed manually
     * in the Google Cloud Healthcare admin console without recreating the FHIR store **
     */
    public readonly enableHistoryImport!: pulumi.Output<boolean | undefined>;
    /**
     * Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update operation
     * to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create
     * operation and attempts to Update a non-existent resource will return errors. Please treat the audit logs with
     * appropriate levels of care if client-specified resource IDs contain sensitive data such as patient identifiers,
     * those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub notifications.
     */
    public readonly enableUpdateCreate!: pulumi.Output<boolean | undefined>;
    /**
     * User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long,
     * have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
     * [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have
     * a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
     * [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list
     * of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource name for the FhirStore. ** Changing this property may recreate the FHIR store (removing all data) **
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A nested object resource
     */
    public readonly notificationConfig!: pulumi.Output<outputs.healthcare.FhirStoreNotificationConfig | undefined>;
    /**
     * The fully qualified name of this dataset
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The FHIR specification version. Supported values include DSTU2, STU3 and R4. Defaults to STU3.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a FhirStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FhirStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FhirStoreArgs | FhirStoreState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FhirStoreState | undefined;
            inputs["dataset"] = state ? state.dataset : undefined;
            inputs["disableReferentialIntegrity"] = state ? state.disableReferentialIntegrity : undefined;
            inputs["disableResourceVersioning"] = state ? state.disableResourceVersioning : undefined;
            inputs["enableHistoryImport"] = state ? state.enableHistoryImport : undefined;
            inputs["enableUpdateCreate"] = state ? state.enableUpdateCreate : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationConfig"] = state ? state.notificationConfig : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as FhirStoreArgs | undefined;
            if (!args || args.dataset === undefined) {
                throw new Error("Missing required property 'dataset'");
            }
            inputs["dataset"] = args ? args.dataset : undefined;
            inputs["disableReferentialIntegrity"] = args ? args.disableReferentialIntegrity : undefined;
            inputs["disableResourceVersioning"] = args ? args.disableResourceVersioning : undefined;
            inputs["enableHistoryImport"] = args ? args.enableHistoryImport : undefined;
            inputs["enableUpdateCreate"] = args ? args.enableUpdateCreate : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FhirStore.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FhirStore resources.
 */
export interface FhirStoreState {
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    readonly dataset?: pulumi.Input<string>;
    /**
     * Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The
     * default value is false, meaning that the API will enforce referential integrity and fail the requests that will
     * result in inconsistent state in the FHIR store. When this field is set to true, the API will skip referential
     * integrity check. Consequently, operations that rely on references, such as Patient.get$everything, will not return
     * all the results if broken references exist. ** Changing this property may recreate the FHIR store (removing all
     * data) **
     */
    readonly disableReferentialIntegrity?: pulumi.Input<boolean>;
    /**
     * Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR
     * store. If set to false, which is the default behavior, all write operations will cause historical versions to be
     * recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If
     * set to true, no historical versions will be kept. The server will send back errors for attempts to read the
     * historical versions. ** Changing this property may recreate the FHIR store (removing all data) **
     */
    readonly disableResourceVersioning?: pulumi.Input<boolean>;
    /**
     * Whether to allow the bulk import API to accept history bundles and directly insert historical resource versions into
     * the FHIR store. Importing resource histories creates resource interactions that appear to have occurred in the past,
     * which clients may not want to allow. If set to false, history bundles within an import will fail with an error. **
     * Changing this property may recreate the FHIR store (removing all data) ** ** This property can be changed manually
     * in the Google Cloud Healthcare admin console without recreating the FHIR store **
     */
    readonly enableHistoryImport?: pulumi.Input<boolean>;
    /**
     * Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update operation
     * to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create
     * operation and attempts to Update a non-existent resource will return errors. Please treat the audit logs with
     * appropriate levels of care if client-specified resource IDs contain sensitive data such as patient identifiers,
     * those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub notifications.
     */
    readonly enableUpdateCreate?: pulumi.Input<boolean>;
    /**
     * User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long,
     * have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
     * [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have
     * a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
     * [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list
     * of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the FhirStore. ** Changing this property may recreate the FHIR store (removing all data) **
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A nested object resource
     */
    readonly notificationConfig?: pulumi.Input<inputs.healthcare.FhirStoreNotificationConfig>;
    /**
     * The fully qualified name of this dataset
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The FHIR specification version. Supported values include DSTU2, STU3 and R4. Defaults to STU3.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FhirStore resource.
 */
export interface FhirStoreArgs {
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     */
    readonly dataset: pulumi.Input<string>;
    /**
     * Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The
     * default value is false, meaning that the API will enforce referential integrity and fail the requests that will
     * result in inconsistent state in the FHIR store. When this field is set to true, the API will skip referential
     * integrity check. Consequently, operations that rely on references, such as Patient.get$everything, will not return
     * all the results if broken references exist. ** Changing this property may recreate the FHIR store (removing all
     * data) **
     */
    readonly disableReferentialIntegrity?: pulumi.Input<boolean>;
    /**
     * Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR
     * store. If set to false, which is the default behavior, all write operations will cause historical versions to be
     * recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If
     * set to true, no historical versions will be kept. The server will send back errors for attempts to read the
     * historical versions. ** Changing this property may recreate the FHIR store (removing all data) **
     */
    readonly disableResourceVersioning?: pulumi.Input<boolean>;
    /**
     * Whether to allow the bulk import API to accept history bundles and directly insert historical resource versions into
     * the FHIR store. Importing resource histories creates resource interactions that appear to have occurred in the past,
     * which clients may not want to allow. If set to false, history bundles within an import will fail with an error. **
     * Changing this property may recreate the FHIR store (removing all data) ** ** This property can be changed manually
     * in the Google Cloud Healthcare admin console without recreating the FHIR store **
     */
    readonly enableHistoryImport?: pulumi.Input<boolean>;
    /**
     * Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update operation
     * to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create
     * operation and attempts to Update a non-existent resource will return errors. Please treat the audit logs with
     * appropriate levels of care if client-specified resource IDs contain sensitive data such as patient identifiers,
     * those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub notifications.
     */
    readonly enableUpdateCreate?: pulumi.Input<boolean>;
    /**
     * User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long,
     * have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
     * [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have
     * a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
     * [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list
     * of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the FhirStore. ** Changing this property may recreate the FHIR store (removing all data) **
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A nested object resource
     */
    readonly notificationConfig?: pulumi.Input<inputs.healthcare.FhirStoreNotificationConfig>;
    /**
     * The FHIR specification version. Supported values include DSTU2, STU3 and R4. Defaults to STU3.
     */
    readonly version?: pulumi.Input<string>;
}
