# ml-deployer

The CI reusable workflow requires that the calling repository contains a makefile with the following targets:
* ci: Runs the repository tests
* lint: Runs the go linter
* helm_lint: Runs helm linter

If the setup input is true, then make target `setup` must exist in the calling repository makefile, to run any non-standard prerequisite setups

## VERSION workflow
The VERSION reusable workflow requires that the calling repository contains a makefile with the following targets:
* bump (or the name specified in the bump_target input): Runs the version bump

## CD workflow
The CD reusable workflow requires that the calling workflow sets the permissions setting with a defined id-token value.

```
permissions:
   id-token: write
```

## Testing

* Create a dev branch off the main branch.
* Create a pull request to merge the dev branch to main.
* Update the workflow with your code changes.
* Push to the dev branch will trigger the ci.yml workflow and the version.yml workflow.
* The cd.yml workflow can be tested manually by deploying the testing branch to dev.
* Merge the dev branch into the main branch.
* Create a new tag for the new commit into the main branch.
