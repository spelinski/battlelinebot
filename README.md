# README #

install golang.

pull down this repo and then set GOPATH and GOBIN to the root directory where you see the src folder (I.E. mine inside .bashrc)

export GOPATH=$HOME/battlelinebot

export GOBIN=$GOPATH/bin

inside of the src folder run "go build" and it should spit out an executable named src

This file can be ran with ./src by the battleline engine i.e.

~/battlelineaiengine$ python Battleline.py --player1-cmd ./src --player1-workdir ../battlelinebot/src