const { JSDOM } = require( "jsdom" );
const { window } = new JSDOM( "" );
const $ = require( "jquery" )( window );

function callPythonScript(url) {
    $.ajax({
        type: "POST",
        url: url,
    }).done(function (o) {
        console.log(o)
    });
}

exports.handler = async function (event, context) {
    console.log("Received event:", event);
    callPythonScript("~/get-credentials.py");
    callPythonScript("~/main.py");
};

