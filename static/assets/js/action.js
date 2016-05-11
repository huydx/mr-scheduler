$(function(){
    $(".nav-input").on('keypress', function(e) {
        if (e.keycode || e.which == 13) {
            var target = $(this).attr("rel");
            console.log("scroll to", target, $("#" + target).offset().top);
            $('html, body').animate({
                scrollTop: $("#" + target).offset().top
            }, 500);
        }
    });
});
