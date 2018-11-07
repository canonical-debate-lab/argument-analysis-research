# Do not edit. bazel-deps autogenerates this file from dependencies.yaml.
def _jar_artifact_impl(ctx):
    jar_name = "%s.jar" % ctx.name
    ctx.download(
        output=ctx.path("jar/%s" % jar_name),
        url=ctx.attr.urls,
        sha256=ctx.attr.sha256,
        executable=False
    )
    src_name="%s-sources.jar" % ctx.name
    srcjar_attr=""
    has_sources = len(ctx.attr.src_urls) != 0
    if has_sources:
        ctx.download(
            output=ctx.path("jar/%s" % src_name),
            url=ctx.attr.src_urls,
            sha256=ctx.attr.src_sha256,
            executable=False
        )
        srcjar_attr ='\n    srcjar = ":%s",' % src_name

    build_file_contents = """
package(default_visibility = ['//visibility:public'])
java_import(
    name = 'jar',
    tags = ['maven_coordinates={artifact}'],
    jars = ['{jar_name}'],{srcjar_attr}
)
filegroup(
    name = 'file',
    srcs = [
        '{jar_name}',
        '{src_name}'
    ],
    visibility = ['//visibility:public']
)\n""".format(artifact = ctx.attr.artifact, jar_name = jar_name, src_name = src_name, srcjar_attr = srcjar_attr)
    ctx.file(ctx.path("jar/BUILD"), build_file_contents, False)
    return None

jar_artifact = repository_rule(
    attrs = {
        "artifact": attr.string(mandatory = True),
        "sha256": attr.string(mandatory = True),
        "urls": attr.string_list(mandatory = True),
        "src_sha256": attr.string(mandatory = False, default=""),
        "src_urls": attr.string_list(mandatory = False, default=[]),
    },
    implementation = _jar_artifact_impl
)

def jar_artifact_callback(hash):
    src_urls = []
    src_sha256 = ""
    source=hash.get("source", None)
    if source != None:
        src_urls = [source["url"]]
        src_sha256 = source["sha256"]
    jar_artifact(
        artifact = hash["artifact"],
        name = hash["name"],
        urls = [hash["url"]],
        sha256 = hash["sha256"],
        src_urls = src_urls,
        src_sha256 = src_sha256
    )
    native.bind(name = hash["bind"], actual = hash["actual"])


