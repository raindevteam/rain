from pyrml import Module

class Pymod(Module):
    def __init__(self):
        Module.__init__(self, "Pymod", "Test module")

    def echo(self, msg, args):
        self.say(msg["Args"][0], " ".join(args))

if __name__ == '__main__':
    m = Pymod()

    m.add_command("echo", {
        "Help": "Repeats arguments",
        "Fun": m.echo
    })

    m.register()
