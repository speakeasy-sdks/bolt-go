# Identifier


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `IdentifierType`                                                                         | [components.IdentifierType](../../models/components/identifiertype.md)                   | :heavy_check_mark:                                                                       | The type of identifier                                                                   | email                                                                                    |
| `IdentifierValue`                                                                        | *string*                                                                                 | :heavy_check_mark:                                                                       | The value of the identifier. The value must be valid for the specified `identifier_type` | alice@example.com                                                                        |