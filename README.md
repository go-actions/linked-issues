# linked-issues docker action

This action prints "Hello World" or "Hello" + the name of a person to greet to the log.

## Inputs

## `who-to-greet`

**Required** The name of the person to greet. Default `"World"`.

## Outputs

## `time`

The time we greeted you.

## Example usage

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
