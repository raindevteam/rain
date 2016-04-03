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