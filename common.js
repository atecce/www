// this is simply to record hits. all it does is log your ip
var xhr = new XMLHttpRequest();
xhr.overrideMimeType("text/plain; charset=x-user-defined");
xhr.open("GET", "https://auth.atec.pub");
xhr.send();