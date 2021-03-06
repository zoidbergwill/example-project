# Go Example "Complex" Project

Example project so I can play with Bazel.

Eventually I wanna make this a microservice architecture, with a couple Go microservices, Docker images, Kubernetes manifests, and Helm charts.

![Diagram of the Dependency Graph](./mygraph.png)

## Useful Links

TODO(zoidbergwill): document how and why we use each of these.

### Bazel Libraries and Tools

#### Go Rules

- [bazelbuild/rules_go](https://github.com/bazelbuild/rules_go) - Bazel rules for Go
- [bazelbuild/bazel-gazelle](https://github.com/bazelbuild/bazel-gazelle) - gazelle is an awesome tool for working with Go projects in Bazel. It can be extended to support other languages, but is awesome because it generates Bazel BUILD files for you.

#### Javascript Rules

- [bazelbuild/rules_nodejs](https://github.com/bazelbuild/rules_nodejs/) - Bazel rules for Node.js
- [Building Javascript and Typescript Outputs (Offical Docs)](https://docs.bazel.build/versions/master/build-javascript.html)

- [Full Stack Development with React and Bazel](https://www.syntaxsuccess.com/viewarticle/full-stack-development-with-react-and-bazel)
- [Scalable React Build with Bazel](https://www.syntaxsuccess.com/viewarticle/scalable-react-build-with-bazel)

#### Docker and Kubernetes Rules

- [bazelbuild/rules_docker](https://github.com/bazelbuild/rules_docker) - Bazel rules for creating Docker images
- [bazelbuild/rules_k8s](https://github.com/bazelbuild/rules_k8s) - Bazel rules for Kubernetes manifests
- [tmc/rules_helm](https://github.com/tmc/rules_helm) - Bazel rules for Helm

#### Remote Execution Services:

> I assumed these were all paid for tools from the docs, but they're actually all open source, it seems.

- [bazelbuild/bazel-buildfarm](https://github.com/bazelbuild/bazel-buildfarm) - WIP official'ish Bazel remote caching and execution service
- [buildgrid/buildgrid](https://gitlab.com/BuildGrid/buildgrid)
- [twitter/scoot](https://github.com/twitter/scoot)
- [buildbarn](https://github.com/buildbarn)

- [Adapting Bazel Rules for Remote Execution](https://docs.bazel.build/versions/master/remote-execution-rules.html)

#### Remote Caching

- [Remote Caching](https://docs.bazel.build/versions/master/remote-caching.html)

- [Unofficial Bazel Remote Cache API](https://github.com/buchgr/bazel-remote/)

### Alternatives to Bazel

- [BuildStream](https://buildstream.build/) - The GNOME projects bazel-inspired alternative?
- Buck
- Pants

### Bazel Blog Posts

[JML's Bazel Correct Reproducible Fast Builds](https://jml.io/2015/07/bazel-correct-reproducible-fast-builds.html)

[BrainTreePayments's Migrating from Gradle to Bazel](https://www.braintreepayments.com/blog/migrating-from-gradle-to-bazel/)

## Pretty things

```
# bazel test '...'

INFO: Analyzed 15 targets (10 packages loaded, 117 targets configured).
INFO: Found 10 targets and 5 test targets...
INFO: Elapsed time: 2.218s, Critical Path: 0.36s
INFO: 10 processes: 10 darwin-sandbox.
INFO: Build completed successfully, 14 total actions
//app-a/cmd/app:go_default_test                                          PASSED in 0.1s
//lib-common/pkg/common:go_default_test                                  PASSED in 0.2s
//lib-d/pkg/d:go_default_test                                            PASSED in 0.1s
//lib-e/pkg/e:go_default_test                                            PASSED in 0.2s
//lib-f/pkg/f:go_default_test                                            PASSED in 0.1s

Executed 5 out of 5 tests: 5 tests pass.
There were tests whose specified size is too big. Use the --test_verbose_timeout_warnings command line option to see wINFO: Build completed successfully, 14 total actions
```

After editing app:

```
# bazel test '...'
INFO: Analyzed 15 targets (1 packages loaded, 5 targets configured).
INFO: Found 10 targets and 5 test targets...
INFO: Elapsed time: 1.632s, Critical Path: 1.39s
INFO: 6 processes: 6 darwin-sandbox.
INFO: Build completed successfully, 7 total actions
//lib-common/pkg/common:go_default_test                         (cached) PASSED in 0.2s
//lib-d/pkg/d:go_default_test                                   (cached) PASSED in 0.1s
//lib-e/pkg/e:go_default_test                                   (cached) PASSED in 0.2s
//lib-f/pkg/f:go_default_test                                   (cached) PASSED in 0.1s
//app-a/cmd/app:a_test                                                   PASSED in 0.5s

Executed 1 out of 5 tests: 5 tests pass.
There were tests whose specified size is too big. Use the --test_verbose_timeout_warnings command line option to see wINFO: Build completed successfully, 7 total actions
```
