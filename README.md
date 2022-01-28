# ml-deployer

## CI workflow
The CI reusable workflow requires that the calling repository contains a makefile with the following targets:
* ci: Runs the repository tests
* lint: Runs the go linter
* helm_lint: Runs helm linter

If the setup input is true, then make target `setup` must exist in the calling repository makefile, to run any non-standard prerequisite setups

## VERSION workflow
The VERSION reusable workflow requires that the calling repository contains a makefile with the following targets:
* bump (or the name specified in the bump_target input): Runs the version bump

## CD workflow
The CD reusable workflow requires that the calling workflow sets the permissions setting with a defined id-token and contents value below:

```
permissions:
  id-token: write
  contents: read
```
When the CD workflow is used with deploy set to false, it sets the ref for the checkout action: https://github.com/actions/checkout to the main branch by default. This default can be overidden by setting the branch input.

## Developing

* Delete any existing dev branch.
* Create a dev branch off the main branch.
* Create a pull request to merge the dev branch to main.
* Update the workflow with your code changes.

## Testing
* Push to the dev branch will trigger and test the **ci.yml** workflow and the **version.yml** workflow.
* Test the **version.yml** workflow fully by adding [bugfix] or [feature] to your commit messages.
* The **cd.yml** workflow must be tested manually by deploying the dev branch to dev environment.
* The release (deployment on merge to main) process can be tested by merging the dev branch to main with [bugfix] or [feature] in the merge commit message.
* Note to add the #none or #minor on merging to main to determine how the reusable workflow tag should be bumped.
* Merge the dev branch into the main branch.
