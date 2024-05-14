install:
	mv CLAI clai
	mkdir -p ${HOME}/.config/clai
	touch ${HOME}/.config/clai/config.yaml

	mkdir -p ${HOME}/.config/clai/chatData
	touch  ${HOME}/.config/clai/chatData/chat.csv

	mv clai ${HOME}/go/bin