const serverUrl = "http://localhost:4000";
let update = {};

$( "#btn-user-changepassword" ).click(function() {
    let data = {
        "oldpassword" : $("#user-changepassword-oldpassword").val(),    
        "newpassword" : $("#user-changepassword-newpassword").val()
    }
    $.ajax({
        method: "PUT",
        url: serverUrl + "/profile/changepassword?auth=" + localStorage.getItem("LRTokenKey"),
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-changepassword-message").text(xhr.responseJSON.Description);
            $("#user-changepassword-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-changepassword-message").text("Password successfully changed.");
        $("#user-changepassword-message").attr("class", "success-message");
    });
});

$("#btn-user-setpassword").click(function() {
    let data = {
        "password" : $("#user-setpassword-password").val()
    }
    $.ajax({
        method: "PUT",
        url: serverUrl + "/profile/setpassword?uid=" + localStorage.getItem("lr-user-uid"),
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr){
            $("#user-setpassword-message").text(xhr.responseJSON.Description);
            $("#user-setpassword-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-setpassword-message").text("Password successfully changed.");
        $("#user-setpassword-message").attr("class", "success-message");
    });
});

$("#btn-user-updateaccount").click(function() {
    let data = {};
    let dataFields = {
        "FirstName": $("#user-updateaccount-firstname").val(),
        "LastName": $("#user-updateaccount-lastname").val(),
        "About": $("#user-updateaccount-about").val()
    }

    for (let key in dataFields) {
        if (dataFields[key] !== "") {
            data[key] = dataFields[key];
        } else {
            data[key] = update[key];
        }
    }

    $.ajax({
        method: "PUT",
        url: serverUrl + "/profile/update?uid=" + localStorage.getItem("lr-user-uid"),
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-updateaccount-message").text(xhr.responseJSON.Description);
            $("#user-updateaccount-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-updateaccount-message").text("Account successfully updated.");
        $("#user-updateaccount-message").attr("class", "success-message");
        profileUpdate();
    });
});

$("#btn-user-createcustomobj").click(function() {
    let data;
    try {
        data = JSON.parse($("#user-createcustomobj-data").val());
    } catch(e) {
        $("#user-createcustomobj-message").text("Please input a valid JSON object in the data field.");
        $("#user-createcustomobj-message").attr("class", "error-message");
    }

    $.ajax({
        method: "POST",
        url: serverUrl + "/customobj?object_name=" + $("#user-createcustomobj-objectname").val() + "&auth=" + localStorage.getItem("LRTokenKey"),
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-createcustomobj-message").text(xhr.responseJSON.Description);
            $("#user-createcustomobj-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-createcustomobj-message").text("Object successfully created.");
        $("#user-createcustomobj-message").attr("class", "success-message");
    });
});

$("#btn-user-updatecustomobj").click(function() {
    let data = {};

    try {
        data = JSON.parse($("#user-updatecustomobj-data").val());
    } catch(e) {
        $("#user-updatecustomobj-message").text("Please input a valid JSON object in the data field.");
        $("#user-updatecustomobj-message").attr("class", "error-message");
    }

    $.ajax({
        method: "PUT",
        url: serverUrl + "/customobj?object_name=" + $("#user-updatecustomobj-objectname").val() + "&auth=" + localStorage.getItem("LRTokenKey") +
            "&object_id=" + $("#user-updatecustomobj-objectrecordid").val(),
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-updatecustomobj-message").text(xhr.responseJSON.Description);
            $("#user-updatecustomobj-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-updatecustomobj-message").text("Object successfully updated.");
        $("#user-updatecustomobj-message").attr("class", "success-message");
    });
});

$("#btn-user-deletecustomobj").click(function() {
    $.ajax({
        method: "DELETE",
        url: serverUrl + "/customobj?object_name=" + $("#user-deletecustomobj-objectname").val() + "&auth=" + localStorage.getItem("LRTokenKey") +
            "&object_id=" + $("#user-deletecustomobj-objectrecordid").val(),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-deletecustomobj-message").text(xhr.responseJSON.Description);
            $("#user-deletecustomobj-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-deletecustomobj-message").text("Custom object deleted successfully.");
        $("#user-deletecustomobj-message").attr("class", "success-message");
    });
});

