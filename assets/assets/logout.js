function deleteAllCookies() {
    var cookies = document.cookie.split(";");

    console.log("deleteAllCookies()\n");
    for (var i = 0; i < cookies.length; i++) {
        var cookie = cookies[i];
        var eqPos = cookie.indexOf("=");
        var name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
        console.log("deleteCookie " + name );
        document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
    }
    console.log("Force logout");
    window.location.href = '/logout.html';
}
