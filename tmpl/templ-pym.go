package tmpl

const PYMTemplate = `from pyrml import Module

class {{.Name}}(Module):
    def __init__(self):
        Module.__init__(self, "{{.Name}}", "Your module")

    def echo(self, msg, args):
        self.say(msg["Args"][0], " ".join(args))

if __name__ == '__main__':
    m = {{.Name}}()

    m.add_command("echo", {
        "Help": "Repeats arguments",
        "Fun": m.echo
    })

    m.register()
`
