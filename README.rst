Format_JSON_Log
===============

*Improve readability of json log in your tty !*


**Get Started:**

> from example_log.json:

```json
{"level":"info","ts":1576167467.5751193,"logger":"leader","msg":"Skipping leader election; not running in a cluster."}
{"level":"info","ts":1576167468.615742,"logger":"controller-runtime.metrics","msg":"metrics server is starting to listen","addr":"0.0.0.0:8383"}
{"level":"error","ts":1576167468.6160917,"logger":"controller-runtime.metrics","msg":"metrics server failed to listen. You may want to disable the metrics server or use another port if it is due to conflicts","error":"error listening on 0.0.0.0:8383: listen tcp 0.0.0.0:8383: bind: address already in use","stacktrace":"github.com/go-logr/zapr.(*zapLogger).Error\n\t/home/igor18/dev/go/k8s_operator_simple/vendor/github.com/go-logr/zapr/zapr.go:128\nsigs.k8s.io/controller-runtime/pkg/metrics.NewListener\n\t/home/igor18/dev/go/k8s_operator_simple/vendor/sigs.k8s.io/controller-runtime/pkg/metrics/listener.go:44\nsigs.k8s.io/controller-runtime/pkg/manager.New\n\t/home/igor18/dev/go/k8s_operator_simple/vendor/sigs.k8s.io/controller-runtime/pkg/manager/manager.go:259\nmain.main\n\t/home/igor18/dev/go/k8s_operator_simple/cmd/manager/main.go:94\nruntime.main\n\t/home/igor18/.go/src/runtime/proc.go:203"}
```

> using ``cat assets/example_log.json |& formatlog``

> get formated output:

  .. image:: https://raw.githubusercontent.com/incubatio/format_json_log/master/assets/formated_log.jpg

wow, much color, such readability !


**Build, Test, Install**: check the Makefile at the root of the project.
