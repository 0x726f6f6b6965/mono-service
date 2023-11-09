# workspace(name = "mono-service")

# load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ############################################################
# # rule go
# ############################################################
# http_archive(
#     name = "io_bazel_rules_go",
#     sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
#     urls = [
#         "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
#         "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
#     ],
# )

# ############################################################
# # gazelle
# ############################################################
# http_archive(
#     name = "bazel_gazelle",
#     sha256 = "d3fa66a39028e97d76f9e2db8f1b0c11c099e8e01bf363a923074784e451f809",
#     urls = [
#         "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
#         "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
#     ],
# )

# ############################################################
# # protobuf
# ############################################################
# http_archive(
#     name = "com_google_protobuf",
#     strip_prefix = "protobuf-25.0",
#     urls = [
#         "https://github.com/protocolbuffers/protobuf/archive/v25.0.tar.gz",
#         "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v25.0.tar.gz",
#     ],
# )


# ############################################################
# # pkg
# ############################################################
# http_archive(
#     name = "rules_pkg",
#     urls = [
#         "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
#         "https://github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
#     ],
#     sha256 = "8f9ee2dc10c1ae514ee599a8b42ed99fa262b757058f65ad3c384289ff70c4b8",
# )

# ############################################################
# # oci
# ############################################################
# http_archive(
#     name = "rules_oci",
#     sha256 = "21a7d14f6ddfcb8ca7c5fc9ffa667c937ce4622c7d2b3e17aea1ffbc90c96bed",
#     strip_prefix = "rules_oci-1.4.0",
#     url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.4.0/rules_oci-v1.4.0.tar.gz",
# )


# load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
# load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
# load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
# load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")
# load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")

# ############################################################
# # Define your own dependencies here using go_repository.
# # Else, dependencies declared by rules_go/gazelle will be used.
# # The first declaration of an external repository "wins".
# ############################################################

# load("//:deps.bzl", "go_dependencies")

# # gazelle:repository_macro deps.bzl%go_dependencies
# go_dependencies()

# go_rules_dependencies()

# go_register_toolchains(version = "1.21.3")

# gazelle_dependencies()

# protobuf_deps()

# rules_pkg_dependencies()

# rules_oci_dependencies()

# load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "LATEST_ZOT_VERSION", "oci_register_toolchains")

# oci_register_toolchains(
#     name = "oci",
#     crane_version = LATEST_CRANE_VERSION,
#     # Uncommenting the zot toolchain will cause it to be used instead of crane for some tasks.
#     # Note that it does not support docker-format images.
#     # zot_version = LATEST_ZOT_VERSION,
# )

# # You can pull your base images using oci_pull like this:
# load("@rules_oci//oci:pull.bzl", "oci_pull")

# oci_pull(
#     name = "distroless_base",
#     digest = "sha256:ccaef5ee2f1850270d453fdf700a5392534f8d1a8ca2acda391fbb6a06b81c86",
#     image = "gcr.io/distroless/base",
#     platforms = [
#         "linux/amd64",
#         "linux/arm64",
#     ],
# )