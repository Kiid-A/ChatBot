package cmd

type SysCmd interface {
	RunBot() error
}