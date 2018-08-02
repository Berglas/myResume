var arrLang = {
    'en': {
        'Itme_HOME': 'HOME',
        'Itme_PROFILE': 'PROFILE',
        'Itme_RESUME': 'RESUME',
        'Itme_PORTFOLIO': 'PORTFOLIO',
        'Itme_CONTACT': 'CONTACT'
    },
    'zh-tw': {
        'Itme_HOME': '首頁',
        'Itme_PROFILE': '自我簡介',
        'Itme_RESUME': '履歷',
        'Itme_PORTFOLIO': '作品集',
        'Itme_CONTACT': '與我聯繫'
    }
};

$(function() {
    $('.translate').click(function() {
        var lang = $(this).attr('id');

        $('.lang').each(function(index, element) {
            $(this).text(arrLang[lang][$(this).attr('key')]);
        });
    });
});