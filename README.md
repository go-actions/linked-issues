# Linked Issue (Github Action)

This action find the Issues linked in a Pull Request. It parses the HTML of the PR page to find the linked issues.

## Inputs

The action has the following inputs:

| Name       | Description                                                      | Type       | Possible Values                             | Default Values |
| ---------- | ---------------------------------------------------------------- | ---------- | ------------------------------------------- | -------------- |
| `pr_url`   | URL of the Pull Request                                          | `Required` | Any valid PR URL                            | `""`           |
| `tag`      | HTML tag that contains the linked Issue URL                      | `Optional` | Any HTML Tag                                | `form`         |
| `attr_key` | Attribute key that will be used to select the desired HTML tag   | `Optional` | Any valid HTML tag attribute                | `aria-label`   |
| `attr_val` | Attribute value that will be used to select the desired HTML tag | `Optional` | Any text                                    | `Link issues`  |
| `format`   | Output format for the linked Issues                              | `Optional` | `IssueNumber`,`IssueURL`,`ExternalIssueRef` | `IssueNumber`  |

## Outputs

The action has the following output:
| Name     | Description                                                                 |
| -------- | --------------------------------------------------------------------------- |
| `issues` | List of issues separated by space and formatted according to `format` input |

For example, if your PR has the following issue linked:

- https://github.com/foo/bar/issues/1
- https://github.com/foo/bar/issues/2
- https://github.com/foo/bar/issues/3

The output of this action will be the following for different formats:

**IssueNumber:**
`1 2 3`

**IssueURL:**

`https://github.com/foo/bar/issues/1 https://github.com/foo/bar/issues/2 https://github.com/foo/bar/issues/3`

**ExternalIssueRef:**

`foo/bar#1 foo/bar#2 foo/bar#3`

## Example usage

Here, is a sample workflow YAML showing how to use this action.

```yaml
on: [pull_request]

jobs:
  linked_issues:
    runs-on: ubuntu-latest
    name: A job to say hello
    steps:
      - name: Find Linked Issues
        id: links
        uses: hossainemruz/linked-issues@main
        with:
          pr_url: ${{github.event.pull_request.html_url}}
          format: IssueNumber

      - name: Output linked Issue list
        run: echo "${{ steps.links.outputs.issues }}"
```

A more practical use of this action can be found in [this workflow](https://github.com/hugo-toha/toha/blob/main/.github/workflows/project-automation-pr.yaml).
