// Copyright (C) 2015  Rodolfo Castillo-Valladares & Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
// Send any inquiries you may have about this program to: rcvallada@gmail.com

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
