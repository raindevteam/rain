package tmpl

const JSMTemplate = `"use strict";

const Module = require('jsrml');

class {{.Name}} {
    constructor() { super("{{.Name}}", "An echo module") }

    echo(irc, msg, args) {
        irc.say(args.join(" "));
	}
}

function main() {
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

if (require.main === module) {
    main();
}
`
