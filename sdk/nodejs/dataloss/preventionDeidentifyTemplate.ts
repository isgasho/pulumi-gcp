// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Allows creation of templates to de-idenfity content.
 *
 * To get more information about DeidentifyTemplate, see:
 *
 * * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.deidentifyTemplates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dlp/docs/concepts-templates)
 *
 * ## Example Usage
 * ### Dlp Deidentify Template Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.dataloss.PreventionDeidentifyTemplate("basic", {
 *     deidentifyConfig: {
 *         infoTypeTransformations: {
 *             transformations: [
 *                 {
 *                     infoTypes: [
 *                         {
 *                             name: "PHONE_NUMBER",
 *                         },
 *                         {
 *                             name: "CREDIT_CARD_NUMBER",
 *                         },
 *                     ],
 *                     primitiveTransformation: {
 *                         replaceConfig: {
 *                             newValue: {
 *                                 integerValue: 9,
 *                             },
 *                         },
 *                     },
 *                 },
 *                 {
 *                     infoTypes: [
 *                         {
 *                             name: "EMAIL_ADDRESS",
 *                         },
 *                         {
 *                             name: "LAST_NAME",
 *                         },
 *                     ],
 *                     primitiveTransformation: {
 *                         characterMaskConfig: {
 *                             charactersToIgnores: [{
 *                                 commonCharactersToIgnore: "PUNCTUATION",
 *                             }],
 *                             maskingCharacter: "X",
 *                             numberToMask: 4,
 *                             reverseOrder: true,
 *                         },
 *                     },
 *                 },
 *                 {
 *                     infoTypes: [{
 *                         name: "DATE_OF_BIRTH",
 *                     }],
 *                     primitiveTransformation: {
 *                         replaceConfig: {
 *                             newValue: {
 *                                 dateValue: {
 *                                     day: 1,
 *                                     month: 1,
 *                                     year: 2020,
 *                                 },
 *                             },
 *                         },
 *                     },
 *                 },
 *             ],
 *         },
 *     },
 *     description: "Description",
 *     displayName: "Displayname",
 *     parent: "projects/my-project-name",
 * });
 * ```
 *
 * ## Import
 *
 * DeidentifyTemplate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/deidentifyTemplates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate default {{parent}}/{{name}}
 * ```
 */
export class PreventionDeidentifyTemplate extends pulumi.CustomResource {
    /**
     * Get an existing PreventionDeidentifyTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PreventionDeidentifyTemplateState, opts?: pulumi.CustomResourceOptions): PreventionDeidentifyTemplate {
        return new PreventionDeidentifyTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataloss/preventionDeidentifyTemplate:PreventionDeidentifyTemplate';

    /**
     * Returns true if the given object is an instance of PreventionDeidentifyTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PreventionDeidentifyTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PreventionDeidentifyTemplate.__pulumiType;
    }

    /**
     * Configuration of the deidentify template
     * Structure is documented below.
     */
    public readonly deidentifyConfig!: pulumi.Output<outputs.dataloss.PreventionDeidentifyTemplateDeidentifyConfig>;
    /**
     * A description of the template.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User set display name of the template.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Name of the information type.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent of the template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    public readonly parent!: pulumi.Output<string>;

    /**
     * Create a PreventionDeidentifyTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PreventionDeidentifyTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PreventionDeidentifyTemplateArgs | PreventionDeidentifyTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PreventionDeidentifyTemplateState | undefined;
            inputs["deidentifyConfig"] = state ? state.deidentifyConfig : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
        } else {
            const args = argsOrState as PreventionDeidentifyTemplateArgs | undefined;
            if (!args || args.deidentifyConfig === undefined) {
                throw new Error("Missing required property 'deidentifyConfig'");
            }
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["deidentifyConfig"] = args ? args.deidentifyConfig : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PreventionDeidentifyTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PreventionDeidentifyTemplate resources.
 */
export interface PreventionDeidentifyTemplateState {
    /**
     * Configuration of the deidentify template
     * Structure is documented below.
     */
    readonly deidentifyConfig?: pulumi.Input<inputs.dataloss.PreventionDeidentifyTemplateDeidentifyConfig>;
    /**
     * A description of the template.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * User set display name of the template.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Name of the information type.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The parent of the template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    readonly parent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PreventionDeidentifyTemplate resource.
 */
export interface PreventionDeidentifyTemplateArgs {
    /**
     * Configuration of the deidentify template
     * Structure is documented below.
     */
    readonly deidentifyConfig: pulumi.Input<inputs.dataloss.PreventionDeidentifyTemplateDeidentifyConfig>;
    /**
     * A description of the template.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * User set display name of the template.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The parent of the template in any of the following formats:
     * * `projects/{{project}}`
     * * `projects/{{project}}/locations/{{location}}`
     * * `organizations/{{organization_id}}`
     * * `organizations/{{organization_id}}/locations/{{location}}`
     */
    readonly parent: pulumi.Input<string>;
}
