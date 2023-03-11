let $ = require( "jquery" );

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

