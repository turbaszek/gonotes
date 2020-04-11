package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

// CompleteCommand sets autocomplete for GoNotes
func CompleteCommand() *cli.Command {
	return &cli.Command{
		Name:      "complete",
		Usage:     "Setup autocomplete",
		ArgsUsage: "SHELL",
		Hidden:    true,
		Action: func(c *cli.Context) error {
			shell := c.Args().First()
			switch shell {
			case "bash":
				fmt.Println(bash)
			case "zsh":
				fmt.Println(zsh)
			default:
				return fmt.Errorf("shell %s not supported. Currently only only zsh and bash are supported :<", shell)
			}
			return nil
		},
	}
}

const bash = `
#! /bin/bash

: ${PROG:=$(basename ${BASH_SOURCE})}

_cli_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}

complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete $PROG
unset PROG`

const zsh = `
#compdef gonotes

export _CLI_ZSH_AUTOCOMPLETE_HACK=1

_cli_zsh_autocomplete() {

  local -a opts
  local cur
  cur=${words[-1]}
  if [[ "$cur" == "-"* ]]; then
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
  else
    opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
  fi

  if [[ "${opts[1]}" != "" ]]; then
    _describe 'values' opts
  fi

  return
}
compdef _cli_zsh_autocomplete gonotes`
