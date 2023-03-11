exports.handler = async function (event, context) {
    console.log("Received event:", event);
    $.ajax({
        type: "POST",
        url: "~/get-credentials.py",
    }).done(function (o) {
        console.log(o)
    });
    $.ajax({
        type: "POST",
        url: "~/main.py",
    }).done(function (o) {
        console.log(o)
    });
};
