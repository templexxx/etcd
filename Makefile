# run from repository root

.PHONY: build
build:
	./build.sh
	./bin/etcd --version
	./bin/etcdctl version

clean:
	rm -f ./codecov
	rm -rf ./agent-*
	rm -rf ./covdir
	rm -f ./*.coverprofile
	rm -f ./*.log
	rm -f ./bin/Dockerfile-release
	rm -rf ./bin/*.etcd
	rm -rf ./default.etcd
	rm -rf ./tests/e2e/default.etcd
	rm -rf ./gopath
	rm -rf ./gopath.proto
	rm -rf ./release
	rm -f ./snapshot/localhost:*
	rm -f ./tools/etcd-dump-metrics/localhost:*
	rm -f ./integration/127.0.0.1:* ./integration/localhost:*
	rm -f ./clientv3/integration/127.0.0.1:* ./clientv3/integration/localhost:*
	rm -f ./clientv3/ordering/127.0.0.1:* ./clientv3/ordering/localhost:*