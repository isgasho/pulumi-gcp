# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Entry(pulumi.CustomResource):
    bigquery_date_sharded_spec: pulumi.Output[dict]
    """
    Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
    https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.

      * `dataset` (`str`)
      * `shardCount` (`float`)
      * `tablePrefix` (`str`)
    """
    bigquery_table_spec: pulumi.Output[dict]
    """
    Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.

      * `tableSourceType` (`str`)
      * `tableSpec` (`dict`)
        * `groupedEntry` (`str`)

      * `viewSpec` (`dict`)
        * `viewQuery` (`str`)
    """
    description: pulumi.Output[str]
    """
    Entry description, which can consist of several sentences or paragraphs that describe entry contents.
    """
    display_name: pulumi.Output[str]
    """
    Display information such as title and description. A short name to identify the entry,
    for example, "Analytics Data - Jan 2011".
    """
    entry_group: pulumi.Output[str]
    """
    The name of the entry group this entry is in.
    """
    entry_id: pulumi.Output[str]
    """
    The id of the entry to create.
    """
    gcs_fileset_spec: pulumi.Output[dict]
    """
    Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.

      * `filePatterns` (`list`) - Patterns to identify a set of files in Google Cloud Storage.
        See [Cloud Storage documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)
        for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:
        * gs://bucket_name/dir/*: matches all files within bucket_name/dir directory.
        * gs://bucket_name/dir/**: matches all files in bucket_name/dir spanning all subdirectories.
        * gs://bucket_name/file*: matches files prefixed by file in bucket_name
        * gs://bucket_name/??.txt: matches files with two characters followed by .txt in bucket_name
        * gs://bucket_name/[aeiou].txt: matches files that contain a single vowel character followed by .txt in bucket_name
        * gs://bucket_name/[a-m].txt: matches files that contain a, b, ... or m followed by .txt in bucket_name
        * gs://bucket_name/a/*/b: matches all files in bucket_name that match a/*/b pattern, such as a/c/b, a/d/b
        * gs://another_bucket/a.txt: matches gs://another_bucket/a.txt
      * `sampleGcsFileSpecs` (`list`) - -
        Sample files contained in this fileset, not all files contained in this fileset are represented here.  Structure is documented below.
        * `filePath` (`str`) - -
          The full file path
        * `sizeBytes` (`float`) - -
          The size of the file, in bytes.
    """
    integrated_system: pulumi.Output[str]
    """
    This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
    """
    linked_resource: pulumi.Output[str]
    """
    The resource this metadata entry refers to.
    For Google Cloud Platform resources, linkedResource is the full name of the resource.
    For example, the linkedResource for a table resource from BigQuery is:
    //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
    Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
    this field is optional and defaults to an empty string.
    """
    name: pulumi.Output[str]
    """
    The Data Catalog resource name of the entry in URL format. Example:
    projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
    child resources may not actually be stored in the location in this name.
    """
    schema: pulumi.Output[str]
    """
    Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
    attached to it. See
    https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
    for what fields this schema can contain.
    """
    type: pulumi.Output[str]
    """
    The type of the entry. Only used for Entries with types in the EntryType enum.
    Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
    """
    user_specified_system: pulumi.Output[str]
    """
    This field indicates the entry's source system that Data Catalog does not integrate with.
    userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
    and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
    """
    user_specified_type: pulumi.Output[str]
    """
    Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
    When creating an entry, users should check the enum values first, if nothing matches the entry
    to be created, then provide a custom value, for example "my_special_type".
    userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
    numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
    """
    def __init__(__self__, resource_name, opts=None, description=None, display_name=None, entry_group=None, entry_id=None, gcs_fileset_spec=None, linked_resource=None, schema=None, type=None, user_specified_system=None, user_specified_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Entry Metadata. A Data Catalog Entry resource represents another resource in Google Cloud Platform
        (such as a BigQuery dataset or a Pub/Sub topic) or outside of Google Cloud Platform. Clients can use
        the linkedResource field in the Entry resource to refer to the original resource ID of the source system.

        An Entry resource contains resource details, such as its schema. An Entry can also be used to attach
        flexible metadata, such as a Tag.

        To get more information about Entry, see:

        * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/data-catalog/docs)

        ## Example Usage
        ### Data Catalog Entry Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        entry_group = gcp.datacatalog.EntryGroup("entryGroup", entry_group_id="my_group")
        basic_entry = gcp.datacatalog.Entry("basicEntry",
            entry_group=entry_group.id,
            entry_id="my_entry",
            user_specified_type="my_custom_type",
            user_specified_system="SomethingExternal")
        ```
        ### Data Catalog Entry Fileset

        ```python
        import pulumi
        import pulumi_gcp as gcp

        entry_group = gcp.datacatalog.EntryGroup("entryGroup", entry_group_id="my_group")
        basic_entry = gcp.datacatalog.Entry("basicEntry",
            entry_group=entry_group.id,
            entry_id="my_entry",
            type="FILESET",
            gcs_fileset_spec={
                "filePatterns": ["gs://fake_bucket/dir/*"],
            })
        ```
        ### Data Catalog Entry Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        entry_group = gcp.datacatalog.EntryGroup("entryGroup", entry_group_id="my_group")
        basic_entry = gcp.datacatalog.Entry("basicEntry",
            entry_group=entry_group.id,
            entry_id="my_entry",
            user_specified_type="my_user_specified_type",
            user_specified_system="Something_custom",
            linked_resource="my/linked/resource",
            display_name="my custom type entry",
            description="a custom type entry for a user specified system",
            schema=\"\"\"{
          "columns": [
            {
              "column": "first_name",
              "description": "First name",
              "mode": "REQUIRED",
              "type": "STRING"
            },
            {
              "column": "last_name",
              "description": "Last name",
              "mode": "REQUIRED",
              "type": "STRING"
            },
            {
              "column": "address",
              "description": "Address",
              "mode": "REPEATED",
              "subcolumns": [
                {
                  "column": "city",
                  "description": "City",
                  "mode": "NULLABLE",
                  "type": "STRING"
                },
                {
                  "column": "state",
                  "description": "State",
                  "mode": "NULLABLE",
                  "type": "STRING"
                }
              ],
              "type": "RECORD"
            }
          ]
        }
        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Entry description, which can consist of several sentences or paragraphs that describe entry contents.
        :param pulumi.Input[str] display_name: Display information such as title and description. A short name to identify the entry,
               for example, "Analytics Data - Jan 2011".
        :param pulumi.Input[str] entry_group: The name of the entry group this entry is in.
        :param pulumi.Input[str] entry_id: The id of the entry to create.
        :param pulumi.Input[dict] gcs_fileset_spec: Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
        :param pulumi.Input[str] linked_resource: The resource this metadata entry refers to.
               For Google Cloud Platform resources, linkedResource is the full name of the resource.
               For example, the linkedResource for a table resource from BigQuery is:
               //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
               Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
               this field is optional and defaults to an empty string.
        :param pulumi.Input[str] schema: Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
               attached to it. See
               https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
               for what fields this schema can contain.
        :param pulumi.Input[str] type: The type of the entry. Only used for Entries with types in the EntryType enum.
               Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
        :param pulumi.Input[str] user_specified_system: This field indicates the entry's source system that Data Catalog does not integrate with.
               userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
               and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        :param pulumi.Input[str] user_specified_type: Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
               When creating an entry, users should check the enum values first, if nothing matches the entry
               to be created, then provide a custom value, for example "my_special_type".
               userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
               numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.

        The **gcs_fileset_spec** object supports the following:

          * `filePatterns` (`pulumi.Input[list]`) - Patterns to identify a set of files in Google Cloud Storage.
            See [Cloud Storage documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)
            for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:
            * gs://bucket_name/dir/*: matches all files within bucket_name/dir directory.
            * gs://bucket_name/dir/**: matches all files in bucket_name/dir spanning all subdirectories.
            * gs://bucket_name/file*: matches files prefixed by file in bucket_name
            * gs://bucket_name/??.txt: matches files with two characters followed by .txt in bucket_name
            * gs://bucket_name/[aeiou].txt: matches files that contain a single vowel character followed by .txt in bucket_name
            * gs://bucket_name/[a-m].txt: matches files that contain a, b, ... or m followed by .txt in bucket_name
            * gs://bucket_name/a/*/b: matches all files in bucket_name that match a/*/b pattern, such as a/c/b, a/d/b
            * gs://another_bucket/a.txt: matches gs://another_bucket/a.txt
          * `sampleGcsFileSpecs` (`pulumi.Input[list]`) - -
            Sample files contained in this fileset, not all files contained in this fileset are represented here.  Structure is documented below.
            * `filePath` (`pulumi.Input[str]`) - -
              The full file path
            * `sizeBytes` (`pulumi.Input[float]`) - -
              The size of the file, in bytes.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['display_name'] = display_name
            if entry_group is None:
                raise TypeError("Missing required property 'entry_group'")
            __props__['entry_group'] = entry_group
            if entry_id is None:
                raise TypeError("Missing required property 'entry_id'")
            __props__['entry_id'] = entry_id
            __props__['gcs_fileset_spec'] = gcs_fileset_spec
            __props__['linked_resource'] = linked_resource
            __props__['schema'] = schema
            __props__['type'] = type
            __props__['user_specified_system'] = user_specified_system
            __props__['user_specified_type'] = user_specified_type
            __props__['bigquery_date_sharded_spec'] = None
            __props__['bigquery_table_spec'] = None
            __props__['integrated_system'] = None
            __props__['name'] = None
        super(Entry, __self__).__init__(
            'gcp:datacatalog/entry:Entry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bigquery_date_sharded_spec=None, bigquery_table_spec=None, description=None, display_name=None, entry_group=None, entry_id=None, gcs_fileset_spec=None, integrated_system=None, linked_resource=None, name=None, schema=None, type=None, user_specified_system=None, user_specified_type=None):
        """
        Get an existing Entry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] bigquery_date_sharded_spec: Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
               https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
        :param pulumi.Input[dict] bigquery_table_spec: Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
        :param pulumi.Input[str] description: Entry description, which can consist of several sentences or paragraphs that describe entry contents.
        :param pulumi.Input[str] display_name: Display information such as title and description. A short name to identify the entry,
               for example, "Analytics Data - Jan 2011".
        :param pulumi.Input[str] entry_group: The name of the entry group this entry is in.
        :param pulumi.Input[str] entry_id: The id of the entry to create.
        :param pulumi.Input[dict] gcs_fileset_spec: Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.  Structure is documented below.
        :param pulumi.Input[str] integrated_system: This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
        :param pulumi.Input[str] linked_resource: The resource this metadata entry refers to.
               For Google Cloud Platform resources, linkedResource is the full name of the resource.
               For example, the linkedResource for a table resource from BigQuery is:
               //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
               Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
               this field is optional and defaults to an empty string.
        :param pulumi.Input[str] name: The Data Catalog resource name of the entry in URL format. Example:
               projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
               child resources may not actually be stored in the location in this name.
        :param pulumi.Input[str] schema: Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
               attached to it. See
               https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
               for what fields this schema can contain.
        :param pulumi.Input[str] type: The type of the entry. Only used for Entries with types in the EntryType enum.
               Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
        :param pulumi.Input[str] user_specified_system: This field indicates the entry's source system that Data Catalog does not integrate with.
               userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
               and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
        :param pulumi.Input[str] user_specified_type: Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
               When creating an entry, users should check the enum values first, if nothing matches the entry
               to be created, then provide a custom value, for example "my_special_type".
               userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
               numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.

        The **bigquery_date_sharded_spec** object supports the following:

          * `dataset` (`pulumi.Input[str]`)
          * `shardCount` (`pulumi.Input[float]`)
          * `tablePrefix` (`pulumi.Input[str]`)

        The **bigquery_table_spec** object supports the following:

          * `tableSourceType` (`pulumi.Input[str]`)
          * `tableSpec` (`pulumi.Input[dict]`)
            * `groupedEntry` (`pulumi.Input[str]`)

          * `viewSpec` (`pulumi.Input[dict]`)
            * `viewQuery` (`pulumi.Input[str]`)

        The **gcs_fileset_spec** object supports the following:

          * `filePatterns` (`pulumi.Input[list]`) - Patterns to identify a set of files in Google Cloud Storage.
            See [Cloud Storage documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)
            for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:
            * gs://bucket_name/dir/*: matches all files within bucket_name/dir directory.
            * gs://bucket_name/dir/**: matches all files in bucket_name/dir spanning all subdirectories.
            * gs://bucket_name/file*: matches files prefixed by file in bucket_name
            * gs://bucket_name/??.txt: matches files with two characters followed by .txt in bucket_name
            * gs://bucket_name/[aeiou].txt: matches files that contain a single vowel character followed by .txt in bucket_name
            * gs://bucket_name/[a-m].txt: matches files that contain a, b, ... or m followed by .txt in bucket_name
            * gs://bucket_name/a/*/b: matches all files in bucket_name that match a/*/b pattern, such as a/c/b, a/d/b
            * gs://another_bucket/a.txt: matches gs://another_bucket/a.txt
          * `sampleGcsFileSpecs` (`pulumi.Input[list]`) - -
            Sample files contained in this fileset, not all files contained in this fileset are represented here.  Structure is documented below.
            * `filePath` (`pulumi.Input[str]`) - -
              The full file path
            * `sizeBytes` (`pulumi.Input[float]`) - -
              The size of the file, in bytes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bigquery_date_sharded_spec"] = bigquery_date_sharded_spec
        __props__["bigquery_table_spec"] = bigquery_table_spec
        __props__["description"] = description
        __props__["display_name"] = display_name
        __props__["entry_group"] = entry_group
        __props__["entry_id"] = entry_id
        __props__["gcs_fileset_spec"] = gcs_fileset_spec
        __props__["integrated_system"] = integrated_system
        __props__["linked_resource"] = linked_resource
        __props__["name"] = name
        __props__["schema"] = schema
        __props__["type"] = type
        __props__["user_specified_system"] = user_specified_system
        __props__["user_specified_type"] = user_specified_type
        return Entry(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop