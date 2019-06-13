var options = {
  redirecturl: {
    afterlogin: "http://localhost:3000/index",
    afterreset: "http://localhost:3000/index",
  },
  socialsquarestyle: false,
  pagesshown: ["login", "signup", "forgotpassword"]
};

LRObject.util.ready(function () {
  LRObject.loginScreen("loginscreen-container", options)
});