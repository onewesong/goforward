build_linux:
	GOOS=linux go build -o ./dist/ ./cmd/goforward
run_test_forward:
	cd ./cmd/goforward && go run . -f '127.0.0.1:12345->240e:379:1aee:c400:cd6:d551:e383:37a2:22,127.0.0.1:1234->1.1.1.1:443'

test_get_all_forward:
	curl 127.1:5668/api/forward |jq .

test_get_specified_forward:
	curl 127.1:5668/api/forward/127.0.0.1:1234 | jq .

test_add_forward:
	curl 127.1:5668/api/forward -v -H 'Content-Type: application/json' -d '{"forward_link": "127.0.0.1:1233->1.1.1.1:443"}'

test_add_forward_with_override:
	curl 127.1:5668/api/forward -v -H 'Content-Type: application/json' -d '{"forward_link": "127.0.0.1:1233->1.1.1.1:443", "override": true}'

test_del_forward:
	curl 127.1:5668/api/forward/127.0.0.1:1233 -X DELETE