$("#btn-user-getcustomobj").click(function() {
    $.ajax({
        method: "GET",
        url: serverUrl + "/customobj?object_name=" + $("#user-getcustomobj-objectname").val() + "&auth=" + localStorage.getItem("LRTokenKey"),
        contentType: "application/json",
        error: function(xhr) {
            $('#table-customobj tr').remove();
            $("#user-getcustomobj-message").text(xhr.responseJSON.Description);
            $("#user-getcustomobj-message").attr("class", "error-message");
        }
    }).done(function(ret) {
        $('#table-customobj tr').remove();
        $('<tr>' +
            '<th>Object ID</th><th>Custom Object</th>' +
            '<tr>').appendTo("#table-customobj > tbody:last-child");

        for (let i = 0; i < ret.data.length; i++) {
            $("<tr><td>" + ret.data[i].Id + "</td></tr>").appendTo("#table-customobj > tbody:last-child");
            $("<td>", {
                text: JSON.stringify(ret.data[i].CustomObject)
            }).appendTo("#table-customobj > tbody:last-child > tr:last-child");
        }
    });
});

$("#btn-user-mfa-resetgoogle").click(function() {
    let data = {
        "googleauthenticator": true
    }

    $.ajax({
        method: "DELETE",
        url: serverUrl + "/mfa/google?auth=" + localStorage.getItem("LRTokenKey"),
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-mfa-message").text(xhr.responseJSON.Description);
            $("#user-mfa-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-mfa-message").text("Google Authenticator settings reset.");
        $("#user-mfa-message").attr("class", "success-message");
    });
});

$("#btn-user-mfaenable").click(function() {
    $.ajax({
        method: "GET",
        url: serverUrl + "/mfa/validate?auth=" + localStorage.getItem("LRTokenKey"),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-mfaenable-message").text(xhr.responseJSON.Description);
            $("#user-mfaenable-message").attr("class", "error-message");
        }
    }).done(function(ret) {
        if (ret.IsGoogleAuthenticatorVerified === false) {
            $("#user-mfaenable-qrcode").append('<img src="' + ret.QRCode + '">');
        }
        $("#user-mfaenable").html('<table><tbody><tr>' +
            '<td>Google Authenticator Code: </td><td><input name="googleauth" type="text" id="user-mfaenable-googleauth"></td>' +
            '</tr></tbody></table>' +
            '<button id="btn-user-mfaenable-googleauth">Submit</button>');
    });
});

$("#user-mfaenable").on("click", "#btn-user-mfaenable-googleauth", function() {
    let data = {
        "googleauthenticatorcode": $("#user-mfaenable-googleauth").val()
    }

    $.ajax({
        method: "PUT",
        url: serverUrl + "/mfa/google/enable?auth=" + localStorage.getItem("LRTokenKey"),
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-mfaenable-message").text(xhr.responseJSON.Description);
            $("#user-mfaenable-message").attr("class", "error-message");
        }
    }).done(function() {
        $("#user-mfaenable-message").text("MFA has been enabled.");
        $("#user-mfaenable-message").attr("class", "success-message");
    });
});

$( "#btn-user-createrole" ).click(function() {
    let data = {
        "roles" : [
            { "Name": $("#user-roles-createrole").val() }
        ]
    }

    $.ajax({
        method: "POST",
        url: serverUrl + "/roles",
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-createrole-message").text(xhr.responseJSON.Description);
            $("#user-createrole-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-createrole-message").text("Role created successfully.");
        $("#user-createrole-message").attr("class", "success-message");
        roleUpdate();
    });
});

$( "#btn-user-deleterole" ).click(function() {
    $.ajax({
        method: "DELETE",
        url: serverUrl + "/roles?role=" + $("#user-roles-deleterole").val(),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-deleterole-message").text(xhr.responseJSON.Description);
            $("#user-deleterole-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-deleterole-message").text("Role deleted successfully.");
        $("#user-deleterole-message").attr("class", "success-message");
        roleUpdate();
    });
});

$( "#btn-user-assignrole" ).click(function() {
    let data = {
        "Roles" : [
            $("#user-roles-assignrole").val()
        ]
    }

    $.ajax({
        method: "PUT",
        url: serverUrl + "/roles?uid=" + localStorage.getItem("lr-user-uid"),
        data: JSON.stringify(data),
        contentType: "application/json",
        error: function(xhr) {
            $("#user-assignrole-message").text(xhr.responseJSON.Description);
            $("#user-assignrole-message").attr("class", "error-message");
        }
    }).done(function() { 
        $("#user-assignrole-message").text("Role added to current user successfully.");
        $("#user-assignrole-message").attr("class", "success-message");
        roleUpdate();
    });
});

