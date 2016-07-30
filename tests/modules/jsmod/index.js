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

"use strict";

const Module = require('jsrml');

class Jsmod extends Module {
    constructor() { 
        super("Jsmod", "A test module");
    }

    echo(msg, args) {
        this.say(args.join(" "));
	}
}

if (require.main === module) {
    const m = new Jsmod();

    m.initialize()
    .then(() => {
        m.addCommand("echo", {
            Help: 'Repeats its arguments',
            Fun: m.echo
        });

        m.register();
    })
    .fail((error) => {
        console.error(error);
    });
}