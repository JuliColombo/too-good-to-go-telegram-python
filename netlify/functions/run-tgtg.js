let XMLHttpRequest = require('xhr2');

function callPythonScript(url) {
    let request = new XMLHttpRequest();
    request.open("POST", url, false);
    request.send();
}

exports.handler = async function (event, context) {
    console.log("Received event:", event);
    callPythonScript("~/get-credentials.py");
    callPythonScript("~/main.py");
};

