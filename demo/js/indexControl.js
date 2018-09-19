const serverUrl = "http://localhost:4000";
let multiFactorAuthToken;

let custom_interface_option = {};
custom_interface_option.templateName = 'loginradiuscustom_tmpl';
LRObject.util.ready(function() {
    LRObject.customInterface(".interfacecontainerdiv", custom_interface_option);
});

let sl_options = {};
sl_options.onSuccess = function(response) {
    console.log(response);
    localStorage.setItem("LRTokenKey", response.access_token);
    localStorage.setItem("lr-user-uid", response.Profile.Uid);
    window.location.replace("profile.html");
};
sl_options.onError = function(errors) {
    console.log(errors);
};
sl_options.container = "sociallogin-container";

LRObject.util.ready(function() {
    LRObject.init('socialLogin', sl_options);
});

$("#btn-minimal-login").click(function() {
    data = {
        "Email" : $("#minimal-login-email").val(),
        "Password" : $("#minimal-login-password").val()
    }
    $.ajax({
        method: "POST",
        data: JSON.stringify(data),
        url: serverUrl + "/login/email",
        contentType: "application/json; charset=utf-8",
        error: function(xhr) {
            $("#minimal-login-message").text(xhr.responseJSON.Description);
            $("#minimal-login-message").attr("class", "error-message");
        }
    }).done(function(ret) {
        localStorage.setItem("LRTokenKey", ret.access_token);
        localStorage.setItem("lr-user-uid", ret.Profile.Uid);
        window.location.replace("profile.html");
    });
});

$("#btn-minimal-mfalogin-next").click(function() {
    data = {
        "Email" : $("#minimal-mfalogin-email").val(),
        "Password" : $("#minimal-mfalogin-password").val()    
    }
    $.ajax({
        method: "POST",
        data: JSON.stringify(data),
        url: serverUrl + "/mfa/login/email",
        contentType: "application/json",
        error: function(xhr) {
            $("#minimal-mfalogin-message").text(xhr.responseJSON.Description);
            $("#minimal-mfalogin-message").attr("class", "error-message");
        }
    }).done(function(ret) {
        $("#minimal-mfalogin-message").text("");
        if (ret.SecondFactorAuthentication.SecondFactorAuthenticationToken != "") {
            if (ret.SecondFactorAuthentication.IsGoogleAuthenticatorVerified === false) {
                $("#minimal-mfalogin-qrcode").append('<img src="' + ret.SecondFactorAuthentication.QRCode + '">');
            }
            $("#minimal-mfalogin-next")
                .html('<table><tbody><tr>' +
                    '<td>Google Authenticator Code: </td><td><input type="text" id="minimal-mfalogin-googlecode"></td>' +
                    '</tr></tbody></table>' + 
                    '<button id="btn-minimal-mfalogin-login">Login</button>');
            multiFactorAuthToken = ret.SecondFactorAuthentication.SecondFactorAuthenticationToken;
        } else {
            localStorage.setItem("LRTokenKey", ret.access_token);
            localStorage.setItem("lr-user-uid", ret.Profile.Uid);
            window.location.replace("profile.html");
        }
    });
});

$("#minimal-mfalogin-next").on('click', "#btn-minimal-mfalogin-login", function() {
    data = {
        "googleauthenticatorcode" : $("#minimal-mfalogin-googlecode").val()    
    }
    
    $.ajax({
        method: "PUT",
        data: JSON.stringify(data),
        url: serverUrl + "/mfa/google/auth?multi_factor_auth_token=" + multiFactorAuthToken,
        contentType: "application/json",
        error: function(xhr) {
            $("#minimal-mfalogin-message").text(xhr.responseJSON.Description);
            $("#minimal-mfalogin-message").attr("class", "error-message");
        }
    }).done(function(ret) {
        localStorage.setItem("LRTokenKey", ret.access_token);
        localStorage.setItem("lr-user-uid", ret.Profile.Uid);
        window.location.replace("profile.html");
    });
});

$("#btn-minimal-pwless").click(function() {
    $.ajax({
        method: "GET",
        url: serverUrl + "/login/passwordless?email=" + $("#minimal-pwless-email").val() + "&verification_url=" + commonOptions.verificationUrl,
        error: function(xhr){
            $("#minimal-pwless-message").text(xhr.responseJSON.Description);
            $("#minimal-pwless-message").attr("class", "error-message");
        }
    }).done(function() {
        $("#minimal-pwless-message").text("Check your email for the login link.");
        $("#minimal-pwless-message").attr("class", "success-message");
    });
});

$("#btn-minimal-signup").click(function() {
    if($("#minimal-signup-password").val() != $("#minimal-signup-confirmpassword").val()) {
        $("#minimal-signup-message").text("Passwords do not match!");
        $("#minimal-signup-message").attr("class", "error-message");
        return
    }
    let data = {
        "Email": [
          {
            "Type": "Primary",
            "Value": $("#minimal-signup-email").val()
          }
        ],
        "Password": $("#minimal-signup-password").val()
    }

    $.ajax({
        method: "POST",
        url: serverUrl + "/register?verification_url=" + commonOptions.verificationUrl,
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#minimal-signup-message").text(xhr.responseJSON.Description);
            $("#minimal-signup-message").attr("class", "error-message");
        }
    }).done(function() {
        $("#minimal-signup-message").text("Check your email to verify your account.");
        $("#minimal-signup-message").attr("class", "success-message");
    });
});

$("#btn-minimal-forgotpassword").click(function() {
    data = {
        "Email" : $("#minimal-forgotpassword-email").val()    
    }
    $.ajax({
        method: "POST",
        data: JSON.stringify(data),
        url: serverUrl + "/forgotpassword?reset_password_url=" + commonOptions.resetPasswordUrl,
        contentType: "application/json",
        error: function(xhr){
            $("#minimal-forgotpassword-message").text(xhr.responseJSON.Description);
            $("#minimal-forgotpassword-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#minimal-forgotpassword-message").text("Check your email to start the password reset process.");
        $("#minimal-forgotpassword-message").attr("class", "success-message");
    });
});
