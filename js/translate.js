var arrLang = {
    'en': {
        'Itme_HOME': 'HOME',
        'Itme_PROFILE': 'PROFILE',
        'Itme_RESUME': 'RESUME',
        'Itme_PORTFOLIO': 'PORTFOLIO',
        'Itme_CONTACT': 'CONTACT',
        'PROFILE_Hi': 'Hi!',
        'PROFILE_myName': 'I’m Kevin Lin',
        'PROFILE_Education_title': 'Education',
        'PROFILE_Education1': ' • Da-Yeh University CSIE.',
        'PROFILE_Personality_title1': '',
        'PROFILE_Personality_content1': '',
        'PROFILE_Personality_title2': '',
        'PROFILE_Personality_content2': '',
        'PROFILE_Personality_title3': '',
        'PROFILE_Personality_content3': '',
        'RESUME_skills':'Skill Sets',
        'RESUME_skills_front_end':'Front End：',
        'RESUME_skills_rear_end':'Rear End：',
        'RESUME_skills_others':'Others：',
        'RESUME_HTML_content':'',
        'RESUME_CSS_content':'',
        'RESUME_JS_content':'',
    },
    'zh-tw': {
        'Itme_HOME': '首頁',
        'Itme_PROFILE': '自我簡介',
        'Itme_RESUME': '履歷',
        'Itme_PORTFOLIO': '作品集',
        'Itme_CONTACT': '與我聯繫',
        'PROFILE_Hi': '嗨!',
        'PROFILE_myName': '我叫林OO',
        'PROFILE_Education_title': '學歷',
        'PROFILE_Education1': ' • 大葉大學資訊工程系',
        'PROFILE_Personality_title1': '熱於學習及分享技術',
        'PROFILE_Personality_content1': '主動學習..........................................................................................',
        'PROFILE_Personality_title2': '跨領域學習',
        'PROFILE_Personality_content2': '擁有前、後端、資料庫管理等...........................................................................................',
        'PROFILE_Personality_title3': '勇於溝通以及表達想法',
        'PROFILE_Personality_content3': '溝通可以使我們更了解客戶需求，表達想法可以使客戶更了解自己的需求，將此兩種可以創造出更友好地設計。',
        'RESUME_skills':'技能',
        'RESUME_skills_front_end':'前端：',
        'RESUME_skills_rear_end':'後端：',
        'RESUME_skills_others':'其他：',
        'RESUME_HTML_content':'<div>切板、<table>切板',
        'RESUME_CSS_content':'能配合HTML排板、熟悉選擇器、美化UI',
        'RESUME_JS_content':'互動功能開發、配合CSS開發模組元件、串接後端、熟悉JQuery',
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