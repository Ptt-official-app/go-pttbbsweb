/* Copyright (c) 2015 ptt.cc */
/* Licensed under The MIT License. */
(function() {

var $ = jQuery;

var InfoRefollow = '原文已被修改，已於結尾重新開始跟隨';

var ErrArticleModified = '原文已被修改，請嘗試重新載入此頁';
var ErrNetwork = '無法更新推文';

function attachPoller(div) {
    var timer = null;
    var autoScroll = false;
    var autoUpdate = false;
    var lastModified = null;
    var etag = null;
    var currSize = parseInt(div.attr('data-offset'));
    var longPollUrl = div.attr('data-longpollurl');
    var pollUrl = div.attr('data-pollurl');
    var currReq = null;
    var fatalError = null;

    function scheduleNextPoll() {
        if (timer) {
            window.clearTimeout(timer);
            timer = null;
        }
        timer = window.setTimeout(longPoll, 1000);
    }

    function longPoll() {
        var headers = {};
        if (lastModified != null) {
            headers['If-Modified-Since'] = lastModified;
            headers['If-None-Match'] = etag;
        }
        currReq = $.ajax({
            url: longPollUrl,
            headers: headers,
            timeout: 28000
        }).success(function(data, textStatus, req) {
            currReq = null;
            console.log(data);
            lastModified = req.getResponseHeader('Last-Modified');
            etag = req.getResponseHeader('Etag');
            if (data.size > currSize) {
                requestContent(data);
            } else {
                scheduleNextPoll();
            }
        }).fail(function(jqXHR, textStatus, errorThrown) {
            currReq = null;
            if (textStatus === 'timeout') {
                scheduleNextPoll();
            } else if (textStatus !== 'abort') {
                setFatalError(ErrNetwork);
            }
        });
    }

    function requestContent(lpdata) {
        currSize = lpdata.size;
        var url = pollUrl
            + '&size=' + encodeURIComponent(lpdata.size.toString())
            + '&size-sig=' + encodeURIComponent(lpdata.sig);
        $.ajax(url).success(function(data) {
            console.log(data);
            if (data.success) {
                receivedPushContent(data.contentHtml);
                pollUrl = data.pollUrl;
                scheduleNextPoll();
            } else {
                refollow(lpdata);
            }
        }).fail(function(jqXHR, textStatus, errorThrown) {
            setFatalError(ErrNetwork);
        });
    }

    function receivedPushContent(contentHtml) {
        if (contentHtml.length > 0) {
            $('#main-content').append(
                    $('<div>').html(contentHtml).addClass('new-push'));
            if (autoScroll)
                $(document.body).animate({scrollTop: document.body.clientHeight - window.innerHeight}, 500);
        }
    }

    function refollow(data) {
        if (!data.cacheKey) {
            return;
        }
        var i = pollUrl.indexOf('?');
        if (i < 0) {
            return;
        }
        pollUrl = pollUrl.substr(0, i)
            + '?cacheKey=' + encodeURIComponent(data.cacheKey)
            + '&offset=' + encodeURIComponent(data.size.toString())
            + '&offset-sig=' + encodeURIComponent(data.sig.toString());
        $('#main-content').append($('<div>').append($('<span>').text(InfoRefollow).addClass('push-stream-info')));
        scheduleNextPoll();
    }

    function setAutoUpdate(enabled) {
        if (currReq != null) {
            currReq.abort();
            currReq = null;
        }
        if (enabled) {
            scheduleNextPoll();
        } else if (timer) {
            window.clearTimeout(timer);
            timer = null;
        }
        autoUpdate = enabled;
        autoScroll = enabled;
        updateStatus();
    }

    function toggleAuto() {
        if (fatalError != null) {
            return;
        }
        setAutoUpdate(!autoUpdate);
    }

    function updateStatus() {
        if (fatalError != null) {
            div.text(fatalError);
            div.addClass('fatal-error');
            return;
        }
        div.text('推文' + (autoUpdate ?
                    ('會自動更新，並' + (autoScroll ? '會' : '不會') + '自動捲動')
                    : '自動更新已關閉'));
    }

    function setFatalError(msg) {
        fatalError = msg;
        updateStatus();
    }

    div.click(toggleAuto);
    setAutoUpdate(false);
}

function attachSearchBar(form) {
    form.submit(function() {
        var q = $('.query', form);
        if (q.val().trim().length == 0)
            return false;
    });
}

function attachDropdown() {
    $('.article-menu > .trigger').click(function(e) {
        var isShown = e.target.parentElement.classList.contains('shown');
        $('.article-menu.shown').toggleClass('shown');
        if (!isShown) {
            e.target.parentElement.classList.toggle('shown');
        }
        e.stopPropagation();
    });
    $(document).click(function(e) {
        $('.article-menu.shown').toggleClass('shown');
    });
}

$(document).ready(function() {
    attachPoller($('#article-polling'));
    attachSearchBar($('#search-bar'));
    attachDropdown();
});

})();