def list_dependencies():
    return [
    {"artifact": "com.fasterxml.jackson.core:jackson-annotations:2.9.0", "lang": "java", "sha1": "07c10d545325e3a6e72e06381afe469fd40eb701", "sha256": "45d32ac61ef8a744b464c54c2b3414be571016dd46bfc2bec226761cf7ae457a", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/com/fasterxml/jackson/core/jackson-annotations/2.9.0/jackson-annotations-2.9.0.jar", "name": "com_fasterxml_jackson_core_jackson_annotations", "actual": "@com_fasterxml_jackson_core_jackson_annotations//jar", "bind": "jar/com/fasterxml/jackson/core/jackson_annotations"},
    {"artifact": "com.fasterxml.jackson.core:jackson-core:2.9.7", "lang": "java", "sha1": "4b7f0e0dc527fab032e9800ed231080fdc3ac015", "sha256": "9e5bc0efabd9f0cac5c1fdd9ae35b16332ed22a0ee19a356de370a18a8cb6c84", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/com/fasterxml/jackson/core/jackson-core/2.9.7/jackson-core-2.9.7.jar", "name": "com_fasterxml_jackson_core_jackson_core", "actual": "@com_fasterxml_jackson_core_jackson_core//jar", "bind": "jar/com/fasterxml/jackson/core/jackson_core"},
    {"artifact": "com.fasterxml.jackson.core:jackson-databind:2.9.7", "lang": "java", "sha1": "e6faad47abd3179666e89068485a1b88a195ceb7", "sha256": "675376decfc070b039d2be773a97002f1ee1e1346d95bd99feee0d56683a92bf", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/com/fasterxml/jackson/core/jackson-databind/2.9.7/jackson-databind-2.9.7.jar", "name": "com_fasterxml_jackson_core_jackson_databind", "actual": "@com_fasterxml_jackson_core_jackson_databind//jar", "bind": "jar/com/fasterxml/jackson/core/jackson_databind"},
    {"artifact": "com.fasterxml.jackson.module:jackson-module-kotlin:2.9.7", "lang": "kotlin", "sha1": "9ec9b84e8af4c4f31efcbc5c21e34da8021419f1", "sha256": "5b313b299717156ee883ef37774f709c8c9942b395edcc1d13368e52a786be28", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/com/fasterxml/jackson/module/jackson-module-kotlin/2.9.7/jackson-module-kotlin-2.9.7.jar", "name": "com_fasterxml_jackson_module_jackson_module_kotlin", "actual": "@com_fasterxml_jackson_module_jackson_module_kotlin//jar:file", "bind": "jar/com/fasterxml/jackson/module/jackson_module_kotlin"},
    {"artifact": "io.javalin:javalin:2.3.0", "lang": "kotlin", "sha1": "73836e9cf29f978e47817584f9cee86b5e1f4c09", "sha256": "3571e83863e1f163854f1b2ee3cbfc1336fcbdfa595ec9c2ed8ab8bfa792e5f4", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/io/javalin/javalin/2.3.0/javalin-2.3.0.jar", "name": "io_javalin_javalin", "actual": "@io_javalin_javalin//jar:file", "bind": "jar/io/javalin/javalin"},
    {"artifact": "javax.servlet:javax.servlet-api:3.1.0", "lang": "java", "sha1": "3cd63d075497751784b2fa84be59432f4905bf7c", "sha256": "af456b2dd41c4e82cf54f3e743bc678973d9fe35bd4d3071fa05c7e5333b8482", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/javax/servlet/javax.servlet-api/3.1.0/javax.servlet-api-3.1.0.jar", "name": "javax_servlet_javax_servlet_api", "actual": "@javax_servlet_javax_servlet_api//jar", "bind": "jar/javax/servlet/javax_servlet_api"},
    {"artifact": "org.eclipse.jetty.websocket:websocket-api:9.4.12.v20180830", "lang": "java", "sha1": "97d6376f70ae6c01112325c5254e566af118bc75", "sha256": "6f7ecb42601058ffe4a6c19c5340cac3ebf0f83e2e252b457558f104238278e3", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/websocket/websocket-api/9.4.12.v20180830/websocket-api-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_websocket_websocket_api", "actual": "@org_eclipse_jetty_websocket_websocket_api//jar", "bind": "jar/org/eclipse/jetty/websocket/websocket_api"},
    {"artifact": "org.eclipse.jetty.websocket:websocket-client:9.4.12.v20180830", "lang": "java", "sha1": "75880b6a90a6eda83fdbfc20a42f23eade4b975d", "sha256": "97c6882c858a75776773eaccc01739757c4e9f60a51613878c1f2b2ba03d91af", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/websocket/websocket-client/9.4.12.v20180830/websocket-client-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_websocket_websocket_client", "actual": "@org_eclipse_jetty_websocket_websocket_client//jar", "bind": "jar/org/eclipse/jetty/websocket/websocket_client"},
    {"artifact": "org.eclipse.jetty.websocket:websocket-common:9.4.12.v20180830", "lang": "java", "sha1": "33997cdafbabb3ffd6947a5c33057f967e10535b", "sha256": "3c35aefa720c51e09532c16fdbfaaebd1af3e07dee699dacaba8e0ab0adf88e5", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/websocket/websocket-common/9.4.12.v20180830/websocket-common-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_websocket_websocket_common", "actual": "@org_eclipse_jetty_websocket_websocket_common//jar", "bind": "jar/org/eclipse/jetty/websocket/websocket_common"},
    {"artifact": "org.eclipse.jetty.websocket:websocket-server:9.4.12.v20180830", "lang": "java", "sha1": "fadf609aec6026cb25f25b6bc0b979821f849fd7", "sha256": "7b1bd39006be8c32d7426a119567d860b3e4a3dc3c01a5c91326450bb0213a03", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/websocket/websocket-server/9.4.12.v20180830/websocket-server-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_websocket_websocket_server", "actual": "@org_eclipse_jetty_websocket_websocket_server//jar", "bind": "jar/org/eclipse/jetty/websocket/websocket_server"},
    {"artifact": "org.eclipse.jetty.websocket:websocket-servlet:9.4.12.v20180830", "lang": "java", "sha1": "8d212616b6ea21b96152ff202c2f53fdca8b8b53", "sha256": "8d43e0882759ecd093bd1a5a0ef2b4db38ac279212488a34edb8d7de7c45cc4d", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/websocket/websocket-servlet/9.4.12.v20180830/websocket-servlet-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_websocket_websocket_servlet", "actual": "@org_eclipse_jetty_websocket_websocket_servlet//jar", "bind": "jar/org/eclipse/jetty/websocket/websocket_servlet"},
    {"artifact": "org.eclipse.jetty:jetty-client:9.4.12.v20180830", "lang": "java", "sha1": "1d329d68f31dce13135243c06013aaf6f708f7e7", "sha256": "62efbbfda88cd4f7644242c4b4df8f3b0a671bfeafea7682dabe00352ba07db7", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-client/9.4.12.v20180830/jetty-client-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_client", "actual": "@org_eclipse_jetty_jetty_client//jar", "bind": "jar/org/eclipse/jetty/jetty_client"},
    {"artifact": "org.eclipse.jetty:jetty-http:9.4.12.v20180830", "lang": "java", "sha1": "1341796dde4e16df69bca83f3e87688ba2e7d703", "sha256": "20547da653be9942cc63f57e632a732608559aebde69753bc7312cfe16e8d9c0", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-http/9.4.12.v20180830/jetty-http-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_http", "actual": "@org_eclipse_jetty_jetty_http//jar", "bind": "jar/org/eclipse/jetty/jetty_http"},
    {"artifact": "org.eclipse.jetty:jetty-io:9.4.12.v20180830", "lang": "java", "sha1": "e93f5adaa35a9a6a85ba130f589c5305c6ecc9e3", "sha256": "ab1784abbb9e0ed0869ab6568fe46f1faa79fb5e948cf96450daecd9d27ba1db", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-io/9.4.12.v20180830/jetty-io-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_io", "actual": "@org_eclipse_jetty_jetty_io//jar", "bind": "jar/org/eclipse/jetty/jetty_io"},
    {"artifact": "org.eclipse.jetty:jetty-security:9.4.12.v20180830", "lang": "java", "sha1": "299e0602a9c0b753ba232cc1c1dda72ddd9addcf", "sha256": "513184970c785ac830424a9c62c2fadfa77a630f44aa0bdd792f00aaa092887e", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-security/9.4.12.v20180830/jetty-security-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_security", "actual": "@org_eclipse_jetty_jetty_security//jar", "bind": "jar/org/eclipse/jetty/jetty_security"},
    {"artifact": "org.eclipse.jetty:jetty-server:9.4.12.v20180830", "lang": "java", "sha1": "b0f25df0d32a445fd07d5f16fff1411c16b888fa", "sha256": "4833644e5c5a09bbddc85f75c53e0c8ed750de120ba248fffd8508028528252d", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-server/9.4.12.v20180830/jetty-server-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_server", "actual": "@org_eclipse_jetty_jetty_server//jar", "bind": "jar/org/eclipse/jetty/jetty_server"},
    {"artifact": "org.eclipse.jetty:jetty-servlet:9.4.12.v20180830", "lang": "java", "sha1": "4c1149328eda9fa39a274262042420f66d9ffd5f", "sha256": "7310d4cccf8abf27fde0c3f1a32e19c75fe33c6f1ab558f0704d915f0f01cb07", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-servlet/9.4.12.v20180830/jetty-servlet-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_servlet", "actual": "@org_eclipse_jetty_jetty_servlet//jar", "bind": "jar/org/eclipse/jetty/jetty_servlet"},
    {"artifact": "org.eclipse.jetty:jetty-util:9.4.12.v20180830", "lang": "java", "sha1": "cb4ccec9bd1fe4b10a04a0fb25d7053c1050188a", "sha256": "60ad53e118a3e7d10418b155b9944d90b2e4e4c732e53ef4f419473288d3f48c", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-util/9.4.12.v20180830/jetty-util-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_util", "actual": "@org_eclipse_jetty_jetty_util//jar", "bind": "jar/org/eclipse/jetty/jetty_util"},
    {"artifact": "org.eclipse.jetty:jetty-webapp:9.4.12.v20180830", "lang": "java", "sha1": "a3e119df2da04fcf5aa290c8c35c5b310ce2dcd1", "sha256": "5301e412a32bf7dddcfad458d952179597c61f8fd531c265873562725c3d4646", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-webapp/9.4.12.v20180830/jetty-webapp-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_webapp", "actual": "@org_eclipse_jetty_jetty_webapp//jar", "bind": "jar/org/eclipse/jetty/jetty_webapp"},
    {"artifact": "org.eclipse.jetty:jetty-xml:9.4.12.v20180830", "lang": "java", "sha1": "e9f1874e9b5edd498f2fe7cd0904405da07cc300", "sha256": "5b8298ab3d43ddaf0941d41f51b82c8ae23a247da055fa161b752ab9495155ed", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/eclipse/jetty/jetty-xml/9.4.12.v20180830/jetty-xml-9.4.12.v20180830.jar", "name": "org_eclipse_jetty_jetty_xml", "actual": "@org_eclipse_jetty_jetty_xml//jar", "bind": "jar/org/eclipse/jetty/jetty_xml"},
    {"artifact": "org.jetbrains.kotlin:kotlin-reflect:1.2.51", "lang": "java", "sha1": "36b719a7b84452dd13eeec979d8c82bfb765c57d", "sha256": "129f42c1ad5c3958856ecf2b2dadcd76e24a0b9b7f85aa2aba383616fcc49c7d", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/jetbrains/kotlin/kotlin-reflect/1.2.51/kotlin-reflect-1.2.51.jar", "name": "org_jetbrains_kotlin_kotlin_reflect", "actual": "@org_jetbrains_kotlin_kotlin_reflect//jar", "bind": "jar/org/jetbrains/kotlin/kotlin_reflect"},
    {"artifact": "org.jetbrains.kotlin:kotlin-stdlib-common:1.2.71", "lang": "java", "sha1": "ba18ca1aa0e40eb6f1865b324af2f4cbb691c1ec", "sha256": "63999687ff2fce8a592dd180ffbbf8f1d21c26b4044c55cdc74ff3cf3b3cf328", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib-common/1.2.71/kotlin-stdlib-common-1.2.71.jar", "name": "org_jetbrains_kotlin_kotlin_stdlib_common", "actual": "@org_jetbrains_kotlin_kotlin_stdlib_common//jar", "bind": "jar/org/jetbrains/kotlin/kotlin_stdlib_common"},
    {"artifact": "org.jetbrains.kotlin:kotlin-stdlib-jdk7:1.2.71", "lang": "java", "sha1": "4ce93f539e2133f172f1167291a911f83400a5d0", "sha256": "b136bd61b240e07d4d92ce00d3bd1dbf584400a7bf5f220c2f3cd22446858082", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib-jdk7/1.2.71/kotlin-stdlib-jdk7-1.2.71.jar", "name": "org_jetbrains_kotlin_kotlin_stdlib_jdk7", "actual": "@org_jetbrains_kotlin_kotlin_stdlib_jdk7//jar", "bind": "jar/org/jetbrains/kotlin/kotlin_stdlib_jdk7"},
    {"artifact": "org.jetbrains.kotlin:kotlin-stdlib-jdk8:1.2.71", "lang": "java", "sha1": "5470d1f752cd342edb77e1062bac07e838d2cea4", "sha256": "ac3c8abf47790b64b4f7e2509a53f0c145e061ac1612a597520535d199946ea9", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib-jdk8/1.2.71/kotlin-stdlib-jdk8-1.2.71.jar", "name": "org_jetbrains_kotlin_kotlin_stdlib_jdk8", "actual": "@org_jetbrains_kotlin_kotlin_stdlib_jdk8//jar", "bind": "jar/org/jetbrains/kotlin/kotlin_stdlib_jdk8"},
# duplicates in org.jetbrains.kotlin:kotlin-stdlib promoted to 1.2.71
# - org.jetbrains.kotlin:kotlin-reflect:1.2.51 wanted version 1.2.51
# - org.jetbrains.kotlin:kotlin-stdlib-jdk8:1.2.71 wanted version 1.2.71
    {"artifact": "org.jetbrains.kotlin:kotlin-stdlib:1.2.71", "lang": "java", "sha1": "d9717625bb3c731561251f8dd2c67a1011d6764c", "sha256": "4c895c270b87f5fec2a2796e1d89c15407ee821de961527c28588bb46afbc68b", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib/1.2.71/kotlin-stdlib-1.2.71.jar", "name": "org_jetbrains_kotlin_kotlin_stdlib", "actual": "@org_jetbrains_kotlin_kotlin_stdlib//jar", "bind": "jar/org/jetbrains/kotlin/kotlin_stdlib"},
    {"artifact": "org.jetbrains:annotations:13.0", "lang": "java", "sha1": "919f0dfe192fb4e063e7dacadee7f8bb9a2672a9", "sha256": "ace2a10dc8e2d5fd34925ecac03e4988b2c0f851650c94b8cef49ba1bd111478", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/jetbrains/annotations/13.0/annotations-13.0.jar", "name": "org_jetbrains_annotations", "actual": "@org_jetbrains_annotations//jar", "bind": "jar/org/jetbrains/annotations"},
    {"artifact": "org.slf4j:slf4j-api:1.7.25", "lang": "java", "sha1": "da76ca59f6a57ee3102f8f9bd9cee742973efa8a", "sha256": "18c4a0095d5c1da6b817592e767bb23d29dd2f560ad74df75ff3961dbde25b79", "repository": "http://central.maven.org/maven2/", "url": "http://central.maven.org/maven2/org/slf4j/slf4j-api/1.7.25/slf4j-api-1.7.25.jar", "name": "org_slf4j_slf4j_api", "actual": "@org_slf4j_slf4j_api//jar", "bind": "jar/org/slf4j/slf4j_api"},
    ]

def maven_dependencies(callback = jar_artifact_callback):
    for hash in list_dependencies():
        callback(hash)
