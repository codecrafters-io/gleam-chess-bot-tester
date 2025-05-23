.PHONY: release build test test_with_bash copy_course_file

current_version_number := $(shell git tag --list "v*" | sort -V | tail -n 1 | cut -c 2-)
next_version_number := $(shell echo $$(($(current_version_number)+1)))

release:
	git tag v$(next_version_number)
	git push origin main v$(next_version_number)

build:
	go build -o dist/main.out ./cmd/tester

test:
	TESTER_DIR=$(shell pwd) go test -v ./internal/

test_and_watch:
	onchange '**/*' -- go test -v ./internal/

copy_course_file:
	hub api \
		repos/codecrafters-io/build-your-own-grep/contents/course-definition.yml \
		| jq -r .content \
		| base64 -d \
		> internal/test_helpers/course_definition.yml

update_tester_utils:
	go get -u github.com/codecrafters-io/tester-utils

test_bot: build
	CODECRAFTERS_REPOSITORY_DIR=./internal/test_helpers/scenarios/test_bot \
	CODECRAFTERS_TEST_CASES_JSON="[ \
		{\"slug\":\"si0\",\"tester_log_prefix\":\"stage_1\",\"title\":\"Stage #2: Respond to move request\"}, \
		{\"slug\":\"xt9\",\"tester_log_prefix\":\"stage_2\",\"title\":\"Stage #3: Respond with an opening\"}, \
		{\"slug\":\"wd4\",\"tester_log_prefix\":\"stage_3\",\"title\":\"Stage #4: Bratko-Kopek test\"}, \
		{\"slug\":\"zz5\",\"tester_log_prefix\":\"stage_4\",\"title\":\"Stage #5: Win at Chess\"}, \
		{\"slug\":\"gd8\",\"tester_log_prefix\":\"stage_5\",\"title\":\"Stage #6: Make a move\"} \
	]" \
	$(shell pwd)/dist/main.out
