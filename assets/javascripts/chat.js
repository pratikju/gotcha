$(document).ready(function(){
    $('.progress').hide();
    var name_color_map = {};
    var jsonObj = JSON.parse($('#data').text());
    socket_addr = 'ws://'+ window.location.host +'/websocket';
    var websocket = new WebSocket(socket_addr);

    websocket.onopen = function(response) {
        websocket.send(jsonObj.name + ' joined the chat.');
    }

    websocket.onmessage = function(response) {
        var regEx = /(.*)~~(.*)$/;
        var uploadregEx = /(.*)```(.*)$/;
        var message_array = regEx.exec(response.data.replace(/\n/g,'<br/>'));
        var div_id = "div" + get_random_divID();

        if (message_array != null) { //handle exchanged messages
            message_name =  message_array[1];
            message_value = message_array[2];
            //handle orientation of messages
            if (message_name != jsonObj.name) { //received messages
                $('#chat_box').append('<div id="{id}" class="messages pull-left">'.interpolate({id: div_id}));
                $('#' + div_id).append('<p style="color:{color}"><strong>{content}</strong></p>'
                                       .interpolate({color: find_suitable_color(message_name, get_random_color()), content: message_name}));
                $('#notify')[0].play();
            } else { //sent messages
                $('#chat_box').append('<div id="{id}" class="mymessages pull-right">'.interpolate({id: div_id}));
            }

            // handle different types of messages -- link, upload(image, other files), normal
            if (is_message_a_link(message_value)) { //link message
                $('#' + div_id).append('<a href="{link}" target="_blank">{link}</a>'.interpolate({link: message_value}));
            } else if (uploadregEx.test(message_value)) {  //upload message
                upload_content = uploadregEx.exec(message_value);
                upload_url = upload_content[1];
                upload_type  = upload_content[2];
                if (/image/.test(upload_type)) { //image files
                    $('#' + div_id).append('<div class="image_link"><a href="/uploads/{link}" title="{link}"><img src="/uploads/{link}" alt="{link}"></a></div>'
                                           .interpolate({link: upload_url}));
                } else {
                    $('#' + div_id).append('<a href="/uploads/{link}" download><span class="glyphicon glyphicon-download"></span>{link}</a>'
                                           .interpolate({link: upload_url}));
                }
            } else { //normal message with emojis
                $('#' + div_id).append('<p>{content}</p>'.interpolate({content: message_value}));
                emojify.setConfig({img_dir : '/assets/images/emojis'});
                emojify.run();
            }

            $('#chat_box').append('</div><hr/>');

        } else { //handle welcome messages
            $('#chat_box').append('<p style="text-align:center;font-weight:bold">{content}</p><br/>'.interpolate({content: response.data}));
            $('#notify')[0].play();
        }

        $('#chat_box').animate({scrollTop: $('#chat_box').prop("scrollHeight")},'fast');
    }

    perform_file_upload(websocket, jsonObj);

    perform_chat_operations(websocket, jsonObj);

    initialize_magnific_popup();

    String.prototype.interpolate = function(object) {
        return this.replace(/{([^{}]*)}/g, function(a, b) {
                    var r = object[b];
                    return typeof r === 'string' || typeof r === 'number' ? r : a;
                });
    };

    var get_random_color = function() {
        return 'rgb(' + (Math.floor(Math.random() * 150)) + ',' + (Math.floor(Math.random() * 150)) + ',' + (Math.floor(Math.random() * 150)) + ')';
    };

    var get_random_divID = function() {
        var min = 0;
        var max = 50000;
        return Math.floor(Math.random() * (max - min + 1)) + min;
    };

    var is_message_a_link = function(message) {
        var regEx = /(http|https|ftp|ftps)\:\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,3}(\/\S*)?/ ;
        if (regEx.test(message)) {
            return true;
        }
        return false;
    };

    var find_suitable_color = function(name, random_color) {
      if (name_color_map[name] === undefined) {
          name_color_map[name] = random_color;
      }
      return name_color_map[name];
    };

});

var perform_file_upload = function (websocket, jsonObj) {
    $('#fileupload').fileupload({
        dataType: 'json',
        progressall: function (e, data) {
            $('.progress').show();
            progress = Math.floor(data.loaded / data.total * 100);
            if (progress == 100) {
                $('.progress-bar').addClass('progress-bar-success');
            }
            $('.progress-bar').attr('aria-valuenow', progress).css('width', progress + '%').text(progress + '%');
        },
        done: function (e, data) {
            $('.progress').delay(1000).slideUp("slow",function(){
                $.each(data.result, function (index, file) {
                    websocket.send(jsonObj.name + '~~' + file.name + '```'+ file.type)
                });
                $('.progress-bar').removeClass('progress-bar-success');
            });
        }
    });
};

var perform_chat_operations = function(websocket, jsonObj) {

    var avatar_src;
    if (jsonObj.avatar_url) {
        avatar_src = jsonObj.avatar_url;
    } else {
        avatar_src = jsonObj.picture;
    }
    $("#avatar").attr("src", avatar_src);
    $('#chat_prompt').val('');
    $('#send').on('click',function(){
        if ($('#chat_prompt').val().trim() === "") {
            return false;
        }
        websocket.send(jsonObj.name + '~~' +$('#chat_prompt').val())
        $('#chat_prompt').val('');
    });

    $('#leave_chat').on('click',function(){
        $('#chat_box').append('<p style="text-align:center;font-weight:bold">you left the chat.</p><br/>');
        websocket.send(jsonObj.name + ' left the chat.');
        websocket.close();
    });

    $('#chat_prompt').keypress(function (e) {
        var key = e.which;
        if (key == 13) {
            $('#send').trigger('click');
            return false;
        }
    });

    $('#clear_chat').on('click',function(){
        $('#chat_box').html('');
    });
};

var initialize_magnific_popup = function() {
    $('.msg_container_base').magnificPopup({
        delegate: '.image_link > a',
        type: 'image',
        closeOnContentClick: true,
        closeBtnInside: false,
        fixedContentPos: true,
        mainClass: 'mfp-no-margins mfp-with-zoom',
        image: {
            verticalFit: true
        },
        zoom: {
            enabled: true,
            duration: 300
        },
        gallery:{
            enabled:true,
            preload: [0,2],
            navigateByImgClick: true,
            tPrev: 'Previous (Left arrow key)',
            tNext: 'Next (Right arrow key)',
            tCounter: '<span class="mfp-counter">%curr% of %total%</span>'
        }
    });
};
