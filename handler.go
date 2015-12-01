package rainbot

import (
    "strings"
    "net/rpc"

    "github.com/sorcix/irc"
)

type Command struct {
    Help   string
    Fun    CommandFun
    PM     bool
    CM     bool
}

type Trigger struct {
    Check TriggerFun
    Fun   Listener
}

type CommandName string
type ModuleName  string
type Numeric     int

type Handler struct {
    Commands map[CommandName]ModuleName
    //Commands map[]
    NumericHooks map[ModuleName][]Numeric

    InternalCommands map[CommandName]*Command
    InternalTriggers map[string]*Trigger

    Modules map[ModuleName]*rpc.Client
}

type IrcData struct {
    Event Event
    Msg   *irc.Message
}

type CommandData struct {
    Msg  *irc.Message
    Name CommandName
    Args []string
}

func NewHandler() *Handler {
    handler := &Handler{
        Commands: make(map[CommandName]ModuleName),
        NumericHooks: make(map[ModuleName][]Numeric),

        InternalCommands: make(map[CommandName]*Command),
        InternalTriggers: make(map[string]*Trigger),

        Modules: make(map[ModuleName]*rpc.Client),
    }
    return handler
}

func (h *Handler) AddModule(mName ModuleName, module *rpc.Client) {
    if _, ok := h.Modules[mName]; !ok {
        h.Modules[mName] = module
    }
}

func (h *Handler) AddCommand(cmd CommandName, mName ModuleName) {
    if _, ok := h.Commands[cmd]; !ok {
        h.Commands[cmd] = mName
    }
}

func (h *Handler) AddInternalCommand(cname CommandName, cfunc *Command) {
    h.InternalCommands[cname] = cfunc
}

func (h *Handler) AddNumeric() {

}

func (h *Handler) IsInternalCommand(cmd CommandName) bool {
    _, ok := h.InternalCommands[cmd]
    return ok
}

func (h *Handler) Invoke(msg *irc.Message, cmd CommandName, args []string) {
    if h.IsInternalCommand(cmd) {
        h.InvokeInternal(msg, cmd, args)
    } else {
        mName, ok := h.Commands[CommandName(strings.ToLower(string(cmd)))]
        if !ok {
            return
        }
        result := ""
        h.Modules[mName].Call(string(mName) + ".InvokeCommand",
                            &CommandData{msg, cmd, args}, &result)
        // Handle Error
    }
}

func (h *Handler) InvokeInternal(msg *irc.Message, cmd CommandName, args []string) {
    hook := h.InternalCommands[cmd]
    hook.Fun(msg, args)
}
