package tmpl

const JSMTemplate = `"use strict";

const Module = require('jsrml');

class {{.Name}} extends Module {
    constructor() { 
        super("{{.Name}}", "An echo module");
    }

    echo(msg, args) {
        this.say(args.join(" "));
	}
}

if (require.main === module) {
    const m = new {{.Name}}();

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
`
