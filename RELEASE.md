# Releasing

This document outlines how to create a release of pluginrpc-go.

1. Clone the repo, ensuring you have the latest main.

2. On a new branch, open [pluginrpc.go](pluginrpc.go) and change the `Version` constant to an
   appropriate [semantic version](https://semver.org/). To select the correct version, look at the
   version number of the [latest release] and the changes that are included in this new release.

- If there are only bug fixes and no new features, remove the `-dev` suffix, set MINOR number to be
  equal to the [latest release], and set the PATCH number to be 1 more than the PATCH number of the
  [latest release].
- If there are features being released, remove the `-dev` suffix, set the MINOR number to be 1 more
  than the MINOR number of the [latest release], and set the PATCH number to `0`. In the common
  case, the diff here will just be to remove the `-dev` suffix.

```patch
-const Version = "1.14.0-dev"
+const Version = "1.14.0"
```

3. Check for any changes in
   [cmd/protoc-gen-pluginrpc-go/main.go](cmd/protoc-gen-pluginrpc-go/main.go) that require a version
   restriction. A constant `IsAtLeastVersionX_Y_Z` should be defined in [pluginrpc.go](pluginrpc.go)
   if generated code has begun to use a new API. Make sure the generated code references this
   constant. If a new constant has been added since the last release, ensure that the name of the
   constant matches the version being released.

4. Open a PR titled "Prepare for vX.Y.Z" and a description tagging all current maintainers. Once
   it's reviewed and CI passes, merge it.

   _Make sure no new commits are merged until the release is complete._

5. Review all commits in the new release and for each PR check an appropriate label is used and edit
   the title to be meaningful to end users.

6. Using the Github UI, create a new release.

   - Under “Choose a tag”, type in “vX.Y.Z” to create a new tag for the release upon publish.
   - Target the main branch.
   - Title the Release “vX.Y.Z”.
   - Click “set as latest release”.
   - Set the last version as the “Previous tag”.
   - Edit the release notes.

7. Publish the release.

8. On a new branch, open [pluginrpc.go](pluginrpc.go) and change the `Version` to increment the
   minor tag and append the `-dev` suffix. Use the next minor release - we never anticipate bugs and
   patch releases.

   ```patch
   -const Version = "1.14.0"
   +const Version = "1.15.0-dev"
   ```

9. Open a PR titled "Back to development" Once it's reviewed and CI passes, merge it.

[latest release]: https://github.com/pluginrpc/pluginrpc-go/releases/latest
