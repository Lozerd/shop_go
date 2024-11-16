// script.js
document.addEventListener('htmx:afterRequest', function(event) {
    // Handle post-request actions here
    console.log("Request completed: ", event.detail.xhr);
});
