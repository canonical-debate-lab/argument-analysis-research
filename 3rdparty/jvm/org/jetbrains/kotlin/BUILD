java_library(
    name = "kotlin_reflect",
    visibility = [
        "//visibility:public",
    ],
    exports = [
        ":kotlin_stdlib",
        "//external:jar/org/jetbrains/kotlin/kotlin_reflect",
    ],
)

java_library(
    name = "kotlin_stdlib",
    visibility = [
        "//visibility:public",
    ],
    exports = [
        ":kotlin_stdlib_common",
        "//3rdparty/jvm/org/jetbrains:annotations",
        "//external:jar/org/jetbrains/kotlin/kotlin_stdlib",
    ],
)

java_library(
    name = "kotlin_stdlib_common",
    visibility = [
        "//visibility:public",
    ],
    exports = [
        "//external:jar/org/jetbrains/kotlin/kotlin_stdlib_common",
    ],
)

java_library(
    name = "kotlin_stdlib_jdk7",
    visibility = [
        "//visibility:public",
    ],
    exports = [
        "//external:jar/org/jetbrains/kotlin/kotlin_stdlib_jdk7",
    ],
)

java_library(
    name = "kotlin_stdlib_jdk8",
    visibility = [
        "//visibility:public",
    ],
    exports = [
        ":kotlin_stdlib",
        ":kotlin_stdlib_jdk7",
        "//external:jar/org/jetbrains/kotlin/kotlin_stdlib_jdk8",
    ],
)
