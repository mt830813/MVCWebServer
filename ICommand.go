package main

type ICommand interface {
	DoCommand(params string)
}
