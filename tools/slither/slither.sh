#!/bin/bash

if [ ! -d "contracts" ]; then 
	echo "error: script needs to be run from project root './tools/slither/slither.sh'"
	exit 1
fi

<<<<<<< HEAD
docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/tools/slither trailofbits/eth-security-toolbox@sha256:6376ca0f1e01cfac40499650e3b5c3c430f7c6fee73fcd2ea71aad4d0fa0038b -c 'solc-select 0.6.11 && slither --solc-args="--optimize" --triage-mode ../../contracts'
=======
docker run --rm -v "$PWD":/contracts -it --workdir=/contracts/tools/slither trailofbits/eth-security-toolbox@sha256:e5e2ebbffcc4c3a334063ba3871ec25a4e1dc4915d77166aef1aa265e0a5978f -c 'solc-select 0.6.4 && slither --solc-args="--optimize" --triage-mode ../../contracts'
>>>>>>> 39b44442... Update tools to latest builds and use --optimize flag in slither
