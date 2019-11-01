local echo_run1 = import "templates/iterate-echo-strings.jsonnet.TEMPLATE";

// The cluster manager deployment for foocorp.
echo_run1 + {
  workflowName:: "casek-iterate-over-template-names-",
  templateNames:: [{name: "stage1", message: "This is stage one !"},
                   {name: "stage-2", message: "Nazdar Karle !!! FungujeTO!?"},
                   {name: "step14", message: "Dalsi stage?? OK :-D"}
                  ],
}
