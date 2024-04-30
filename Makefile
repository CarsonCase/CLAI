install:
	mv CLAI clai
	mkdir -p ${HOME}/.config/clai
	touch ${HOME}/.config/clai/config.yaml
	mv clai ${HOME}/go/bin