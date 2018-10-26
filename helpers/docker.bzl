# image returns the image prefix for the command.
#
# Concretely, image("foo") returns "{STABLE_PROW_REPO}/foo"
# which usually becomes gcr.io/k8s-prow/foo
# (See hack/print-workspace-status.sh)
def prefix(cmd):
  return "{REGISTRY}/{PROJECT}/%s" % cmd[1]

# target returns the image target for the command.
#
# Concretely, target("foo") returns "//prow/cmd/foo:image"
def target(cmd):
  return "//cmd/%s:image" % cmd[0]

def docker_tags(**names):
  outs = {}
  for img, target in names.items():
    outs['%s:{STABLE_VERSION}' % img] = target
    outs['%s:latest' % img] = target
  return outs

def tags(*cmds):
  # Create :version(/commit) :latest tags
  return docker_tags(**{prefix(cmd): target(cmd) for cmd in cmds})