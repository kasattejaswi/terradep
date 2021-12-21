# Versioning design
This document explains the versioning of terradep binary and terraform module and template artefacts.

## Terradep versioning scheme
Terradep will use 3 number versioning scheme. It will be in format MAJOR.MINOR.PATCH.
MAJOR - When breaking changes will be published.
MINOR - When backward compatible changes will be published.
PATCH - When bug fixes with backward compatible changes will be published.

The first version will be 0.1.0
During development, it will be labelled as 0.1.0-alpha
Once all the development is completed, it will labelled as 0.1.0-beta
At the end, when there are minimal bugs, first major version will be released labelled as 1.0.0