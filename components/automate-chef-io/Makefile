# we use pushd/popd here, and /bin/sh of our chefes/buildkite image is not bash
# so we have to override the default shell here
SHELL=bash
SWAGGER_RESULT_FILE=static/api-docs/all-apis.swagger.json
SWAGGER_DIR=data/docs/api_chef_automate
SWAGGER_FILES=$(shell find $(SWAGGER_DIR) -name '*.swagger.json')
STATIC_SWAGGER_FILES=$(shell find data/docs/api-static/ -type f | sort -n)

themes/chef:
	scripts/clone_hugo.sh

clean:
	rm -f $(SWAGGER_RESULT_FILE)
	rm -rf themes/
	rm -rf public/

sync: themes/chef
	pushd themes/chef && git fetch && git reset --hard origin/master && popd

clean_swagger_files:
	rm -rf ./$(SWAGGER_DIR)/*

sync_swagger_files: clean_swagger_files
	pushd ../../api/external && \
		find . -name '*.swagger.json' | while read i ; do \
			mkdir -p "../../components/automate-chef-io/$(SWAGGER_DIR)/`dirname $$i`" && \
			cp $$i "../../components/automate-chef-io/$(SWAGGER_DIR)/$$i" \
		; done && \
	popd

	pushd ../automate-gateway/api/ && \
		find . -name '*.swagger.json' | while read i ; do \
			mkdir -p "../../automate-chef-io/$(SWAGGER_DIR)/`dirname $$i`" && \
			cp $$i "../../automate-chef-io/$(SWAGGER_DIR)/$$i" \
		; done && \
	popd

$(SWAGGER_RESULT_FILE): $(SWAGGER_FILES)

# This jq script is more complicated than it needs to be to workaround an issue where jq treats empty input as valid input:
# https://github.com/stedolan/jq/issues/1142
generate_swagger: $(SWAGGER_RESULT_FILE)
	swagger mixin -c=0 $(STATIC_SWAGGER_FILES) $(SWAGGER_FILES) | \
		jq -es 'if . == [] then null else .[] | .paths = (.paths | with_entries( select( all( .value[].tags[]; . != "hidden") ) ) ) end' \
		> $(SWAGGER_RESULT_FILE)

serve: sync generate_swagger
	hugo server --buildDrafts --noHTTPCache

lint: themes/chef
	hugo -D

build: themes/chef generate_swagger
	hugo

# Consult scripts/spellcheck.sh in the root Automate directory for notes on using cspell.
# This does files in automate-chef-io plus exposed proto files.
# Alternately, you can use the common `make spell` to skip the proto files.
spell:
	@pushd ../.. > /dev/null; \
		cspell \
			"api/external/**/*.proto" \
			"components/automate-chef-io/**/*.md" \
			"components/automate-chef-io/data/docs/api-static/**/*" \
			"components/automate-chef-io/layouts/**/*.html" \
			"components/automate-gateway/**/*.proto" \
		; EXIT_CODE=$$?; \
		popd > /dev/null; \
		exit $$EXIT_CODE
