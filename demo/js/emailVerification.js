const url = location.href;
const params = url.split("?")[1];
const serverUrl = "http://localhost:4000";
let paramsObj = {};

if (params) {
    paramsObj = JSON.parse('{"' + decodeURI(params.replace(/&/g, "\",\"").replace(/=/g,"\":\"")) + '"}');

    if (paramsObj.vtype === "emailverification") {
        $.ajax({
            method: "GET",
            url: serverUrl + "/register/verify/email?verification_token=" + paramsObj.vtoken,
            contentType: "application/json",
            error: function(xhr){
                $("#minimal-verification-message").text(xhr.responseJSON.Description);
                $("#minimal-verification-message").attr("class", "error-message");
            }
        }).done(function() { 
            $("#minimal-verification-message").text("Email successfully verified.");
            $("#minimal-verification-message").attr("class", "success-message");
        });
    } else if (paramsObj.vtype === "oneclicksignin") {
        $.ajax({
            method: "GET",
            url: serverUrl + "/login/passwordless/auth?verification_token=" + paramsObj.vtoken,
            contentType: "application/json",
            error: function(xhr) {
                $("#minimal-verification-message").text(xhr.responseJSON.Description);
                $("#minimal-verification-message").attr("class", "error-message");
            }
        }).done(function(ret) {
            localStorage.setItem("LRTokenKey", ret.access_token);
            localStorage.setItem("lr-user-uid", ret.Profile.Uid);
            window.location.replace("profile.html");
        });
    } else {
        window.location.replace("index.html");
    }
} else {
    window.location.replace("index.html");
}
