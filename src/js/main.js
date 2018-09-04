window.ga = function() {
  ga.q.push(arguments);
};
ga.q = [];
ga.l = +new Date();
ga("create", "UA-XXXXX-Y", "auto");
ga("send", "pageview");

$(document).ready(function() {
  $.getJSON("/votd/", function(data) {
    console.log(data);
    $("#verseBox").append(data.verse.details.text);
  });
});