let profileUpdate = function() {
    if(localStorage.getItem("LRTokenKey") === null) {
        window.location.replace("index.html");
        return;
    }

    $.ajax({
        method: "GET",
        url: serverUrl + "/profile?auth=" + localStorage.getItem("LRTokenKey"),
        error: function(){
            localStorage.removeItem("LRTokenKey");
            localStorage.removeItem("lr-user-uid");
            window.location.replace("index.html");
        }
    }).done(function(ret) {
        $("#profile-name").html("<b>" + ret.FullName + "</b>");
        $("#profile-provider").text("Provider: " + ret.Provider);
        $("#profile-email").text(ret.Email[0].Value);
        $("#profile-lastlogin").text("Last Login Date: " + ret.LastLoginDate);
        update.firstName = ret.FirstName;
        update.lastName = ret.LastName;
        update.about = ret.About;
    });
}

let roleUpdate = function() {
    $.ajax({
        method: "GET",
        url: serverUrl + "/roles",
        error: function(xhr) {
            $("#minimal-verification-message").text(xhr.responseJSON.Description);
            $("#minimal-verification-message").attr("class", "error-message");
        }
    }).done(function(ret) {
        $('#table-allroles tr:not(:first)').remove();
        for (let i = 0; i < ret.data.length; i++) {
            $("<tr></tr>").appendTo("#table-allroles > tbody:last-child");
            $("<td>", {
                text: ret.data[i].Name
            }).appendTo('#table-allroles > tbody:last-child > tr:last-child');
        }
    });


    $.ajax({
        method: "GET",
        url: serverUrl + "/roles/get?uid=" + localStorage.getItem("lr-user-uid"),
        error: function(xhr) {
            $("#minimal-verification-message").text(xhr.responseJSON.Description);
            $("#minimal-verification-message").attr("class", "error-message");
        }
    }).done(function(ret) {
        $('#table-userroles tr:not(:first)').remove();
        if (ret.Roles) {
            for (let i = 0; i < ret.Roles.length; i++) {
                $("<tr></tr>").appendTo("#table-userroles > tbody:last-child");
                $("<td>", {
                    text: ret.Roles[i]
                }).appendTo('#table-userroles > tbody:last-child > tr:last-child');
            }
        }
    });
}

let script = $(
    '<script type="text/html" id="loginradiuscustom_tmpl_link">' +
    '<# if(isLinked) { #>' +
    '<div class="lr-linked">' +
    '<a class="lr-provider-label" href="javascript:void(0)" title="<#= Name #>" alt="Connected" onclick=\'return LRObject.util.unLinkAccount(\"<#= Name.toLowerCase() #>\",\"<#= providerId #>\")\'><#=Name#> is connected | Delete</a>' +
    '</div>' +
    '<# }  else {#>' +
    '<div class="lr-unlinked">' +
    '<a class="lr-provider-label" href="javascript:void(0)" onclick="return LRObject.util.openWindow(\'<#= Endpoint #>\');" title="<#= Name #>" alt="Sign in with <#=Name#>">' +
    '<#=Name#></a></div>' +
    '<# } #>' +
    '</script>'
);

$("#script-accountlinking").append(script);

let la_options = {};
la_options.container = "interfacecontainerdiv";
la_options.templateName = 'loginradiuscustom_tmpl_link';
la_options.onSuccess = function() {
    $("#interfacecontainerdiv").empty();
    LRObject.util.ready(function() {
        LRObject.init("linkAccount", la_options);
    });
}
la_options.onError = function(errors) {
    $("#user-accountlinking-message").text(errors[0].Description);
    $("#user-accountlinking-message").attr("class", "error-message");
}

let unlink_options = {};
unlink_options.onSuccess = function() {
    $("#interfacecontainerdiv").empty();
    LRObject.util.ready(function() {
        LRObject.init("linkAccount", la_options);
    });
}
unlink_options.onError = function(errors) {
    $("#user-accountlinking-message").text(errors[0].Description);
    $("#user-accountlinking-message").attr("class", "error-message");
}

LRObject.util.ready(function() {
    LRObject.init("linkAccount", la_options);
    LRObject.init("unLinkAccount", unlink_options);
});

profileUpdate();
roleUpdate();
