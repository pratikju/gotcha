$(document).ready(function(){
  $('.progress').hide();
  var name_color_map = {};
  var jsonObj = JSON.parse($('#data').text());
  socket_addr = 'ws://'+ jsonObj.context +'/websocket';
  var websocket = new WebSocket(socket_addr);

  console.log("Websocket - status: " + websocket.readyState);
  websocket.onopen = function(res) {
    console.log("CONNECTION opened..." + this.readyState);
    websocket.send(jsonObj.name + ' joined the chat.');
  }
  websocket.onmessage = function(res) {
    var regEx = /(.*)~~(.*)$/;
    var uploadregEx = /(.*)```(.*)$/;
    var dataArray = regEx.exec(res.data.replace(/\n/g,'<br/>'));
    var div_id = "div" + getRandomIntInclusive(0,50000);
    var random_color = 'rgb(' + (Math.floor(Math.random() * 150)) + ',' + (Math.floor(Math.random() * 150)) + ',' + (Math.floor(Math.random() * 150)) + ')';
    if(dataArray != null){
      if(dataArray[1] != jsonObj.name){
        $('#chat_box').append('<div id="{id}" class="messages pull-left">'.interpolate({id: div_id}));
        $('#{id}'.interpolate({id: div_id})).append('<p style="color:{color}"><strong>{content}</strong></p>'.interpolate({color: find_suitable_color(dataArray, random_color), content: dataArray[1]}));
        $('#notify')[0].play();
      }else{
        $('#chat_box').append('<div id="{id}" class="mymessages pull-right">'.interpolate({id: div_id}));
      }

      if(checkForLink(dataArray[2])){
        $('#{id}'.interpolate({id: div_id})).append('<a href="{link}" target="_blank">{link}</a>'.interpolate({link: dataArray[2]}));
      }else if(uploadregEx.exec(dataArray[2]) != null){
          uploadedUrl = uploadregEx.exec(dataArray[2])[1];
          $('#{id}'.interpolate({id: div_id})).append('<a href="/assets/uploads/{link}" target="_blank" download><span class="glyphicon glyphicon-download"></span>{link}</a>'.interpolate({link: uploadedUrl}));
      }else{
        $('#{id}'.interpolate({id: div_id})).append('<p>{content}</p>'.interpolate({content: dataArray[2]}));
        emojify.setConfig({img_dir : '/assets/images/emojis'});
        emojify.run();
      }
      $('#chat_box').append('</div><hr/>');
    }else{
      $('#chat_box').append('<p style="text-align:center;font-weight:bold">{content}</p><br/>'.interpolate({content: res.data}));
      $('#notify')[0].play();
    }

    $('#chat_box').animate({scrollTop: $('#chat_box').prop("scrollHeight")},'fast');
  }

  websocket.onerror = function(res) {
    console.log("Error occured sending..." + m.data);
  }
  websocket.onclose = function(res) {
    console.log("Disconnected - status " + this.readyState);
  }

  String.prototype.interpolate = function (o) {
    return this.replace(/{([^{}]*)}/g,
    function (a, b) {
      var r = o[b];
      return typeof r === 'string' || typeof r === 'number' ? r : a;
    });
  };

  var getRandomIntInclusive = function (min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  var checkForLink = function(text){
    var regEx = /(http|https|ftp|ftps)\:\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,3}(\/\S*)?/ ;
    if(regEx.test(text)){
      return true;
    }
    return false;
  }

  var find_suitable_color = function(dataArray, random_color){
    if(name_color_map[dataArray[1]] === undefined){
      name_color_map[dataArray[1]] = random_color;
    }
    return name_color_map[dataArray[1]];
  }

  $('#fileupload').fileupload({
    dataType: 'json',
    progressall: function (e, data) {
      $('.progress').show();
      progress = Math.floor(data.loaded / data.total * 100);
      if(progress == 100){
        $('.progress-bar').addClass('progress-bar-success');
      }
      $('.progress-bar').attr('aria-valuenow', progress).css(
        'width',
        progress + '%'
      ).text(progress + '%');

    },
    done: function (e, data) {

      $('.progress').delay(2000).slideUp("slow",function(){
        $.each(data.result, function (index, file) {
          websocket.send(jsonObj.name + '~~' + file.name + '```link')
        });
        $('.progress-bar').removeClass('progress-bar-success');
      });
    }
  });



  $('#chat_prompt').val('');
  $('#send').on('click',function(){
    if($('#chat_prompt').val().trim() === ""){
      return false;
    }
    websocket.send(jsonObj.name + '~~' +$('#chat_prompt').val())
    $('#chat_prompt').val('');
  });
  $('#clear_chat').on('click',function(){
    $('#chat_box').html('');
  });
  $('#leave_chat').on('click',function(){
    $('#chat_box').append('<p style="text-align:center;font-weight:bold">you left the chat.</p><br/>');
    websocket.send(jsonObj.name + ' left the chat.');
    websocket.close();
  });

  $('#chat_prompt').keypress(function (e) {
    var key = e.which;
    if(key == 13){
      $('#send').trigger('click');
      return false;
    }
  });
});
