function page_load() {
    $.ajax({
        url: '/selectMenu',
        type: 'post',
        dataType: 'json',
        success: function(result) {
            var MenuHTML = "<div><ul>";
            for (var i = 0; i < result.length; i++) {                
                // MenuHTML += "<li onclick=\"gotoURL(\'" + result[i]["html_url"] + "\')\">" + result[i]["chinese_name"] + "</li>"
                MenuHTML += "<li onclick=\"gotoURL(\'" + result[i]["html_url"] + "\')\"><a href='#' class='animBtn themeD'>" + result[i]["chinese_name"] + "</a></li>"
            }
            MenuHTML += "</div></ul>";
            $('#menu-item').append(MenuHTML);
        },
        error: function(request, ajaxOptions, thrownError) {}
    });
}

function gotoURL(url) {
    $('#myiframe').attr("src", "../" + url);
}