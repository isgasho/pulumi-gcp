// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    /// <summary>
    /// A Hl7V2Store is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/hl7V2/STU3/)
    /// standard for Healthcare information exchange
    /// 
    /// To get more information about Hl7V2Store, see:
    /// 
    /// * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.hl7V2Stores)
    /// * How-to Guides
    ///     * [Creating a HL7v2 Store](https://cloud.google.com/healthcare/docs/how-tos/hl7v2)
    /// 
    /// ## Example Usage
    /// ### Healthcare Hl7 V2 Store Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var topic = new Gcp.PubSub.Topic("topic", new Gcp.PubSub.TopicArgs
    ///         {
    ///         });
    ///         var dataset = new Gcp.Healthcare.Dataset("dataset", new Gcp.Healthcare.DatasetArgs
    ///         {
    ///             Location = "us-central1",
    ///         });
    ///         var store = new Gcp.Healthcare.Hl7Store("store", new Gcp.Healthcare.Hl7StoreArgs
    ///         {
    ///             Dataset = dataset.Id,
    ///             NotificationConfigs = 
    ///             {
    ///                 new Gcp.Healthcare.Inputs.Hl7StoreNotificationConfigsArgs
    ///                 {
    ///                     PubsubTopic = topic.Id,
    ///                 },
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "label1", "labelvalue1" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Healthcare Hl7 V2 Store Parser Config
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dataset = new Gcp.Healthcare.Dataset("dataset", new Gcp.Healthcare.DatasetArgs
    ///         {
    ///             Location = "us-central1",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var store = new Gcp.Healthcare.Hl7Store("store", new Gcp.Healthcare.Hl7StoreArgs
    ///         {
    ///             Dataset = dataset.Id,
    ///             ParserConfig = new Gcp.Healthcare.Inputs.Hl7StoreParserConfigArgs
    ///             {
    ///                 AllowNullHeader = false,
    ///                 SegmentTerminator = "Jw==",
    ///                 Schema = @"{
    ///   ""schemas"": [{
    ///     ""messageSchemaConfigs"": {
    ///       ""ADT_A01"": {
    ///         ""name"": ""ADT_A01"",
    ///         ""minOccurs"": 1,
    ///         ""maxOccurs"": 1,
    ///         ""members"": [{
    ///             ""segment"": {
    ///               ""type"": ""MSH"",
    ///               ""minOccurs"": 1,
    ///               ""maxOccurs"": 1
    ///             }
    ///           },
    ///           {
    ///             ""segment"": {
    ///               ""type"": ""EVN"",
    ///               ""minOccurs"": 1,
    ///               ""maxOccurs"": 1
    ///             }
    ///           },
    ///           {
    ///             ""segment"": {
    ///               ""type"": ""PID"",
    ///               ""minOccurs"": 1,
    ///               ""maxOccurs"": 1
    ///             }
    ///           },
    ///           {
    ///             ""segment"": {
    ///               ""type"": ""ZPD"",
    ///               ""minOccurs"": 1,
    ///               ""maxOccurs"": 1
    ///             }
    ///           },
    ///           {
    ///             ""segment"": {
    ///               ""type"": ""OBX""
    ///             }
    ///           },
    ///           {
    ///             ""group"": {
    ///               ""name"": ""PROCEDURE"",
    ///               ""members"": [{
    ///                   ""segment"": {
    ///                     ""type"": ""PR1"",
    ///                     ""minOccurs"": 1,
    ///                     ""maxOccurs"": 1
    ///                   }
    ///                 },
    ///                 {
    ///                   ""segment"": {
    ///                     ""type"": ""ROL""
    ///                   }
    ///                 }
    ///               ]
    ///             }
    ///           },
    ///           {
    ///             ""segment"": {
    ///               ""type"": ""PDA"",
    ///               ""maxOccurs"": 1
    ///             }
    ///           }
    ///         ]
    ///       }
    ///     }
    ///   }],
    ///   ""types"": [{
    ///     ""type"": [{
    ///         ""name"": ""ZPD"",
    ///         ""primitive"": ""VARIES""
    ///       }
    /// 
    ///     ]
    ///   }],
    ///   ""ignoreMinOccurs"": true
    /// }
    /// ",
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Healthcare Hl7 V2 Store Unschematized
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dataset = new Gcp.Healthcare.Dataset("dataset", new Gcp.Healthcare.DatasetArgs
    ///         {
    ///             Location = "us-central1",
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///         var store = new Gcp.Healthcare.Hl7Store("store", new Gcp.Healthcare.Hl7StoreArgs
    ///         {
    ///             Dataset = dataset.Id,
    ///             ParserConfig = new Gcp.Healthcare.Inputs.Hl7StoreParserConfigArgs
    ///             {
    ///                 AllowNullHeader = false,
    ///                 SegmentTerminator = "Jw==",
    ///                 Version = "V2",
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             Provider = google_beta,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Hl7V2Store can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/hl7Store:Hl7Store default {{dataset}}/hl7V2Stores/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/hl7Store:Hl7Store default {{dataset}}/{{name}}
    /// ```
    /// </summary>
    public partial class Hl7Store : Pulumi.CustomResource
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Output("dataset")]
        public Output<string> Dataset { get; private set; } = null!;

        /// <summary>
        /// User-supplied key-value pairs used to organize HL7v2 stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name for the Hl7V2Store.
        /// ** Changing this property may recreate the Hl7v2 store (removing all data) **
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional, Deprecated)
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Output("notificationConfig")]
        public Output<Outputs.Hl7StoreNotificationConfig?> NotificationConfig { get; private set; } = null!;

        /// <summary>
        /// A list of notification configs. Each configuration uses a filter to determine whether to publish a
        /// message (both Ingest &amp; Create) on the corresponding notification destination. Only the message name
        /// is sent as part of the notification. Supplied by the client.
        /// Structure is documented below.
        /// </summary>
        [Output("notificationConfigs")]
        public Output<ImmutableArray<Outputs.Hl7StoreNotificationConfigs>> NotificationConfigs { get; private set; } = null!;

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Output("parserConfig")]
        public Output<Outputs.Hl7StoreParserConfig?> ParserConfig { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of this dataset
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a Hl7Store resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hl7Store(string name, Hl7StoreArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7Store:Hl7Store", name, args ?? new Hl7StoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hl7Store(string name, Input<string> id, Hl7StoreState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7Store:Hl7Store", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Hl7Store resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hl7Store Get(string name, Input<string> id, Hl7StoreState? state = null, CustomResourceOptions? options = null)
        {
            return new Hl7Store(name, id, state, options);
        }
    }

    public sealed class Hl7StoreArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Input("dataset", required: true)]
        public Input<string> Dataset { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize HL7v2 stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the Hl7V2Store.
        /// ** Changing this property may recreate the Hl7v2 store (removing all data) **
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// -
        /// (Optional, Deprecated)
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.Hl7StoreNotificationConfigArgs>? NotificationConfig { get; set; }

        [Input("notificationConfigs")]
        private InputList<Inputs.Hl7StoreNotificationConfigsArgs>? _notificationConfigs;

        /// <summary>
        /// A list of notification configs. Each configuration uses a filter to determine whether to publish a
        /// message (both Ingest &amp; Create) on the corresponding notification destination. Only the message name
        /// is sent as part of the notification. Supplied by the client.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.Hl7StoreNotificationConfigsArgs> NotificationConfigs
        {
            get => _notificationConfigs ?? (_notificationConfigs = new InputList<Inputs.Hl7StoreNotificationConfigsArgs>());
            set => _notificationConfigs = value;
        }

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Input("parserConfig")]
        public Input<Inputs.Hl7StoreParserConfigArgs>? ParserConfig { get; set; }

        public Hl7StoreArgs()
        {
        }
    }

    public sealed class Hl7StoreState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Input("dataset")]
        public Input<string>? Dataset { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize HL7v2 stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the Hl7V2Store.
        /// ** Changing this property may recreate the Hl7v2 store (removing all data) **
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// -
        /// (Optional, Deprecated)
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.Hl7StoreNotificationConfigGetArgs>? NotificationConfig { get; set; }

        [Input("notificationConfigs")]
        private InputList<Inputs.Hl7StoreNotificationConfigsGetArgs>? _notificationConfigs;

        /// <summary>
        /// A list of notification configs. Each configuration uses a filter to determine whether to publish a
        /// message (both Ingest &amp; Create) on the corresponding notification destination. Only the message name
        /// is sent as part of the notification. Supplied by the client.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.Hl7StoreNotificationConfigsGetArgs> NotificationConfigs
        {
            get => _notificationConfigs ?? (_notificationConfigs = new InputList<Inputs.Hl7StoreNotificationConfigsGetArgs>());
            set => _notificationConfigs = value;
        }

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Input("parserConfig")]
        public Input<Inputs.Hl7StoreParserConfigGetArgs>? ParserConfig { get; set; }

        /// <summary>
        /// The fully qualified name of this dataset
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public Hl7StoreState()
        {
        }
    }
}
