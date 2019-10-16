local echo_run = import "templates/echo-string.jsonnet.TEMPLATE";

// The cluster manager deployment for foocorp.
echo_run + {
  workflowName:: "casek-echo-test",
  sayArgument:: "Nazdar Karle !!!",
}
