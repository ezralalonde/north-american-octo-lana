$(document).ready(function() {
	$('a').click(function (e) {
		loadPage()
	});	
});

function loadPage()
{
    $.ajax({
		type: "POST",
		url: "/lunch/",
		dataType: "text",
		success: function(msg) {
			$('#change').html(msg);
		}
	});
}
