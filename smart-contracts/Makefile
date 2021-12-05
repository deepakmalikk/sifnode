# Makefile

SHELL:=/bin/bash

include Makefile.smartContracts
# smartContracts=contracts/Blocklist.sol ... contracts/BridgeRegistry.sol

include Makefile.artifacts
# artifacts=artifacts/contracts/Blocklist.sol/Blocklist.json ... artifacts/contracts/BridgeRegistry.sol/BridgeRegistry.json

# "Finds whitespace-separated words in text that match pattern and replaces them with replacement.
# Here pattern may contain a ‘%’ which acts as a wildcard, matching any number of any characters within a word.
# If replacement also contains a ‘%’, the ‘%’ is replaced by the text that matched the ‘%’ in pattern. Only the first
# ‘%’ in the pattern and replacement is treated this way; any subsequent ‘%’ is unchanged." (GNU docs)
# So this is: gofiles=artifacts/contracts/Blocklist.sol/Blocklist.go ... artifacts/contracts/BridgeRegistry.sol/BridgeRegistry.go
gofiles=$(artifacts:.json=.go)

all: ${artifacts}

install: all

# x:; echo artifcats: ${artifacts}

artifacts/%.go: artifacts/%.json
	@# scripts/do_abigen.sh accepts two parameters.
	@# The first parameter ("$<") evaluates to the name of a prerequisite, i.e. artifacts/...something....json
	@# Second parameter is literal "../cmd/ebrelayer/contract/generated".
	@# do_abigen.sh will convert .json to .go using jq and abigen
	@# bash scripts/do_abigen.sh $< ../cmd/ebrelayer/contract/generated
	@echo $<
	@echo $@
	bash scripts/do_abigen.sh $< ../cmd/ebrelayer/contract/generated

${artifacts}: .hardhat-compile

all: ${gofiles}

.hardhat-compile: node_modules
	npx hardhat compile --force
	touch $@

node_modules: package.json
	npm install

.PHONY: clean clean-smartcontracts clean-node
clean: clean-node clean-smartcontracts
	rm -f .hardhat-compile

clean-smartcontracts:
	rm -rf build .openzepplin

clean-node:
	rm -rf node_modules
